package account

import (
	"time"
)

type Lctb_userInfoEx struct {
	Id              int64 `xorm:"pk autoincr"`
	Lc_userInfoId   int64
	Lc_identityCard string `xorm:"varchar(50)"`
	Lc_qq           string `xorm:"varchar(20)"`
	Lc_weibo        string `xorm:"varchar(100)"`
	Lc_weixin       string `xorm:"varchar(100)"`
	Lc_mobilePhone1 string `xorm:"varchar(20)"`
	Lc_mobilePhone2 string `xorm:"varchar(20)"`
	Lc_language     int
	Lc_language1    int
	Lc_language2    int
	Lc_email1       string `xorm:"varchar(100)"`
	Lc_email2       string `xorm:"varchar(100)"`
	Lc_telephone    string `xorm:"varchar(25)"`
	Lc_telephone1   string `xorm:"varchar(25)"`
	Lc_telephone2   string `xorm:"varchar(25)"`
	Lc_birthday     time.Time
	Lc_hometown     string `xorm:"varchar(100)"`
}
