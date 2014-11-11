package main

import (
	. "commonPackage"
	"commonPackage/model/account"
	"commonPackage/model/meeting"
	"service/db"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
)

func StartMeeting(ws *websocket.Conn, body []byte) {
	var mi meetingInfo              //会议结果
	var mt meeting.Lctb_meetingInfo //数据库会议信息

	//reJson := simplejson.New()

	JSONUnmarshal(body, &mi)

	has, err := db.GetMeetingById(mi.MEETINGID, &mt)
	CheckError(err)
	Printf(has)
	if has {
		//NewGMeetings(mt.Id, mt.Lc_limitCount)
		_, mhas := G[mt.Id]
		if !mhas { //没有才创建
			Println("new meeting")
			NewGMeetings(mt.Id, mt.Lc_limitCount)
			Println("new meeting finish")
		}
		mi.CONTENT = mt
		mi.OK = true
	} else {
		mi.OK = false
		mi.ERR = ErrMeetingId
	}

	remsg := GetReComm(reStartMeeting, mi)
	ws.WriteMessage(websocket.BinaryMessage, remsg)
}

func LoginMeeting(ws *websocket.Conn, body []byte) {
	//commonPackage.Println(ws.RemoteAddr().String())
	ut := account.Lctb_userInfo{}
	m, _ := simplejson.NewJson(body)
	mid := m.Get(JSON_MEETINGID).MustInt64()
	email, _ := m.Get(JSON_EMAIL).String()
	has, err := db.GetAccountByEmail(email, &ut)
	CheckError(err)
	rm := simplejson.New()
	//设置userinfo
	ui := simplejson.New()
	ui.Set(JSON_EMAIL, email)
	if has { //设置用户信息
		if ut.Lc_realName == "" {
			ui.Set(JSON_NAME, ut.Lc_realName)
		} else {
			ui.Set(JSON_NAME, ut.Lc_nickName)
		}
		ui.Set(JSON_ROLE, OfficialUser)
		rm.Set(JSON_USERINFO, ut)
	} else {
		ui.Set(JSON_NAME, "游客")
		ui.Set(JSON_ROLE, TrialUsers)
	}

	//目前始终都是登录成功的
	rm.Set(JSON_OK, true)
	uib, _ := ui.MarshalJSON()
	mm, has := G[mid]
	if has {
		G[mid].AddClient(email, uib, ws)
		BoardMsg := simplejson.New()
		BoardMsg.Set(JSON_MEETINGID, mid)
		BoardMsg.Set(JSON_CONTENT, ui)
		BoardMsgB, err := BoardMsg.MarshalJSON()
		if err != nil {
			Println(err.Error())
		}
		bmsg := ByteJoin([]byte(boardLoginMeeting), BoardMsgB)
		mm.BoardMsg(ws, bmsg)
		rm.Set(JSON_OK, true)
	} else {
		rm.Set(JSON_ERR, ErrMeetingId)
	}
	rmb, _ := rm.MarshalJSON()
	remsg := ByteJoin([]byte(reLoginMeeting), rmb)
	ws.WriteMessage(websocket.BinaryMessage, remsg)
}

func LogoutMeeting(ws *websocket.Conn, body []byte) {
	m, _ := simplejson.NewJson(body)
	rm := simplejson.New()
	bmg := simplejson.New()
	mid := m.Get(JSON_MEETINGID).MustInt64()
	email, _ := m.Get(JSON_EMAIL).String()

	mm, has := G[mid]
	if has {
		c, has := mm.EmailClients[email]
		if has {
			ui := c.UserInfo
			bmg.Set(JSON_CONTENT, ui)
			bmg.Set(JSON_MEETINGID, mid)

			bmgb, _ := bmg.MarshalJSON()
			bbmsg := ByteJoin([]byte(boardLogoutMeeting), bmgb)
			mm.RemoveClient(email)
			mm.BoardMsg(ws, bbmsg)
			rm.Set(JSON_OK, true)
		} else {
			rm.Set(JSON_ERR, ErrUserEmail)
			rm.Set(JSON_OK, false)
		}

	} else {
		rm.Set(JSON_ERR, ErrMeetingId)
		rm.Set(JSON_OK, false)
	}

	rmb, _ := rm.MarshalJSON()
	remsg := ByteJoin([]byte(reLogoutMeeting), rmb)
	ws.WriteMessage(websocket.BinaryMessage, remsg)
}

func MouseMove(ws *websocket.Conn, body []byte) {
	var mp mousePoints
	JSONUnmarshal(body, &mp)
	bf := ByteJoin([]byte(boradMouseMove), body)
	m, has := G[mp.MEETINGID]
	if has {
		m.BoardMsg(ws, bf)
	}
	rm := simplejson.New()
	rm.Set(JSON_OK, true)
	rmb, _ := rm.MarshalJSON()
	remsg := ByteJoin([]byte(reMouseMove), rmb)
	ws.WriteMessage(websocket.BinaryMessage, remsg)
}

func SynRoomPermissionService(ws *websocket.Conn, body []byte) {
	var rp userPower
	JSONUnmarshal(body, &rp)
	m, has := G[rp.MEETINGID]
	if has {
		m.SynUserPower(rp)
		bf := ByteJoin([]byte(boardRoomPermission), body)
		m.BoardMsg(ws, bf)
	}

	rm := simplejson.New()
	rm.Set(JSON_OK, true)
	rmb, _ := rm.MarshalJSON()
	remsg := ByteJoin([]byte(reSynRoomPermission), rmb)
	ws.WriteMessage(websocket.BinaryMessage, remsg)

}

func SysMeetingFileInfo(ws *websocket.Conn, body []byte) {
	var mf meetingFileInfo
	JSONUnmarshal(body, &mf)
	m, has := G[mf.MEETINGID]
	if has {
		m.SynMeetingFileInfo(mf)
		bf := ByteJoin([]byte(boardMeetingFileInfo), body)
		m.BoardMsg(ws, bf)
	}

	rm := simplejson.New()
	rm.Set(JSON_OK, true)
	rmb, _ := rm.MarshalJSON()
	remsg := ByteJoin([]byte(reSynRoomPermission), rmb)
	ws.WriteMessage(websocket.BinaryMessage, remsg)

}

func GetMeetingUserInfos(ws *websocket.Conn, body []byte) {
	m, _ := simplejson.NewJson(body)
	rm := simplejson.New()
	mid := m.Get(JSON_MEETINGID).MustInt64()

	mm, has := G[mid]
	uis := [][]byte{}
	if has {
		for _, v := range mm.IpClients {
			if v.Conn != ws {
				uis = append(uis, v.UserInfo)
			}
		}
		rm.Set(JSON_CONTENT, uis)
		rm.Set(JSON_OK, true)
	} else {
		rm.Set(JSON_ERR, ErrSys)
	}
	rmb, _ := rm.MarshalJSON()
	remsg := ByteJoin([]byte(reSynRoomPermission), rmb)
	ws.WriteMessage(websocket.BinaryMessage, remsg)
}
