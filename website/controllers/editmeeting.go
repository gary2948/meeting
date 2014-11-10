package controllers

import (
	"commonPackage"
	"commonPackage/model/account"
	"commonPackage/model/meeting"
	"github.com/astaxie/beego"
	"service/db"
	"strconv"
)

type EditMeetingController struct {
	beego.Controller
}

func (this *EditMeetingController) Post() {
	beego.Info(this.Ctx.Input.Param(":id"))                                 //获取会议的id
	meetingId, err := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64) //数据转换
	if err != nil {
		//在此处实现错误地址处理逻辑
	}
	beego.Info(meetingId)
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo) //获取当前用户
	meetingInfo := meeting.Lctb_meetingInfo{}
	result, err := db.GetMeetingById(meetingId, &meetingInfo)
	if !result {
		//在此处添加没有相关会议的处理逻辑
	}
	var canEdit bool
	beego.Info(meetingInfo)
	if userinfo.Id == int64(meetingInfo.Lc_userInfoId) && meetingInfo.Lc_status == commonPackage.PlanMeeting {
		canEdit = true
	} else {
		canEdit = false
	}
	this.Data["canEdit"] = !canEdit
	this.Data["meetingInfo"] = meetingInfo
	this.Data["userInfo"] = userinfo
	if meetingInfo.Lc_status == commonPackage.PlanMeeting {
		this.Data["MeetingState"] = "预约会议"
	} else if meetingInfo.Lc_status == commonPackage.StartedMeeting {
		this.Data["MeetingState"] = "正在进行"
	} else {
		this.Data["MeetingState"] = "已经结束"
	}

	this.TplNames = "pages/editMeeting.html"
}

func (this *EditMeetingController) Get() {
	this.Post()
}
