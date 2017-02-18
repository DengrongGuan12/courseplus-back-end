package teacher

import (
	// "fmt"
	"encoding/json"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/base"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/err"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/models"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/util"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/vo"
	"time"
)

type LiveController struct {
	base.BaseController
}

// @router /detail [get]
func (this *LiveController) GetDetailById() {
	id, _ := this.GetInt("live_id")
	if id <= 0 {
		this.RetError(err.ErrInputData, "live_id")
	}
	livePo := models.GetLiveById(id)
	if livePo == nil {
		this.RetError(err.ErrInputData, "live nil")
	}
	// 判断是否是这个老师的直播
	teacherId := this.GetSession("teacherId").(int)
	teacher := models.GetTeacherByCourseId(livePo.Period.Course.Id)
	if teacher == nil || teacher.Id != teacherId {
		this.RetError(err.ErrPermission, "not your course")
	}
	var liveVo vo.Live
	util.PoToVo(livePo, &liveVo)

	//必须是预计开始时间前15分钟内才能获取推流地址
	intervalSec := util.GetIntervalToNow(livePo.Start_time)
	// fmt.Println("intervalSec:%d", intervalSec)
	if intervalSec <= -15*60 || livePo.State == "END" {
		this.Data["json"] = map[string]interface{}{"data": liveVo}
		this.ServeJSON()
		return
	}
	if liveVo.Push_url == "" {
		//地址的有效时间暂时设置为预计开始时间当天的23点59分
		liveVo.Push_url, livePo.Stream_id = util.GetPushUrl(teacherId, livePo.Id, util.GetEndOfThatDay(livePo.Start_time))
		livePo.Push_url = liveVo.Push_url
		livePo.Live_url = util.GetPlayUrl(teacherId, livePo.Id)
		models.UpdateLive(livePo)
	}
	this.Data["json"] = map[string]interface{}{"data": liveVo}
	this.ServeJSON()

}

// @router /tecentCallBack [post]
func (this *LiveController) TecentCallBack() {
	//腾讯云回调接口
	sign := this.Ctx.Input.Header("sign")
	t := this.Ctx.Input.Header("t")
	if !util.VerifyCallBackSign(sign, t) {
		this.Data["json"] = map[string]int{"code": 0}
		this.ServeJSON()
		return
	}
	event_type := this.Ctx.Input.Header("event_type")
	stream_id := this.Ctx.Input.Header("stream_id")
	if event_type == "100" {
		// 新录制文件生成
		var videoInfo models.Video
		json.Unmarshal(this.Ctx.Input.RequestBody, &videoInfo)
		models.AddVideo(videoInfo)
		live := models.GetLiveByStreamId(stream_id)
		video_duration := videoInfo.End_time - videoInfo.Start_time
		live.Video_duration = live.Video_duration + int(video_duration)
		if live.State == "END" {
			// 已结束直播需要判断是否转码完成
			if live.Video_duration == live.Duration {
				// 转码完成
				live.State = "CODED"
				models.UpdateLive(live)
			}
		}

	} else if event_type == "0" {
		// 断流, 1 教师下课 2 断网
		live := models.GetLiveByStreamId(stream_id)
		live.Duration = live.Duration + util.GetIntervalToNow(live.Last_push_time)
		models.UpdateLive(live)
	} else if event_type == "1" {
		// 推流,判断是不是第一次推
		live := models.GetLiveByStreamId(stream_id)
		if live.State == "BEFORE" {
			// 第一次推
			live.Real_start_time = time.Now()
			live.Last_push_time = time.Now()
			live.State = "ING"
			models.UpdateLive(live)
		} else if live.State == "ING" {
			// 中断后的推流
			live.Last_push_time = time.Now()
			models.UpdateLive(live)
		}
	} else if event_type == "200" {
		// 新截图文件
	}
	this.Data["json"] = map[string]int{"code": 0}
	this.ServeJSON()
}

// @router /finish [get]
func (this *LiveController) Finish() {
	// 下课
	liveId, _ := this.GetInt("liveId")
	if liveId <= 0 {
		this.RetError(err.ErrInputData, "live_id")
	}
	live := models.GetLiveById(liveId)
	if live == nil || live.State == "END" {
		this.RetError(err.ErrInputData, "live_id")
	}
	teacherId := this.GetSession("teacherId").(int)
	teacher := models.GetTeacherByCourseId(live.Period.Course.Id)
	if teacher == nil || teacher.Id != teacherId {
		this.RetError(err.ErrPermission, "not your course")
	}
	// 先更新记录
	live.State = "END"
	models.UpdateLive(live)
	util.CloseStream(live.Stream_id)
	this.Data["json"] = map[string]int{"data": 1}
	this.ServeJSON()
}
