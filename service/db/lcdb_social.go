package db

import (
	"commonPackage"
	"commonPackage/model/account"
	"commonPackage/model/social"
	"commonPackage/viewModel"
	"fmt"
	"time"

	"github.com/go-xorm/xorm"
)

const (
	MaxGroupNumber = 100
)

var fk *social.Lctb_followKind = &social.Lctb_followKind{}
var gm *social.Lctb_groupMember = &social.Lctb_groupMember{}
var gms *social.Lctb_groupMessage = &social.Lctb_groupMessage{}
var gmi *social.Lctb_groupMessageImg = &social.Lctb_groupMessageImg{}
var gmts *social.Lctb_groupMessageStatus = &social.Lctb_groupMessageStatus{}
var gmis *social.Lctb_groupMessageImgStatus = &social.Lctb_groupMessageImgStatus{}
var pm *social.Lctb_personMessage = &social.Lctb_personMessage{}
var pmi *social.Lctb_personMessageImg = &social.Lctb_personMessageImg{}
var rp *social.Lctb_reply = &social.Lctb_reply{}
var tg *social.Lctb_talkGroup = &social.Lctb_talkGroup{}
var ur *social.Lctb_userRelation = &social.Lctb_userRelation{}
var wb *social.Lctb_weibo = &social.Lctb_weibo{}

func InitSocialTables() {
	fmt.Println("-----InitSocialTables -----")
	engine, err := postgresEngine(socialDB)
	checkError(err)
	defer engine.Close()

	err = engine.Sync(fk)
	checkError(err)

	err = engine.Sync(gm)
	checkError(err)

	err = engine.Sync(gms)
	checkError(err)

	err = engine.Sync(pm)
	checkError(err)

	err = engine.Sync(rp)
	checkError(err)

	err = engine.Sync(tg)
	checkError(err)

	err = engine.Sync(ur)
	checkError(err)

	err = engine.Sync(wb)
	checkError(err)

	err = engine.Sync(pmi)
	checkError(err)

	err = engine.Sync(gmi)
	checkError(err)

	err = engine.Sync(gmts)
	checkError(err)

	err = engine.Sync(gmis)
	checkError(err)

}

//关注用户
func FollowUser(userId int64, kindId int64, followUIds []int64) error {
	engine, err := GetSocialEng()
	commonPackage.CheckError(err)

	n := len(followUIds)
	if n > 0 {
		urs := make([]social.Lctb_userRelation, n)
		for i, v := range followUIds {
			ur := social.Lctb_userRelation{
				Lc_userInfoId:       userId,
				Lc_followUserInfoId: v,
				Lc_followKindId:     kindId,
			}
			urs[i] = ur
		}
		_, err = engine.Insert(&urs)
		commonPackage.CheckError(err)
	}
	return err
}

//取消关注用户
func UnFollowUser(userId, followUID int64) error {
	engine, err := GetSocialEng()
	commonPackage.CheckError(err)

	ur := social.Lctb_userRelation{}
	_, err = engine.Where("lc_user_info_id = ? and lc_follow_user_info_id = ?", userId, followUID).Delete(ur)

	return err
}

func AddFollowKind(userId int64, kindName string) (int64, error) {
	engine, err := GetSocialEng()
	commonPackage.CheckError(err)

	ur := social.Lctb_followKind{}
	ur.Lc_followKindName = kindName
	ur.Lc_userInfoId = userId

	_, err = engine.Insert(&ur)
	return ur.Id, err
}

func RenameFollowKind(kindId int64, kindName string) error {
	engine, err := GetSocialEng()
	commonPackage.CheckError(err)

	ur := social.Lctb_followKind{}
	ur.Lc_followKindName = kindName

	_, err = engine.Id(kindId).Cols("lc_follow_kind_name").Update(&ur)
	return err

}

func DelFollowKind(kindId int64) error {
	engine, err := GetSocialEng()
	commonPackage.CheckError(err)

	uk := social.Lctb_followKind{}
	_, err = engine.Id(kindId).Delete(&uk)

	return err

}

func MoveUserToFollowKind(kindId, userId int64, userIds []int64) error {
	engine, err := GetSocialEng()
	commonPackage.CheckError(err)

	ur := social.Lctb_userRelation{}
	ur.Lc_followKindId = kindId
	_, err = engine.Where("lc_user_info_id = ?", userId).In("lc_follow_user_info_id", userIds).Cols("lc_follow_kind_id").Update(&ur)

	return err

}

