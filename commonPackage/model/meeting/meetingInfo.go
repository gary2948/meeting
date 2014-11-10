package meeting

import (
	"time"
)

type Lctb_meetingInfo struct {
	Id              int64 `xorm:"pk autoincr"`
	Lc_preMeetingId int64
	Lc_userInfoId   int64
	Lc_userName     string `xorm:"varchar(100)"`
	Lc_createTime   time.Time
	Lc_modifyTime   time.Time `xorm:"updated"`
	Lc_planTime     time.Time
	Lc_planTimespan time.Time
	Lc_topic        string `xorm:"varchar(200)"`
	Lc_schema       string `xorm:"text`
	Lc_remindTime   int
	Lc_remindType   int
	Lc_code         string `xorm:"varchar(50)"`
	Lc_passwd       string `xorm:"varchar(50)"`
	Lc_limitCount   int
	Lc_beginTime    time.Time
	Lc_endTime      time.Time
	Lc_status       int
	Lc_delete       int
	Lc_shareInfoId  int64
	Lc_cloudFilesId int64
}
