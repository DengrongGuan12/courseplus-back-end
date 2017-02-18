package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"time"
)

type User_buy_record struct {
	Id		int
	User		*User `orm:"rel(fk)"`
	Type		string
	Commented	int8
	Object_id	int
	Pay_time	time.Time `orm:"null;type(datetime)"`
}

func AddUserBuyRecord(record User_buy_record) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&record)
	if err == nil {
		return id, nil
	} else {
		return 0, err
	}
}

func GetUserBuyRecordByRecordId(recordId int) *User_buy_record {
	o := orm.NewOrm()
	record := User_buy_record{
		Id: recordId,
	}
	err := o.Read(&record)
	if err != nil {
		return nil
	}
	return &record
}

func IsBuy(t string, userId int, objectId int) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("user_buy_record").Filter("user_id", userId).Filter("type", t).Filter("object_id", objectId).Exist()
	return exist
}

func (user_buy_record *User_buy_record) GetOrderInfo() string {
	switch user_buy_record.Type {
	case "DOC":
		doc := GetDocumentById(user_buy_record.Object_id)
		return fmt.Sprintf("购买资料 %v", doc.Name)
	case "COURSE":
		course := GetCourseById(user_buy_record.Object_id)
		return fmt.Sprintf("购买课程 %v", course.Name)
	case "PERIOD":
		period := GetPeriodById(user_buy_record.Object_id)
		return fmt.Sprintf("购买课时 %v", period.Name)
	case "ANSWER":
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
