package clouddisk

type Lctb_sharedFiles struct {
	Id              int64 `xorm:"pk autoincr"`
	Lc_cloudFilesId int64
	Lc_sharedInfoId int64
}
