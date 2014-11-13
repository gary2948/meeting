package main

import (
	. "commonPackage"
	"commonPackage/model/account"
	"commonPackage/model/social"
	"commonPackage/viewModel"
	"service/db"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
)

func GetFriendList(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64(0)
	tag := jsonBody.Get(JSON_TAG).MustString("")
	uType := jsonBody.Get(JSON_TYPE).MustInt(0)

	reJson := simplejson.New()
	var err error
	us := []viewModel.FollowKind{}
	if uType == Follow {
		err = db.GetFollows(uId, &us)
	} else if uType == Fans {
		err = db.GetFans(uId, &us)
	}

	if err == nil {
		reJson.Set(JSON_CONTENT, us)
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_TAG, tag)
	} else {
		reJson.Set(JSON_ERR, err.Error())
		reJson.Set(JSON_OK, false)
	}

	return reJson.Encode()
}

func FollowUser(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64(0)
	flId := jsonBody.Get(JSON_KINDID).MustInt64(0)
	flIds, err := Int64Array(jsonBody.Get(JSON_FOLLOWUIDS))

	reJson := simplejson.New()
	err = db.FollowUser(uId, flId, flIds)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
	}

	return reJson.Encode()
}

func SendPersonMsg(ws *websocket.Conn, jsonBody *simplejson.Json, msgType int) ([]byte, error) {
	bm, _ := jsonBody.Encode()
	reJson := simplejson.New()
	uId := jsonBody.Get(JSON_USERID).MustInt64(0)
	toUId := jsonBody.Get(JSON_TOUSERID).MustInt64(0)
	if len(bm)+cmdLeg >= bufLeg {
		reJson.Set(JSON_OK, false)
		reJson.Set(JSON_ERR, ErrTooLong)
	} else {
		postTime := time.Now()
		jsonBody.Set(JSON_POSTTIME, postTime.Format(TimeFormat))
		bmm, err := jsonBody.Encode()
		client, ok := G.IdClients[toUId]

		if ok {
			if msgType == TextMsg {
				err = client.Conn.WriteMessage(websocket.BinaryMessage, ByteJoin([]byte(reseive), bmm))
			} else if msgType == ImgMsg {
				err = client.Conn.WriteMessage(websocket.BinaryMessage, ByteJoin([]byte(reseiveImg), bmm))
			}
			if err != nil {
				_, err = db.SendPersonMessage(uId, toUId, postTime, string(bmm), UnreadMsg, msgType)
			} else {
				_, err = db.SendPersonMessage(uId, toUId, postTime, string(bmm), ReadedMsg, msgType)
			}

		} else {
			_, err = db.SendPersonMessage(uId, toUId, postTime, string(bmm), UnreadMsg, msgType)
		}
		if err == nil {
			reJson.Set(JSON_OK, true)
		} else {
			reJson.Set(JSON_OK, false)
			reJson.Set(JSON_ERR, err.Error())
		}
	}

	return reJson.Encode()

}

func GetUnreadMsgs(ws *websocket.Conn, jsonBody *simplejson.Json) {
	uId := jsonBody.Get(JSON_USERID).MustInt64(0)
	txtMsg, imgMsg := db.GetAllUnreadMsg(uId)
	for _, v := range txtMsg {
		ws.WriteMessage(websocket.BinaryMessage, []byte(reUnreadTxtMsg+v))
	}
	for _, v := range imgMsg {
		ws.WriteMessage(websocket.BinaryMessage, []byte(reUnreadImgMsg+v))
	}

}

func GetUsers(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64(0)
	limit := jsonBody.Get(JSON_LIMIT).MustInt(0)
	start := jsonBody.Get(JSON_START).MustInt(0)
	reJson := simplejson.New()
	us := []account.Lctb_userInfo{}
	err := db.GetAccountByPages(uId, limit, start, &us)
	if err == nil {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, us)
	} else {
		reJson.Set(JSON_OK, false)
		reJson.Set(JSON_ERR, err.Error())
	}

	return reJson.Encode()
}

func CreateGroup(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64(0)
	gName, _ := jsonBody.Get(JSON_GROUPNAME).String()
	uIds, _ := Int64Array(jsonBody.Get(JSON_USERIDS))
	reJson := simplejson.New()
	gid, err := db.CreateGroup(uId, gName)
	err = db.AddUserToGroup(gid, uIds)
	if err != nil {
		reJson.Set(JSON_OK, false)
		reJson.Set(JSON_ERR, err.Error())
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_GROUPID, gid)
	}

	return reJson.Encode()
}

func AddUsersToGroup(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	gId := jsonBody.Get(JSON_GROUPID).MustInt64(0)
	uIds, err := Int64Array(jsonBody.Get(JSON_USERIDS))
	reJson := simplejson.New()
	CheckError(err)
	err = db.AddUserToGroup(gId, uIds)

	var ok bool
	if err != nil {
		ok = false
		reJson.Set(JSON_ERR, err.Error())
	} else {
		ok = true
	}

	reJson.Set(JSON_OK, ok)
	if ok { //给在线人员转发消息

		var group social.Lctb_talkGroup
		var us []account.Lctb_userInfo
		db.GetGroupInfo(gId, &group, &us)

		var users []account.Lctb_userInfo
		db.GetAccountsById(uIds, &users)
		jsonBody.Set(JSON_USERINFOS, users)
		body, _ := jsonBody.Encode()
		bmsg := ByteJoin([]byte(boardAddUsersToGroup), body)
		for _, v := range us {
			c, has := G.IdClients[v.Id]
			if has {
				c.Conn.WriteMessage(websocket.BinaryMessage, bmsg)
			}
		}
		for _, v := range uIds {
			c, has := G.IdClients[v]
			if has {
				c.Conn.WriteMessage(websocket.BinaryMessage, bmsg)
			}
		}
	}

	return reJson.Encode()
}

