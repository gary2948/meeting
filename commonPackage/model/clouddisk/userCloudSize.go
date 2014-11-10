package clouddisk

type Lctb_userCloudSize struct {
	Id               int64 `xorm:"pk autoincr"`
	Lc_userInfoId    int64
	Lc_usedCloudSize int64
	Lc_maxCloudSize  int64
}
