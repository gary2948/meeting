package controllers

import ()

type ClouddiskController struct {
	BaseController
}

func (h *ClouddiskController) MyDoc() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/clouddisk/mydoc.html"
	}
}
