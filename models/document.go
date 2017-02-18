package models

import (
	"time"

	"gogs.mebox.wiki/course-plus/courseplus-back-end/util"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Document struct {
	Id           int
	Name         string `orm:"size(63)"`
	Price        int
	Ext          string `orm:"size(7)"`
	Remote_path  string `orm:"size(255)"`
	Preview_path string `orm:"size(255)"`
	Cover_path   string `orm:"size(255)"`
	Origin_price int
	Buyer_num    int
	Create_time  time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time  time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time  time.Time `orm:"type(datetime)"`
	Course       *Course   `orm:"rel(fk)"`
	Is_buy       int       `orm:"-"`
}

func GetCourseDocument(courseId int, limit int, offset int) []Document {
	o := orm.NewOrm()
	var documents []Document
	o.QueryTable("document").Filter("course_id", courseId).Offset(offset).Limit(limit).All(&documents)
	return documents
}

func GetAllDocByCourseId(courseId int) []Document {
	o := orm.NewOrm()
	var documents []Document
	o.QueryTable("document").Filter("course_id", courseId).All(&documents)
	return documents
}
func GetDocNumByCourseId(courseId int) int64 {
	num, _ := orm.NewOrm().QueryTable("document").Filter("course_id", courseId).Count()
	return num
}

func GetPreviewDocument(id int, userId int) Document {
	o := orm.NewOrm()

	var document Document
	o.QueryTable("document").Filter("id", id).One(&document)

	domain := beego.AppConfig.String("qiniu_doc_url")
	fmt.Printf("document: %v\n", document)
	if IsBuy("DOC", userId, document.Id) {
		document.Is_buy = 1
	}
	document.Preview_path = util.GetResourceUrl(domain, document.Preview_path, true)
	if !(document.Is_buy == 1 || document.Price <= 0) {
		document.Remote_path = ""
	} else {
		document.Remote_path = util.GetResourceUrl(domain, document.Remote_path, true)
	}

	return document
}

func GetDocumentById(id int) *Document {
	o := orm.NewOrm()
	var document Document
	err := o.QueryTable("document").Filter("id", id).One(&document)
	if err != nil {
		return nil
	}
	return &document
}

func GetLeftDocPriceByCourseId(courseId int, userId int) int{
	documents := GetAllDocByCourseId(courseId)
	var totalPrice int
	for _,v:=range documents{
		if !IsBuy("DOC", userId, v.Id){
			totalPrice += v.Price
		}
	}
	return totalPrice
}