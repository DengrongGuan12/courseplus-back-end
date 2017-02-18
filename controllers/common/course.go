package common

import (
	"encoding/json"
	"fmt"
	"html/template"

	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/base"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/err"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/models"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/util"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/vo"
)

type CourseController struct {
	base.BaseController
}

// @router /search [get]
func (c *CourseController) Search() {
	// result["data"] = data
	// c.Data["json"] = result
	// c.ServeJSON()

	limit, err := c.GetInt("limit")
	if err != nil {
		limit = 5
	}
	offset, _ := c.GetInt("offset")
	major_id, _ := c.GetInt("major_id")
	school_id, _ := c.GetInt("school_id")
	var courses vo.Courses
	courses.Limit = limit
	courses.Offset = offset
	if major_id != 0 {
		//专业过滤
		courseList := models.GetCoursesByMajorId(major_id, offset, limit)
		var resultList []vo.Course
		util.PoListToVoList(courseList, &resultList)
		courses.List = resultList
		courses.Count = models.GetCourseCountByMajorId(major_id)
	} else if school_id != 0 {
		//学校过滤
		courseList := models.GetCoursesBySchoolId(school_id, offset, limit)
		var resultList []vo.Course
		util.PoListToVoList(courseList, &resultList)
		courses.List = resultList
		courses.Count = models.GetCourseCountBySchoolId(school_id)
	} else {
		courseList := models.GetCourses(offset, limit)
		var resultList []vo.Course
		util.PoListToVoList(courseList, &resultList)
		courses.List = resultList
		courses.Count = models.GetCourseCount()
	}
	result := make(map[string]vo.Courses)
	result["data"] = courses
	c.Data["json"] = result
	c.ServeJSON()

}

