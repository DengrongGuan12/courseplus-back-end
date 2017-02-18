package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type About struct {
	Id          int
	Type        string
	Content     string    `orm:"type(text)"`
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time `orm:"type(datetime)"`
}

func GetAboutContent() string {
	about := GetAboutById(1)
	return about.Content
}
func GetProblemContent() string {
	about := GetAboutById(2)
	return about.Content
}

func GetAboutById(id int) *About {
	o := orm.NewOrm()
	about := About{
		Id: id,
	}
	o.Read(&about)
	return &about
}
