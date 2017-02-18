package models

import (
	"time"
)

type Message struct {
	Id          int
	Live        *Live `orm:"rel(fk)"`
	User        *User `orm:"rel(fk)"`
	Time        int
	Text        string    `orm:"size(255)"`
	Color       string    `orm:"size(15)"`
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time `orm:"type(datetime)"`
}
