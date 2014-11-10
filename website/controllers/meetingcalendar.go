package controllers

import (
	"commonPackage/model/account"
	"github.com/astaxie/beego"
)

type MeetingCalendarController struct {
	beego.Controller
}

func (m *MeetingCalendarController) Post() {

}

func (m *MeetingCalendarController) Get() {
	userinfo, _ := m.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	m.Data["userinfo"] = userinfo
	m.TplNames = "pages/meetingCalendar.html"
}
