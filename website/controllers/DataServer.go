package controllers

import (
	"commonPackage"
	"commonPackage/model/account"
	"commonPackage/model/clouddisk"
	"commonPackage/model/meeting"
	"commonPackage/viewModel"
	"github.com/astaxie/beego"
	"service/db"
	"strconv"
	"time"
	"website/models"
)

type DataServerController struct {
	beego.Controller
}

func (this *DataServerController) Get() {
	this.Post()
}

func (this *DataServerController) Post() {
	actionType := this.GetString("ActionType")
	beego.Info(actionType)
	var result interface{}
	switch actionType {
	case "Test":
		beego.Info("Test response")
		result = "this is a test action"
	case "CreateMeeting": //创建会议
		result = CreateMeeting(this)
	case "GetFileList": //获取我的文档数据
		result = GetFileList(this)
	case "GetShareList": //获取我的分享数据集合
		result = GetShareFileList(this)
	case "GetRecycledList": //获取回收站中的数据集合
		result = GetRecycledFileList(this)
	case "SearchFile": //查找文件
		result = SearchFile(this)
	case "RenameFile": //重命名文件或文件夹
		result = RenameFile(this)
	case "AddFolder": //添加文件夹
		result = AddFolder(this)
	case "DeleteFile": //删除文件
		result = DeleteFile(this)
	case "ShareFile": //分享文件
		result = ShareFile(this)
	case "TimeSpanMeetingList": //通过时间段查会议记录（日历中用到）
		result = TimeSpanMeetingList(this)
	case "StartMeetingList": //正在进行中的会议（和我有关的）
		result = StartMeetingList(this)
	case "UnStartMeetingList": //计划中的会议（和我有关的）
		result = UnStartMeetingList(this)
	case "EndMeetingList": //已经结束的会议
		result = EndMeetingList(this)
	case "AllMeetingList": //所有的会议记录
		result = AllMeetingList(this)
	case "DeleteMeeting": //删除会议
		result = DeleteMeeting(this)
	case "CancelShareRecord": //取消分享
		result = CancelShareRecord(this)
	case "ShiftDeleteFile": //彻底删除文件
		result = ShiftDeleteFile(this)
	case "ReBackFile": //从回收站中还原文件
		result = ReBackFile(this)
	case "ClearRecycleBin": //清空回收站
		result = ClearRecycleBin(this)
	case "EditUserInfo": //编辑用户信息
		result = EditUserInfo(this)
	case "DynamicInfo": //获取动态信息（首页）
		result = DynamicInfo(this)
	case "GetSharedFile": //获取共享的文件
		result = GetSharedFile(this)
	case "EditMeetingInfo":
		result = EditMeetingInfo(this)
	}
	this.Data["json"] = result
	this.ServeJson()

}

func CreateMeeting(this *DataServerController) models.ResultModel {
	var result = models.ResultModel{}
	result.RequestResult = true
	meeting := meeting.Lctb_meetingInfo{}
	value, _ := this.GetInt("BeginTime")
	meeting.Lc_planTime = time.Unix(value/1000, 0)
	value, _ = this.GetInt("EndTime")
	meeting.Lc_planTimespan = time.Unix(value/1000, 0)

	meeting.Lc_topic = this.GetString("Topic")
	meeting.Lc_schema = this.GetString("Schema")
	meeting.Lc_passwd = this.GetString("PassWord")

	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)

	meeting.Lc_userInfoId = userinfo.Id

	beego.Info(meeting)
	db.AddMeeting(&meeting)
	result.RequestResult = true
	result.ErrorMsg = ""
	result.Data = nil
	return result
}

func GetFileList(this *DataServerController) models.ResultModel {
	var result = models.ResultModel{}
	result.RequestResult = true
	FileID, err := this.GetInt("FileID")
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
	}
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	var cfiles []clouddisk.Lctb_cloudFiles
	err = db.GetUserFolderViewFiles(userinfo.Id, FileID, &cfiles)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
		return result
	}
	result.Data = cfiles
	return result
}

