package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Course struct {
	Id          int
	Code        string `orm:"size(15)"`
	Name        string `orm:"size(63)"`
	Major       *Major `orm:"rel(fk)"`
	Buyer_num   int
	Star        float32
	Img_url     string   `orm:"size(255)"`
	Teacher     *Teacher `orm:"rel(fk)"`
	School      *School  `orm:"rel(fk)"`
	Subject     *Subject `orm:"rel(fk)"`
	Active      int8
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time `orm:"type(datetime)"`
}

func GetCoursesByMajorId(majorId int, offset int, limit int) []Course {
	o := orm.NewOrm()
	var courses []Course
	qs := o.QueryTable("course")

	// qs.RelatedSel("subject")
	qs.Filter("major_id", majorId).Filter("active", 1).RelatedSel("teacher", "teacher__user", "subject", "major", "major__school", "major__category").Limit(limit, offset).OrderBy("-star").All(&courses)
	// qs.RelatedSel("School")
	// o.LoadRelated(&courses, "School")
	return courses
}

func GetCourseCountByMajorId(majorId int) int64 {
	o := orm.NewOrm()
	count, _ := o.QueryTable("course").Filter("major_id", majorId).Filter("active", 1).Count()
	return count
}
func GetCoursesBySchoolId(schoolId int, offset int, limit int) []Course {
	o := orm.NewOrm()
	var courses []Course
	o.QueryTable("course").Filter("school_id", schoolId).Filter("active", 1).RelatedSel("teacher", "teacher__user", "subject", "major", "major__school", "major__category").Limit(limit, offset).OrderBy("-star").All(&courses)
	return courses
}
func GetCourseCountBySchoolId(schoolId int) int64 {
	o := orm.NewOrm()
	count, _ := o.QueryTable("course").Filter("school_id", schoolId).Filter("active", 1).Count()
	return count
}
func GetCourses(offset int, limit int) []Course {
	o := orm.NewOrm()
	var courses []Course
	o.QueryTable("course").Filter("active", 1).RelatedSel("teacher", "teacher__user", "subject", "major", "major__school", "major__category").Limit(limit, offset).OrderBy("-star").All(&courses)
	return courses
}
func GetCourseCount() int64 {
	o := orm.NewOrm()
	count, _ := o.QueryTable("course").Filter("active", 1).Count()
	return count
}
func GetCourseById(id int) *Course {
	o := orm.NewOrm()
	var course Course
	err := o.QueryTable("course").Filter("id", id).RelatedSel("teacher", "teacher__user", "major", "subject", "major__school", "major__category").One(&course)
	if err != nil {
		return nil
	}
	return &course
}
func CourseExist(courseId int) bool {
	o := orm.NewOrm()
	return o.QueryTable("course").Filter("id", courseId).Exist()
}

func GetTeacherByCourseId(courseId int) *Teacher {
	o := orm.NewOrm()
	var course Course
	err := o.QueryTable("course").Filter("id", courseId).RelatedSel("teacher").One(&course)
	if err != nil {
		return nil
	}
	return course.Teacher
}

func GetCoursesBySubjectId(subjectId int) []Course {
	o := orm.NewOrm()
	var courses []Course
	o.QueryTable("course").Filter("subject_id", subjectId).RelatedSel("teacher", "teacher__user").All(&courses)
	return courses
}
