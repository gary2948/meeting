package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	BaseController
}

func (this *LogoutController) Post() {
	beego.Info("logout...................................................")
	this.DelSession("Lctb_userInfo")
	this.Redirect("/", 301)
}

func (this *LogoutController) Get() {
	this.Post()
}
