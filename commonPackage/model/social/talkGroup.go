package social

import (
	"time"
)

type Lctb_talkGroup struct {
	Id              int64     `xorm:"pk autoincr"`
	Lc_createTime   time.Time `xorm:"created"`
	Lc_userInfoId   int64
	Lc_groupName    string `xorm:"varchar(100)"`
	Lc_filesMapId   int64
	Lc_cloudFileId  int64
	Lc_groupStatus  int
	Lc_maxUserCount int
}
