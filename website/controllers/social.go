package controllers

import ()

type SocialController struct {
	BaseController
}

func (h *SocialController) Weibo() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/weibo.html"
	}
}

func (h *SocialController) Message() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/message.html"
	}
}
func (h *SocialController) Group() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/group.html"
	}
}
func (h *SocialController) AddGroup() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/addgroup.html"
	}
}
func (h *SocialController) Mycard() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/mycard.html"
	}
}
func (h *SocialController) Mydynamic() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/mydynamic.html"
	}
}
func (h *SocialController) Mycontact() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/mycontact.html"
	}
}

func (h *SocialController) Companycontact() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/companycontact.html"
	}
}
func (h *SocialController) Groupcontact() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo
		h.TplNames = "pages/social/groupcontact.html"
	}
}
