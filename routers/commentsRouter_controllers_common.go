package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:AboutController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:AboutController"],
		beego.ControllerComments{
			Method: "GetAboutContent",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:AboutController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:AboutController"],
		beego.ControllerComments{
			Method: "GetProblemContent",
			Router: `/problem`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:AboutController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:AboutController"],
		beego.ControllerComments{
			Method: "FeedBack",
			Router: `/feedback`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CarouselController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CarouselController"],
		beego.ControllerComments{
			Method: "AddCarousel",
			Router: `/addCarousel`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CarouselController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CarouselController"],
		beego.ControllerComments{
			Method: "DeleteCarousel",
			Router: `/deleteCarousel`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CarouselController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CarouselController"],
		beego.ControllerComments{
			Method: "UpdateCarousel",
			Router: `/updateCarousel`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CarouselController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CarouselController"],
		beego.ControllerComments{
			Method: "GetAllCarousels",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CarouselController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CarouselController"],
		beego.ControllerComments{
			Method: "GetCarouselById",
			Router: `/getCarouselById`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "Search",
			Router: `/search`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "GetDocument",
			Router: `/documents`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "GetPreviewDocument",
			Router: `/previewDocument`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "GetCommentByCourseId",
			Router: `/comments`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "GetCommentReplyByCommentId",
			Router: `/commentsReply`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "GetQuestionsById",
			Router: `/questions`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "PostComment",
			Router: `/postComment`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "MakeGrade",
			Router: `/makeGrade`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "GetPeriodsByCourseId",
			Router: `/period`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "GetShoppingList",
			Router: `/shoppingList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "GetAllQualityComment",
			Router: `/qualityComment`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "CourseDetail",
			Router: `/courseDetail`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "FreeAskNum",
			Router: `/freeAskNum`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "PeriodDetail",
			Router: `/periodDetail`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "LiveDetail",
			Router: `/liveDetail`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:CourseController"],
		beego.ControllerComments{
			Method: "GetCoursesBySubjectId",
			Router: `/getCoursesBySubjectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:ExecController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:ExecController"],
		beego.ControllerComments{
			Method: "ExecGitSh",
			Router: `/execGitSh`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:ExecController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:ExecController"],
		beego.ControllerComments{
			Method: "ExecWhoAnI",
			Router: `/execWhoAmI`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:OrderController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:OrderController"],
		beego.ControllerComments{
			Method: "GenerateOrder",
			Router: `/generateOrder`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:OrderController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:OrderController"],
		beego.ControllerComments{
			Method: "GetOrderByOrderId",
			Router: `/getOrder`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:PayController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:PayController"],
		beego.ControllerComments{
			Method: "PayPeriod",
			Router: `/period`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:PayController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:PayController"],
		beego.ControllerComments{
			Method: "PayDocument",
			Router: `/document`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:PayController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:PayController"],
		beego.ControllerComments{
			Method: "PayAnswer",
			Router: `/document`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:PayController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:PayController"],
		beego.ControllerComments{
			Method: "GetCharge",
			Router: `/getCharge`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:PayController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:PayController"],
		beego.ControllerComments{
			Method: "PayOrder",
			Router: `/payOrder`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:PayController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:PayController"],
		beego.ControllerComments{
			Method: "Webhook",
			Router: `/webhook`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:QiniuController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:QiniuController"],
		beego.ControllerComments{
			Method: "GetUploadToken",
			Router: `/getUploadToken`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:SchoolController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:SchoolController"],
		beego.ControllerComments{
			Method: "GetSchools",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:SchoolController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:SchoolController"],
		beego.ControllerComments{
			Method: "GetMajorBySchoolId",
			Router: `/major`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:TeacherController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:TeacherController"],
		beego.ControllerComments{
			Method: "GetAllQualityTeacher",
			Router: `/qualityTeacher`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:TestController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:TestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:UserController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:UserController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:UserController"],
		beego.ControllerComments{
			Method: "SendVerifyCode",
			Router: `/sendVerifyCode`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:UserController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:UserController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:UserController"],
		beego.ControllerComments{
			Method: "ChangePassword",
			Router: `/changePassword`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:UserController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:UserController"],
		beego.ControllerComments{
			Method: "ResetPassword",
			Router: `/resetPassword`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:UserController"] = append(beego.GlobalControllerRouter["gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/common:UserController"],
		beego.ControllerComments{
			Method: "UpdateInfo",
			Router: `/updateInfo`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
