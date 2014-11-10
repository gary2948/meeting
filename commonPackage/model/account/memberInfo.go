package account

type Lctb_memberInfo struct {
	Id            int64 `xorm:"pk autoincr"`
	Lc_userInfoId int64
	Lc_memberRank int
}
