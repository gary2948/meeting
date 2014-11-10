package db

import (
	"commonPackage/model/meeting"
	"fmt"
	//	"strconv"
	"commonPackage"
	"commonPackage/viewModel"
	"testing"
	"time"
)

//func TestAddMeeting(t *testing.T) {

//	meet := &meeting.Lctb_meetingInfo{Lc_userInfoId: 1, Lc_planTime: time.Now(), Lc_topic: "test"}
//	_, code, err := AddMeeting(meet)

//	fmt.Println(code)

//	if err != nil {
//		t.Errorf("err")
//	}

//}

func TestGetMeetingsByPlanTime(t *testing.T) {

	meetings := make([]meeting.Lctb_meetingInfo, 0)
	begingTime := time.Date(2014, 01, 01, 0, 0, 0, 0, time.Local)
	endTime := time.Date(2015, 01, 01, 0, 0, 0, 0, time.Local)
	err := GetMeetingsByPlanTime(1, begingTime, endTime, &meetings)
	if len(meetings) == 0 {
		t.Errorf("err not found meetings")
		t.Errorf(begingTime.Format("2006-01-02 15:04:05"))
	}
	if err != nil {
		t.Errorf(err.Error())
	}

}

//func TestAddInvitedPersonEmail(t *testing.T) {
//	email := "zhao@163.com"

//	_, err := AddInvitedPersonEmail(1, email)
//	if err != nil {
//		t.Errorf(err.Error())
//	}
//}

//func TestAddAttendedPersonEmail(t *testing.T) {
//	email := "zhao@163.com"
//	_, err := AddAttendedPersonEmail(1, email, 1)
//	if err != nil {
//		t.Errorf(err.Error())
//	}
//}

func TestGetInvitedPersonEmails(t *testing.T) {

	emails := make([]string, 0)
	_, err := GetInvitedPersonEmails(1, &emails)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(emails) == 0 {
		t.Errorf("err not found email")
	}
	fmt.Print(emails)
}

func TestGetAttendPerson(t *testing.T) {

	persons := make([]meeting.Lctb_attendMeeting, 0)
	n, err := GetAttendedPerson(1, &persons)
	if err != nil {
		t.Errorf(err.Error())
	}
	if n == 0 {
		t.Errorf("err not found email")
	}
	fmt.Println(n)
	for _, v := range persons {
		fmt.Println(v)
	}
}

//func TestGetMeetingById(t *testing.T) {
//	meet := &meeting.Lctb_meetingInfo{}
//	has, err := GetMeetingById(1, meet)
//	if !has {
//		t.Errorf("err not found meeting")
//	}
//	if err != nil {
//		t.Errorf(err.Error())
//	}
//	fmt.Print(meet)
//}

func TestGetSimpleMeetingsByDate(t *testing.T) {

	beginTime := "2014-06-01"
	endTime := "2014-06-30"

	r, err := GetSimpleMeetingsByDate(1, beginTime, endTime)
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, v := range r {
		fmt.Printf("%v\n", v)
	}

}

func TestGetUserMeetingByDate(t *testing.T) {

	meetingTime := "2014-06-23"
	meet := []meeting.Lctb_meetingInfo{}
	err := GetUserMeetingByDate(1, meetingTime, &meet)

	if err != nil {
		t.Errorf(err.Error())
	}

	for _, v := range meet {
		fmt.Println(v)
	}

}

//func TestAddInvitedPersonEmails(t *testing.T) {
//	emails := []string{"1403836686@163.com"}
//	meeting := &meeting.Lctb_meetingInfo{Lc_userInfoId: 1, Lc_planTime: time.Now(), Lc_topic: "test1"}
//	//for i := 0; i < 1; i++ {
//	//	emails[i] = "zhao@" + strconv.Itoa(i) + ".com"
//	//}
//	code, _, err := AddMeetingAndInvPer(meeting, emails)
//	fmt.Println(code)
//	if err != nil {
//		t.Errorf(err.Error())
//	}
//}

//func TestGetMeetingsByPlanStatus(t *testing.T) {
//	//_, err := GetMeetingsCountByPlanStatus(1, -1)
//	//if err != nil {
//	//	t.Errorf(err.Error())
//	//}
//	meetings := &[]meeting.Lctb_meetingInfo{}
//	pageSize := 5
//	pageCount := 1
//	_, err := GetMeetingsByPlanStatus(1, -1, pageSize, pageCount, meetings)
//	if err != nil {
//		t.Errorf(err.Error())
//	}
//}

//func TestDelMeeting(t *testing.T) {
//	err := DelMeeting(1)
//	if err != nil {
//		t.Errorf(err.Error())
//	}
//}

func TestGetDynamicState(t *testing.T) {
	dynamicState := []viewModel.DynamicState{}
	err := GetDynamicState(1, &dynamicState)
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, v := range dynamicState {
		commonPackage.Printf(v)
	}
}
