package account

import (
	"time"
)

type Lctb_userInfo struct {
	Id              int64  `xorm:"pk autoincr"`
	Lc_email        string `xorm:"varchar(100) unique"`
	Lc_nickName     string `xorm:"varchar(100)"`
	Lc_realName     string `xorm:"varchar(50)"`
	Lc_mobilePhone  string `xorm:"varchar(20)"`
	Lc_passwd       string `xorm:"varchar(50)"`
	Lc_active       int
	Lc_introduction string    `xorm:"varchar(200)"`
	Lc_regTime      time.Time `xorm:"created"`
	Lc_sex          string    `xorm:"varchar(5)"`
	Lc_postAddress  string    `xorm:"varchar(200)"`
	Lc_photoFile    string    `xorm:"varchar(500)"`
	Lc_UUID         string    `xorm:"'lc_UUID' varchar(50)"`
}

type Lctb_userClient struct {
	Id             int64 `xorm:"pk autoincr"`
	Lc_userInfoId  int64
	Lc_userAddr    string
	Lc_userStatus  int
	Lc_OnLineTime  time.Time
	Lc_OffLineTime time.Time
}
