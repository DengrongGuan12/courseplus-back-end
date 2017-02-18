package common

import (
	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/base"
	"fmt"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/models"
)

type TestController struct {
	base.BaseController
}

// @router / [get]
func (c *TestController) GetAll() {
	//util.GetDanmakuToken(1)
	//_ , e := c.GetSession("userId").(int)
	//fmt.Println(e)
	//_ , price := models.GetLeftPriceByCourseId(1, 1)
	//fmt.Println(price)
	price := models.GetLeftDocPriceByCourseId(1,8)
	fmt.Println(price)
}