func DelUesrsFromGroup(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	gId := jsonBody.Get(JSON_GROUPID).MustInt64(0)
	uIds, _ := Int64Array(jsonBody.Get(JSON_USERIDS))
	reJson := simplejson.New()

	err := db.DelUserFromGroup(gId, uIds)
	var ok bool
	if err != nil {
		ok = false
		reJson.Set(JSON_ERR, err.Error())
	} else {
		ok = true
	}
	reJson.Set(JSON_OK, ok)
	if ok { //给在线人员转发消息
		var group social.Lctb_talkGroup
		var us []account.Lctb_userInfo
		db.GetGroupInfo(gId, &group, &us)
		body, _ := jsonBody.Encode()
		bmsg := ByteJoin([]byte(boardDelUserFromGroup), body)
		for _, v := range us {
			c, has := G.IdClients[v.Id]
			if has {
				c.Conn.WriteMessage(websocket.BinaryMessage, bmsg)
			}
		}
	}

	return reJson.Encode()
}

func GetUserGroups(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64(0)
	gms := []social.Lctb_talkGroup{}
	err := db.GetUserGroups(uId, &gms)
	reJson := simplejson.New()
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, gms)
	}

	return reJson.Encode()
}

func GetGroupInfo(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	gm := social.Lctb_talkGroup{}
	us := []account.Lctb_userInfo{}

	gid := jsonBody.Get(JSON_GROUPID).MustInt64()
	err := db.GetGroupInfo(gid, &gm, &us)

	reJson := simplejson.New()
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_GROUPINFO, gm)
		reJson.Set(JSON_USERINFOS, us)
	}

	return reJson.Encode()
}

func SendGroupMsg(ws *websocket.Conn, jsonBody *simplejson.Json, msgType int) ([]byte, error) {
	gid := jsonBody.Get(JSON_GROUPID).MustInt64()
	uid := jsonBody.Get(JSON_USERID).MustInt64()
	msg, _ := jsonBody.Get(JSON_CONTENT).String()
	body, _ := jsonBody.Encode()
	msgId, userIds, err := db.SendGroupMsg(gid, uid, msg, msgType)

	for _, v := range userIds {
		c, has := G.IdClients[v]
		boardMsg := ByteJoin([]byte(BoardGroupTxtMsg), body)
		if has {
			err = db.InsertGroupMsgStatus(gid, msgId, v, ReadedMsg, msgType)
			c.Conn.WriteMessage(websocket.BinaryMessage, boardMsg)
		} else {
			err = db.InsertGroupMsgStatus(gid, msgId, v, UnreadMsg, msgType)
		}
	}

	reJson := simplejson.New()
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
	}

	return reJson.Encode()
}

func GetUnreadGroupMsg(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	groupId := jsonBody.Get(JSON_GROUPID).MustInt64()
	userId := jsonBody.Get(JSON_USERID).MustInt64()

	txtmsg, imgMsg, err := db.GetUnreadGroupMsg(groupId, userId)

	reJson := simplejson.New()
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_TXTMSG, txtmsg)
		reJson.Set(JSON_IMGMSG, imgMsg)
	}

	return reJson.Encode()
}

func SearchUserByEmail(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	email, _ := jsonBody.Get(JSON_EMAIL).String()
	u := account.Lctb_userInfo{}
	has, err := db.GetAccountByEmail(email, &u)

	reJson := simplejson.New()
	if has && err == nil {
		reJson.Set(JSON_CONTENT, u)
		reJson.Set(JSON_OK, true)
	} else {
		reJson.Set(JSON_OK, false)
		if err != nil {
			reJson.Set(JSON_ERR, err.Error())
		}
	}

	return reJson.Encode()
}

func AddFollowKind(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64(0)
	kname, err := jsonBody.Get(JSON_KINDNAME).String()
	CheckError(err)
	reJson := simplejson.New()
	kid, err := db.AddFollowKind(uId, kname)
	if err != nil {
		reJson.Set(JSON_OK, false)
		reJson.Set(JSON_ERR, ErrSys)
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, kid)
	}

	return reJson.Encode()
}

func RenameFollowKind(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	kId := jsonBody.Get(JSON_KINDID).MustInt64(0)
	kname, err := jsonBody.Get(JSON_KINDNAME).String()
	CheckError(err)
	reJson := simplejson.New()
	err = db.RenameFollowKind(kId, kname)
	if err != nil {
		reJson.Set(JSON_OK, false)
		reJson.Set(JSON_ERR, ErrSys)
	} else {
		reJson.Set(JSON_OK, true)
	}

	return reJson.Encode()
}

func DelFollowKind(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	kId := jsonBody.Get(JSON_KINDID).MustInt64(0)
	reJson := simplejson.New()

	err := db.DelFollowKind(kId)
	if err != nil {
		reJson.Set(JSON_OK, false)
		reJson.Set(JSON_ERR, ErrSys)
	} else {
		reJson.Set(JSON_OK, true)
	}

	return reJson.Encode()

}

func MoveFollowKind(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	kId := jsonBody.Get(JSON_KINDID).MustInt64(0)
	uId := jsonBody.Get(JSON_USERID).MustInt64(0)
	uIds, err := Int64Array(jsonBody.Get(JSON_USERIDS))
	CheckError(err)
	reJson := simplejson.New()

	err = db.MoveUserToFollowKind(kId, uId, uIds)

	return reJson.Encode()
}
