package common

import (
	"fmt"

	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/base"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/models"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/util"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/vo"
)

type SchoolController struct {
	base.BaseController
}

// @router / [get]
func (c *SchoolController) GetSchools() {
	schoolList := models.GetAllSchools()
	var resultList []vo.School
	util.PoListToVoList(schoolList, &resultList)
	result := make(map[string]interface{})
	result["data"] = resultList
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /major [get]
func (c *SchoolController) GetMajorBySchoolId() {
	id, _ := c.GetInt("school_id")
	majorList := models.GetMajorBySchoolId(id)
	fmt.Printf("%v %T", majorList, majorList)
	var resultList []vo.Major
	util.PoListToVoList(majorList, &resultList)
	result := make(map[string][]vo.Major)
	result["data"] = resultList
	c.Data["json"] = result
	c.ServeJSON()
}
