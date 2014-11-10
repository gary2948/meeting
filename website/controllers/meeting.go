package controllers

import ()

type MeetingController struct {
	BaseController
}

func (h *MeetingController) MeetingInfo() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/meeting/meetinginfo.html"
	}
}

func (h *MeetingController) AddMeeting() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/meeting/addmeeting.html"
	}
}

func (h *MeetingController) MeetingLog() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/meeting/meetinglog.html"
	}
}

func (h *MeetingController) ScheduleManager() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/meeting/schedulemanager.html"
	}
}
