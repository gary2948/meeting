package clouddisk

type Lctb_filesMap struct {
	Id              int64  `xorm:"pk autoincr"`
	Lc_fileSystemId string `xorm:"varchar(100) unique"`
	Lc_usedCount    int
}
