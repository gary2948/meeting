package controllers

import (
	"commonPackage/model/account"
	"fmt"
	"service/db"
	"time"

	"github.com/astaxie/beego"
)

type SocialController struct {
	BaseController
}

func (h *SocialController) Weibo() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/weibo.html"
	}
}

func (h *SocialController) Message() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/message.html"
	}
}

func (f *SocialController) SendMessages() {
	mystruct := ""
	if f.userinfo != nil {
		f.Data["userinfo"] = f.userinfo
		var mes = f.GetString("mes")
		//var resUserId = 146
		var postTime = time.Now()
		//var status = 0
		//var msgType = 0
		userid, err := db.SendPersonMessage(f.userinfo.Id, 146, postTime, mes, 0, 0)
		beego.Info(userid)
		if err != nil {
			mystruct = `{result:false}`
		} else {
			mystruct = `{result:true}`
		}
		fmt.Println(err)
	} else {
		//用户未登录，返回发送失败
		mystruct = `{result:false}`
	}
	f.Data["json"] = &mystruct
	f.ServeJson()

}

//通过邮箱搜索用户
func (f *SocialController) SearchAccountByEmail() {
	if f.userinfo != nil {
		var searchuser account.Lctb_userInfo
		var userEmail = f.GetString("email")
		ret, _ := db.GetAccountByEmail(userEmail, &searchuser)
		fmt.Println(ret)
		f.Data["userinfo"] = f.userinfo
		f.Data["ret"] = ret
		if ret {
			f.Data["searchuser"] = searchuser
			fmt.Println(searchuser)
		}
		f.TplNames = "pages/social/message.html"
	}
}

func (h *SocialController) Group() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/group.html"
	}
}
func (h *SocialController) AddGroup() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/addgroup.html"
	}
}
func (h *SocialController) Mycard() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/mycard.html"
	}
}
func (h *SocialController) Mydynamic() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/mydynamic.html"
	}
}
func (h *SocialController) Mycontact() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/mycontact.html"
	}
}

func (h *SocialController) Companycontact() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/companycontact.html"
	}
}
func (h *SocialController) Groupcontact() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/groupcontact.html"
	}
}
