package viewModel

import (
	"time"
)

type EditUserInfoExModel struct {
	Lc_identityCard string
	Lc_qq           string
	Lc_weibo        string
	Lc_weixin       string
	Lc_mobilePhone1 string
	Lc_mobilePhone2 string
	Lc_language     string
	Lc_language1    string
	Lc_language2    string
	Lc_email1       string
	Lc_email2       string
	Lc_telephone    string
	Lc_telephone1   string
	Lc_telephone2   string
	Lc_birthday     time.Time
	Lc_hometown     string
}
