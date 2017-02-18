package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Order struct {
	Id          int
	Orderno     string
	User        *User `orm:"rel(fk)"`
	Type        string
	Object_id   int
	Charge_id   string
	Description string
	Price       int
	Commented   int8       `orm:"size(4)"`
	Pay_time    *time.Time `orm:"null;type(datetime)"`
	Create_time time.Time  `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time  `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time  `orm:"type(datetime)"`
}

func AddOrder(order Order) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&order)
	if err == nil {
		return id, nil
	} else {
		return 0, err
	}
}

func (order Order) GetOrderDescription() string {
	switch order.Type {
	case "DOC":
		doc := GetDocumentById(order.Object_id)
		return fmt.Sprintf("购买资料 %v", doc.Name)
	case "COURSE":
		course := GetCourseById(order.Object_id)
		return fmt.Sprintf("购买课程 %v", course.Name)
	case "PERIOD":
		period := GetPeriodById(order.Object_id)
		return fmt.Sprintf("购买课时 %v", period.Name)
	case "ASK":
		// answer := GetAnswerById(order.Object_id)
		return fmt.Sprintf("购买回答 %v", "undefined")
	case "DOC_PACK":
		return fmt.Sprintf("购买资料: 打包")
	case "PEEK":
		// answer := GetAnswerById(order.Object_id)
		return fmt.Sprintf("购买偷听 %v", "undefined")
	default:
		return "unknown order type"
	}
}

func (order Order) GetOrderPrice() int {
	switch order.Type {
	case "DOC":
		doc := GetDocumentById(order.Object_id)
		return doc.Price
	case "COURSE":
		_, price := GetLeftPriceByCourseId(order.Object_id, order.User.Id)
		return price
	case "PERIOD":
		period := GetPeriodById(order.Object_id)
		return period.Price
	case "ASK":
		question_id := order.Object_id
		question := GetQuestionById(question_id)
		//课程下老师的价格
		price := question.Course.Teacher.Price
		return price
	case "DOC_PACK":
		courseId := order.Object_id
		userId := order.User.Id
		return GetLeftDocPriceByCourseId(courseId, userId)
	case "PEEK":
		return 100
	default:
		return -1
	}
}

func GetDocByOrderId(orderId int) Document {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("document.*").
		From("`order`").
		InnerJoin("document").
		On("`order`.object_id = document.id").
		Where("`order`.id = ?")

	sql := qb.String()
	var document Document
	o.Raw(sql, orderId).QueryRow(&document)
	return document
}

func GetBuyedPeriods(userId int) []Order {
	o := orm.NewOrm()
	var orders []Order
	o.QueryTable("order").Filter("type", "PERIOD").Filter("user_id", userId).Filter("pay_time__isnull", false).All(&orders)
	return orders
}
func GetOrderByOrderId(orderId int) *Order {
	o := orm.NewOrm()
	order := Order{
		Id: orderId,
	}
	err := o.Read(&order)
	if err != nil {
		return nil
	}
	return &order
}
func UpdateOrder(order Order) bool {
	o := orm.NewOrm()
	order.Update_time = time.Now()
	_, err := o.Update(&order)
	if err == nil {
		return true
	}
	return false
}
func GetOrderByObjectId(orderType string, objectId int) *Order {
	o := orm.NewOrm()
	var order Order
	o.QueryTable("order").Filter("type", orderType).Filter("object_id", objectId).One(&order)
	return &order
}
