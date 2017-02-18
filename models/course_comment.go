package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Course_comment struct {
	Id              int
	Course          *Course `orm:"rel(fk)"`
	Content         string  `orm:"size(255)"`
	User            *User   `orm:"rel(fk)"`
	High_quality    int8
	Reply_count     int
	Star            float32
	User_buy_record *User_buy_record `orm:"rel(fk)"`
	Order_info      string           `orm:"-"`
	Create_time     time.Time        `orm:"auto_now_add;type(datetime)"`
	Update_time     time.Time        `orm:"auto_now_add;type(datetime)"`
	Delete_time     time.Time        `orm:"type(datetime)"`
}

func GetAllQualityComment() []Course_comment {
	o := orm.NewOrm()
	var comments []Course_comment
	o.QueryTable("course_comment").Filter("high_quality", 1).RelatedSel(3).All(&comments)
	return comments
}

func GetCommentByCourseId(courseId int, offset int, limit int) []Course_comment {
	o := orm.NewOrm()
	var comments []Course_comment
	o.QueryTable("course_comment").Filter("course_id", courseId).RelatedSel(2, "user", "User_buy_record").Offset(offset).Limit(limit).All(&comments)
	return comments
}
func AddCourseComment(courseId int, content string, userId int, orderId int, star float32) (int64, error) {
	o := orm.NewOrm()
	var courseComment Course_comment
	courseComment.Content = content
	courseComment.Course = &Course{
		Id: courseId,
	}
	courseComment.User = &User{
		Id: userId,
	}
	if orderId > 0 {
		// courseComment.Order = &Order{
		// 	Id: orderId,
		// }
	}
	if star > 0 {
		courseComment.Star = star
	}
	id, err := o.Insert(&courseComment)
	if err == nil {
		return id, nil
	}
	return 0, err
}
func CourseCommentExist(commentId int) bool {
	o := orm.NewOrm()
	return o.QueryTable("course_comment").Filter("id", commentId).Exist()
}
func CourseCommentExistByOrderId(orderId int) bool {
	o := orm.NewOrm()
	return o.QueryTable("course_comment").Filter("order_id", orderId).Exist()
}
func AddCommentReplyCount(commentId int) bool {
	comment := GetCourseCommentById(commentId)
	comment.Update_time = time.Now()
	comment.Reply_count++
	_, err := orm.NewOrm().Update(comment)
	if err == nil {
		return true
	}
	return false
}
func GetCourseCommentById(id int) *Course_comment {
	c := Course_comment{Id: id}
	err := orm.NewOrm().Read(&c)
	if err != nil {
		return nil
	}
	return &c
}
