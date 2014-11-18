package db

import (
	"commonPackage"
	"commonPackage/model/account"
	"commonPackage/model/clouddisk"
	"commonPackage/viewModel"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"service/uuid"
)

func StringMD5Value(value string) (result string) {
	var md = md5.New()
	md.Write([]byte(value))
	result = hex.EncodeToString(md.Sum(nil))
	return
}

var u *account.Lctb_userInfo = &account.Lctb_userInfo{}
var m *account.Lctb_memberInfo = &account.Lctb_memberInfo{}
var pe *account.Lctb_personExperience = &account.Lctb_personExperience{}
var ue *account.Lctb_userInfoEx = &account.Lctb_userInfoEx{}
var uc *account.Lctb_userClient = &account.Lctb_userClient{}

// initial all the tables of account db,if not exsit, create one
func InitAccountTables() (err error) {
	//	init()

	fmt.Println("-----InitAccountTables-----")
	engine, err := postgresEngine(accountDB)
	checkError(err)
	defer engine.Close()

	err = engine.Sync(u)
	//err = engine.CreateTables(u)
	checkError(err)

	err = engine.Sync(m)
	//err = engine.CreateTables(m)
	checkError(err)

	err = engine.Sync(pe)
	//err = engine.CreateTables(pe)
	checkError(err)

	err = engine.Sync(ue)
	//err = engine.CreateTables(ue)
	checkError(err)

	err = GetUserSqliteEng().Sync(uc)
	checkError(err)

	return err
}

// Add a new user account
func AddAccount(user *account.Lctb_userInfo) (int64, error) {
	user.Lc_passwd = StringMD5Value(user.Lc_passwd)
	user.Lc_UUID, _ = uuid.GenUUID()
	engine, err := GetAccountEng()
	checkError(err)
	_, err = engine.Insert(user)

	if err == nil {
		uc := &clouddisk.Lctb_userCloudSize{}
		uc.Lc_userInfoId = user.Id
		uc.Lc_maxCloudSize = 1024 * 1024 * 100 //默认100M

		addUserCloudSize(uc)
	}
	return user.Id, err
}

func ExchangePwd(userId int64, pwd string) error {
	engine, err := GetAccountEng()
	checkError(err)

	user := account.Lctb_userInfo{}
	user.Lc_passwd = StringMD5Value(pwd)

	_, err = engine.Id(userId).Cols("lc_passwd").Update(&user)

	return err
}

func UpdatePhotoFile(userId int64, photoFile string) error {
	engine, err := GetAccountEng()
	checkError(err)

	user := account.Lctb_userInfo{}
	user.Lc_photoFile = photoFile

	_, err = engine.Id(userId).Cols("lc_photo_file").Update(&user)

	return err

}

func UpdateAccount(userId int64, userInfo *viewModel.EditUserInfoModel) error {
	engine, err := GetAccountEng()
	checkError(err)
	user := account.Lctb_userInfo{}
	has, err := engine.Id(userId).Get(&user)
	if has {
		user.Lc_email = userInfo.Lc_email
		user.Lc_nickName = userInfo.Lc_nickName
		user.Lc_realName = userInfo.Lc_realName
		user.Lc_mobilePhone = userInfo.Lc_mobilePhone
		user.Lc_introduction = userInfo.Lc_introduction
		user.Lc_sex = userInfo.Lc_sex
		user.Lc_postAddress = userInfo.Lc_postAddress
		_, err = engine.Id(userId).Update(user)
	} else {
		err = commonPackage.NewErr(commonPackage.ErrUserId)
	}
	return err
}

