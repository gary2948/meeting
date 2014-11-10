package account

import (
	"time"
)

type Lctb_personExperience struct {
	Id            int64 `xorm:"pk autoincr"`
	Lc_userInfoId int64
	Lc_experKind  int
	Lc_unitType   int
	Lc_unitName   string `xorm:"varchar(100)"`
	Lc_beginTime  time.Time
	Lc_endTime    time.Time
}
