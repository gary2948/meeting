package controllers

import (
	"commonPackage/viewModel"
	"github.com/astaxie/beego"
	"service/db"
	"strconv"
)

type ShareFileViewController struct {
	beego.Controller
}

func (this *ShareFileViewController) Post() {

}

func (this *ShareFileViewController) Get() {
	beego.Info(this.Ctx.Input.Param(":id"))
	Id, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	beego.Info(Id)
	this.Data["id"] = this.Ctx.Input.Param(":id")
	var shareRed viewModel.ShareViewModel
	db.GetShareInfo(Id, &shareRed)
	this.Data["ShareInfo"] = shareRed
	this.Data["ShareTime"] = shareRed.ShareTime.Format("2006-01-02")
	//this.Data["ShareInfo"]
	this.TplNames = "pages/shareFileView.html"
}
