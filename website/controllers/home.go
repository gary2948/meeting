package controllers

import (
	"service/db"
	"website/utils"
)

type HomeController struct {
	BaseController
}

func (h *HomeController) Get() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		//获得用户云盘空间
		mSize, uSize, _ := db.GetUserSizeByUserId(h.userinfo.Id)
		h.Data["mSize"] = utils.SizeConvert(mSize)
		h.Data["uSize"] = utils.SizeConvert(uSize)
		h.Data["Sizep"] = utils.SizePercent(uSize, mSize)

		//获得用户文档总数-暂无接口
		h.Data["doctotal"] = 0

		h.TplNames = "pages/home/home.html"
	}

}
