package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/admin:TestController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/admin:TestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
