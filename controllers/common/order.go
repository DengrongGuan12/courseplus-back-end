package common

import (
	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/base"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/err"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/models"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/util"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/vo"
	"encoding/json"
)

type OrderController struct {
	base.BaseController
}

type generateOrderInfo struct {
	Object_id	int
	Order_type	string
}

// @description 生成订单
// @router /generateOrder [post]
func (c *OrderController) GenerateOrder(){
	var data generateOrderInfo
	json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	objectId := data.Object_id
	if objectId <=0 {
		c.RetError(err.ErrInputData, "object_id")
	}
	order_type := data.Order_type
	if order_type =="" {
		c.RetError(err.ErrInputData, "order_type")
	}

	var order models.Order
	order.Type = order_type
	order.Object_id = objectId
	order.User = &models.User{ Id: c.Ctx.Input.Session("userId").(int) }
//	order.User = &models.User{ Id: 1 }
	order.Price = order.GetOrderPrice()
	order.Orderno = util.GenerateOrderno()
	order.Description = order.GetOrderDescription()
	order.Create_time = util.GetCurrentTime()
	order.Update_time = util.GetCurrentTime()

	id, e := models.AddOrder(order)
	if e != nil{
		c.RetError(err.ErrDatabase, "添加order失败")
	}
	order.Id = int(id)

	var orderVo vo.Order
	util.PoToVo(order, &orderVo)
	result := make(map[string]interface{})
	result["data"] = orderVo
	c.Data["json"] = result
	c.ServeJSON()
}

// @description 根据订单号获取订单
// @router /getOrder [get]
func (c *OrderController) GetOrderByOrderId(){
	order_id, e:= c.GetInt("order_id")
	if e != nil || order_id <= 0{
		c.RetError(err.ErrInputData, "order_id")
	}
	order := models.GetOrderByOrderId(order_id)
	if order == nil{
		c.RetError(err.ErrNil, "order")
	}
	var orderVo vo.Order
	util.PoToVo(order, &orderVo)
	result := make(map[string]interface{})
	result["data"] = orderVo
	c.Data["json"] = result
	c.ServeJSON()
}