package main

import (
	"commonPackage/model/account"
	"commonPackage/model/meeting"
	"encoding/json"
)

const (
	//会议相关接口
	startMeeting         = "6001000" //启动会议
	reStartMeeting       = "6001010" //返回启动会议
	loginMeeting         = "6002000" //登录会议室
	reLoginMeeting       = "6002010" //返回结果
	boardLoginMeeting    = "6002020" //转发登录结果
	logoutMeeting        = "6003000" //登出会议室
	reLogoutMeeting      = "6003010" //返回结果
	boardLogoutMeeting   = "6003020" //装发登出结果
	mouseMove            = "6004000" //鼠标移动
	reMouseMove          = "6004010" //返回鼠标移动
	boradMouseMove       = "6004020" //转发鼠标消息
	getOtherUserInfo     = "6005000" //获取其他用户信息
	reGetOtherUserInfo   = "6005010" //返回其他用户信息
	synGeModel           = "6006000" //同步图形
	reSynGeModel         = "6006010" // 返回图形
	boardGeModel         = "6006020" //转发图形
	synRoomPermission    = "6007000" //同步用户设置
	reSynRoomPermission  = "6007010" // 返回用户设置
	boardRoomPermission  = "6007020" // 同步用户设置
	synMeetingFileInfo   = "6008000" //获取会议文件信息设置
	reSynMeetingFileInfo = "6008010" //返回会议文件设置
	boardMeetingFileInfo = "6008020" //广播文件设置
)

//会议结构体
type meetingInfo struct {
	MEETINGID int64
	OK        bool
	CONTENT   meeting.Lctb_meetingInfo
	ERR       string
}

type userInfo struct {
	MEETINGID int64
	EMAIL     string
	OK        bool
	CONTENT   account.Lctb_userInfo
	ERR       string
}

type mousePoints struct {
	MEETINGID int64
	EMAIL     string
	CONTENT   json.RawMessage
}

type geModel struct {
	MEETINGID int64
	EMAIL     string
	GeID      string
	CONTENT   json.RawMessage
}

type userPower struct {
	MEETINGID int64
	EMAIL     string
	POWERID   int64
	CONTENT   json.RawMessage
}

type meetingFileInfo struct {
	MEETINGID int64
	EMAIL     string
	FILEID    int64
	CONTENT   json.RawMessage
}