func UpdateAccountEx(userId int64, userInfo *viewModel.EditUserInfoExModel) error {
	engine, err := GetAccountEng()
	checkError(err)
	user := account.Lctb_userInfoEx{}
	has, err := engine.Where("lc_user_info_id = ?", userId).Get(&user)
	if has {
		user.Lc_identityCard = userInfo.Lc_identityCard
		user.Lc_qq = userInfo.Lc_qq
		user.Lc_weibo = userInfo.Lc_weibo
		user.Lc_weixin = userInfo.Lc_weixin
		user.Lc_mobilePhone1 = userInfo.Lc_mobilePhone1
		user.Lc_mobilePhone2 = userInfo.Lc_mobilePhone2
		user.Lc_language = commonPackage.GetLanguageInt(userInfo.Lc_language)
		user.Lc_language1 = commonPackage.GetLanguageInt(userInfo.Lc_language1)
		user.Lc_language2 = commonPackage.GetLanguageInt(userInfo.Lc_language2)
		user.Lc_email1 = userInfo.Lc_email1
		user.Lc_email2 = userInfo.Lc_email2
		user.Lc_telephone = userInfo.Lc_telephone
		user.Lc_telephone1 = userInfo.Lc_telephone1
		user.Lc_telephone2 = userInfo.Lc_telephone2
		user.Lc_birthday = userInfo.Lc_birthday
		user.Lc_hometown = userInfo.Lc_hometown
		_, err = engine.Where("lc_user_info_id = ?", userId).Update(user)
	} else {
		err = commonPackage.NewErr(commonPackage.ErrUserId)
	}

	return err
}

// Add a new user account ex info
func AddAccountEx(userEx *account.Lctb_userInfoEx) (int64, error) {
	engine, err := GetAccountEng()
	checkError(err)

	return engine.Insert(userEx)
}

//Get a helper account, if not exsit create a new one
//Helper account is used for share meeting files or something else.
func GetHelperAccount(user *account.Lctb_userInfo) error {
	engine, err := GetAccountEng()
	checkError(err)

	has, err := engine.Where("lc_email=?", helperEmail).Get(user)

	if has {
		return err
	} else {
		user.Lc_email = helperEmail
		user.Lc_passwd = helpPasswd
		user.Lc_nickName = "会议小助手"
		_, err := AddAccount(user)
		checkError(err)
		return err
	}
	return err
}

//login by email,return is successed, If login fail, the eroor will
//return the reasons
func LoginByEmail(email, passwd string, user *account.Lctb_userInfo) (bool, error) {
	has, err := GetAccountByEmail(email, user)
	pwd := StringMD5Value(passwd)

	if !has {
		err = errors.New(commonPackage.ErrUserInfo)
	} else if user.Lc_passwd != pwd {
		err = errors.New(commonPackage.ErrPasswd)
	} else {
		return true, err
	}

	return false, err
}

func GetAccountsById(userIds []int64, users *[]account.Lctb_userInfo) error {
	engine, err := GetAccountEng()
	checkError(err)

	err = engine.In("id", userIds).Find(users)

	return err
}

//Get a user account by userId
func GetAccountById(userId int64, user *account.Lctb_userInfo) (bool, error) {
	engine, err := GetAccountEng()
	checkError(err)

	has, err := engine.Id(userId).Get(user)

	if !has {
		err = errors.New(commonPackage.ErrUserId)
	}

	return has, err
}

//Get a user account by userId
func GetAccountExById(userId int64, userEx *account.Lctb_userInfoEx) (bool, error) {
	engine, err := GetAccountEng()
	checkError(err)

	has, err := engine.Where("lc_user_info_id = ?", userId).Get(userEx)

	if !has {
		err = errors.New(commonPackage.ErrUserId)
	}

	return has, err
}

func GetPersonExperience(userId int64, pe *[]account.Lctb_personExperience) error {
	engine, err := GetAccountEng()
	checkError(err)
	err = engine.Where("lc_user_info_id=?", userId).Find(pe)
	return err
}

func AddPersonExperience(pe account.Lctb_personExperience) (int64, error) {
	engine, err := GetAccountEng()
	checkError(err)

	n, err := engine.Insert(pe)
	if n == 0 || err != nil {
		return int64(-1), errors.New(commonPackage.ErrSys)
	}

	return pe.Id, err
}

func UpdatePersonExperience(Id int64, pe account.Lctb_personExperience) error {
	engine, err := GetAccountEng()
	checkError(err)

	n, err := engine.Id(Id).Update(&pe)
	if n == 0 || err != nil {
		return errors.New(commonPackage.ErrSys)
	}

	return err
}

//Get a user account by email
func GetAccountByEmail(email string, user *account.Lctb_userInfo) (bool, error) {
	engine, err := GetAccountEng()
	checkError(err)
	has, err := engine.Where("lc_email=?", email).Get(user)
	return has, err
}

func UserClientOnline(userId int, uc *account.Lctb_userClient) {

}

func UserClientOffline(userId int, uc *account.Lctb_userClient) {

}
