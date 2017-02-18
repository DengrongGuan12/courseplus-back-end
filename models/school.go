package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type School struct {
	Id          int
	Name        string
	Img_url     string
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time `orm:"type(datetime)"`
}

func GetAllSchools() []*School {
	o := orm.NewOrm()

	var schools []*School
	_, err := o.QueryTable("school").All(&schools)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return schools
}
