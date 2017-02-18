package teacher

import (
	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/base"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/err"
)

type TestController struct {
	base.BaseController
}

// @router / [get]
func (c *TestController) GetAll() {
	c.RetError(err.ErrInputData, "userId未传！")
}
