package db

import (
	"commonPackage"
	"testing"
	//	"time"
)

func TestSocial(t *testing.T) {
	TFollowUser(t)
	//TGetFollows(t)
	//TGetFanes(t)
	//TSendPersonMessage(t)
}

func TFollowUser(t *testing.T) {
	followIds := []int64{1}
	err := FollowUser(8, 0, followIds)
	if err != nil {
		t.Error(err)
	}
}

//func TGetFollows(t *testing.T) {
//	users := []account.Lctb_userInfo{}
//	err := GetFollows(1, &users)
//	if err != nil {
//		t.Error(err)
//	}
//	commonPackage.Printf(users)
//}

//func TGetFanes(t *testing.T) {
//	users := []account.Lctb_userInfo{}
//	err := GetFans(8, &users)
//	if err != nil {
//		t.Error(err)
//	}
//	commonPackage.Printf(users)
//}

func TSendPersonMessage(t *testing.T) {
	//msg := "123test"
	//_, err := SendPersonMessage(1, 8, msg)
	//if err != nil {
	//	t.Error(err)
	//}
}

//func TestCreateGroup(t *testing.T) {
//	_, err := CreateGroup(1, "test")
//	if err != nil {
//		t.Error(err)
//	}
//}

//func TestAddGroupMember(t *testing.T) {
//	userIds := []int64{2, 3, 4, 5}
//	err := AddUserToGroup(1, userIds)
//	if err != nil {
//		t.Error(err)
//	}
//}

//func TestDelUserFromGroup(t *testing.T) {
//	userIds := []int64{4, 5}
//	err := DelUserFromGroup(2, userIds)
//	if err != nil {
//		t.Error(err)
//	}
//}

//func TestUserGroups(t *testing.T) {

//	gms := []social.Lctb_talkGroup{}
//	err := GetUserGroups(5, &gms)
//	if err != nil {
//		t.Error(err)
//	}

//	commonPackage.Printf(gms)
//}

//func TestGetGroupInfo(t *testing.T) {
//	gm := social.Lctb_talkGroup{}
//	us := []account.Lctb_userInfo{}
//	err := GetGroupInfo(1, &gm, &us)
//	if err != nil {
//		t.Error(err)
//	}

//	commonPackage.Printf(gm)
//	commonPackage.Printf(us)
//}

//func TestSendGroupMsg(t *testing.T) {
//	msg := `{"USERID":1,"TOUSERID":8,"COLOR":"#000000","CONTENT":"[[{\"INFO\":\"sdadsav\",\"TYPE\":2}]]","FONT":"Microsoft YaHei,9,-1,5,50,0,0,0,0,0","MSGID":"d112433d-a1c6-4c97-a2d9-2f5768c61a09","TIME":"2014-07-09 15:47:14","ADJUSTTIME":"2014-07-09 15:47:13"}`
//	msgId, userIds, err := SendGroupMsg(1, 1, msg, commonPackage.TextMsg)
//	if err != nil {
//		t.Error(err)
//	}

//	commonPackage.Printf(msgId)
//	commonPackage.Printf(userIds)
//}

func TestGetUnreadGroupMsg(t *testing.T) {
	txtMsg, imgMsg, err := GetUnreadGroupMsg(1, 4)
	if err != nil {
		t.Error(err)
	}

	commonPackage.Printf(txtMsg)
	commonPackage.Printf(imgMsg)
}

//func TestTimeFmt(t *testing.T) {

//	tt := time.Now().Format("2006-01-02 15:04:05")
//	commonPackage.Println(tt)
//	commonPackage.Println(time.UTC.String())
//}

//func TestFollowKind(t *testing.T) {
//	id, err := AddFollowKind(1, "test")
//	if err != nil {
//		t.Error(err)
//		commonPackage.Println(err.Error())
//		return
//	}

//	commonPackage.Println("add ok")

//	err = RenameFollowKind(id, "rename")
//	if err != nil {
//		t.Error(err)
//	}
//	commonPackage.Println("rename ok")

//	uids := []int64{8, 9, 10}
//	err = MoveUserToFollowKind(id, 1, uids)
//	if err != nil {
//		t.Error(err)
//	}
//	commonPackage.Println("move ok")

//}
