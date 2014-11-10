package clouddisk

import "time"

type Lctb_recycleBinFiles struct {
	Id              int64 `xorm:"pk autoincr"`
	Lc_cloudFilesId int64
	Lc_userInfoId   int64
	Lc_delTime      time.Time `xorm:"created"`
}