func GetShareFileList(this *DataServerController) models.ResultModel {

	var result = models.ResultModel{}
	result.RequestResult = true
	var cfiles []viewModel.ShareViewModel
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	db.GetShareCloudFiles(userinfo.Id, &cfiles)
	result.Data = cfiles
	return result
}
func GetRecycledFileList(this *DataServerController) models.ResultModel {
	var result = models.ResultModel{}
	result.RequestResult = true
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	var cfiles []clouddisk.Lctb_cloudFiles
	err := db.GetInRecycelCloudFile(int64(userinfo.Id), &cfiles)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
		return result
	}
	result.Data = cfiles
	return result
}

func SearchFile(this *DataServerController) models.ResultModel {
	var result = models.ResultModel{}
	result.RequestResult = true
	searchKey := this.GetString("SearchKey")
	beego.Info(searchKey)
	return result
}

func RenameFile(this *DataServerController) models.ResultModel {
	var result = models.ResultModel{}
	result.RequestResult = true
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	RecordID, _ := this.GetInt("RecordID")
	Rename := this.GetString("Rename")
	ParentID, _ := this.GetInt("ParentID")
	db.RenameCloudFile(userinfo.Id, ParentID, RecordID, Rename)
	beego.Info(RecordID)
	beego.Info(Rename)

	var cfiles []clouddisk.Lctb_cloudFiles
	err := db.GetUserFolderFiles(userinfo.Id, ParentID, &cfiles)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
	}
	result.Data = cfiles

	return result
}

func AddFolder(this *DataServerController) models.ResultModel {
	var result = models.ResultModel{}
	result.RequestResult = true
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	parentid, err := this.GetInt("ParentID")
	if err != nil {

	}
	folderName := this.GetString("FolderName")
	_, err = db.AddFloder(userinfo.Id, parentid, folderName)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
	}
	var cfiles []clouddisk.Lctb_cloudFiles
	db.GetUserFolderViewFiles(userinfo.Id, parentid, &cfiles)
	result.Data = cfiles
	return result
}

func DeleteFile(this *DataServerController) models.ResultModel {
	var result = models.ResultModel{}
	result.RequestResult = true
	FileIdList, err := ConvertStringArr2IntArr(this.GetStrings("FileIdList[]"))
	beego.Info(FileIdList)
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)

	err = db.InRecycelCloudFile(userinfo.Id, FileIdList)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
		return result
	}
	ParentID, err := this.GetInt("ParentID")
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
		return result
	}

	var cfiles []clouddisk.Lctb_cloudFiles
	err = db.GetUserFolderViewFiles(userinfo.Id, ParentID, &cfiles)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
		return result
	}
	result.RequestResult = true
	result.Data = cfiles
	beego.Info(cfiles)
	return result
}

func ShareFile(this *DataServerController) models.ResultModel {
	var result = models.ResultModel{}
	result.RequestResult = true
	FileIdList, err := ConvertStringArr2IntArr(this.GetStrings("FileIdList[]"))
	beego.Info(FileIdList)
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	shareName := this.GetString("shareName")
	fileExt := this.GetString("fileExt")
	fileSize, _ := this.GetInt("fileSize")
	fileType, _ := this.GetInt("fileType")
	//userId int64, fileIds []int64, shareName, fileExt string, fileSize int64
	sharecode, err := db.ShareCloudFileByCode(userinfo.Id, int(fileType), FileIdList, shareName, fileExt, fileSize)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
		return result
	}
	result.RequestResult = true
	result.Data = sharecode
	return result
}

func ConvertStringArr2IntArr(value []string) (result []int64, err error) {
	err = nil
	result = make([]int64, len(value))
	for i := 0; i < len(value); i++ {
		result[i], err = strconv.ParseInt(value[i], 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return
}

func AddFile(this *DataServerController) models.ResultModel {
	var result = models.ResultModel{}
	result.RequestResult = true
	return result
}

func TimeSpanMeetingList(this *DataServerController) models.ResultModel {
	beego.Info(this.GetString("StartTime"))
	start, _ := this.GetInt("StartTime")
	end, _ := this.GetInt("EndTime")
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)

	starttime := time.Unix(start, 0)
	endtime := time.Unix(end, 0)
	var meetings []meeting.Lctb_meetingInfo
	db.GetMeetingsByPlanTime(userinfo.Id, starttime, endtime, &meetings)

	var result = models.ResultModel{}
	result.RequestResult = true
	result.Data = meetings
	return result
}

func StartMeetingList(this *DataServerController) models.ResultModel {
	pagenum, _ := this.GetInt("CurrentPage")
	rowsperpage, _ := this.GetInt("RowPerPage")
	var result = models.ResultModel{}
	result.RequestResult = true
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	var meetings []meeting.Lctb_meetingInfo

	count, err := db.GetMeetingsByPlanStatus(userinfo.Id, commonPackage.StartedMeeting, int(pagenum), int(rowsperpage), &meetings)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
		return result
	}
	var rows models.DataGridMeetingResult
	result.RequestResult = true
	rows.TotalRow = count
	rows.Rows = meetings
	result.Data = rows
	return result
}

