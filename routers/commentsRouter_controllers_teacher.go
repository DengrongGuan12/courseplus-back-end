package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/teacher:LiveController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/teacher:LiveController"],
		beego.ControllerComments{
			Method: "GetDetailById",
			Router: `/detail`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/teacher:LiveController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/teacher:LiveController"],
		beego.ControllerComments{
			Method: "TecentCallBack",
			Router: `/tecentCallBack`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/teacher:LiveController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/teacher:LiveController"],
		beego.ControllerComments{
			Method: "Finish",
			Router: `/finish`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/teacher:TestController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/teacher:TestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
