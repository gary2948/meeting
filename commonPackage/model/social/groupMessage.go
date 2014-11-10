package social

import (
	"time"
)

type Lctb_groupMessage struct {
	Id                int64 `xorm:"pk autoincr"`
	Lc_talkGoupId     int64
	Lc_userInfoId     int64
	Lc_messageContext string    `xorm:"TEXT"`
	Lc_postTime       time.Time `xorm:"created"`
}

type Lctb_groupMessageImg struct {
	Id                int64 `xorm:"pk autoincr"`
	Lc_talkGoupId     int64
	Lc_userInfoId     int64
	Lc_messageContext string    `xorm:"TEXT"`
	Lc_postTime       time.Time `xorm:"created"`
}

type Lctb_groupMessageStatus struct {
	Id                int64 `xorm:"pk autoincr"`
	Lc_talkGoupId     int64
	Lc_groupMessageId int64
	Lc_messageStatus  int
	Lc_userInfoId     int64
}

type Lctb_groupMessageImgStatus struct {
	Id                   int64 `xorm:"pk autoincr"`
	Lc_talkGoupId        int64
	Lc_groupMessageImgId int64
	Lc_messageStatus     int
	Lc_userInfoId        int64
}