func UnStartMeetingList(this *DataServerController) models.ResultModel {
	pagenum, _ := this.GetInt("CurrentPage")
	rowsperpage, _ := this.GetInt("RowPerPage")
	var result = models.ResultModel{}
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	var meetings []meeting.Lctb_meetingInfo
	count, err := db.GetMeetingsByPlanStatus(userinfo.Id, commonPackage.PlanMeeting, int(pagenum), int(rowsperpage), &meetings)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
		return result
	}
	var rows models.DataGridMeetingResult
	result.RequestResult = true
	rows.TotalRow = count
	rows.Rows = meetings
	result.Data = rows
	return result
}

func EndMeetingList(this *DataServerController) models.ResultModel {
	pagenum, _ := this.GetInt("CurrentPage")
	rowsperpage, _ := this.GetInt("RowPerPage")
	var result = models.ResultModel{}
	result.RequestResult = true
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	var meetings []meeting.Lctb_meetingInfo
	count, err := db.GetMeetingsByPlanStatus(userinfo.Id, commonPackage.FinishMeeting, int(pagenum), int(rowsperpage), &meetings)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
		return result
	}
	var rows models.DataGridMeetingResult
	result.RequestResult = true
	rows.TotalRow = count
	rows.Rows = meetings
	result.Data = rows
	return result
}

func AllMeetingList(this *DataServerController) models.ResultModel {
	pagenum, _ := this.GetInt("CurrentPage")
	rowsperpage, _ := this.GetInt("RowPerPage")
	beego.Info(pagenum)
	beego.Info(rowsperpage)
	var result = models.ResultModel{}
	result.RequestResult = true
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	var meetings []meeting.Lctb_meetingInfo
	count, err := db.GetMeetingsByPlanStatus(userinfo.Id, commonPackage.AllMeetings, int(pagenum), int(rowsperpage), &meetings)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
		return result
	}
	beego.Info(count)
	var rows models.DataGridMeetingResult
	result.RequestResult = true
	rows.TotalRow = count
	rows.Rows = meetings
	result.Data = rows
	return result
}

func DeleteMeeting(this *DataServerController) models.ResultModel {
	var result = models.ResultModel{}
	result.RequestResult = true
	meetingid, err := this.GetInt("meetingID")
	beego.Info(meetingid)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
		return result
	}
	err = db.DelMeeting(meetingid)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
		return result
	}
	result.RequestResult = true
	return result
}

func CancelShareRecord(this *DataServerController) models.ResultModel {
	beego.Info("CancelShareRecord")
	var result = models.ResultModel{}
	result.RequestResult = true
	FileIdList, _ := ConvertStringArr2IntArr(this.GetStrings("FileIdList[]"))
	beego.Info(FileIdList)
	var err error
	for i := 0; i < len(FileIdList); i++ {
		err = db.UnshareCloudFile(FileIdList[i])
		beego.Info(err)
	}
	var cfiles []viewModel.ShareViewModel
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	db.GetShareCloudFiles(userinfo.Id, &cfiles)
	result.Data = cfiles
	return result
}

func ShiftDeleteFile(this *DataServerController) models.ResultModel {
	beego.Info("ShiftDeleteFile")
	var result = models.ResultModel{}
	result.RequestResult = true
	FileIdList, _ := ConvertStringArr2IntArr(this.GetStrings("FileIdList[]"))
	err := db.RecycelCloudFile(FileIdList)
	if err != nil {
		result.ErrorMsg = err.Error()
		result.RequestResult = false
		return result
	}
	return result
}

