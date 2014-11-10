package meeting

import (
	"time"
)

type meetingActivity struct {
	Lc_userInfoId  int
	Lc_actUserName string
	Lc_actTime     time.Time
}
