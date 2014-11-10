package db

import (
	"commonPackage"
	"commonPackage/model/account"
	"commonPackage/model/clouddisk"
	"commonPackage/model/meeting"
	"commonPackage/viewModel"
	"fmt"
	"github.com/go-xorm/xorm"
	"service/uuid"
	"strconv"
	"time"
)

var am *meeting.Lctb_attendMeeting = &meeting.Lctb_attendMeeting{}
var ip *meeting.Lctb_invitedPersons = &meeting.Lctb_invitedPersons{}
var mi *meeting.Lctb_meetingInfo = &meeting.Lctb_meetingInfo{}
var oh *meeting.Lctb_operateHistory = &meeting.Lctb_operateHistory{}

//func GetMeetingActivity(userId int) string error {

//}

func InitMeetingTables() {
	fmt.Println("-----InitMeetingTables -----")
	engine, err := postgresEngine(meetingDB)
	checkError(err)
	defer engine.Close()

	err = engine.Sync(am)
	checkError(err)

	err = engine.Sync(ip)
	checkError(err)

	err = engine.Sync(mi)
	checkError(err)

	err = engine.Sync(oh)
	checkError(err)

}

//邀请参与人员
func AddMeetingAndInvPer(meetingInfo *meeting.Lctb_meetingInfo, emails []string) (meetingCode, url string, err error) {
	fmt.Println(emails)
	_, code, err := AddMeeting(meetingInfo)
	checkError(err)
	if len(emails) > 0 {
		err = AddInvitedPersonEmails(meetingInfo, emails)
	}

	checkError(err)
	return code, "", err
}

// Add a new meeting
func AddMeeting(meetingInfo *meeting.Lctb_meetingInfo) (meetingId int64, meetingCode string, err error) {
	engine, err := GetMeetingEng()
	cengine, err := GetCloudDiskEng()
	checkError(err)

	//get helper account
	helper := account.Lctb_userInfo{}
	err = GetHelperAccount(&helper)
	commonPackage.CheckError(err)

	//add cloud file folder
	fId, err := AddFloder(helper.Id, 0, meetingInfo.Lc_topic)
	commonPackage.CheckError(err)

	//add shareinfo
	si := &clouddisk.Lctb_sharedInfos{}
	si.Lc_userInfoId = helper.Id
	si.Lc_sharedName = meetingInfo.Lc_topic + "(" + time.Now().Format(commonPackage.TimeFormat) + ")"
	si.Lc_sharedType = commonPackage.ShareByUser
	_, err = cengine.Insert(si)

	//add meeting
	meetingInfo.Lc_code, err = uuid.GenUUID()
	meetingInfo.Lc_createTime = time.Now()
	meetingInfo.Lc_shareInfoId = si.Id
	meetingInfo.Lc_userName = helper.Lc_nickName
	meetingInfo.Lc_cloudFilesId = fId
	_, err = engine.Insert(meetingInfo)

	return meetingInfo.Id, meetingInfo.Lc_code, err
}

func GetMeetingById(meetingId int64, meetingInfo *meeting.Lctb_meetingInfo) (bool, error) {
	engine, err := GetMeetingEng()
	checkError(err)

	has, err := engine.Id(meetingId).Get(meetingInfo)

	return has, err
}

func GetUserMeetingByDate(userId int64, date string, meetings *[]meeting.Lctb_meetingInfo) error {
	engine, err := GetMeetingEng()
	checkError(err)

	meetingIds := getMeetingsAboutMe(userId, engine)
	if len(meetingIds) > 0 {
		///
		err = engine.In("id", meetingIds).Where("to_char(lc_plan_time, 'YYYY-MM-DD') = ? and lc_delete = 0", date).Or("lc_user_info_id = ? and to_char(lc_plan_time, 'YYYY-MM-DD') = ? ", userId, date).Find(meetings)
	}

	return err
}

func GetMeetingsByPlanStatus(userId int64, status, pageIndex, pageCount int, meetings *[]meeting.Lctb_meetingInfo) (count int64, err error) {
	engine, err := GetMeetingEng()
	checkError(err)

	if status == commonPackage.AllMeetings {
		err = engine.Where("lc_user_info_id = ? and lc_delete = 0 ", userId).Limit(pageCount, (pageIndex-1)*pageCount).Find(meetings)
		count, err = engine.Where("lc_user_info_id = ?  and lc_delete = 0", userId).Count(new(meeting.Lctb_meetingInfo))
	} else {
		err = engine.Where("lc_user_info_id = ? and lc_status =? and lc_delete = 0", userId, status).Limit(pageCount, (pageIndex-1)*pageCount).Find(meetings)
		count, err = engine.Where("lc_user_info_id = ? and lc_status =? and lc_delete = 0", userId, status).Count(new(meeting.Lctb_meetingInfo))
	}

	checkError(err)
	return count, err
}

