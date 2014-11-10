package controllers

import ()

type ChangePasswordController struct {
	BaseController
}

func (this *ChangePasswordController) Post() {
	this.TplNames = "pages/home/changePassword.html"
}

func (h *ChangePasswordController) Get() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo

		h.Data["doctotal"] = 0

		h.TplNames = "pages/home/changePassword.html"
	}
}
