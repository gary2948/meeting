package controllers

import (
	"commonPackage/model/account"
	"github.com/astaxie/beego"
	"service/db"
)

type RegistController struct {
	beego.Controller
}

func (r *RegistController) Get() {
	_, err := r.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	//beego.Info(userInfo)
	if err {
		r.Redirect("/home", 301)
	} else {
		r.TplNames = "pages/register.html"
	}
}

func (r *RegistController) Post() {
	var user account.Lctb_userInfo

	user.Lc_email = r.GetString("email")
	user.Lc_nickName = r.GetString("nickname")
	user.Lc_realName = r.GetString("realname")
	user.Lc_mobilePhone = r.GetString("mobilephone")
	user.Lc_passwd = r.GetString("password")
	user.Lc_introduction = r.GetString("introduction")
	user.Lc_sex = r.GetString("sex")

	beego.Info(r.GetString("email"))
	beego.Info(r.GetString("nickname"))
	beego.Info(r.GetString("realname"))
	beego.Info(r.GetString("mobilephone"))
	beego.Info(r.GetString("password"))
	beego.Info(r.GetString("introduction"))
	beego.Info(r.GetString("sex"))

	userid, err := db.AddAccount(&user)
	beego.Info(userid)
	if err != nil {
		r.Data["err"] = "注册失败"
		r.TplNames = "pages/register.html"
		//r.Redirect("http://www.meetinware.com", 301)
	} else {
		r.Redirect("/registersuccess", 301)
		//r.TplNames = "pages/registsucced.html"
	}
}

func (r *RegistController) ResigterSuccess() {
	r.TplNames = "pages/registersuccess.html"
}
