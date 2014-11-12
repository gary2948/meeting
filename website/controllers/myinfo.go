package controllers

import (
	"commonPackage/model/account"
	"commonPackage/viewModel"
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

//扩展信息修改
func (f *MyinfoController) UpdateExContactInfo() {
	mystruct := ""
	if f.userinfo != nil {
		var vm viewModel.EditUserInfoExModel
		vm.Lc_mobilePhone1 = f.GetString("Lc_mobilePhone1")
		vm.Lc_qq = f.GetString("Lc_qq")
		vm.Lc_weixin = f.GetString("Lc_weixin")
		vm.Lc_weibo = f.GetString("Lc_weibo")
		err := db.UpdateAccountEx(f.userinfo.Id, &vm)
		if err != nil {
			mystruct = `{result:false}`
		} else {
			mystruct = `{result:true}`
		}
	} else {
		//用户未登录，返回修改失败
		mystruct = `{result:false}`
	}
	f.Data["json"] = &mystruct
	f.ServeJson()

}
func (f *MyinfoController) UpdateExBaseInfo() {
	mystruct := ""
	if f.userinfo != nil {
		var vm viewModel.EditUserInfoExModel
		vm.Lc_mobilePhone1 = f.GetString("Lc_mobilePhone1")
		vm.Lc_qq = f.GetString("Lc_qq")
		vm.Lc_weixin = f.GetString("Lc_weixin")
		vm.Lc_weibo = f.GetString("Lc_weibo")
		err := db.UpdateAccountEx(f.userinfo.Id, &vm)
		if err != nil {
			mystruct = `{result:false}`
		} else {
			mystruct = `{result:true}`
		}
	} else {
		//用户未登录，返回修改失败
		mystruct = `{result:false}`
	}
	f.Data["json"] = &mystruct
	f.ServeJson()

}