// @router /documents [get]
func (c *CourseController) GetDocument() {
	limit, e := c.GetInt("limit")
	if e != nil {
		limit = 5
	}
	offset, _ := c.GetInt("offset")
	courseId, e := c.GetInt("course_id")
	if e != nil {
		c.RetError(err.ErrInputData, "course_id")
	}

	documentList := models.GetCourseDocument(courseId, limit, offset)
	fmt.Printf("document: %v\n", documentList[0].Course)
	var resultList []vo.Document
	util.PoListToVoList(documentList, &resultList)

	userId, isLogin := c.GetSession("userId").(int)

	for idx := range resultList {
		if isLogin {
			if models.IsBuy("DOC", userId, resultList[idx].Id) {
				resultList[idx].Is_buy = 1
			} else {
				resultList[idx].Is_buy = 0
			}
		} else {
			resultList[idx].Is_buy = 0
		}

		if resultList[idx].Is_buy == 0 {
			resultList[idx].Remote_path = ""
		}
	}

	data := make(map[string]interface{})
	data["offset"] = offset
	data["limit"] = limit
	data["count"] = len(resultList)
	data["list"] = resultList
	result := make(map[string]interface{})
	result["data"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /previewDocument [get]
func (c *CourseController) GetPreviewDocument() {
	id, _ := c.GetInt("document_id")
	userId, ok := c.GetSession("userId").(int)
	if !ok {
		userId = -1
	}
	document := models.GetPreviewDocument(id, userId)
	var documentVO vo.DocumentPreviewInfo
	util.PoToVo(document, &documentVO)
	result := make(map[string]interface{})
	result["data"] = documentVO
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /comments [get]
func (c *CourseController) GetCommentByCourseId() {
	limit, e := c.GetInt("limit")
	if e != nil {
		limit = 5
	}
	offset, _ := c.GetInt("offset")
	courseId, e := c.GetInt("course_id")
	if e != nil {
		c.RetError(err.ErrInputData, "course_id")
	}

	commentList := models.GetCommentByCourseId(courseId, offset, limit)
	fmt.Printf("commentList %v \n", commentList)
	for idx := range commentList {
		if commentList[idx].User_buy_record != nil {
			commentList[idx].Order_info = commentList[idx].User_buy_record.GetOrderInfo()
		}
	}
	var resultList []vo.Course_comment
	util.PoListToVoList(commentList, &resultList)
	result := make(map[string]interface{})
	result["data"] = resultList
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /commentsReply [get]
func (c *CourseController) GetCommentReplyByCommentId() {
	limit, e := c.GetInt("limit")
	if e != nil {
		limit = 5
	}
	offset, _ := c.GetInt("offset")
	commentId, e := c.GetInt("comment_id")
	commentReplyList := models.GetCommentReplyByCommentId(commentId, offset, limit)

	var resultList []vo.Comment_reply
	util.PoListToVoList(commentReplyList, &resultList)

	result := make(map[string]interface{})
	result["data"] = resultList
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /questions [get]
func (c *CourseController) GetQuestionsById() {
	limit, e := c.GetInt("limit")
	if e != nil {
		limit = 5
	}
	offset, _ := c.GetInt("offset")
	courseId, e := c.GetInt("course_id")
	if e != nil {
		c.RetError(err.ErrInputData, "course_id")
	}
	var questionsVo vo.Questions
	questionsVo.Limit = limit
	questionsVo.Offset = offset
	questionsVo.Count = models.GetQuestionCountByCourseId(courseId)
	questions := models.GetQuestionByCourseId(courseId, offset, limit)
	var resultList []vo.Question
	util.PoListToVoList(questions, &resultList)
	questionsVo.List = resultList
	result := make(map[string]interface{})
	result["data"] = questionsVo
	c.Data["json"] = result
	c.ServeJSON()
}

type CommentInfo struct {
	Content    string
	Course_id  int
	Comment_id int
	Reply_id   int
}

// @router /postComment [post]
func (this *CourseController) PostComment() {
	userIdOb := this.GetSession("userId")
	if userIdOb == nil {
		this.RetError(err.ErrUnauthorized, "未登录")
	}
	userId := userIdOb.(int)
	var commentInfo CommentInfo
	json.Unmarshal(this.Ctx.Input.RequestBody, &commentInfo)
	if commentInfo.Course_id <= 0 || !models.CourseExist(commentInfo.Course_id) {
		this.RetError(err.ErrInputData, "course_id")
	}
	if len(commentInfo.Content) == 0 {
		this.RetError(err.ErrInputData, "content")
	}
	commentInfo.Content = template.HTMLEscapeString(commentInfo.Content)
	if commentInfo.Comment_id <= 0 || !models.CourseCommentExist(commentInfo.Comment_id) {
		//根评论
		models.AddCourseComment(commentInfo.Course_id, commentInfo.Content, userId, 0, 0)
	} else {
		//评论的回复
		if commentInfo.Reply_id <= 0 || !models.CommentReplyExist(commentInfo.Reply_id, commentInfo.Comment_id) {
			models.AddCommentReply(commentInfo.Comment_id, commentInfo.Content, 0, userId)
		} else {
			models.AddCommentReply(commentInfo.Comment_id, commentInfo.Content, commentInfo.Reply_id, userId)
		}
		//增加评论数量
		models.AddCommentReplyCount(commentInfo.Comment_id)

	}
	result := make(map[string]int)
	result["data"] = 1
	this.Data["json"] = result
	this.ServeJSON()

}

type GradeInfo struct {
	Content   string
	Course_id int
	Order_id  int
	Star      float32
}

// @router /makeGrade [post]
func (this *CourseController) MakeGrade() {
	userIdOb := this.GetSession("userId")
	if userIdOb == nil {
		this.RetError(err.ErrUnauthorized, "未登录")
	}
	userId := userIdOb.(int)
	var gradeInfo GradeInfo
	json.Unmarshal(this.Ctx.Input.RequestBody, &gradeInfo)
	if gradeInfo.Course_id <= 0 || !models.CourseExist(gradeInfo.Course_id) {
		this.RetError(err.ErrInputData, "course_id")
	}
	if gradeInfo.Order_id <= 0 {
		this.RetError(err.ErrInputData, "order_id <= 0")
	}
	order := models.GetOrderByOrderId(gradeInfo.Order_id)

	if order.Pay_time == nil {
		this.RetError(err.ErrInputData, "order no pay")
	}
	if order.Type != "COURSE" && order.Type != "PERIOD" {
		this.RetError(err.ErrInputData, "order type")
	}
	if order.Commented == 1 {
		this.RetError(err.ErrInputData, "order commented")
	}
	if len(gradeInfo.Content) == 0 {
		this.RetError(err.ErrInputData, "content")
	}
	if gradeInfo.Star <= 0 {
		this.RetError(err.ErrInputData, "star")
	}
	gradeInfo.Content = template.HTMLEscapeString(gradeInfo.Content)
	models.AddCourseComment(gradeInfo.Course_id, gradeInfo.Content, userId, gradeInfo.Order_id, gradeInfo.Star)
	order.Commented = 1
	models.UpdateOrder(*order)
	result := make(map[string]int)
	result["data"] = 1
	this.Data["json"] = result
	this.ServeJSON()
}

// @router /period [get]
func (this *CourseController) GetPeriodsByCourseId() {
	courseId, _ := this.GetInt("course_id")
	if courseId <= 0 {
		this.RetError(err.ErrInputData, "course_id")
	}
	periods := models.GetPeriodsByCourseId(courseId)
	var periodVos []vo.SimplePeriod
	userId := this.GetSession("userId")
	if userId == nil {
		util.PoListToVoList(periods, &periodVos)
	} else {
		buyedPeriods := models.GetBuyedPeriods(userId.(int))
		for _, period := range periods {
			var periodVo vo.SimplePeriod
			util.PoToVo(period, &periodVo)
			for _, order := range buyedPeriods {
				if order.Object_id == period.Id {
					periodVo.Is_buy = 1
					break
				}
			}
			periodVos = append(periodVos, periodVo)
		}
	}
	result := make(map[string]interface{})
	result["data"] = periodVos
	this.Data["json"] = result
	this.ServeJSON()
}

// @router /shoppingList [get]
func (this *CourseController) GetShoppingList() {
	courseId, _ := this.GetInt("course_id")
	if courseId <= 0 {
		this.RetError(err.ErrInputData, "course_id")
	}
	var periodVos []vo.SimplePeriod
	userId, ok := this.GetSession("userId").(int)
	if !ok {
		userId = -1
	}
	periodVos, totalPrice := models.GetLeftPriceByCourseId(courseId, userId)
	documents := models.GetAllDocByCourseId(courseId)
	var simpleDocs []vo.SimpleDocument
	util.PoListToVoList(documents, &simpleDocs)
	shoppingList := vo.ShoppingList{
		Period_list:    periodVos,
		Document_list:  simpleDocs,
		Price:          totalPrice,
		Question_count: 5,
	}
	result := make(map[string]interface{})
	result["data"] = shoppingList
	this.Data["json"] = result
	this.ServeJSON()
}

// @router /qualityComment [get]
func (c *CourseController) GetAllQualityComment() {
	commentList := models.GetAllQualityComment()
	var resultList []vo.Course_comment
	util.PoListToVoList(commentList, &resultList)
	result := make(map[string]interface{})
	result["data"] = resultList
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /courseDetail [get]
func (this *CourseController) CourseDetail() {
	courseId, _ := this.GetInt("course_id")
	if courseId <= 0 {
		this.RetError(err.ErrInputData, "course_id")
	}
	coursePo := models.GetCourseById(courseId)
	var courseVo vo.Course
	util.PoToVo(coursePo, &courseVo)
	result := make(map[string]vo.Course)
	userId := this.GetSession("userId")
	if userId == nil {
		//未登录
		courseVo.Is_buy = 0
		courseVo.Origin_price = models.GetTotalPriceByCourseId(courseVo.Id)
		courseVo.Current_price = courseVo.Origin_price - 10
	} else {
		if buyed := models.IsBuy("COURSE", userId.(int), courseVo.Id); buyed {
			// 已买断
			courseVo.Is_buy = 1
			courseVo.Origin_price = 0
			courseVo.Current_price = 0
		} else {
			//计算已买的课时和总课时的价格差
			courseVo.Is_buy = 0
			totalPrice := 0
			periods := models.GetPeriodsByCourseId(courseVo.Id)
			buyedOrders := models.GetBuyedPeriods(userId.(int))
			for _, period := range periods {
				//判断有没有买过
				fmt.Println("period:%d", period.Id)
				buyed := false
				for _, order := range buyedOrders {
					fmt.Println("order object_id:%d", order.Object_id)
					if order.Object_id == period.Id {
						buyed = true
						break
					}
				}
				if !buyed {
					totalPrice += period.Price
				}

			}
			courseVo.Origin_price = totalPrice
			courseVo.Current_price = totalPrice - 10
			// len := periods.Len()

		}
	}
	courseVo.Free_ask_num = 2
	courseVo.Doc_num = models.GetDocNumByCourseId(courseId)
	courseVo.Period_num = models.GetPeriodNumByCourseId(courseId)
	var simpleCourses []vo.SimpleCourse
	courses := models.GetCoursesBySubjectId(courseVo.Subject.Id)
	for _, course := range courses {
		if course.Id == courseVo.Id {
			continue
		}
		var simpleCourse vo.SimpleCourse
		util.PoToVo(course, &simpleCourse)
		simpleCourses = append(simpleCourses, simpleCourse)
	}
	courseVo.Other_courses = simpleCourses
	// courseVo.Current_price = 0
	// courseVo.Origin_price = 100
	result["data"] = courseVo
	this.Data["json"] = result
	this.ServeJSON()
}

// @router /freeAskNum [get]
func (this *CourseController) FreeAskNum() {
	courseId, _ := this.GetInt("course_id")
	userId, _ := this.GetSession("userId").(int)
	freeNum := models.GetFreeNum(userId, courseId)
	this.Data["json"] = map[string]int{
		"data": freeNum,
	}
	this.ServeJSON()
}

// @router /periodDetail [get]
func (this *CourseController) PeriodDetail() {
	periodId, _ := this.GetInt("period_id")
	if periodId == 0 {
		this.RetError(err.ErrInputData, "period_id")
	}
	period := models.GetPeriodById(periodId)
	live := models.GetLiveByPeriodId(periodId)
	userId, ok := this.GetSession("userId").(int)
	if !ok {
		userId = 0
	}
	var videoPos []models.Video
	if models.IsBuy("PERIOD", userId, periodId) {
		videoPos = models.GetVideosByLiveId(live.Id, 0, 0)
	} else {
		videoPos = models.GetVideosByLiveId(live.Id, 0, 1)
	}
	var videoVos []vo.Video
	util.PoListToVoList(videoPos, &videoVos)
	var periodVo vo.Period
	var liveVo vo.Live
	util.PoToVo(period, &periodVo)
	util.PoToVo(live, &liveVo)
	liveVo.Push_url = ""
	liveVo.Video_list = videoVos
	periodVo.Live = liveVo
	result := make(map[string]vo.Period)
	result["data"] = periodVo
	this.Data["json"] = result
	this.ServeJSON()
}

// @router /liveDetail [get]
func (this *CourseController) LiveDetail() {
	liveId, _ := this.GetInt("live_id")
	live := models.GetLiveById(liveId)
	if live == nil {
		this.RetError(err.ErrInputData, "直播id不存在")
		return
	}
	course := models.GetCourseById(live.Period.Course.Id)
	var liveVo vo.Live
	var courseVo vo.Course
	result := make(map[string]interface{})
	util.PoToVo(live, &liveVo)
	liveVo.Push_url = ""
	util.PoToVo(course, &courseVo)
	result["live"] = liveVo
	result["course"] = courseVo
	data := map[string]interface{}{"data": result}
	this.Data["json"] = data
	this.ServeJSON()
}

// @router /getCoursesBySubjectId [get]
func (this *CourseController) GetCoursesBySubjectId() {
	subjectId, _ := this.GetInt("subject_id")
	coursePos := models.GetCoursesBySubjectId(subjectId)
	var courses []vo.SimpleCourse
	util.PoListToVoList(coursePos, &courses)
	this.Data["json"] = map[string]interface{}{
		"data": courses,
	}
	this.ServeJSON()
}

func GetCoursePrice() {

}
