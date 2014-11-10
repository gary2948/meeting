package clouddisk

import (
	"time"
)

type Lctb_cloudFiles struct {
	Id                int64 `xorm:"pk autoincr"`
	Lc_filesMapId     int64
	Lc_fileName       string `xorm:"varchar(260)"`
	Lc_fileExtension  string `xorm:"varchar(10)"`
	Lc_fileType       int
	Lc_parentId       int64
	Lc_fileStatus     int
	Lc_createTime     time.Time
	Lc_linkFilesMapId int64
	Lc_userInfoId     int64
	Lc_fileSize       int64
}
