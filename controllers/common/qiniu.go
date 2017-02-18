package common

import (
	"github.com/astaxie/beego"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/base"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

type QiniuController struct {
	base.BaseController
}

// @router /getUploadToken [get]
func (this *QiniuController) GetUploadToken() {
	bucket := this.GetString("bucket")

	conf.ACCESS_KEY = beego.AppConfig.String("qiniu_access_key")
	conf.SECRET_KEY = beego.AppConfig.String("qiniu_secret_key")
	c := kodo.New(0, nil)
	policy := &kodo.PutPolicy{
		Scope:   bucket,
		Expires: 3600,
	}
	token := c.MakeUptoken(policy)
	this.Data["json"] = map[string]map[string]string{"data": {"token": token}}
	this.ServeJSON()

}
