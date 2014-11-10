package meeting

type Lctb_invitedPersons struct {
	Id               int64 `xorm:"pk autoincr"`
	Lc_meetingInfoId int64
	Lc_invitedEmail  string `xorm:"varchar(100)"`
}
