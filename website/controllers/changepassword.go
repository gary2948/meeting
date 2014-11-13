package controllers

import (
	"fmt"
	"service/db"
)

type ChangePasswordController struct {
	BaseController
}

func (this *ChangePasswordController) Post() {
	mystruct := ""
	if this.userinfo != nil {
		this.Data["userinfo"] = this.userinfo
		this.Data["doctotal"] = 0
		var Lc_passwd = this.GetString("userpassword")
		//var Old_passwd = this.GetString("oldpassword")
		err := db.ExchangePwd(this.userinfo.Id, Lc_passwd)
		if err != nil {
			mystruct = `修改失败`
		} else {
			mystruct = `修改成功`
		}
		fmt.Println(err)
	} else {
		//用户未登录，返回修改失败
		mystruct = `修改失败`
	}
	this.Data["err"] = mystruct
	fmt.Println(mystruct)
	this.TplNames = "pages/home/changePassword.html"
}

func (h *ChangePasswordController) Get() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo

		h.Data["doctotal"] = 0

		h.TplNames = "pages/home/changePassword.html"
	}
}
