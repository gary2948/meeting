package db

import (
	"commonPackage"
	"commonPackage/model/account"
	"service/uuid"
	"testing"
)

func TestAddAccount(t *testing.T) {

	//	u := &account.Lctb_userInfo{}
	//
	//	u.Lc_nickName = commonPackage.GetUnixTime()
	//	u.Lc_email = u.Lc_nickName + `@163.com`
	//	u.Lc_passwd = StringMD5Value("123456")
	//	u.Lc_mobilePhone = "1388603558"
	//	u.Lc_active = 1
	//	u.Lc_introduction = `测试数据`
	//	u.Lc_photoFile = `http://img.t.sinajs.cn/t5/style/images/face/face_friend.png`
	//	u.Lc_postAddress = `浦东大道1200号巨洋大厦1902室`
	//	u.Lc_realName = `测试的`
	//	_, err := AddAccount(u)
	//	if err != nil {
	//		t.Error(err)
	//	}
}

func TestGetAccount(t *testing.T) {
	users := &[]account.Lctb_userInfo{}
	err := GetAccountByPages(1, 5, 0, users)
	if err != nil {
		t.Error(err)
	}
	commonPackage.Printf(users)
}

func TestGetPersonEx(t *testing.T) {
	pe := &[]account.Lctb_personExperience{}
	err := GetPersonExperience(1, pe)
	if err != nil {
		t.Error(err)
	}
	commonPackage.Printf(pe)
}

func TestUUID(t *testing.T) {
	engine, _ := GetAccountEng()
	//checkError(err)
	users := &[]account.Lctb_userInfo{}

	engine.Find(users)

	n := len(*users)

	for i := 0; i < n; i++ {
		user := (*users)[i]
		user.Lc_UUID, _ = uuid.GenUUID()
		engine.Id(user.Id).Cols("lc_UUID").Update(&user)
	}

}
