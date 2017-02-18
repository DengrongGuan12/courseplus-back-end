package models

import (
	"time"
)

type Major_category struct {
	Id          int
	Code        string    `orm:"size(15)"`
	Name        string    `orm:"size(15)"`
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time `orm:"type(datetime)"`
}