//得到关注列表
func GetFollows(userId int64, userKinds *[]viewModel.FollowKind) error {
	uks := make([]viewModel.FollowKind, 0)
	engine, err := GetSocialEng()
	commonPackage.CheckError(err)
	uengine, err := GetAccountEng()
	commonPackage.CheckError(err)

	//先取分类
	fks := []social.Lctb_followKind{}
	err = engine.Where("lc_user_info_id = ?", userId).Find(&fks)
	//取已分类的用户
	n := len(fks)
	if n > 0 {
		for _, v := range fks {

			vfk := viewModel.FollowKind{}
			vfk.KindId = v.Id
			vfk.KindName = v.Lc_followKindName
			//users := []account.Lctb_userInfo{}
			err = getUsersByKindId(engine, uengine, userId, v.Id, &vfk.CONTENT)
			commonPackage.CheckError(err)
			uks = append(uks, vfk)
		}
	}

	//取未分类的用户
	uk := viewModel.FollowKind{}
	uk.KindId = 0
	uk.KindName = "未分类"
	//	users := []account.Lctb_userInfo{}
	err = getUsersByKindId(engine, uengine, userId, 0, &uk.CONTENT)
	commonPackage.CheckError(err)
	uks = append(uks, uk)
	*userKinds = uks
	return err
}

func getUsersByKindId(engine, uengine *xorm.Engine, userId, kindId int64, users *[]account.Lctb_userInfo) error {
	urs := []social.Lctb_userRelation{}
	err := engine.Where("lc_user_info_id = ? and lc_follow_kind_id = ?", userId, kindId).Find(&urs)
	n := len(urs)
	if n > 0 {
		userIds := make([]int64, n)
		for i, v := range urs {
			userIds[i] = v.Lc_followUserInfoId
		}
		err = uengine.In("id", userIds).Find(users)
		commonPackage.CheckError(err)
	}
	return err
}

//得到粉丝列表
func GetFans(userId int64, userKinds *[]viewModel.FollowKind) error {
	uks := make([]viewModel.FollowKind, 0)
	engine, err := GetSocialEng()
	uengine, err := GetAccountEng()
	commonPackage.CheckError(err)

	uk := viewModel.FollowKind{}
	uk.KindId = 0
	uk.KindName = "粉丝"
	urs := []social.Lctb_userRelation{}
	err = engine.Where("lc_follow_user_info_id = ?", userId).Find(&urs)
	commonPackage.CheckError(err)
	n := len(urs)
	if n > 0 {
		userIds := make([]int64, n)
		for i, v := range urs {
			userIds[i] = v.Lc_userInfoId
		}
		users := []account.Lctb_userInfo{}
		err = uengine.In("id", userIds).Find(&users)
		uk.CONTENT = users
		commonPackage.CheckError(err)
	}
	uks = append(uks, uk)
	*userKinds = uks
	return err
}

//发私信
func SendPersonMessage(sendUserId, resUserId int64, postTime time.Time, mes string, status, msgType int) (int64, error) {
	engine, err := GetSocialEng()
	commonPackage.CheckError(err)

	var id int64
	if msgType == commonPackage.TextMsg {
		pm := social.Lctb_personMessage{
			Lc_fromUserInfoId: sendUserId,
			Lc_toUserInfoId:   resUserId,
			Lc_messageContext: mes,
			Lc_postTime:       postTime,
			Lc_messageStatus:  status,
		}
		_, err = engine.Insert(&pm)
		commonPackage.CheckError(err)
		id = pm.Id
	} else if msgType == commonPackage.ImgMsg {
		pm := social.Lctb_personMessageImg{
			Lc_fromUserInfoId: sendUserId,
			Lc_toUserInfoId:   resUserId,
			Lc_messageContext: mes,
			Lc_postTime:       postTime,
			Lc_messageStatus:  status,
		}
		_, err = engine.Insert(&pm)
		commonPackage.CheckError(err)
		id = pm.Id
	}

	return id, err
}

