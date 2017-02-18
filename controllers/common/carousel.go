package common

import (
	"encoding/json"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/base"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/models"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/util"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/vo"
)

type CarouselController struct {
	base.BaseController
}

// @router /addCarousel [post]
func (c *CarouselController) AddCarousel() {
	var carousel models.Carousel
	json.Unmarshal(c.Ctx.Input.RequestBody, &carousel)
	id, err := models.AddCarousel(carousel)
	result := make(map[string]int64)
	if err == nil {
		result["data"] = id
		c.Data["json"] = result
		c.ServeJSON()
	} else {
		result["data"] = id
		c.Data["json"] = result
		c.ServeJSON()
	}
}

// @router /deleteCarousel [post]
func (c *CarouselController) DeleteCarousel() {
	var carousel models.Carousel
	json.Unmarshal(c.Ctx.Input.RequestBody, &carousel)
	id, err := models.DeleteCarousel(carousel.Id)
	result := make(map[string]int64)
	if err == nil {
		result["data"] = id
		c.Data["json"] = result
		c.ServeJSON()
	} else {
		result["data"] = id
		c.Data["json"] = result
		c.ServeJSON()
	}
}

// @router /updateCarousel [post]
func (c *CarouselController) UpdateCarousel() {
	var carousel models.Carousel
	json.Unmarshal(c.Ctx.Input.RequestBody, &carousel)
	id, err := models.UpdateCarousel(carousel)
	result := make(map[string]int64)
	if err == nil {
		result["data"] = id
		c.Data["json"] = result
		c.ServeJSON()
	} else {
		result["data"] = id
		c.Data["json"] = result
		c.ServeJSON()
	}
}

// @router / [get]
func (c *CarouselController) GetAllCarousels() {
	carousels := models.GetAllCarousels()
	var resultList []vo.Carousel
	for _, carouselPo := range carousels {
		// fmt.Println(carousel)
		carouselVo := vo.Carousel{}
		util.PoToVo(carouselPo, &carouselVo)
		resultList = append(resultList, carouselVo)
	}
	result := make(map[string][]vo.Carousel)
	result["data"] = resultList
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /getCarouselById [get]
func (c *CarouselController) GetCarouselById() {
	id, _ := c.GetInt("id")
	carousel := models.GetCarouselById(id)
	result := make(map[string]models.Carousel)
	result["data"] = *carousel
	c.Data["json"] = result
	c.ServeJSON()
}
