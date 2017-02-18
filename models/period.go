package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/vo"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/util"
)

type Period struct {
	Id           int
	Course       *Course   `orm:"rel(fk)"`
	Name         string    `orm:"size(255)"`
	Date         time.Time `orm:"auto_now_add;type(date)"`
	Price        int
	Origin_price int
	Content      string `orm:"type(text)"`
	Number       int
	Create_time  time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time  time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time  time.Time `orm:"type(datetime)"`
}

func GetPeriodById(id int) *Period {
	o := orm.NewOrm()
	var period Period
	err := o.QueryTable("period").Filter("id", id).One(&period)
	if err != nil {
		return nil
	}
	return &period
}
func GetTotalPriceByCourseId(courseId int) int {
	o := orm.NewOrm()
	type result struct {
		Total int
	}
	var sum result
	err := o.Raw("select sum(price) as total from period where course_id = ? group by course_id", courseId).QueryRow(&sum)
	if err != nil {
		return -1
	}
	return sum.Total
}

func GetLeftPriceByCourseId(courseId int, userId int) ([]vo.SimplePeriod,int){
	totalPrice := 0
	var periodVos []vo.SimplePeriod
	periods := GetPeriodsByCourseId(courseId)
	if userId <= 0{
		totalPrice = GetTotalPriceByCourseId(courseId)
		util.PoListToVoList(periods, &periodVos)
	}else {
		buyedPeriods := GetBuyedPeriods(userId)
		for _, period := range periods {
			var periodVo vo.SimplePeriod
			util.PoToVo(period, &periodVo)
			for _, order := range buyedPeriods {
				if order.Object_id == period.Id {
					periodVo.Is_buy = 1
					break
				}
			}
			if periodVo.Is_buy == 0 {
				totalPrice += period.Price
			}
			periodVos = append(periodVos, periodVo)
		}
	}
	return periodVos, totalPrice
}


func GetPeriodsByCourseId(courseId int) []Period {
	o := orm.NewOrm()
	var periods []Period
	qs := o.QueryTable("period")
	qs.Filter("course_id", courseId).All(&periods)
	return periods
}

func GetPeriodNumByCourseId(courseId int) int64 {
	num, _ := orm.NewOrm().QueryTable("period").Filter("course_id", courseId).Count()
	return num
}
