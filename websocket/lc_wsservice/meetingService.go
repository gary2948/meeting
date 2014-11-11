package main

import (
	. "commonPackage"
	"commonPackage/model/clouddisk"
	"commonPackage/model/meeting"
	"service/db"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
)

func GetMeetingInfo(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	mId := jsonBody.Get(JSON_MEETINGID).MustInt64()
	mm := &meeting.Lctb_meetingInfo{}
	reJson := simplejson.New()
	ok, err := db.GetMeetingById(mId, mm)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else if ok {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, *mm)
	}
	return reJson.Encode()
}

func GetMeetingInvPer(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	mId := jsonBody.Get(JSON_MEETINGID).MustInt64()
	persons := &[]string{}
	reJson := simplejson.New()
	count, err := db.GetInvitedPersonEmails(mId, persons)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else if count > 0 {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, *persons)
	}
	return reJson.Encode()
}

func GetAttendPerson(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	mId := jsonBody.Get(JSON_MEETINGID).MustInt64()
	reJson := simplejson.New()
	attendPerson := make([]meeting.Lctb_attendMeeting, 0)
	count, err := db.GetAttendedPerson(mId, &attendPerson)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else if count > 0 {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, attendPerson)
	}
	return reJson.Encode()
}

func GetSimpleMeetingsByDate(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64()
	beginTime, err := jsonBody.Get(JSON_BEGINTIME).String()
	CheckError(err)
	endTime, err := jsonBody.Get(JSON_ENDTIME).String()
	CheckError(err)
	tag := jsonBody.Get(JSON_TAG).MustString("")
	CheckError(err)
	reJson := simplejson.New()

	r, err := db.GetSimpleMeetingsByDate(uId, beginTime, endTime)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_TAG, tag)
		reJson.Set(JSON_OK, true)
		content := make([]simplejson.Json, len(r))
		for i, v := range r {
			j := simplejson.New()
			j.Set(JSON_MEETINGDATE, string(v["meetingdate"]))
			j.Set(JSON_MEETINGDATE, string(v["meetingstatus"]))
			content[i] = *j
		}
		reJson.Set(JSON_CONTENT, content)
	}
	return reJson.Encode()
}

func GetUserMeetingByDate(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64()
	meetingDate, err := jsonBody.Get(JSON_MEETINGDATE).String()
	CheckError(err)
	tag := jsonBody.Get(JSON_TAG).MustString("")
	CheckError(err)
	meetings := []meeting.Lctb_meetingInfo{}
	reJson := simplejson.New()
	err = db.GetUserMeetingByDate(uId, meetingDate, &meetings)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, meetings)
		reJson.Set(JSON_TAG, tag)
	}

	return reJson.Encode()
}

func PlanMeetingByDate(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	isQuickStart, err := jsonBody.Get(JSON_QUICKSTART).Bool()
	CheckError(err)
	topic, err := jsonBody.Get(JSON_TOPIC).String()
	CheckError(err)
	uId := jsonBody.Get(JSON_USERID).MustInt64()
	emails, err := jsonBody.Get(JSON_EMAILS).StringArray()
	meetingInfo := &meeting.Lctb_meetingInfo{}
	reJson := simplejson.New()
	if isQuickStart {
		meetingInfo.Lc_planTime = time.Now()
		meetingInfo.Lc_beginTime = meetingInfo.Lc_planTime
	} else {
		planDate, err := jsonBody.Get(JSON_PLANDATE).String()
		CheckError(err)
		planSpan, err := jsonBody.Get(JSON_PLANSPAN).String()
		CheckError(err)
		shema, err := jsonBody.Get(JSON_SCHEMA).String()
		CheckError(err)
		password, err := jsonBody.Get(JSON_PASSWORD).String()
		CheckError(err)
		meetingInfo.Lc_planTime, err = time.Parse(TimeFormat, planDate)
		meetingInfo.Lc_planTimespan, err = time.Parse(TimeFormat, planSpan)
		meetingInfo.Lc_schema = shema
		meetingInfo.Lc_passwd = password
	}
	meetingInfo.Lc_topic = topic
	meetingInfo.Lc_userInfoId = uId
	code, url, err := db.AddMeetingAndInvPer(meetingInfo, emails)

	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, *meetingInfo)
		reJson.Set(JSON_CODE, code)
		reJson.Set(JSON_URL, url)
	}
	return reJson.Encode()
}

func GetMeetingFiles(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	mId := jsonBody.Get(JSON_MEETINGID).MustInt64()
	tag := jsonBody.Get(JSON_TAG).MustString("")
	reJson := simplejson.New()
	files := []clouddisk.Lctb_cloudFiles{}
	meetingInfo := meeting.Lctb_meetingInfo{}
	has, err := db.GetMeetingById(mId, &meetingInfo)
	if has {
		reJson.Set(JSON_TAG, tag)
		err = db.GetShareFiles(meetingInfo.Lc_shareInfoId, &files)
		if err == nil {
			reJson.Set(JSON_OK, true)
			reJson.Set(JSON_CONTENT, files)
		} else {
			reJson.Set(JSON_ERR, ErrSys)
			reJson.Set(JSON_OK, false)
		}
	} else {
		reJson.Set(JSON_ERR, ErrMeetingId)
		reJson.Set(JSON_OK, false)

	}

	return reJson.Encode()
}