func getMeetingsAboutMe(userId int64, engine *xorm.Engine) (meetingIds []int64) {
	user := &account.Lctb_userInfo{}
	has, err := GetAccountById(userId, user)

	if !has {
		err = commonPackage.NewErr(commonPackage.ErrUserId)
	}

	if err != nil {
		err = commonPackage.NewErr(commonPackage.ErrSys)
	}

	commonPackage.Println(user.Lc_email)
	////别人邀请的会议
	invitedPersons := make([]meeting.Lctb_invitedPersons, 0)
	err = engine.Where("lc_invited_email = ? ", user.Lc_email).Find(&invitedPersons)
	if err != nil {
		err = commonPackage.NewErr(commonPackage.ErrSys)
	} else {
		for _, v := range invitedPersons {
			meetingIds = append(meetingIds, v.Lc_meetingInfoId)
		}
	}
	//我创建的会议
	//
	meetingInfos := []meeting.Lctb_meetingInfo{}
	err = engine.Where("lc_user_info_id = ?", userId).Find(&meetingInfos)
	if err != nil {
		err = commonPackage.NewErr(commonPackage.ErrSys)
	} else {
		for _, v := range meetingInfos {
			meetingIds = append(meetingIds, v.Id)
		}
	}

	return meetingIds
}

//func GetMeetingsCountByPlanStatus(userId int64, status int) (count int64, err error) {
//	engine, err := mysqlEngine(lcdb_meeting)
//	checkError(err)
//	defer engine.Close()
//	if status == commonPackage.AllMeetings {
//		count, err = engine.Where("lc_user_info_id = ? ", userId).Count(new(meeting.Lctb_meetingInfo))
//	} else {
//		count, err = engine.Where("lc_user_info_id = ? and lc_status =?", userId, status).Count(new(meeting.Lctb_meetingInfo))
//	}

//	checkError(err)
//	return count, err
//}