//获得未读消息
func GetAllUnreadMsg(userId int64) (txtMsg, imgMsg []string) {
	engine, err := GetSocialEng()
	commonPackage.CheckError(err)

	pmts := []social.Lctb_personMessage{}
	err = engine.Where("lc_to_user_info_id = ? and lc_message_status = ?", userId, commonPackage.UnreadMsg).Find(&pmts)
	commonPackage.CheckError(err)
	for _, v := range pmts {
		txtMsg = append(txtMsg, v.Lc_messageContext)
	}

	pmis := []social.Lctb_personMessageImg{}
	err = engine.Where("lc_to_user_info_id = ? and lc_message_status = ?", userId, commonPackage.UnreadMsg).Find(&pmis)
	commonPackage.CheckError(err)
	for _, v := range pmts {
		imgMsg = append(imgMsg, v.Lc_messageContext)
	}

	//设置为已读
	pmt := social.Lctb_personMessage{
		Lc_messageStatus: commonPackage.ReadedMsg,
	}
	engine.Where("lc_to_user_info_id = ? and lc_message_status = ?", userId, commonPackage.UnreadMsg).Cols("lc_message_status").Update(&pmt)

	pmi := social.Lctb_personMessageImg{
		Lc_messageStatus: commonPackage.ReadedMsg,
	}
	engine.Where("lc_to_user_info_id = ? and lc_message_status = ?", userId, commonPackage.UnreadMsg).Cols("lc_message_status").Update(&pmi)

	return
}

func GetAccountByPages(userId int64, limit, count int, users *[]account.Lctb_userInfo) error {
	engine, err := GetAccountEng()
	checkError(err)

	return engine.Where("id != ?", userId).Asc("id").Limit(limit, count).Find(users)
}

//create talk group
func CreateGroup(userid int64, groupName string) (groupId int64, err error) {
	engine, err := GetSocialEng()
	checkError(err)

	group := &social.Lctb_talkGroup{}
	group.Lc_userInfoId = userid
	group.Lc_groupName = groupName
	//group.Lc_cloudFileId 暂时不处理
	group.Lc_maxUserCount = MaxGroupNumber
	_, err = engine.Insert(group) //insert data
	if err != nil {
		return 0, err
	}

	//add creater to group member
	gm := &social.Lctb_groupMember{}
	gm.Lc_talkGroupId = group.Id
	gm.Lc_userInfoId = userid
	gm.Lc_userInfoRole = commonPackage.GroupOwer
	_, err = engine.Insert(gm)
	return group.Id, err
}

//add user to talk group
func AddUserToGroup(groupId int64, userIds []int64) error {
	engine, err := GetSocialEng()
	checkError(err)

	group := &social.Lctb_talkGroup{}
	has, err := engine.Id(groupId).Get(group)
	if !has || err != nil {
		return commonPackage.NewErr(commonPackage.ErrSys)
	}

	gms := []social.Lctb_groupMember{}
	for _, v := range userIds {
		gm := social.Lctb_groupMember{}
		gm.Lc_talkGroupId = groupId
		gm.Lc_userInfoId = v
		gm.Lc_userInfoRole = commonPackage.GroupMember
		gms = append(gms, gm)
	}

	_, err = engine.Insert(&gms)
	return err
}

func DelUserFromGroup(groupId int64, userIds []int64) error {
	engine, err := GetSocialEng()
	checkError(err)

	group := &social.Lctb_talkGroup{}
	has, err := engine.Id(groupId).Get(group)
	if !has || err != nil {
		return commonPackage.NewErr(commonPackage.ErrSys)
	}

	for _, v := range userIds {
		gm := social.Lctb_groupMember{}
		gm.Lc_talkGroupId = groupId
		gm.Lc_userInfoId = v
		_, err = engine.Delete(&gm)
	}

	return err
}

func GetUserGroups(userId int64, groups *[]social.Lctb_talkGroup) error {
	engine, err := GetSocialEng()
	checkError(err)

	gms := []social.Lctb_groupMember{}
	err = engine.Cols("lc_talk_group_id").Where("lc_user_info_id = ?", userId).Find(&gms)
	if err != nil {
		return commonPackage.NewErr(commonPackage.ErrSys)
	}
	gids := []int64{}
	for _, v := range gms {
		gids = append(gids, v.Lc_talkGroupId)
	}
	err = engine.In("id", gids).Find(groups)
	return err
}

