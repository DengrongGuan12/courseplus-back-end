package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CarouselController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CarouselController"],
		beego.ControllerComments{
			Method:           "AddCarousel",
			Router:           `/addCarousel`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CarouselController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CarouselController"],
		beego.ControllerComments{
			Method:           "DeleteCarousel",
			Router:           `/deleteCarousel`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CarouselController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CarouselController"],
		beego.ControllerComments{
			Method:           "UpdateCarousel",
			Router:           `/updateCarousel`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CarouselController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CarouselController"],
		beego.ControllerComments{
			Method:           "GetAllCarousels",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CarouselController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CarouselController"],
		beego.ControllerComments{
			Method:           "GetCarouselById",
			Router:           `/getCarouselById`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "Search",
			Router:           `/search`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "GetDocument",
			Router:           `/documents`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "GetPreviewDocument",
			Router:           `/previewDocument`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "GetCommentByCourseId",
			Router:           `/comments`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "GetCommentReplyByCommentId",
			Router:           `/commentsReply`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "GetQuestionsById",
			Router:           `/questions`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "PostComment",
			Router:           `/postComment`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "MakeGrade",
			Router:           `/makeGrade`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "GetPeriodsByCourseId",
			Router:           `/period`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "GetShoppingList",
			Router:           `/shoppingList`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "GetAllQualityComment",
			Router:           `/qualityComment`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "CourseDetail",
			Router:           `/courseDetail`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "PeriodDetail",
			Router:           `/periodDetail`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "LiveDetail",
			Router:           `/liveDetail`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:ExecController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:ExecController"],
		beego.ControllerComments{
			Method:           "ExecGitSh",
			Router:           `/execGitSh`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:ExecController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:ExecController"],
		beego.ControllerComments{
			Method:           "ExecWhoAnI",
			Router:           `/execWhoAmI`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:PayController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:PayController"],
		beego.ControllerComments{
			Method:           "GetChargeByPeriod",
			Router:           `/period`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:PayController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:PayController"],
		beego.ControllerComments{
			Method:           "Webhook",
			Router:           `/webhook`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:QiniuController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:QiniuController"],
		beego.ControllerComments{
			Method:           "GetUploadToken",
			Router:           `/getUploadToken`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:SchoolController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:SchoolController"],
		beego.ControllerComments{
			Method:           "GetSchools",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:SchoolController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:SchoolController"],
		beego.ControllerComments{
			Method:           "GetMajorBySchoolId",
			Router:           `/major`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:TeacherController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:TeacherController"],
		beego.ControllerComments{
			Method:           "GetAllQualityTeacher",
			Router:           `/qualityTeacher`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:TestController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:TestController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:UserController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:UserController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:UserController"],
		beego.ControllerComments{
			Method:           "SendVerifyCode",
			Router:           `/sendVerifyCode`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:UserController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Register",
			Router:           `/register`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:UserController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:UserController"],
		beego.ControllerComments{
			Method:           "ChangePassword",
			Router:           `/changePassword`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:UserController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:UserController"],
		beego.ControllerComments{
			Method:           "ResetPassword",
			Router:           `/resetPassword`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:UserController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UpdateInfo",
			Router:           `/updateInfo`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

}
