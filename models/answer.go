package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Answer struct {
	Id      int
	Content string `orm:type(text)`
	//enum 类型
	Type        string
	Length      int
	Teacher     *Teacher  `orm:"rel(fk)"`
	Question    *Question `orm:"rel(one)"`
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time `orm:"type(datetime)"`
	Buyer_num   int
}

func GetAnswerById(id int) *Answer {
	o := orm.NewOrm()
	var answer Answer
	err := o.QueryTable("answer").Filter("id", id).One(&answer)
	if err != nil {
		return nil
	}
	return &answer
}
