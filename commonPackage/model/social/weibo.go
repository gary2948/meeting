package social

import (
	"time"
)

type Lctb_weibo struct {
	Id                int64 `xorm:"pk autoincr"`
	Lc_userInfoId     int
	Lc_weibo          string
	Lc_resFiles       string `xorm:"varchar(1000)"`
	Lc_resFilesCount  int
	Lc_postTime       time.Time
	Lc_status         int
	Lc_forwardWeiboId int
}
