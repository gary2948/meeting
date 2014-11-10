package controllers

import ()

type ChangeheadController struct {
	BaseController
}

func (h *ChangeheadController) Get() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo

		h.Data["doctotal"] = 0

		h.TplNames = "pages/home/changehead.html"
	}

}
