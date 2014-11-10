package clouddisk

type Lctb_sharedUsers struct {
	Id              int64 `xorm:"pk autoincr"`
	Lc_sharedInfoId int64
	Lc_userInfoId   int64
	Lc_toUserInfoId int64
}
