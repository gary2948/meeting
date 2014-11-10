package commonPackage

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/bitly/go-simplejson"
	. "github.com/qiniu/api/conf"
	"github.com/qiniu/api/rs"
)

const TimeFormat = "2006-01-02 15:04:05"

func init() {
	ACCESS_KEY = "spVfjluxAEzJTIDFQ-uNqORCKU-1GpJi8XMpA_ki"
	SECRET_KEY = "OTOSxnutPuolb6ZKpZiumjPlZolxeUuikCYyJlxb"
	AK = ACCESS_KEY
	SK = SECRET_KEY
	BucketName = "xybstone"
	Domain = `xybstone.qiniudn.com`
}

var BucketName string
var AK string
var SK string
var Domain string

//json字段定义
const (
	JSON_USERID        = "USERID"
	JSON_USERIDS       = "USERIDS"
	JSON_TOUSERID      = "TOUSERID"
	JSON_USERNAME      = "USERNAME"
	JSON_USERINFO      = "USERINFO"
	JSON_USERINFOS     = "USERINFOS"
	JSON_PASSWORD      = "PASSWORD"
	JSON_GROUPID       = "GROUPID"
	JSON_GROUPNAME     = "GROUPNAME"
	JSON_GROUPINFO     = "GROUPINFO"
	JSON_CONTENT       = "CONTENT"
	JSON_EMAIL         = "EMAIL"
	JSON_EMAILS        = "EMAILS"
	JSON_OK            = "OK"
	JSON_ERR           = "ERR"
	JSON_TXTMSG        = "TXTMSG"
	JSON_IMGMSG        = "IMGMSG"
	JSON_NAME          = "NAME"
	JSON_ROLE          = "ROLE"
	JSON_MEETINGID     = "MEETINGID"
	JSON_MEETINGDATE   = "MEETINGDATE"
	JSON_MEETINGSTATUS = "MEETINGSTATUS"
	JSON_BEGINTIME     = "BEGINTIME"
	JSON_ENDTIME       = "ENDTIME"
	JSON_TAG           = "TAG"
	JSON_QUICKSTART    = "QUICKSTART"
	JSON_PLANDATE      = "PLANDATE"
	JSON_PLANSPAN      = "PLANSPAN"
	JSON_SCHEMA        = "SCHEMA"
	JSON_TOPIC         = "TOPIC"
	JSON_CODE          = "CODE"
	JSON_URL           = "URL"
	JSON_PARENTID      = "PARENTID" //文件夹ID
	JSON_COUNT         = "COUNT"
	JSON_FOLDERNAME    = "FOLDERNAME"
	JSON_FLODERID      = "FLODERID"
	JSON_FILEID        = "FILEID"
	JSON_FILEIDS       = "FILEIDS" //[]int64
	JSON_FILEMAPID     = "FILEMAPID"
	JSON_LINKFILEMAPID = "LINKFILEMAPID"
	JSON_FILENAME      = "FILENAME"
	JSON_SHAREFILETYPE = "SHAREFILETYPE"
	JSON_SHARENAME     = "SHARENAME"
	JSON_FILEEXT       = "FILEEXT"
	JSON_FILESIZE      = "FILESIZE"
	JSON_SHAREINFOID   = "SHAREINFOID"
	JSON_FOLLOWUID     = "FOLLOWUID"
	JSON_FOLLOWUIDS    = "FOLLOWUIDS"
	JSON_TYPE          = "TYPE"
	JSON_POSTTIME      = "POSTTIME"
	JSON_LIMIT         = "LIMIT"
	JSON_START         = "START"
	JSON_KINDNAME      = "KINDNAME"
	JSON_KINDID        = "KINDID"
)

const (
	//错误信息
	ErrUserInfo  = "101" //账号不存在
	ErrPasswd    = "102" //密码错误
	ErrUserState = "103" //用户状态错误
	ErrUserId    = "104" //用户id不存在
	ErrFileId    = "105" //文件或文件夹id不存在
	ErrFileName  = "106" //无效的文件名
	ErrFileSize  = "107" //文件大小异常
	ErrMeetingId = "108" //无效的会议Id
	ErrTooLong   = "109" //文件过大
	ErrUserEmail = "110" //错误的email
	ErrSys       = "999" //系统错误
	//语言
	zh  = 1  //中文
	en  = 2  //英语
	jp  = 3  //日语
	cx  = 4  //朝鲜语
	fy  = 5  //法语
	xby = 6  //西班牙语
	ty  = 7  //泰语
	alb = 8  //阿拉伯语
	ey  = 9  //俄语
	pty = 10 //葡萄牙语
	//会议状态
	AllMeetings    = -1 //全部会议
	PlanMeeting    = 0  //计划会议
	StartedMeeting = 1  //进行中
	FinishMeeting  = 2  //已结束
	DeletedMeeting = 3  //已删除
	InvitedMeeting = 4  //邀请的会议I
	//文件or文件夹
	File   = 0
	Folder = 1
	//文件状态
	CommFile    = 0 //常规文件
	RecycleFile = 1 //回收文件
	DelFile     = 2 //删除文件
	ShareFile   = 3 //分享文件
	//分享类型
	ShareByCode = 0 //提取码分享
	ShareByUser = 1 //分享给用户
	//消息状态
	UnreadMsg = 0 //未读消息
	ReadedMsg = 1 //已读消息
	//朋友列表类型
	Follow = 0 //关注列表
	Fans   = 1 //粉丝列表
	//分享文件类型 0表示单个文件 1表示单个文件夹 2表示多个文件或文件一起分享
	OneFileShare    = 0 //单个文件
	OneFolderShare  = 1 //单个文件夹
	MultiFilesShare = 2 //多个文件或文件一起分享
	//消息类型
	TextMsg = 0 //文本消息
	ImgMsg  = 1 //图片消息
	//talk group role
	GroupOwer    = 0
	GroupMember  = 1
	GroupManager = 2
	//user role
	Guest        = 0 //游客
	OfficialUser = 1 //正式用户
	TrialUsers   = 2 //试用用户

)

