package main

import (
	. "commonPackage"

	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

func WsServer(ws *websocket.Conn) {
	for {
		_, buf, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		//70秒连接
		//ws.SetWriteDeadline(time.Now().Add(time.Duration(70) * time.Second))
		head := string(buf[:cmdLeg]) //命令头
		body := buf[cmdLeg:]         //命令数据
		log.Print(head)
		i, err := strconv.Atoi(head)
		if err != nil {
			log.Print(err.Error())
			ws.WriteMessage(websocket.BinaryMessage, []byte(err.Error()))
			break
		}
		if i < 7000000 && i > 6000000 {
			go doServer(ws, head, body)
		} else if i != 1000000 {
			ws.WriteMessage(websocket.BinaryMessage, []byte("无效命令"))
			break
		}

	}
	fmt.Printf("service finish %s\n", time.Now().String())
}

func doServer(ws *websocket.Conn, head string, body []byte) {
	//	jsonBody, err := simplejson.NewJson(body)
	//	CheckError(err)
	Println(head)
	Printb(body)
	switch head {
	case startMeeting:
		log.Print("start meeting room")
		StartMeeting(ws, body)
		break
	case loginMeeting:
		log.Print("login meeting")
		LoginMeeting(ws, body)
		break
	case logoutMeeting:
		log.Print("logout meeting")
		LogoutMeeting(ws, body)
		break
	case mouseMove:
		log.Print("mouse move")
		MouseMove(ws, body)
		break
	case synGeModel:
		log.Print("synGeModel")
		SynGeModel(ws, body)
		break
	case synRoomPermission:
		log.Print("synRoomPermission")
		SynRoomPermissionService(ws, body)
		break
	case synMeetingFileInfo:
		log.Print("synMeetingFileInfo")

		break
	case getOtherUserInfo:
		log.Print("getOtherUserInfo")
		GetMeetingUserInfos(ws, body)
		break
	default:
		log.Print("无效命令")
		//ws.Write([]byte("无效命令"))
		break

	}
}
