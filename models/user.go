package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/util"
	"math/rand"
	"time"
)

type User struct {
	Id          int
	Type        int8
	Nickname    string    `orm:"size(63)"`
	Img_url     string    `orm:"size(255)"`
	Password    string    `orm:"size(64)"`
	Account     string    `orm:"size(63)"`
	Gender      int8      `orm:"size(4)"`
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time `orm:"type(datetime)"`
}

// Login 登录
func Login(username string, password string) (bool, User) {
	var user User
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("account", username).One(&user)
	if err == orm.ErrNoRows {
		//没有找到记录
		return false, user
	}
	cryptoPasswd := util.Sha256(password)
	if user.Password == cryptoPasswd {
		return true, user
	}
	return false, user
}

// AccountExist 判断账户是否存在
func AccountExist(account string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("user").Filter("account", account).Exist()
	return exist
}

// AddUser 增加用户
func AddUser(account string, password string, nickname string) (int64, error) {
	o := orm.NewOrm()
	var user User
	user.Account = account
	user.Password = util.Sha256(password)
	user.Nickname = nickname
	user.Gender = 3
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 随机 1-8
	n := r.Intn(7) + 1
	nStr := fmt.Sprintf("%d", n)
	basePicURL := beego.AppConfig.String("qiniu_pic_url")
	user.Img_url = basePicURL + "userhead/defaulticon" + nStr + ".png"
	id, err := o.Insert(&user)
	if err == nil {
		return id, nil
	}
	return 0, err
}

// UpdateUser 更新用户
func UpdateUser(user User) bool {
	o := orm.NewOrm()
	user.Update_time = time.Now()
	_, err := o.Update(&user)
	if err == nil {
		return true
	}
	return false
}

// GetUserById 根据id获取user
func GetUserById(id int) *User {
	c := User{Id: id}
	err := orm.NewOrm().Read(&c)
	if err != nil {
		return nil
	}
	return &c
}

func GetUserByAccount(account string) *User {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("account", account).One(&user)
	if err == orm.ErrNoRows {
		return nil
	}
	return &user
}
