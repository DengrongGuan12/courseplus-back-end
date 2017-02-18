package common

import (
	// "encoding/json"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/base"
	// "gogs.mebox.wiki/course-plus/courseplus-back-end/err"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/models"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/util"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/vo"
)

type TeacherController struct {
	base.BaseController
}

// @router /qualityTeacher [get]
func (c *TeacherController) GetAllQualityTeacher() {
	teacherList := models.GetAllQualityTeacher()
	var resultList []vo.Teacher
	util.PoListToVoList(teacherList, &resultList)
	result := make(map[string]interface{})
	result["data"] = resultList
	c.Data["json"] = result
	c.ServeJSON()
}
