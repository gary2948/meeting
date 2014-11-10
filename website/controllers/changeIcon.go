package controllers

import (
	"github.com/astaxie/beego"
)

type ChangeIconController struct {
	beego.Controller
}

func (this *ChangeIconController) Post() {
	files, fileshead, err := this.GetFile("file")
	beego.Info(files)
	beego.Info(fileshead.Filename)
	beego.Info(err)
	err = this.SaveToFile("file", this.GetString("filename"))
	beego.Info(err)
	this.TplNames = "pages/changeIcon.html"
}

func (this *ChangeIconController) Get() {
	this.TplNames = "pages/changeIcon.html"
}
