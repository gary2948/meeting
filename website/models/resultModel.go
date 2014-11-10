package models

import (
	"commonPackage/model/meeting"
	"time"
)

type ResultModel struct {
	RequestResult bool
	ErrorMsg      string
	Data          interface{}
}

type FileInfoModel struct {
	FileName   string
	FileSize   int64
	FileID     int64
	FileType   int32
	CreateTime time.Time
}

type DataGridMeetingResult struct {
	TotalRow int64
	Rows     []meeting.Lctb_meetingInfo
}

const (
	Result_SessionLose = iota //session丢失
	Result_ParmError          //参数错误
)

type UploadFileResponseModel struct {
	Code int
	Data interface{}
}

const (
	UploadingPercent = iota
	Finished
)
