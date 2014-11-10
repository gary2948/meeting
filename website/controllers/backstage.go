package controllers

import ()

type BackstageController struct {
	BaseController
}

func (h *BackstageController) Backstage() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/backstage/backstage.html"
	}
}

func (h *BackstageController) Useraccount() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/backstage/useraccount.html"
	}
}

func (h *BackstageController) Adminaccount() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/backstage/adminaccount.html"
	}
}
func (h *BackstageController) Teamaccount() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/backstage/teamaccount.html"
	}
}
func (h *BackstageController) Memberinfo() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/backstage/memberinfo.html"
	}
}
func (h *BackstageController) Dynamicmanager() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/backstage/dynamicmanager.html"
	}
}
func (h *BackstageController) Ipmanager() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/backstage/ipmanager.html"
	}
}
