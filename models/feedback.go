package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Feedback struct {
	Id          int
	Title       string    `orm:"size(25)"`
	Content     string    `orm:"type(text)"`
	Contact     string    `orm:"size(25)"`
	Img_url     string    `orm:"size(255)"`
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time `orm:"type(datetime)"`
}

func AddFeedback(title string, content string, contact string, img_url string) (int64, error) {
	o := orm.NewOrm()
	var feedback Feedback
	feedback.Contact = contact
	feedback.Content = content
	feedback.Title = title
	feedback.Img_url = img_url
	feedback.Create_time = time.Now()
	feedback.Update_time = time.Now()
	id, err := o.Insert(&feedback)
	if err == nil {
		return id, nil
	}
	return 0, err
}
