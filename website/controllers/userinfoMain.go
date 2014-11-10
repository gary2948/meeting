//进行个人信息主页跳转
package controllers

import (
	"commonPackage/model/account"
	"github.com/astaxie/beego"
	"service/db"
	"time"
)

type UserInfoController struct {
	beego.Controller
}

func (u *UserInfoController) Post() {

}

func (this *UserInfoController) Get() {
	beego.Info("userinfo")
	userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	beego.Info(userinfo)
	var userEx account.Lctb_userInfoEx
	db.GetAccountExById(userinfo.Id, &userEx)
	this.Data["userInfo"] = userinfo
	this.Data["userEx"] = userEx
	this.Data["Lc_regTime"] = userinfo.Lc_regTime.Format("2006-01-02")
	this.Data["zodiac"] = chineseZodiac(userEx.Lc_birthday)
	this.Data["constellation"] = constellation(userEx.Lc_birthday)
	this.Data["birthday"] = userEx.Lc_birthday.Format("2006年01月02日")
	this.Data["bitthSecond"] = userEx.Lc_birthday.Unix()
	//this.Data["language"] = Language(userEx.Lc_language) + Language(userEx.Lc_language1)
	this.TplNames = "pages/personInfo.html"
}

//生肖
func chineseZodiac(birth time.Time) string {
	arr := []string{"猴", "鸡", "狗", "猪", "鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊"}
	return arr[birth.Year()%12]
}

//星座
func constellation(birth time.Time) (result string) {
	month := birth.Month()
	day := birth.Day()
	if month == 3 && day >= 21 || month == 4 && day <= 19 {
		result = "白羊座"
	} else if month == 4 && day >= 20 || month == 5 && day <= 20 {
		result = "金牛座"
	} else if month == 5 && day >= 21 || month == 6 && day <= 21 {
		result = "双子座"
	} else if month == 6 && day >= 22 || month == 7 && day <= 22 {
		result = "巨蟹座"
	} else if month == 7 && day >= 23 || month == 8 && day <= 22 {
		result = "狮子座"
	} else if month == 8 && day >= 23 || month == 9 && day <= 22 {
		result = "处女座"
	} else if month == 9 && day >= 23 || month == 10 && day <= 23 {
		result = "天秤座"
	} else if month == 10 && day >= 24 || month == 11 && day <= 22 {
		result = "天蝎座"
	} else if month == 11 && day >= 23 || month == 12 && day <= 21 {
		result = "射手座"
	} else if month == 12 && day >= 22 || month == 1 && day <= 19 {
		result = "摩羯座"
	} else if month == 1 && day >= 20 || month == 2 && day <= 18 {
		result = "水瓶座"
	} else if month == 2 && day >= 19 || month == 3 && day <= 20 {
		result = "双鱼座"
	}
	return
}

func Language(value string) string {
	switch value {
	case "zh":
		return "中文"
	case "jp":
		return "日语"
	case "spa":
		return "西班牙语"
	case "th":
		return "泰语"
	case "ru":
		return "俄罗斯语"
	case "yue":
		return "粤语"
	case "en":
		return "英语"
	case "kor":
		return "韩语"
	case "fra":
		return "法语"
	case "ara":
		return "阿拉伯语"
	case "pt":
		return "葡萄牙语"
	case "wyw":
		return "文言文"
	default:
		return ""
	}
}
