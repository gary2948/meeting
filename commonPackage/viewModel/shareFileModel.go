package viewModel

import (
	"time"
)

type ShareViewModel struct {
	ShareName     string    `xorm:"'lc_shared_name' <-"`      //分享的名称
	ShareCode     string    `xorm:"'lc_share_code' <-"`       //分享码
	DownloadCount int       `xorm:"'lc_download_count' <-"`   //下载次数
	ShareSize     int64     `xorm:"'lc_shared_size' <-"`      //分享的文件的大小
	ShareTime     time.Time `xorm:"'lc_shared_time' <-"`      //分享时间
	ShareId       int64     `xorm:"'id' <-"`                  //分享的记录的id
	ShareUser     int64     `xorm:"'lc_user_info_id' <-"`     //分享的人 Lc_sharedFileType
	ShareFileType int       `xorm:"'lc_shared_file_type' <-"` //分享记录是否可以展开 0表示单个文件 1表示单个文件夹 2表示多个文件或文件一起分享
	FILEEXT       string    `xorm:"'lc_shared_ext' <-"`       //分享后缀
}
