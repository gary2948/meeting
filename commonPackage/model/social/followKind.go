package social

type Lctb_followKind struct {
	Id                int64 `xorm:"pk autoincr"`
	Lc_userInfoId     int64
	Lc_followKindName string
}