func ReBackFile(this *DataServerController) models.ResultModel {
	var result = models.ResultModel{}
	result.RequestResult = true
	FileIdList, _ := ConvertStringArr2IntArr(this.GetStrings("FileIdList[]"))
	err := db.RecycelCloudFile(FileIdList)
	if err != nil {
		result.ErrorMsg = err.Error()
		result.RequestResult = false
		return result
	}
	return result
}

func ClearRecycleBin(this *DataServerController) models.ResultModel {
	var result = models.ResultModel{}
	result.RequestResult = true
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	err := db.ClearAllRecycle(userinfo.Id)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
	}
	result.RequestResult = true
	return result
}

func EditUserInfo(this *DataServerController) models.ResultModel {
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	var result = models.ResultModel{}
	result.RequestResult = true
	beego.Info(this.GetString("list[0][name]"))
	result.RequestResult = true
	edituserinfo := viewModel.EditUserInfoModel{}
	edituserinfoEx := viewModel.EditUserInfoExModel{}

	edituserinfo.Lc_email = this.GetString("email")
	edituserinfo.Lc_introduction = this.GetString("introduction")
	edituserinfo.Lc_mobilePhone = this.GetString("mobilePhone")
	edituserinfo.Lc_nickName = this.GetString("nickName")
	edituserinfo.Lc_postAddress = this.GetString("postAddress")
	edituserinfo.Lc_realName = this.GetString("realName")
	edituserinfo.Lc_sex = this.GetString("sex")

	timevalue, _ := this.GetInt("birthday")
	edituserinfoEx.Lc_birthday = time.Unix(timevalue/1000, 0)
	edituserinfoEx.Lc_email1 = this.GetString("email1")
	edituserinfoEx.Lc_email2 = this.GetString("email2")
	edituserinfoEx.Lc_hometown = this.GetString("hometown")
	edituserinfoEx.Lc_identityCard = this.GetString("identityCard")
	edituserinfoEx.Lc_language = this.GetString("language")
	edituserinfoEx.Lc_language1 = this.GetString("language2")
	edituserinfoEx.Lc_language2 = this.GetString("")
	edituserinfoEx.Lc_mobilePhone1 = this.GetString("mobilePhone1")
	edituserinfoEx.Lc_mobilePhone2 = this.GetString("mobilePhone2")
	edituserinfoEx.Lc_qq = this.GetString("qq")
	edituserinfoEx.Lc_telephone = this.GetString("telephone")
	edituserinfoEx.Lc_telephone1 = this.GetString("telephone1")
	edituserinfoEx.Lc_telephone2 = this.GetString("telephone2")
	edituserinfoEx.Lc_weibo = this.GetString("weibo")
	edituserinfoEx.Lc_weixin = this.GetString("weixin")

	beego.Info(edituserinfo)
	beego.Info(edituserinfoEx)
	err := db.UpdateAccount(userinfo.Id, &edituserinfo)
	beego.Info(err)
	err = db.UpdateAccountEx(userinfo.Id, &edituserinfoEx)
	beego.Info(err)

	return result

}

func DynamicInfo(this *DataServerController) models.ResultModel {
	var result models.ResultModel
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	var info []viewModel.DynamicState
	err := db.GetDynamicState(userinfo.Id, &info)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
		return result
	} else {
	}
	result.RequestResult = true
	result.Data = info
	return result
}

func GetSharedFile(this *DataServerController) models.ResultModel {
	var result models.ResultModel
	result.RequestResult = true
	id, err := this.GetInt("Id")
	if err != nil {
		result.RequestResult = true
		result.ErrorMsg = err.Error()
		return result
	}
	var files []clouddisk.Lctb_cloudFiles
	err = db.GetShareFiles(id, &files)
	if err != nil {
		result.RequestResult = false
		result.ErrorMsg = err.Error()
		return result
	}
	result.RequestResult = true
	result.Data = files
	return result
}

func EditMeetingInfo(this *DataServerController) models.ResultModel {
	var result models.ResultModel
	result.RequestResult = true
	return result
}