func GetGroupInfo(groupId int64, group *social.Lctb_talkGroup, userInfos *[]account.Lctb_userInfo) error {
	gEngine, err := GetSocialEng()
	checkError(err)
	uEngine, err := GetAccountEng()
	checkError(err)

	has, err := gEngine.Id(groupId).Get(group)
	if !has {
		return commonPackage.NewErr(commonPackage.ErrSys)
	}

	gms := []social.Lctb_groupMember{}
	err = gEngine.Cols("lc_user_info_id").Where("lc_talk_group_id = ?", groupId).Find(&gms)

	uids := []int64{}
	for _, v := range gms {
		uids = append(uids, v.Lc_userInfoId)
	}

	err = uEngine.In("id", uids).Find(userInfos)
	return err
}

func SendGroupMsg(userId, groupId int64, msg string, msgType int) (msgid int64, userIds []int64, err error) {
	engine, err := GetSocialEng()
	checkError(err)

	var msgId int64
	if msgType == commonPackage.TextMsg {
		gmsg := social.Lctb_groupMessage{
			Lc_talkGoupId:     groupId,
			Lc_userInfoId:     userId,
			Lc_messageContext: msg,
		}
		_, err = engine.Insert(&gmsg)
		msgId = gmsg.Id
	} else if msgType == commonPackage.ImgMsg {
		gmsg := social.Lctb_groupMessageImg{
			Lc_talkGoupId:     groupId,
			Lc_userInfoId:     userId,
			Lc_messageContext: msg,
		}
		_, err = engine.Insert(&gmsg)
		msgId = gmsg.Id
	}

	userIds = []int64{}
	gms := []social.Lctb_groupMember{}
	err = engine.Where("lc_talk_group_id = ?", groupId).Find(&gms)
	for _, v := range gms {
		if v.Lc_userInfoId != userId {
			userIds = append(userIds, v.Lc_userInfoId)
		}
	}

	commonPackage.Println("DDD")
	commonPackage.Printf(msgId)

	return msgId, userIds, err

}

func InsertGroupMsgStatus(groupId, msgId, userId int64, status, msgType int) error {
	engine, err := GetSocialEng()
	checkError(err)

	if msgType == commonPackage.TextMsg {
		gtms := social.Lctb_groupMessageStatus{
			Lc_groupMessageId: msgId,
			Lc_messageStatus:  status,
			Lc_userInfoId:     userId,
			Lc_talkGoupId:     groupId,
		}
		_, err = engine.Insert(&gtms)
	} else if msgType == commonPackage.ImgMsg {
		gims := social.Lctb_groupMessageImgStatus{
			Lc_groupMessageImgId: msgId,
			Lc_messageStatus:     status,
			Lc_userInfoId:        userId,
			Lc_talkGoupId:        groupId,
		}
		_, err = engine.Insert(&gims)
	}

	return err
}

func GetUnreadGroupMsg(groupId, userId int64) (txtMsg, imgMsg []map[string][]byte, err error) {
	engine, err := GetSocialEng()
	commonPackage.CheckError(err)

	txtMsg, err = engine.Query(`select * from lctb_group_message where id in (select lc_group_message_id from lctb_group_message_status where lc_talk_goup_id = ? and lc_user_info_id = ? and lc_message_status = ?)`, groupId, userId, commonPackage.UnreadMsg)
	commonPackage.CheckError(err)
	imgMsg, err = engine.Query(`select * from lctb_group_message where id in (select lc_group_message_id from lctb_group_message_status where lc_talk_goup_id = ? and lc_user_info_id = ? and lc_message_status = ?)`, groupId, userId, commonPackage.UnreadMsg)
	commonPackage.CheckError(err)
	//设置为已读
	pmt := social.Lctb_groupMessageStatus{
		Lc_messageStatus: commonPackage.ReadedMsg,
	}
	engine.Where("lc_talk_goup_id = ? and lc_user_info_id = ? and lc_message_status = ?", groupId, userId, commonPackage.UnreadMsg).Cols("lc_message_status").Update(&pmt)

	pmi := social.Lctb_groupMessageImgStatus{
		Lc_messageStatus: commonPackage.ReadedMsg,
	}
	engine.Where("lc_talk_goup_id = ? and lc_user_info_id = ? and lc_message_status = ?", groupId, userId, commonPackage.UnreadMsg).Cols("lc_message_status").Update(&pmi)

	return
}
