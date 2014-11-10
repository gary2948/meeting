package social

type Lctb_groupMember struct {
	Id               int64 `xorm:"pk autoincr"`
	Lc_talkGroupId   int64
	Lc_userInfoId    int64
	Lc_userInfoRole  int
	Lc_userGroupName string
}
