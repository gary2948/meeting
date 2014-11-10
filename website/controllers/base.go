package controllers

import (
	"commonPackage/model/account"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	userinfo *account.Lctb_userInfo
}

func (b *BaseController) Prepare() {
	var err bool
	b.userinfo, err = b.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	if !err {
		b.Redirect("/", 301)
	}
}
