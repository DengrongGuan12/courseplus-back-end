package models

import (
	"time"
)

type Subject struct {
	Id          int
	Name        string    `orm:"size(15)"`
	Code        string    `orm:"size(15)"`
	Img_url     string    `orm:"size(255)"`
	School      *School   `orm:"rel(fk)"`
	Major       *Major    `orm:"rel(fk)"`
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time `orm:"type(datetime)"`
}
