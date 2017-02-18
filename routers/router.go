// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/admin"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/teacher"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/common",
			beego.NSNamespace("/about",
				beego.NSInclude(
					&common.AboutController{},
				),
			),
			beego.NSNamespace("/carousel",
				beego.NSInclude(
					&common.CarouselController{},
				),
			),
			beego.NSNamespace("/exec",
				beego.NSInclude(
					&common.ExecController{},
				),
			),
			beego.NSNamespace("/course",
				beego.NSInclude(
					&common.CourseController{},
				),
			),
			beego.NSNamespace("/school",
				beego.NSInclude(
					&common.SchoolController{},
				),
			),
			beego.NSNamespace("/teacher",
				beego.NSInclude(
					&common.TeacherController{},
				),
			),
			beego.NSNamespace("/pay",
				beego.NSInclude(
					&common.PayController{},
				),
			),
			beego.NSNamespace("/order",
				beego.NSInclude(
					&common.OrderController{},
				),
			),
			beego.NSNamespace("/test",
				beego.NSInclude(
					&common.TestController{},
				),
			),
			beego.NSNamespace("/user",
				beego.NSInclude(
					&common.UserController{},
				),
			),
			beego.NSNamespace("/qiniu",
				beego.NSInclude(
					&common.QiniuController{},
				),
			),
		),
		beego.NSNamespace("teacher",
			beego.NSNamespace("/test",
				beego.NSInclude(
					&teacher.TestController{},
				),
			),
			beego.NSNamespace("/live",
				beego.NSInclude(
					&teacher.LiveController{},
				)),
		),
		beego.NSNamespace("admin",
			beego.NSNamespace("/test",
				beego.NSInclude(
					&admin.TestController{},
				),
			),
		),
	)
	beego.AddNamespace(ns)
}