func GetMeetingsByPlanTime(userId int64, beginTime, endTime time.Time, meetings *[]meeting.Lctb_meetingInfo) error {
	engine, err := GetMeetingEng()
	checkError(err)

	err = engine.Where("lc_user_info_id = ? and Lc_plan_time >= ? and Lc_plan_time <=? and lc_delete = 0", userId, beginTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05")).Find(meetings)
	checkError(err)
	return err
}

//得到会议简要信息
func GetSimpleMeetingsByDate(userId int64, beginTime, endTime string) (resultsSlice []map[string][]byte, err error) {
	engine, err := GetMeetingEng()
	checkError(err)

	meetingIds := getMeetingsAboutMe(userId, engine)
	sql := `SELECT to_char(lc_plan_time, 'YYYY-MM-DD') as MEETINGDATE,  array_to_string(array_agg(DISTINCT lc_status), ',') AS MEETINGSTATUS FROM lctb_meeting_info where to_char(lc_plan_time, 'YYYY-MM-DD') >= '` + beginTime + `' and to_char(lc_plan_time, 'YYYY-MM-DD') <= '` + endTime + `' and lc_user_info_id = ` + strconv.FormatInt(userId, 10) + ` and lc_delete = 0 GROUP BY to_char(lc_plan_time, 'YYYY-MM-DD')`
	return engine.In("id", meetingIds).Query(sql)
}

//邀请与会人员多人
func AddInvitedPersonEmails(meetingInfo *meeting.Lctb_meetingInfo, invitedEmails []string) error {
	engine, err := GetMeetingEng()
	cengine, err := GetCloudDiskEng()
	uengine, err := GetAccountEng()
	checkError(err)

	helper := account.Lctb_userInfo{}
	err = GetHelperAccount(&helper)
	commonPackage.CheckError(err)

	//get userids
	users := []account.Lctb_userInfo{}
	err = uengine.In("lc_email", invitedEmails).Cols("id").Find(&users)
	commonPackage.CheckError(err)
	n := len(users)
	userIds := make([]int64, len(users))
	if n > 0 {
		for i, v := range users {
			userIds[i] = v.Id
		}
	}

	//add to share list
	n = len(userIds)
	if n > 0 {
		sus := make([]clouddisk.Lctb_sharedUsers, n)
		for i, v := range userIds {
			su := clouddisk.Lctb_sharedUsers{}
			su.Lc_sharedInfoId = meetingInfo.Lc_shareInfoId
			su.Lc_userInfoId = helper.Id
			su.Lc_toUserInfoId = v
			sus[i] = su
		}
		_, err = cengine.Insert(&sus)
	}

	//Add invite person
	n = len(invitedEmails)
	if n > 0 {
		personEmails := make([]meeting.Lctb_invitedPersons, n)
		for i, v := range invitedEmails {
			personEmails[i] = meeting.Lctb_invitedPersons{Lc_meetingInfoId: meetingInfo.Id, Lc_invitedEmail: v}
		}
		_, err = engine.Insert(personEmails)
	}

	return err
}

//邀请与会人员
func AddInvitedPersonEmail(meetingId int64, invitedEmail string) (int64, error) {
	engine, err := GetMeetingEng()
	checkError(err)

	//InvitedPersonEmail := &meeting.Lctb_invitedPersons{Lc_meetingInfoId: meetingId, Lc_invitedEmail: invitedEmail}
	personEmail := new(meeting.Lctb_invitedPersons)
	personEmail.Lc_meetingInfoId = meetingId
	personEmail.Lc_invitedEmail = invitedEmail
	return engine.Insert(personEmail)
}

//添加实际与会人员
func AddAttendedPersonEmail(meetingId int64, invitedEmail string, userType int) (int64, error) {
	engine, err := GetMeetingEng()
	checkError(err)

	//InvitedPersonEmail := &meeting.Lctb_invitedPersons{Lc_meetingInfoId: meetingId, Lc_invitedEmail: invitedEmail}
	attend := new(meeting.Lctb_attendMeeting)
	attend.Lc_meetingInfoId = meetingId
	attend.Lc_joinedEmail = invitedEmail
	attend.Lc_userType = userType
	return engine.Insert(attend)
}

//得到邀请人员
func GetInvitedPersonEmails(meetingId int64, invitedEmails *[]string) (int, error) {
	engine, err := GetMeetingEng()
	checkError(err)

	emails := make([]meeting.Lctb_invitedPersons, 0)
	err = engine.Where("lc_meeting_info_id = ?", meetingId).Cols("lc_invited_email").Find(&emails)
	checkError(err)

	leng := len(emails)
	for i := 0; i < leng; i++ {
		*invitedEmails = append(*invitedEmails, emails[i].Lc_invitedEmail)
	}

	return leng, err
}

//得到实际与会人员
func GetAttendedPerson(meetingId int64, attendedPerson *[]meeting.Lctb_attendMeeting) (int, error) {
	engine, err := GetMeetingEng()
	checkError(err)

	err = engine.Where("lc_meeting_info_id = ?", meetingId).Find(attendedPerson)
	checkError(err)

	leng := len(*attendedPerson)

	return leng, err
}

//删除会议，会议相关文件暂时保留
func DelMeeting(meetingId int64) error {
	engine, err := GetMeetingEng()
	commonPackage.CheckError(err)

	m := meeting.Lctb_meetingInfo{Lc_delete: commonPackage.DeletedMeeting}
	_, err = engine.Id(meetingId).Cols("lc_delete").Update(m)
	return err
}

//添加会议文件
func AddMeetingFiles(meetingId int64, fileIds []int64) error {
	engine, err := GetMeetingEng()
	commonPackage.CheckError(err)
	cengine, err := GetCloudDiskEng()
	commonPackage.CheckError(err)

	mi := meeting.Lctb_meetingInfo{}
	has, err := engine.Id(meetingId).Get(mi)
	commonPackage.CheckError(err)
	if !has {
		err = commonPackage.NewErr(commonPackage.ErrMeetingId)
	} else {
		sfs := make([]clouddisk.Lctb_sharedFiles, len(fileIds))
		for i, v := range fileIds {
			sf := clouddisk.Lctb_sharedFiles{}
			sf.Lc_cloudFilesId = v
			sf.Lc_sharedInfoId = mi.Lc_shareInfoId
			sfs[i] = sf
		}
		_, err = cengine.Insert(&sfs)
	}
	return err
}

//get a meeting code,
//return a meeting code ,if error is nil
//func GetMeetingCode() (string, error) {

//}

//
//func ReserveMeeting(mCode string, userId int, rTime time.Time) {

//}

func GetDynamicState(userId int64, dynamicState *[]viewModel.DynamicState) error {
	engine, err := GetMeetingEng()
	uengine, err := GetAccountEng()
	mids := getMeetingsAboutMe(userId, engine)
	meeting := meeting.Lctb_meetingInfo{}
	//get dynamicState
	engine.In("id", mids).Asc("lc_modify_time").Limit(10).Table(meeting).Find(dynamicState)
	//get uids
	n := len(*dynamicState)
	if n > 0 {
		for i, v := range *dynamicState {
			//get user
			user := account.Lctb_userInfo{}
			has, _ := uengine.Id(v.UserID).Cols("lc_photo_file", "lc_nick_name").Get(&user)
			if has {
				(*dynamicState)[i].IconAddr = user.Lc_photoFile
			}
			//创建人不是自己，切没有其他操作，就是被邀请的会议
			if v.UserID != userId && (*dynamicState)[i].MeetingAction == commonPackage.PlanMeeting {
				(*dynamicState)[i].MeetingAction = commonPackage.InvitedMeeting
			}
			(*dynamicState)[i].NickName = user.Lc_nickName
		}
	}

	return err
}
