package controllers

import (
	"commonPackage/model/account"
	"commonPackage/viewModel"
	"fmt"
	"service/db"
	"strconv"
	"strings"
	"time"
)

type MyinfoController struct {
	BaseController
}

func (h *MyinfoController) Get() {
	if h.userinfo != nil {

		//取用户联系信息
		var userinfoex account.Lctb_userInfoEx
		var userinfo account.Lctb_userInfo
		var userExpe = make([]account.Lctb_personExperience, 0)
		ret, _ := db.GetAccountExById(h.userinfo.Id, &userinfoex)
		ret, _ = db.GetAccountById(h.userinfo.Id, &userinfo)
		_ = db.GetPersonExperience(h.userinfo.Id, &userExpe)
		if ret {
			h.Data["userinfoex"] = userinfoex
			h.Data["userinfo"] = userinfo
			h.Data["userExpe"] = userExpe
		}
		//fmt.Println(userExpe)
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

//教育信息和职业信息的新增
func (f *MyinfoController) AddPersonexp() {
	mystruct := ""
	if f.userinfo != nil {
		var userExpe account.Lctb_personExperience
		userExpe.Lc_experKind, err: = f.GetInt("Lc_experKind")
		userExpe.Lc_unitType, err: = f.GetInt("Lc_unitType")
		userExpe.Lc_unitName = f.GetString("Lc_unitName")
		userExpe.Lc_beginTime = f.GetString("Lc_beginTime")
		userExpe.Lc_endTime = f.GetString("Lc_endTime")
		exps := [...]account.Lctb_personExperience{userExpe}
		err := db.AddPersonExperience(exps)
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

func (f *MyinfoController) UpdateBaseInfo() {
	mystruct := ""
	if f.userinfo != nil {
		var vm1 viewModel.EditUserInfoExModel
		var vm2 viewModel.EditUserInfoModel
		vm2.Lc_sex = f.GetString("Lc_sex")
		fmt.Println(f.GetString("Lc_sex"))
		if f.GetString("Lc_birthday") != "" {
			dates := strings.Split(f.GetString("Lc_birthday"), "-")
			year, _ := strconv.Atoi(dates[0])
			month, _ := strconv.Atoi(dates[1])
			day, _ := strconv.Atoi(dates[2])
			birthday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

			vm1.Lc_birthday = birthday
		}
		vm1.Lc_language = f.GetString("Lc_language")
		vm2.Lc_introduction = f.GetString("Lc_introduction")
		err := db.UpdateAccountEx(f.userinfo.Id, &vm1)
		if err != nil {
			mystruct = `{result:false}`
		} else {
			err = db.UpdateAccount(f.userinfo.Id, &vm2)
			if err != nil {
				mystruct = `{result:false}`
			} else {
				mystruct = `{result:true}`
			}
		}
	} else {
		//用户未登录，返回修改失败
		mystruct = `{result:false}`
	}
	f.Data["json"] = &mystruct
	f.ServeJson()

}
