package main

import (
	. "commonPackage"
	"commonPackage/model/account"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
	//	"log"
	"service/db"
	"time"
)

func Login(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	reJson := simplejson.New()
	uu := &account.Lctb_userInfo{}
	userName, err := jsonBody.Get(JSON_USERNAME).String()
	CheckError(err)
	password, err := jsonBody.Get(JSON_PASSWORD).String()
	CheckError(err)
	ok, err := db.LoginByEmail(userName, password, uu)

	var isLogin bool
	if err != nil {
		fmt.Println(err)
		reJson.Set(JSON_ERR, err.Error())
		isLogin = false
	} else if err == nil && ok {
		isLogin = true
	}

	reJson.Set(JSON_OK, isLogin)

	if isLogin {
		reJson.Set(JSON_USERID, uu.Id)
		//add to group
		client := G.AddClient(uu.Id, ws)
		if client == nil {
			Printf(client.Addr)
		}
	}
	return reJson.Encode()
}

func Logout(ws *websocket.Conn, jsonBody *simplejson.Json) {
	uid := jsonBody.Get(JSON_USERID).MustInt64()
	G.RemoveClient(uid)
	ws.Close()
}

func GetUserInfo(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uid := jsonBody.Get(JSON_USERID).MustInt64(0)
	reJson := simplejson.New()
	uu := &account.Lctb_userInfo{}
	ok, err := db.GetAccountById(uid, uu)
	if err != nil {
		fmt.Println(err)
		reJson.Set(JSON_ERR, err.Error())
		reJson.Set(JSON_OK, false)
	} else if ok {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, *uu)
		reJson.Set(JSON_USERID, uid)
	}

	return reJson.Encode()

}

func GetUserInfoEx(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uid := jsonBody.Get(JSON_USERID).MustInt64(0)
	reJson := simplejson.New()
	uu := &account.Lctb_userInfoEx{}
	_, err := db.GetAccountExById(uid, uu)
	if err != nil {
		fmt.Println(err)
		reJson.Set(JSON_ERR, err.Error())
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, *uu)
		reJson.Set(JSON_USERID, uid)
	}
	return reJson.Encode()
}

func GetPsersonExp(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uid := jsonBody.Get(JSON_USERID).MustInt64(0)
	reJson := simplejson.New()
	p := &[]account.Lctb_personExperience{}
	err := db.GetPersonExperience(uid, p)
	if err != nil {
		fmt.Println(err)
		reJson.Set(JSON_ERR, err.Error())
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, *p)
		reJson.Set(JSON_USERID, uid)
	}
	return reJson.Encode()
}

func GetSystime(ws *websocket.Conn) ([]byte, error) {
	reJson := simplejson.New()
	reJson.Set(JSON_CONTENT, time.Now().Format(TimeFormat))
	return reJson.Encode()
}

func GetRTMPURL(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uid := jsonBody.Get(JSON_USERID).MustInt64(0)
	user := &account.Lctb_userInfo{}
	has, err := db.GetAccountById(uid, user)

	reJson := simplejson.New()
	if has && err == nil {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, RTMP_PRIX+user.Lc_UUID)
	} else {
		reJson.Set(JSON_OK, false)
		reJson.Set(JSON_ERR, ErrUserId)
	}

	return reJson.Encode()

}

func AskVedio(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uid := jsonBody.Get(JSON_USERID).MustInt64(0)
	uids, _ := Int64Array(jsonBody.Get(JSON_USERIDS))
	user := &account.Lctb_userInfo{}
	has, err := db.GetAccountById(uid, user)

	reJson := simplejson.New()
	if has && err == nil {

		users := []account.Lctb_userInfo{}
		err = db.GetAccountsById(uids, &users)
		if err != nil {
			reJson.Set(JSON_OK, false)
			reJson.Set(JSON_ERR, ErrUserId)
		} else {
			urls := []simplejson.Json{}

			for _, v := range users {
				rurl := simplejson.New()
				rurl.Set(JSON_USERID, v.Id)
				rurl.Set(JSON_CONTENT, RTMP_PRIX+v.Lc_UUID)
				urls = append(urls, *rurl)

				otherUrls := []simplejson.Json{}
				otherUrl := simplejson.New()
				otherUrl.Set(JSON_USERID, user.Id)
				otherUrl.Set(JSON_CONTENT, RTMP_PRIX+user.Lc_UUID)
				otherUrls = append(otherUrls, *otherUrl)

				boardJson := simplejson.New()
				boardJson.Set(JSON_MYRTMPURL, RTMP_PRIX+v.Lc_UUID)
				boardJson.Set(JSON_USERID, uid)
				for _, vv := range users {

					if v != vv {
						otherUrl2 := simplejson.New()
						otherUrl2.Set(JSON_USERID, vv.Id)
						otherUrl2.Set(JSON_CONTENT, RTMP_PRIX+vv.Lc_UUID)
						otherUrls = append(otherUrls, *otherUrl2)
					}
				}
				boardJson.Set(JSON_CONTENT, otherUrls)
				boardMsg, _ := boardJson.Encode()

				G.SendMsgToUser(v.Id, ByteJoin([]byte(boardAskVedio), boardMsg))
			}
			reJson.Set(JSON_OK, true)
			reJson.Set(JSON_CONTENT, urls)
			reJson.Set(JSON_MYRTMPURL, RTMP_PRIX+user.Lc_UUID)
		}
	} else {
		reJson.Set(JSON_OK, false)
		reJson.Set(JSON_ERR, ErrUserId)
	}

	return reJson.Encode()
}
