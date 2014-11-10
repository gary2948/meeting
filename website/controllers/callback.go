package controllers

import (
	"github.com/astaxie/beego"
)

type CallBackController struct {
	beego.Controller
}

func (this *CallBackController) Post() {
	//this.Ctx.Output.Context.WriteString("this is a test example")

}

func (this *CallBackController) Get() {

}
