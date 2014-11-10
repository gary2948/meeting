package social

type Lctb_userRelation struct {
	Id                  int64 `xorm:"pk autoincr"`
	Lc_userInfoId       int64
	Lc_followUserInfoId int64
	Lc_followUserName   string `xorm:"varchar(100)"`
	Lc_followKindId     int64
}
