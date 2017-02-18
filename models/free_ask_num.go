package models

import (
	// "fmt"
	"github.com/astaxie/beego/orm"
)

type Free_ask_num struct {
	User_id   int `orm:"pk"`
	Course_id int
	Free_num  int
}

func (f *Free_ask_num) TableUnique() [][]string {
	return [][]string{
		[]string{"user_id", "course_id"},
	}
}
func (f *Free_ask_num) TableIndex() [][]string {
	return [][]string{
		[]string{"user_id", "course_id"},
	}
}
func GetFreeNum(userId int, courseId int) int {
	o := orm.NewOrm()
	freeNum := Free_ask_num{
		User_id:   userId,
		Course_id: courseId,
	}
	err := o.Read(&freeNum)
	if err != nil {
		return 0
	} else {
		return freeNum.Free_num
	}
}
