package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Live struct {
	Id              int
	Period          *Period `orm:"rel(one)"`
	Preview_url     string  `orm:"size(255)"`
	Complete_url    string  `orm:"size(255)"`
	Live_url        string  `orm:"size(255)"`
	Push_url        string  `orm:"size(255)"`
	Stream_id       string  `orm:"size(20)"`
	State           string
	Start_time      time.Time `orm:"auto_now_add;type(datetime)"`
	Real_start_time time.Time `orm:"auto_now_add;type(datetime)"`
	Last_push_time  time.Time `orm:"null;type(datetime)"`
	Create_time     time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time     time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time     time.Time `orm:"type(datetime)"`
	Duration        int
	Video_duration  int
}

func GetLiveById(liveId int) *Live {
	o := orm.NewOrm()
	var live Live
	err := o.QueryTable("live").Filter("id", liveId).RelatedSel("period").One(&live)
	// err := o.Read(&live)
	if err != nil {
		return nil
	}
	return &live
}

func GetLiveByPeriodId(periodId int) *Live {
	o := orm.NewOrm()
	var live Live
	err := o.QueryTable("live").Filter("period_id", periodId).One(&live)
	if err != nil {
		return nil
	}
	return &live
}

func UpdateLive(live *Live) bool {
	o := orm.NewOrm()
	live.Update_time = time.Now()
	_, err := o.Update(live)
	if err == nil {
		return true
	}
	return false
}
func GetLiveByStreamId(streamId string) *Live {
	o := orm.NewOrm()
	var live Live
	err := o.QueryTable("live").Filter("stream_id", streamId).One(&live)
	if err != nil {
		return nil
	}
	return &live
}
