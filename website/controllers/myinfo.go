package controllers

import (
	"commonPackage/model/account"
	"service/db"
)

type MyinfoController struct {
	BaseController
}

func (h *MyinfoController) Get() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo

		//取用户联系信息
		var userinfoex account.Lctb_userInfoEx
		ret, _ := db.GetAccountExById(h.userinfo.Id, &userinfoex)
		if ret {
			h.Data["userinfoex"] = userinfoex
		}
		h.TplNames = "pages/home/myinfo.html"
	}

}

func (h *MyinfoController) PayInfo() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo

		h.TplNames = "pages/home/payinfo.html"
	}

}
