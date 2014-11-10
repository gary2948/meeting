package meeting

import (
	"time"
)

type Lctb_operateHistory struct {
	Id               int64 `xorm:"pk autoincr"`
	Lc_meetingInfoId int
	Lc_operater      string `xorm:"varchar(100)"`
	Lc_operateType   int
	Lc_operateObject string `xorm:"varchar(100)"`
	Lc_operateTime   time.Time
}
