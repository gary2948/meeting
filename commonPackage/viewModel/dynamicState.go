package viewModel

import (
	"time"
)

type DynamicState struct {
	IconAddr          string
	UserID            int64     `xorm:"'lc_user_info_id' <-"`
	NickName          string    `xorm:"'lc_user_name' <-"`
	MeetingId         int64     `xorm:"'id' <-"`
	MeetingAction     int       `xorm:"'lc_status' <-"`
	MeetingTopic      string    `xorm:"'lc_topic' <-"`
	MeetingCode       string    `xorm:"'lc_code' <-"`
	MeetingActionDate time.Time `xorm:"'lc_modify_time' <-"`
}
