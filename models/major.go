package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Major struct {
	Id          int
	Name        string          `orm:"size(63)"`
	Code        string          `orm:"size(15)"`
	School      *School         `orm:"rel(fk)"`
	Category    *Major_category `orm:"rel(fk)"`
	Create_time time.Time       `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time       `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time       `orm:"type(datetime)"`
}

func GetMajorBySchoolId(schoolId int) []Major {
	o := orm.NewOrm()

	var majors []Major
	_, err := o.QueryTable("major").Filter("school_id", schoolId).RelatedSel("School", "Category").All(&majors)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return majors
}
