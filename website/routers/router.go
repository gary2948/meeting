package routers

import (
	"commonPackage/model/account"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	//"net/http"
	"website/controllers"
	"website/models"
)

//在此处进行所有的路由注册工作
func init() {
	//beego.InsertFilter("", beego.BeforeExec, BeforeRouter)
	//registStaticResource()
	registRoter()
}

func BeforeRouter(ctx *context.Context) {

	beego.Info("请求的url地址" + ctx.Request.URL.Path)
	if ctx.Request.URL.Path == "/" {
		return
	} else if ctx.Request.URL.Path == "/regist" {
		return
	}

	///供第三方调用  不需要进行登录
	if ctx.Request.URL.Path == "/callback" {
		///在此处添加处理逻辑
		///ctx.Output.Json(result, true, false)
		beego.Info("call back")
		ctx.Output.Json("this is a test", true, false)
		return
	}

	beego.Info("用户信息验证")
	user, err := ctx.Input.CruSession.Get("Lctb_userInfo").(*account.Lctb_userInfo)
	if !err {
		beego.Info("用户session丢失" + ctx.Request.URL.Path)
		if ctx.Input.Header("X-Requested-With") == "XMLHttpRequest" {
			result := models.ResultModel{}
			result.RequestResult = false
			result.ErrorMsg = ""
			result.Data = models.Result_SessionLose
			ctx.Output.Json(result, true, false)
		} else {
			ctx.Redirect(301, "/")
		}
	} else {
		ctx.Input.Data["nickName"] = user.Lc_nickName
	}
}

//注册所有的静态资源
func registStaticResource() {
	//beego.SetStaticPath("/", "")
	//beego.SetStaticPath("/pages", "views/pages")
	//beego.SetStaticPath("/favicon.ico", "static/favicon.ico")
}

func registRoter() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/meeting", &controllers.MeetingController{})
	beego.Router("/calendar", &controllers.MeetingCalendarController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/register", &controllers.RegistController{})
	beego.Router("/registersuccess", &controllers.RegistController{}, "*:ResigterSuccess")

	/////////////////////
	//
	//首页
	//
	/////////////////////
	beego.Router("/home", &controllers.HomeController{})
	//个人信息
	beego.Router("/myinfo", &controllers.MyinfoController{})
	//修改头像
	beego.Router("/changehead", &controllers.ChangeheadController{})
	//修改密码
	beego.Router("/changepassword", &controllers.ChangePasswordController{})
	//支付信息
	beego.Router("/payinfo", &controllers.MyinfoController{}, "*:PayInfo")
	//修改联系信息
	beego.Router("/updateexcontactinfo", &controllers.MyinfoController{}, "post:UpdateExContactInfo")
	//修改基本信息
	beego.Router("/updatebaseinfo", &controllers.MyinfoController{}, "post:UpdateBaseInfo")
	//新增个人经历
	beego.Router("/addpersonexp", &controllers.MyinfoController{}, "post:AddPersonexp")
	//修改个人经历
	beego.Router("/updatepersonexp", &controllers.MyinfoController{}, "post:UpdatePersonexp")

	/////////////////////
	//
	//会议
	//
	/////////////////////
	beego.Router("/meetinginfo", &controllers.MeetingController{}, "*:MeetingInfo")
	//新增会议
	beego.Router("/addmeeting", &controllers.MeetingController{}, "*:AddMeeting")
	//会议日志
	beego.Router("/meetinglog", &controllers.MeetingController{}, "*:MeetingLog")
	//日程管理
	beego.Router("/schedulemanager", &controllers.MeetingController{}, "*:ScheduleManager")

	/////////////////////
	//
	//社交
	//
	/////////////////////
	//微博
	beego.Router("/weibo", &controllers.SocialController{}, "*:Weibo")
	//收件箱
	beego.Router("/message", &controllers.SocialController{}, "*:Message")
	//小组
	beego.Router("/group", &controllers.SocialController{}, "*:Group")
	beego.Router("/addgroup", &controllers.SocialController{}, "*:AddGroup")
	//我的名片
	beego.Router("/mycard", &controllers.SocialController{}, "*:Mycard")
	//我的动态
	beego.Router("/mydynamic", &controllers.SocialController{}, "*:Mydynamic")
	//我的联系人
	beego.Router("/mycontact", &controllers.SocialController{}, "*:Mycontact")
	//企业通讯录
	beego.Router("/companycontact", &controllers.SocialController{}, "*:Companycontact")
	//小组通讯录
	beego.Router("/groupcontact", &controllers.SocialController{}, "*:Groupcontact")
	//发私信
	beego.Router("/sendmsg", &controllers.SocialController{}, "post:SendMessages")

	/////////////////////
	//
	//后台
	//
	/////////////////////
	//基本信息
	beego.Router("/backstage", &controllers.BackstageController{}, "*:Backstage")
	//成员账号
	beego.Router("/useraccount", &controllers.BackstageController{}, "*:Useraccount")
	//管理员账号
	beego.Router("/adminaccount", &controllers.BackstageController{}, "*:Adminaccount")
	//团队账号
	beego.Router("/teamaccount", &controllers.BackstageController{}, "*:Teamaccount")
	//成员资料管理
	beego.Router("/memberinfo", &controllers.BackstageController{}, "*:Memberinfo")
	//动态设置
	beego.Router("/dynamicmanager", &controllers.BackstageController{}, "*:Dynamicmanager")
	//IP访问控制
	beego.Router("/ipmanager", &controllers.BackstageController{}, "*:Ipmanager")

	/////////////////////
	//
	//网盘
	//
	/////////////////////
	//我的文档
	beego.Router("/mydoc", &controllers.ClouddiskController{}, "*:MyDoc")

	beego.Router("/clouddisk", &controllers.ClouddiskController{})
	beego.Router("/social", &controllers.SocialController{})
	beego.Router("/regist", &controllers.RegistController{})
	beego.Router("/userinfo", &controllers.UserInfoController{})
	beego.Router("/upload", &controllers.WebSocketController{})
	beego.Router("/dataserver", &controllers.DataServerController{})
	beego.Router("/editmeeting/:id", &controllers.EditMeetingController{})
	beego.Router("/changeicon", &controllers.ChangeIconController{})
	beego.Router("/sharefileview/:id", &controllers.ShareFileViewController{})
	beego.Router("/callback", &controllers.CallBackController{})
}
