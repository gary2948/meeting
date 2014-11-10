package meeting

type Lctb_attendMeeting struct {
	Id               int64 `xorm:"pk autoincr"`
	Lc_meetingInfoId int64
	Lc_joinedEmail   string `xorm:"varchar(100)"`
	Lc_userType      int
}
