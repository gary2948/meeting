package social

import (
	"time"
)

type Lctb_personMessage struct {
	Id                int64 `xorm:"pk autoincr"`
	Lc_fromUserInfoId int64
	Lc_toUserInfoId   int64
	Lc_messageContext string    `xorm:"TEXT"`
	Lc_postTime       time.Time `xorm:"created"`
	Lc_messageStatus  int
	Lc_UUID           string `xorm:"'lc_UUID'"`
}

type Lctb_personMessageImg struct {
	Id                int64 `xorm:"pk autoincr"`
	Lc_fromUserInfoId int64
	Lc_toUserInfoId   int64
	Lc_messageContext string    `xorm:"TEXT"`
	Lc_postTime       time.Time `xorm:"created"`
	Lc_messageStatus  int
	Lc_UUID           string `xorm:"'lc_UUID'"`
}
