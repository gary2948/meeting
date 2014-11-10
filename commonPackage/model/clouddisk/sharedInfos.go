package clouddisk

import "time"

type Lctb_sharedInfos struct {
	Id                int64 `xorm:"pk autoincr"`
	Lc_userInfoId     int64
	Lc_shareCode      string    `xorm:"varchar(100)"`
	Lc_sharedLink     string    `xorm:"varchar(100)"`
	Lc_sharedName     string    `xorm:"varchar(100)"`
	Lc_sharedTime     time.Time `xorm:"created"`
	Lc_sharedSize     int64
	Lc_sharedType     int
	Lc_sharedFileType int
	Lc_sharedExt      string `xorm:"varchar(10)"`
	Lc_downloadCount  int
	Lc_saveCount      int
	Lc_viewCount      int
}
