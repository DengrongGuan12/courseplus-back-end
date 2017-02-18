package common

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/base"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/err"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/models"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/util"
)

type AboutController struct {
	base.BaseController
}

// @router / [get]
func (this *AboutController) GetAboutContent() {
	this.Data["json"] = map[string]string{
		"data": models.GetAboutContent(),
	}
	this.ServeJSON()
}

// @router /problem [get]
func (this *AboutController) GetProblemContent() {
	this.Data["json"] = map[string]string{
		"data": models.GetProblemContent(),
	}
	this.ServeJSON()
}

type Feedback struct {
	Title   string
	Content string
	Img_key string
	Contact string
}

// @router /feedback [post]
func (this *AboutController) FeedBack() {
	var data Feedback
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	if data.Contact == "" || data.Title == "" || data.Content == "" {
		this.RetError(err.ErrInputData, "内容不全")
	}
	var img_url string
	if data.Img_key != "" {
		domain := beego.AppConfig.String("qiniu_pic_url")
		img_url = util.GetResourceUrl(domain, data.Img_key, false)
	}
	id, error := models.AddFeedback(data.Title, data.Content, data.Contact, img_url)
	if error != nil {
		this.RetError(err.ErrDatabase, "插入失败")
	}
	this.Data["json"] = map[string]int64{
		"data": id,
	}
	this.ServeJSON()
}
