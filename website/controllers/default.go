package controllers

import (
	"commonPackage/model/account"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"service/db"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	_, err := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	//beego.Info(userInfo)
	if err {
		this.Redirect("/home", 301)
	} else {
		this.TplNames = "login.html"
	}
}

func (this *MainController) Post() {
	_, err := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	//beego.Info(userInfo)
	if err {
		this.Redirect("/home", 301)
	} else {
		var Account = this.GetString("email")
		var password = this.GetString("password")
		result, _, user := loginValid(Account, password)
		if result {
			this.SetSession("Lctb_userInfo", user)
			this.Redirect("/home", 301)
		} else {
			this.Data["err"] = "用户账号或密码错误"
			this.TplNames = "login.html"
		}
	}
}

//登录验证 验证账号和密码
func loginValid(accout, password string) (result bool, err error, user *account.Lctb_userInfo) {
	user = new(account.Lctb_userInfo)
	result, err = db.LoginByEmail(accout, password, user)
	return
}

func StringMD5Value(value string) (result string) {
	var md = md5.New()
	md.Write([]byte(value))
	result = hex.EncodeToString(md.Sum(nil))
	return
}
