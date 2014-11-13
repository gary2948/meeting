package main

import (
	. "commonPackage"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
)

func WsServer(ws *websocket.Conn) {
	for {
		_, buf, err := ws.ReadMessage()
		if err != nil {
			Println(err.Error())
			break
		}

		//ws.SetWriteDeadline(time.Now().Add(time.Duration(70) * time.Second))
		head := string(buf[:cmdLeg])
		body := buf[cmdLeg:]

		Println(head)
		i, err := strconv.Atoi(head)
		if err != nil {
			Println(err.Error())
			//ws.Write([]byte(err.Error()))
			ws.WriteMessage(websocket.BinaryMessage, []byte(err.Error()))
			break
		}

		if i < 5000000 && i > 1000000 {
			go doServer(ws, head, body)
		} else if i != 1000000 {
			//ws.Write([]byte("无效命令"))
			ws.WriteMessage(websocket.BinaryMessage, []byte("无效命令"))
			break
		}

	}
	G.RemoveClientAddr(ws.RemoteAddr().String())
}

func reDoServer(ws *websocket.Conn, recmd, jsonBytes []byte) {
	ws.WriteMessage(websocket.BinaryMessage, ByteJoin(recmd, jsonBytes))
}

func doServer(ws *websocket.Conn, head string, body []byte) {
	jsonBody, err := simplejson.NewJson(body)
	if err != nil {
		return
	}

	switch head {
	case login:
		jsonBytes, err := Login(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reLogin), jsonBytes)
		break
	case logout:
		Logout(ws, jsonBody)
		break
	case getSystime:
		jsonBytes, err := GetSystime(ws)
		CheckError(err)
		reDoServer(ws, []byte(reGetSystime), jsonBytes)
		break
	case getUserInfo:
		jsonBytes, err := GetUserInfo(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetUserInfo), jsonBytes)
		break
	case getMyInfo:
		jsonBytes, err := GetUserInfo(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetMyInfo), jsonBytes)
		break
	case getUserInfoEx:
		jsonBytes, err := GetUserInfoEx(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetUserInfoEx), jsonBytes)
		break
	case getMeetingInfo:
		jsonBytes, err := GetMeetingInfo(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetMeetingInfo), jsonBytes)
		break
	case getMeetingInvPer:
		jsonBytes, err := GetMeetingInvPer(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetMeetingInvPer), jsonBytes)
		break
	case getMeetingPerson:
		jsonBytes, err := GetAttendPerson(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetMeetingPerson), jsonBytes)
		break
	case getSimpleMeetingsByDate:
		jsonBytes, err := GetSimpleMeetingsByDate(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetSimpleMeetingsByDate), jsonBytes)
		break
	case getDetailMeetingsByDate:
		jsonBytes, err := GetUserMeetingByDate(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetDetailMeetingsByDate), jsonBytes)
		break
	case getMeetingFiles:
		jsonBytes, err := GetMeetingFiles(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetMeetingFiles), jsonBytes)
		break
	case planMeeting:
		jsonBytes, err := PlanMeetingByDate(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(rePlanMeeting), jsonBytes)
		break
	case getFilesCountByFolder:
		jsonBytes, err := GetFilesCount(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetFilesCountByFolder), jsonBytes)
		break
	case getFolderVolumes:
		jsonBytes, err := GetFolderSize(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetFolderVolumes), jsonBytes)
		break
	case createFolder:
		jsonBytes, err := AddFolder(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reCreateFolder), jsonBytes)
		break
	case addFileInfo:
		jsonBytes, err := AddFileInfo(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reAddFileInfo), jsonBytes)
		break
	case updateLinkFileMapId:
		jsonBytes, err := UpdateLinFileMapId(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reUpdateLinkFileMapId), jsonBytes)
		break
	case getCloudFileInfo:
		jsonBytes, err := GetCouldFileInfo(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetCloudFileInfo), jsonBytes)
	case getUploadtoken:
		jsonBytes, err := GetUploadtoken(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetUploadtoken), jsonBytes)
		break
	case getDownloadtoken:
		jsonBytes, err := GetDownLoadUrl(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetDownloadtoken), jsonBytes)
	case getFilesByFolder:
		jsonBytes, err := GetUserFolderViewFiles(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetFilesByFolder), jsonBytes)
		break
	case getSharedFiles:
		jsonBytes, err := GetSharedFiles(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetSharedFiles), jsonBytes)
		break
	case shareFiles:
		jsonBytes, err := ShareFiles(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reShareFiles), jsonBytes)
		break
	case unshareFiles:
		jsonBytes, err := UnshareFiles(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reUnshareFiles), jsonBytes)
		break
	case getRecycleFiles:
		jsonBytes, err := GetRecycleFiles(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetRecycleFiles), jsonBytes)
		break
	case inRecycleBin:
		jsonBytes, err := InRecycleFiles(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reInRecycleBin), jsonBytes)
		break
	case outRecycleBin:
		jsonBytes, err := OutRecycleFiles(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reOutRecycleBin), jsonBytes)
		break
	case delFiles:
		jsonBytes, err := DelRecycleFiles(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reDelFiles), jsonBytes)
		break
	case addFollow:
		jsonBytes, err := FollowUser(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reAddFollow), jsonBytes)
		break
	case getFollowOrFansList:
		jsonBytes, err := GetFriendList(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetFollowFansList), jsonBytes)
		break
	case sendMessage:
		jsonBytes, err := SendPersonMsg(ws, jsonBody, TextMsg)
		CheckError(err)
		reDoServer(ws, []byte(reSendMessage), jsonBytes)
		break
	case sendImgMessage:
		jsonBytes, err := SendPersonMsg(ws, jsonBody, ImgMsg)
		CheckError(err)
		reDoServer(ws, []byte(reSendImgMessage), jsonBytes)
		break
	case getAllUnreadMsg:
		GetUnreadMsgs(ws, jsonBody)
		break
	case getUsers:
		jsonBytes, err := GetUsers(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetUsers), jsonBytes)
		break
	case getPersonExperience:
		jsonBytes, err := GetPsersonExp(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetPersonExperience), jsonBytes)
		break
	case createGroup:
		jsonBytes, err := CreateGroup(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reCreateGroup), jsonBytes)
		break
	case addUsersToGroup:
		jsonBytes, err := AddUsersToGroup(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reAddUsersToGroup), jsonBytes)
		break
	case delUsersFromGroup:
		jsonBytes, err := DelUesrsFromGroup(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reDelUserFromGroup), jsonBytes)
		break
	case getUserGroups:
		jsonBytes, err := GetUserGroups(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetUserGroups), jsonBytes)
		break
	case getGroupInfo:
		jsonBytes, err := GetGroupInfo(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetGroupInfo), jsonBytes)
		break
	case sendGroupTxtMsg:
		jsonBytes, err := SendGroupMsg(ws, jsonBody, TextMsg)
		CheckError(err)
		reDoServer(ws, []byte(reSendGroupTxtMsg), jsonBytes)
		break
	case sendGroupImgMsg:
		jsonBytes, err := SendGroupMsg(ws, jsonBody, ImgMsg)
		CheckError(err)
		reDoServer(ws, []byte(reSendGroupImgMsg), jsonBytes)
		break
	case getUnreadGroupMsg:
		jsonBytes, err := GetUnreadGroupMsg(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetUnreadGroupMsg), jsonBytes)
		break
	case searchUserByEmail:
		jsonBytes, err := SearchUserByEmail(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reSearchUserByEmail), jsonBytes)
		break
	case addFollowKind:
		jsonBytes, err := AddFollowKind(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reAddFollowKind), jsonBytes)
		break
	case renameFollowKind:
		jsonBytes, err := RenameFollowKind(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reReanemFollowKind), jsonBytes)
		break
	case delFollowKind:
		jsonBytes, err := DelFollowKind(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reDelFollowKind), jsonBytes)
		break
	case moveFollowKind:
		jsonBytes, err := MoveFollowKind(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reMoveFollowKind), jsonBytes)
		break
	case getRTMPURL:
		jsonBytes, err := GetRTMPURL(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reGetRTMPURL), jsonBytes)
	case askVedio:
		jsonBytes, err := AskVedio(ws, jsonBody)
		CheckError(err)
		reDoServer(ws, []byte(reAskVedio), jsonBytes)
	default:
		Println("无效命令")
		//ws.Write([]byte("无效命令"))
		break

	}

}
