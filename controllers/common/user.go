package common

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/base"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/err"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/models"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/util"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/util/sms"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/vo"
)

type UserController struct {
	base.BaseController
}
type loginInfo struct {
	Account  string
	Password string
}
type sendVerifyCodeInfo struct {
	Phone    string
	Use_page int
}
type registerInfo struct {
	Account     string
	Password    string
	Verify_code string
	Nickname    string
}
type changePasswordInfo struct {
	Old_password string
	New_password string
}
type resetPasswordInfo struct {
	Account      string
	Verify_code  string
	New_password string
}
type newUserInfo struct {
	Nickname string
	Gender   int8
	Img_key  string
}

// @router /login [post]
func (this *UserController) Login() {
	var data loginInfo
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	result, user := models.Login(data.Account, data.Password)
	if result {
		//登录成功
		this.SetSession("userId", user.Id)
		// 获取弹幕token
		util.GetDanmakuToken(user.Id)
		// type 0是普通用户 1是教师
		this.SetSession("type", user.Type)
		if user.Type == 1 {
			//教师，记录teacher表的主键
			teacher := models.GetTeacherByUserId(user.Id)
			this.SetSession("teacherId", teacher.Id)
		}
		var userVo vo.User
		util.PoToVo(user, &userVo)
		result := make(map[string]vo.User)
		result["data"] = userVo
		this.Data["json"] = result
		this.ServeJSON()
	} else {
		this.RetError(err.ErrNoUserPass, "")
	}
}

// @router /sendVerifyCode [post]
func (this *UserController) SendVerifyCode() {
	var data sendVerifyCodeInfo
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)

	phonenumber := data.Phone
	fmt.Printf("phoneNumber %v\n", phonenumber)
	if !util.MatchPhoneNum(phonenumber) {
		this.RetError(err.ErrInputData, "手机号格式不正确")
	}
	use_page := data.Use_page
	exist := models.AccountExist(phonenumber)
	if use_page == 0 {
		//注册
		if exist {
			this.RetError(err.ErrDupUser, "手机号已存在")
		}
		result, msg := sms.SendSmsCode(phonenumber)
		if !result {
			this.RetError(err.ErrSendMessage, msg)
		}
		this.Data["json"] = map[string]string{"data": ""}
		this.ServeJSON()
	}
	//忘记密码
	if !exist {
		this.RetError(err.ErrNoUser, "手机号不存在")
	}
	result, msg := sms.SendSmsCode(phonenumber)
	if !result {
		this.RetError(err.ErrSendMessage, msg)
	}
	this.Data["json"] = map[string]string{"data": ""}
	this.ServeJSON()
}

// @router /register [post]
func (this *UserController) Register() {
	var registerInfo registerInfo
	json.Unmarshal(this.Ctx.Input.RequestBody, &registerInfo)
	if result, msg := sms.VerifyCode(registerInfo.Account, registerInfo.Verify_code); !result {
		this.RetError(err.ErrVerifyCode, msg)
		return
	}
	uid, e := models.AddUser(registerInfo.Account, registerInfo.Password, registerInfo.Nickname)
	if e != nil {
		this.RetError(err.ErrDupUser, "该账户已存在")
		return
	}
	this.Data["json"] = map[string]int64{"data": uid}
	this.ServeJSON()
}

// @router /changePassword [post]
func (this *UserController) ChangePassword() {
	userIdOb := this.GetSession("userId")
	if userIdOb == nil {
		this.RetError(err.ErrUnauthorized, "未登录")
	}
	userId := userIdOb.(int)
	user := models.GetUserById(userId)
	var changePasswdInfo changePasswordInfo
	json.Unmarshal(this.Ctx.Input.RequestBody, &changePasswdInfo)
	if changePasswdInfo.New_password == "" || changePasswdInfo.Old_password == "" {
		this.RetError(err.ErrInputData, "参数不全")
		return
	}
	if util.Sha256(changePasswdInfo.Old_password) == user.Password {
		user.Password = util.Sha256(changePasswdInfo.New_password)
		if !models.UpdateUser(*user) {
			this.RetError(err.ErrDatabase, "更新失败")
			return
		}
	} else {
		this.RetError(err.ErrInputData, "原密码不正确")
		return
	}
	this.Data["json"] = map[string]int64{"data": 1}
	this.ServeJSON()
}

// @router /resetPassword [post]
func (this *UserController) ResetPassword() {
	var resetPasswordInfo resetPasswordInfo
	json.Unmarshal(this.Ctx.Input.RequestBody, &resetPasswordInfo)
	if resetPasswordInfo.Account == "" || resetPasswordInfo.Verify_code == "" || resetPasswordInfo.New_password == "" {
		this.RetError(err.ErrInputData, "参数不全")
		return
	}
	if result, msg := sms.VerifyCode(resetPasswordInfo.Account, resetPasswordInfo.Verify_code); !result {
		this.RetError(err.ErrVerifyCode, msg)
		return
	}
	user := models.GetUserByAccount(resetPasswordInfo.Account)
	if user == nil {
		this.RetError(err.ErrNoUser, "账户不存在")
		return
	}
	user.Password = util.Sha256(resetPasswordInfo.New_password)
	if models.UpdateUser(*user) {
		this.Data["json"] = map[string]int{"data": 1}
		this.ServeJSON()
	} else {
		this.RetError(err.ErrDatabase, "更新失败")
		return
	}

}

// @router /updateInfo [post]
func (this *UserController) UpdateInfo() {
	var newUserInfo newUserInfo
	userId := this.GetSession("userId").(int)
	user := models.GetUserById(userId)
	json.Unmarshal(this.Ctx.Input.RequestBody, &newUserInfo)
	if newUserInfo.Gender != 0 {
		user.Gender = newUserInfo.Gender
	}
	if newUserInfo.Img_key != "" {
		domain := beego.AppConfig.String("qiniu_pic_url")
		user.Img_url = util.GetResourceUrl(domain, newUserInfo.Img_key, false)
	}
	if newUserInfo.Nickname != "" {
		user.Nickname = newUserInfo.Nickname
	}
	if models.UpdateUser(*user) {
		var userVo vo.User
		result := make(map[string]vo.User)
		util.PoToVo(user, &userVo)
		result["data"] = userVo
		this.Data["json"] = result
		this.ServeJSON()
	} else {
		this.RetError(err.ErrDatabase, "更新失败")
		return
	}
}
