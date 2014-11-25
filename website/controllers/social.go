package controllers

import (
	"commonPackage/model/account"
	"commonPackage/model/social"
	"fmt"
	"service/db"
	"strconv"
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
func (h *SocialController) Followlist() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/followlist.html"
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
			f.Data["json"] = searchuser
			f.ServeJson()
			fmt.Println(searchuser)
		} else {
			f.Data["json"] = `{result:false}`
			f.ServeJson()
		}
		//f.TplNames = "pages/social/message.html"
	}
}

//关注用户
func (f *SocialController) FollowtheUsers() {
	mystruct := ""
	if f.userinfo != nil {
		Followid := f.GetString("followid")
		tempid, _ := strconv.Atoi(Followid)
		var exps = []int64{int64(tempid)}
		err := db.FollowUser(f.userinfo.Id, 1, exps)
		fmt.Println(err)
		if err != nil {
			mystruct = `{result:false}`
		} else {
			mystruct = `{result:true}`
		}
	} else {
		//用户未登录，返回失败
		mystruct = `{result:false}`
	}
	f.Data["json"] = &mystruct
	f.ServeJson()
}

//取消关注
func (f *SocialController) UfollowtheUsers() {
	mystruct := ""
	if f.userinfo != nil {
		Followid := f.GetString("followid")
		tempid, _ := strconv.Atoi(Followid)
		err := db.UnFollowUser(f.userinfo.Id, int64(tempid))
		fmt.Println(err)
		if err != nil {
			mystruct = `{result:false}`
		} else {
			mystruct = `{result:true}`
		}
	} else {
		//用户未登录，返回失败
		mystruct = `{result:false}`
	}
	f.Data["json"] = &mystruct
	f.ServeJson()
}

//创建小组
func (f *SocialController) CreatnewGroup() {
	mystruct := ""
	if f.userinfo != nil {
		groupName := f.GetString("groupName")
		group_id, err := db.CreateGroup(f.userinfo.Id, groupName)
		fmt.Println(group_id)
		if err != nil {
			mystruct = `创建失败`
		} else {
			mystruct = `创建成功`
		}
		var userGroup = make([]social.Lctb_talkGroup, 0)
		_ = db.GetUserGroups(f.userinfo.Id, &userGroup)
		f.Data["userGroup"] = userGroup
		f.Data["mystruct"] = mystruct
		f.Data["userinfo"] = f.userinfo
		f.TplNames = "pages/success.html"
	} else {
		//用户未登录
		mystruct = `创建失败`
		f.Data["mystruct"] = mystruct
		f.Data["userinfo"] = f.userinfo
		f.TplNames = "login.html"
	}
}

func (h *SocialController) Group() {
	if h.userinfo != nil {
		var userGroup = make([]social.Lctb_talkGroup, 0)

		_ = db.GetUserGroups(h.userinfo.Id, &userGroup)
		fmt.Println(userGroup)
		h.Data["userinfo"] = h.userinfo
		h.Data["userGroup"] = userGroup
		//fmt.Println(userExpe)
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