func GetLanguageInt(languageStr string) (languageInt int) {

	switch languageStr {
	case "中文":
		languageInt = 1
		break
	case "英语":
		languageInt = 2 //
		break
	case "日语":
		languageInt = 3 //
		break
	case "朝鲜语":
		languageInt = 4
		break
	case "法语":
		languageInt = 5 //
		break
	case "西班牙语":
		languageInt = 6 //
		break
	case "泰语":
		languageInt = 7 //
		break
	case "阿拉伯语":
		languageInt = 8 //
		break
	case "俄语":
		languageInt = 9 //
		break
	case "葡萄牙语":
		languageInt = 10 //
		break
	}
	return
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func Println(msg string) {
	fmt.Println(msg)
}

func Printf(i interface{}) {
	fmt.Printf("Object:%v \n", i)
}

func Printb(b []byte) {
	fmt.Println(string(b))
}

func GetUnixTime() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func NewErr(err string) error {
	return errors.New(err)
}

func JSONMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	CheckError(err)
	return b
}

func JSONUnmarshal(data []byte, v interface{}) {
	CheckError(json.Unmarshal(data, v))
}

func ByteJoin(a, b []byte) []byte {
	bf := &bytes.Buffer{}
	bf.Write(a)
	bf.Write(b)
	return bf.Bytes()
}

func GetReComm(comm string, v interface{}) []byte {
	return ByteJoin([]byte(comm), JSONMarshal(v))
}

func OpenFile(path string) (has bool, size int64, context []byte) {
	fileInfo, err := os.Stat(path)
	CheckError(err)
	has = true
	size = fileInfo.Size()
	if has {
		fin, err := os.Open(path)
		CheckError(err)
		defer fin.Close()

		buffs := bytes.NewBuffer([]byte{})
		buf := make([]byte, 1024)
		for {
			n, _ := fin.Read(buf)
			if 0 == n {
				break
			} else {
				buffs.Write(buf[:n])
			}
		}
		context = buffs.Bytes()
	}

	return has, size, context
}

func MergeJson(newJson []byte, OldJson *simplejson.Json) ([]byte, error) {
	newModel, err := simplejson.NewJson(newJson)
	CheckError(err)
	mpp := newModel.MustMap()

	for key, v := range mpp {
		OldJson.Set(key, v)
	}

	return OldJson.Encode()
}

func Uptoken(fileKey ...string) string {
	var callbackBody string
	switch len(fileKey) {
	case 0:
		callbackBody = "key=$(key)"
	case 1:
		callbackBody = fmt.Sprintf("key=%s", fileKey)
	default:
	}

	putPolicy := rs.PutPolicy{
		Scope:        BucketName,
		CallbackUrl:  "http://58.246.49.158:7001/callback",
		CallbackBody: callbackBody,
		//ReturnUrl:   returnUrl,
		//ReturnBody: `{"success":true,"name":"sunflowerb.jpg"}`,
		//AsyncOps:    asyncOps,
		//EndUser:     endUser,
		//Expires:     expires,
	}
	putPolicy.Expires = 36000
	return putPolicy.Token(nil)
}

func DownloadUrl(key string) string {
	baseUrl := rs.MakeBaseUrl(Domain, key)
	policy := rs.GetPolicy{}
	return policy.MakeRequest(baseUrl, nil)
}

type CB_MeetingParams struct {
	Email     string
	MeetingId int64
	FileGuid  string
	FileName  string
}

func Int64Array(j *simplejson.Json) ([]int64, error) {
	arr, err := j.Array()
	if err != nil {
		return nil, err
	}
	retArr := make([]int64, 0, len(arr))
	for _, a := range arr {
		if a == nil {
			retArr = append(retArr, 0)
			continue
		}
		i, err := a.(json.Number).Int64()
		if err != nil {
			return nil, err
		}
		retArr = append(retArr, i)
	}
	return retArr, nil
}
