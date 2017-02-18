package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Video struct {
	Id          int
	Live        *Live  `orm:"rel(fk)"`
	Video_id    string `orm:"size(127)"`
	Video_url   string `orm:"size(255)"`
	File_size   int64
	Start_time  int64
	End_time    int64
	Stream_id   string    `orm:"size(20)"`
	File_format string    `orm:"size(7)"`
	File_id     string    `orm:"size(63)"`
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time `orm:"type(datetime)"`
}

func AddVideo(video Video) (int64, error) {
	video.Create_time = time.Now()
	video.Update_time = time.Now()
	live := GetLiveByStreamId(video.Stream_id)
	video.Live = live
	o := orm.NewOrm()
	id, err := o.Insert(&video)
	if err == nil {
		return id, nil
	}
	return 0, err
}
func GetVideosByLiveId(liveId int, offset int, limit int) []Video {
	var videos []Video
	o := orm.NewOrm()
	if limit == 0 {
		// 表示不限制
		o.QueryTable("video").Filter("live_id", liveId).All(&videos)
	} else {
		o.QueryTable("video").Filter("live_id", liveId).Limit(limit, offset).All(&videos)
	}
	return videos
}
