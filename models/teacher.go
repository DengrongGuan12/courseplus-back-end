package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Teacher struct {
	Id int
	// Nickname     string `orm:"size(63)"`
	Img_url      string `orm:"size(255)"`
	Introduction string `orm:"size(255)"`
	Description  string `orm:"size(255)"`
	// Password     string `orm:"size(64)"`
	// Account      string `orm:"size(63)"`
	// Gender       int8   `orm:"size(4)"`
	User         *User  `orm:"rel(fk)"`
	Education    string `orm:"size(63)"`
	Price        int
	Peek_price   int
	High_quality int8      `orm:"size(4)"`
	Create_time  time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time  time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time  time.Time `orm:"type(datetime)"`
}

func GetAllQualityTeacher() []*Teacher {
	o := orm.NewOrm()

	var teachers []*Teacher
	_, err := o.QueryTable("teacher").RelatedSel("user").Filter("high_quality", 1).All(&teachers)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return teachers
}

func GetTeacherByUserId(userId int) *Teacher {
	o := orm.NewOrm()
	var teacher Teacher
	err := o.QueryTable("teacher").Filter("user_id", userId).One(&teacher)
	if err != nil {
		return nil
	}
	return &teacher
}
