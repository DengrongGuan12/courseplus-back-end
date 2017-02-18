package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Question struct {
	Id          int
	Content     string    `orm:"size(255)"`
	Course      *Course   `orm:"rel(fk)"`
	User        *User     `orm:"rel(fk)"`
	Answer      *Answer   `orm:"null;reverse(one)"`
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time `orm:"type(datetime)"`
	State       string    `orm:"size(255)"`
}

func GetQuestionByCourseId(courseId int, offset int, limit int) []Question {
	o := orm.NewOrm()
	var questions []Question
	o.QueryTable("question").Filter("course_id", courseId).RelatedSel(2).Offset(offset).Limit(limit).All(&questions)
	for index := range questions {
		o.LoadRelated(&questions[index], "Answer", true)
		fmt.Println("answer", questions[index].Answer)
	}
	return questions
}

func GetQuestionById(questionId int) *Question {
	q := Question{Id: questionId}
	err := orm.NewOrm().QueryTable("question").RelatedSel(3).One(&q)
	if err != nil {
		return nil
	}
	return &q
}

func GetQuestionCountByCourseId(courseId int) int64 {
	o := orm.NewOrm()
	num, _ := o.QueryTable("question").Filter("course_id", courseId).Count()
	return num
}

func PayQuestionById(questionId int){
	o := orm.NewOrm()
	question := Question{Id:questionId}
	o.Read(&question)
	question.State = "PAIED"
	o.Update(&question, "state")
}
