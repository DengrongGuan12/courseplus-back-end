package base

import (
	"github.com/astaxie/beego"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/err"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) RetError(e *err.Error, moreInfo string) {
	if mode := beego.AppConfig.String("runmode"); mode == "prod" {
		e.DevInfo = ""
	}
	//more info
	e.MoreInfo = moreInfo

	errorResult := make(map[string]err.Error)
	errorResult["error"] = *e

	c.Data["json"] = errorResult
	c.Ctx.ResponseWriter.WriteHeader(e.Status)
	c.ServeJSON()
	c.Abort("")
}
