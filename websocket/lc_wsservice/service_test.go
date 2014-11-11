package main

import (
	. "commonPackage"
	"fmt"
	"log"
	"code.google.com/p/go.net/websocket"
	"github.com/bitly/go-simplejson"
	//	"math/rand"
	"net"
	"net/http/httptest"
	//"service/uuid"
	"sync"
	"testing"
	//	"time"
)

var (
	serverAddr string
	once       sync.Once
)

//var origin = "http://localhost/"
//var url = "ws://localhost:8001/"

func startServer() {
	G = NewGroup(ConnectionMax)
	server := httptest.NewServer(new(myServer))
	serverAddr = server.Listener.Addr().String()
	log.Print("Test WebSocket server listening on ", serverAddr)

}

func newConfig(t *testing.T, path string) *websocket.Config {
	config, _ := websocket.NewConfig(fmt.Sprintf("ws://%s%s", serverAddr, path), "http://localhost")
	return config
}

func newConn(t *testing.T) (ws *websocket.Conn, err error) {
	// websocket.Dial()
	client, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return nil, err
	}

	return websocket.NewClient(newConfig(t, "/"), client)

}

func TestLogin(t *testing.T) {
	once.Do(startServer)

	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}

	//message := []byte("1000000")
	//_, err = ws.Write(message)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//var msg = make([]byte, 512)
	//n, err := ws.Read(msg)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//if string(msg[:n]) != "1000000" {
	//	t.Errorf("链接错误")
	//}

	cmd := `1001000{"USERNAME":"zhao@163.com","PASSWORD":"123456"}`

	Println(cmd)
	//login ok

	ws.Write([]byte(cmd))
	var msg []byte
	websocket.Message.Receive(ws, &msg)
	if err != nil {
		t.Error(err)
	}

	//var ru reUserLogin
	//if err := json.Unmarshal(msg[7:n], &ru); err != nil {
	//	t.Error(err)
	//}

	Printb(msg)

	//cmd = login + `{"USERNAME":"zhao@163.com","PASSWORD":"111111"}`

	//ws.Write([]byte(cmd))

	//n, err = ws.Read(msg)
	//if err != nil {
	//	t.Error(err)
	//}

	//if err := json.Unmarshal(msg[7:n], &ru); err != nil {
	//	t.Error(err)
	//}
	//if ru.OK {
	//	 Printf(ru)
	//	t.Error("login ok")
	//}
	//login fail
	ws.Close()
}

func TestGetUserInfo(t *testing.T) {
	once.Do(startServer)

	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}
	msg := make([]byte, 512)
	cmd := "1002100" + `{"USERID":1}`

	ws.Write([]byte(cmd))
	n, err := ws.Read(msg)
	if err != nil {
		t.Error(err)
	}

	Printb(msg[7:n])
	if err != nil {
		t.Error(err)
	}

}

func TestGetMyInfo(t *testing.T) {
	once.Do(startServer)

	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}
	msg := make([]byte, 512)
	cmd := "1012100" + `{"USERID":1}`

	ws.Write([]byte(cmd))
	n, err := ws.Read(msg)
	if err != nil {
		t.Error(err)
	}

	Printb(msg[7:n])
	if err != nil {
		t.Error(err)
	}

}

func TestGetUserInfoEx(t *testing.T) {
	once.Do(startServer)

	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}
	msg := make([]byte, 512)
	cmd := getUserInfoEx + `{"USERID":1}`
	Println(cmd)
	ws.Write([]byte(cmd))
	_, err = ws.Read(msg)
	if err != nil {
		t.Error(err)
	}

}

//func TestGetMeetingInfo(t *testing.T) {
//	once.Do(startServer)

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}
//	msg := make([]byte, 1024)
//	cmd := "2001000" + `{"MEETINGID":1}`

//	ws.Write([]byte(cmd))
//	n, err := ws.Read(msg)
//	if err != nil {
//		t.Error(err)
//	}

//	var ui userInfo
//	 JSONUnmarshal(msg[7:n], &ui)

//	if err != nil {
//		t.Error(err)
//	}
//	if ui.ERR != "" {
//		t.Errorf(ui.ERR)
//	}

//}

//func TestGetInvitedPersons(t *testing.T) {
//	once.Do(startServer)
//	var mm meetingPersons

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	msg := make([]byte, 1024)

//	cmd := "2002000" + `{"MEETINGID":214}`

//	ws.Write([]byte(cmd))
//	n, err := ws.Read(msg)

//	if string(msg[:7]) != "2002010" {
//		t.Errorf("cmd error")
//	}
//	 Println(string(msg[7:n]))
//	err = json.Unmarshal(msg[7:n], &mm)
//	if err != nil {
//		t.Error(err)
//	}

//}

//func TestGetAttendPersons(t *testing.T) {
//	once.Do(startServer)
//	var mm attendPersons

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	cmd := "2007000" + `{"MEETINGID":1}`

//	ws.Write([]byte(cmd))

//	msg := make([]byte, 1024)
//	err = websocket.Message.Receive(ws, &msg)
//	//n, err := ws.Read(msg)

//	if string(msg[:7]) != "2007010" {
//		t.Errorf("cmd error")
//	}
//	err = json.Unmarshal(msg[7:], &mm)
//	if err != nil {
//		t.Error(err)
//	}

//}

//func TestGetSimpleMeetingsByDate(t *testing.T) {
//	once.Do(startServer)
//	var mm meetingSimpleInfo

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	cmd := "2003000" + `{"USERID":1,"BEGINTIME":"2014-06-01","ENDTIME":"2014-06-30"}`

//	ws.Write([]byte(cmd))
//	msg := make([]byte, 1024)
//	err = websocket.Message.Receive(ws, &msg)
//	//n, err := ws.Read(msg)
//	if string(msg[:7]) != "2003010" {
//		t.Errorf("cmd error")
//	}
//	err = json.Unmarshal(msg[7:], &mm)
//	 Printf(mm)
//	if err != nil {
//		t.Error(err)
//	}
//}

func TestGetUserMeetingByDate(t *testing.T) {
	once.Do(startServer)
	//	var mm userMeetingInfo

	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}

	cmd := "2004000" + `{"USERID":1,"MEETINGDATE":"2014-06-06"}`

	websocket.Message.Send(ws, cmd)

	msg := make([]byte, 1024)
	err = websocket.Message.Receive(ws, &msg)

	//err = json.Unmarshal(msg[7:], &mm)
	Printb(msg)
	if err != nil {
		t.Error(err)
	}

}

//func TestPlanMeetingByDate(t *testing.T) {
//	once.Do(startServer)
//	var mm planMeetingInfo
//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	//cmd := "2005000" + `{ "USERID": 1, "PLANDATE": "2014-06-10 00:00:00", "QUICKSTART": true, "EMAILS": [ "zhao1@1.com", "zhao1@2.com" , "zhao3@3.com", "zhao4@4.com"] }`
//	cmd := "2005000" + `{ "USERID": 1, "TOPIC": "TEST", "QUICKSTART": true}`
//	websocket.Message.Send(ws, cmd)

//	msg := make([]byte, 1024)
//	err = websocket.Message.Receive(ws, &msg)
//	 Println("22222222222222")
//	 Printf(string(msg[7:]))
//	err = json.Unmarshal(msg[7:], &mm)
//	if err != nil {
//		t.Error(err)
//	}

//}

//func TestGetFilesCount(t *testing.T) {
//	once.Do(startServer)
//	var cc couldFolderCount

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	cmd := "3001000" + `{"USERID":1,"FOLDERID":0}`

//	websocket.Message.Send(ws, cmd)

//	msg := make([]byte, 1024)
//	err = websocket.Message.Receive(ws, &msg)

//	err = json.Unmarshal(msg[7:], &cc)
//	if err != nil {
//		t.Error(err)
//	}
//}

//func TestGetFilesSize(t *testing.T) {
//	once.Do(startServer)
//	var cc couldFolderCount

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	cmd := "3002000" + `{"USERID":1,"FOLDERID":0}`

//	websocket.Message.Send(ws, cmd)

//	msg := make([]byte, 1024)
//	err = websocket.Message.Receive(ws, &msg)

//	err = json.Unmarshal(msg[7:], &cc)
//	if err != nil {
//		t.Error(err)
//	}
//}

//func TestAddFolder(t *testing.T) {
//	once.Do(startServer)
//	var cc couldFolderCount

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}
//	s, _ := uuid.GenUUID()
//	//r := rand.New(rand.NewSource(time.Now().UnixNano()))

//	cmd := "3004000" + fmt.Sprintf(`{"USERID":1,"PARENTID": %d, "FLODERNAME":"%s"}`, 1, s)
//	 Println(cmd)
//	websocket.Message.Send(ws, cmd)

//	msg := make([]byte, 1024)
//	err = websocket.Message.Receive(ws, &msg)

//	err = json.Unmarshal(msg[7:], &cc)
//	if err != nil {
//		t.Error(err)
//	}
//}

func TestGetMeetingFiles(t *testing.T) {
	once.Do(startServer)

	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}

	cmd := "2006020" + fmt.Sprintf(`{"MEETINGID":%d}`, 214)
	Println(cmd)
	websocket.Message.Send(ws, cmd)

	msg := make([]byte, 1024)
	err = websocket.Message.Receive(ws, &msg)

	if err != nil {
		t.Error(err)
	}

}

//func TeatAddCash(t *testing.T) {
//	once.Do(startServer)
//	var cc couldFolderCount

//	ws, err := newConn(t)
//	f := &UploadingFile{}
//	f.FileKey =  GetUnixTime()
//	f.FileSize = 0
//	f.FileName = "aaaa"
//	err = AddCash(f)
//	if err != nil {
//		t.Error(err)
//	}

//}

//func TestPrepareUploading(t *testing.T) {
//	once.Do(startServer)

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	pp := &prepareUploading{}
//	rpp := &rePrepareUploading{}
//	pp.CRC =  GetUnixTime()
//	pp.FOLDERID = 0
//	pp.SIZE = 0
//	pp.USERID = 9
//	b :=  JSONMarshal(pp)
//	cmd := "3005000" + string(b)
//	websocket.Message.Send(ws, cmd)

//	msg := make([]byte, 1024)
//	err = websocket.Message.Receive(ws, &msg)

//	 JSONUnmarshal(msg[7:], &rpp)

//	if err != nil {
//		t.Error(err)
//	}

//	 Printf(rpp)

//}

func TestFriendList(t *testing.T) {
	once.Do(startServer)
	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}

	cmd := "4001000" + `{"USERID":8,"TYPE":1,"TAG":"TAG"}`
	Println(cmd)
	websocket.Message.Send(ws, cmd)

	msg := make([]byte, 1024)
	err = websocket.Message.Receive(ws, &msg)

	if err != nil {
		t.Error(err)
	}
	Printb(msg[7:])
}

//func TestAddFollow(t *testing.T) {
//	once.Do(startServer)
//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	var rmi reFollowUser
//	cmd := "4002000" + `{"USERID":8,"FOLLOWUIDS":[1,12,13]}`
//	websocket.Message.Send(ws, cmd)

//	msg := make([]byte, 1024)
//	err = websocket.Message.Receive(ws, &msg)

//	err = json.Unmarshal(msg[7:], &rmi)
//	if err != nil {
//		t.Error(err)
//	}
//	 Printf(rmi)
//}

func TestGetFolderFiles(t *testing.T) {
	once.Do(startServer)
	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}

	cmd := "3003000" + `{"USERID":1,"FILEID":90}`

	websocket.Message.Send(ws, cmd)

	msg := make([]byte, 1024)
	err = websocket.Message.Receive(ws, &msg)

	if err != nil {
		t.Error(err)
	}
}

func TestGetSharedFiles(t *testing.T) {
	once.Do(startServer)
	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}

	cmd := "3006000" + `{"USERID":1}`

	websocket.Message.Send(ws, cmd)

	msg := make([]byte, 1024)
	err = websocket.Message.Receive(ws, &msg)
	Println(string(msg[7:]))
	if err != nil {
		t.Error(err)
	}
}

//func TestSharedFiles(t *testing.T) {
//	once.Do(startServer)
//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}
//	var rmi reShareFilesInfo
//	cmd := "3006020" + `{"USERID":1,"SHAREFILETYPE":1,"FILEIDS":[9,11],"SHARENAME":"分享文件","FILEEXT":"","FILESIZE":0}`

//	websocket.Message.Send(ws, cmd)

//	msg := make([]byte, 1024)
//	err = websocket.Message.Receive(ws, &msg)
//	 Printb(msg[7:])
//	err = json.Unmarshal(msg[7:], &rmi)
//	if err != nil {
//		t.Error(err)
//	}
//	 Printf(rmi)
//}

//func TestUnSharedFiles(t *testing.T) {
//	once.Do(startServer)
//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}
//	var rmi reShareFilesInfo
//	cmd := "3006040" + `{"SHAREINFOID":58}`

//	websocket.Message.Send(ws, cmd)

//	msg := make([]byte, 1024)
//	err = websocket.Message.Receive(ws, &msg)
//	err = json.Unmarshal(msg[7:], &rmi)
//	if err != nil {
//		t.Error(err)
//	}
//	 Printf(rmi)
//}

//func TestUnSharedFiles(t *testing.T) {
//	once.Do(startServer)
//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}
//	var rmi reShareFilesInfo
//	cmd := "3006040" + `{"SHAREINFOID":58}`

//	websocket.Message.Send(ws, cmd)

//	msg := make([]byte, 1024)
//	err = websocket.Message.Receive(ws, &msg)
//	err = json.Unmarshal(msg[7:], &rmi)
//	if err != nil {
//		t.Error(err)
//	}
//	 Printf(rmi)
//}

func TestGetRecycleFiles(t *testing.T) {
	once.Do(startServer)
	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}
	cmd := "3007000" + `{"USERID":1}`

	websocket.Message.Send(ws, cmd)

	msg := make([]byte, 1024)
	err = websocket.Message.Receive(ws, &msg)
	if err != nil {
		t.Error(err)
	}
}

//func TestInRecycleFiles(t *testing.T) {
//	once.Do(startServer)
//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}
//	var rmi reNoContentInfo
//	cmd := "3007020" + `{"USERID":1,"FILEIDS":[9,10,11]}`

//	websocket.Message.Send(ws, cmd)

//	msg := make([]byte, 1024)
//	err = websocket.Message.Receive(ws, &msg)
//	err = json.Unmarshal(msg[7:], &rmi)
//	if err != nil {
//		t.Error(err)
//	}
//	 Printf(rmi)
//}

//func TestOutRecycleFiles(t *testing.T) {
//	once.Do(startServer)
//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}
//	var rmi reNoContentInfo
//	cmd := "3007040" + `{"FILEIDS":[9,10,11]}`

//	websocket.Message.Send(ws, cmd)

//	msg := make([]byte, 1024)
//	err = websocket.Message.Receive(ws, &msg)
//	err = json.Unmarshal(msg[7:], &rmi)
//	if err != nil {
//		t.Error(err)
//	}
//	 Printf(rmi)
//}

//func TestDelRecycleFiles(t *testing.T) {
//	once.Do(startServer)
//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}
//	var rmi reNoContentInfo
//	cmd := "3007060" + `{"FILEIDS":[9,10,11]}`

//	websocket.Message.Send(ws, cmd)

//	msg := make([]byte, 1024)
//	err = websocket.Message.Receive(ws, &msg)
//	err = json.Unmarshal(msg[7:], &rmi)
//	if err != nil {
//		t.Error(err)
//	}
//	 Printf(rmi)
//}

//func TestSendMsg(t *testing.T) {
//	once.Do(startServer)
//	ws1, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	cmd1 := `1001000{"USERNAME":"zhao@163.com","PASSWORD":"123456"}`
//	ws1.Write([]byte(cmd1))
//	msg := make([]byte, 1024)
//	_, err = ws1.Read(msg)
//	if err != nil {
//		t.Error(err)
//	}

//	ws2, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}
//	cmd2 := `1001000{"USERNAME":"1403784744@163.com","PASSWORD":"123456"}`
//	ws2.Write([]byte(cmd2))

//	_, err = ws2.Read(msg)
//	if err != nil {
//		t.Error(err)
//	}

//	cmd3 := `4003000{"USERID":1,"TOUSERID":4,"CONTENTS":"231TEST"}`
//	ws1.Write([]byte(cmd3))

//	n, err := ws1.Read(msg)
//	 Printb(msg[0:n])
//	if err != nil {
//		t.Error(err)
//	}
//	n, err = ws2.Read(msg)
//	 Printb(msg[0:n])
//	if err != nil {
//		t.Error(err)
//	}
//}

//func TestUploading(t *testing.T) {

//	once.Do(startServer)

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	cmd := "4004000" + `{"USERID":1}`
//	ws.Write([]byte(cmd))

//	msg := make([]byte, 1024)
//	websocket.Message.Receive(ws, msg)
//	 Printb(msg)
//	if err != nil {
//		t.Error(err)
//	}

//}

//func TestPingPong(t *testing.T) {
//	once.Do(startServer)

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}
//}

func TestGetUsers(t *testing.T) {
	once.Do(startServer)

	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}

	cmd := "4006000" + `{"USERID":1,"LIMIT":1,"START":0}`
	ws.Write([]byte(cmd))

	msg := make([]byte, 1024)
	websocket.Message.Receive(ws, &msg)
	Println(string(msg[7:]))

}

func TestGetPersonExp(t *testing.T) {
	once.Do(startServer)

	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}

	cmd := "1004000" + `{"USERID":1}`
	ws.Write([]byte(cmd))

	msg := make([]byte, 1024)
	websocket.Message.Receive(ws, &msg)
	Println(string(msg[7:]))

}

//func TestCreateGroup(t *testing.T) {
//	once.Do(startServer)

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	var gi groupInfo
//	gi.GROUPNAME = "wstest"
//	gi.USERID = 1
//	gi.USERIDS = []int64{2, 3, 4, 5}

//	b :=  JSONMarshal(gi)
//	cmd := "4007300" + string(b)
//	websocket.Message.Send(ws, cmd)

//	msg := make([]byte, 1024)
//	websocket.Message.Receive(ws, &msg)
//	 Println(string(msg[7:]))
//}

func TestUserGroups(t *testing.T) {
	once.Do(startServer)

	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}

	m := simplejson.New()
	m.Set(JSON_USERID, 1)
	b, err := m.Encode()
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}
	cmd := getUserGroups + string(b)
	Println(cmd)
	websocket.Message.Send(ws, cmd)

	msg := make([]byte, 1024)
	websocket.Message.Receive(ws, &msg)
	Println(string(msg[7:]))
}

//func TestGroupInfo(t *testing.T) {
//	once.Do(startServer)

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	m := simplejson.New()
//	m.Set(JSON_GROUPID, 1)
//	b, err := m.Encode()
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}
//	cmd := getGroupInfo + string(b)
//	Println(cmd)
//	websocket.Message.Send(ws, cmd)

//	msg := make([]byte, 1024)
//	websocket.Message.Receive(ws, &msg)
//	Println(string(msg[7:]))
//}

//func TestSendGroupMsg(t *testing.T) {
//	once.Do(startServer)

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	msg := `{"USERID":1,"TOUSERID":8,"COLOR":"#000000","CONTENT":"[[{\"INFO\":\"sdadsav\",\"TYPE\":2}]]","FONT":"Microsoft YaHei,9,-1,5,50,0,0,0,0,0","MSGID":"d112433d-a1c6-4c97-a2d9-2f5768c61a09","TIME":"2014-07-09 15:47:14","ADJUSTTIME":"2014-07-09 15:47:13"}`
//	m := simplejson.New()
//	m.Set(JSON_GROUPID, 1)
//	m.Set(JSON_USERID, 1)
//	m.Set(JSON_CONTENT, msg)
//	b, _ := m.Encode()
//	mmsg := ByteJoin([]byte(sendGroupTxtMsg), b)
//	websocket.Message.Send(ws, mmsg)

//	remsg := make([]byte, 1024)
//	websocket.Message.Receive(ws, &remsg)
//	Println(string(remsg[7:]))
//}

//func TestGetUnreadGroupMsg(t *testing.T) {

//	once.Do(startServer)

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	m := simplejson.New()
//	m.Set(JSON_GROUPID, 1)
//	m.Set(JSON_USERID, 5)

//	b, _ := m.Encode()
//	mmsg := ByteJoin([]byte(getUnreadGroupMsg), b)

//	websocket.Message.Send(ws, mmsg)

//	remsg := make([]byte, 1024)
//	websocket.Message.Receive(ws, &remsg)
//	Println(string(remsg[7:]))
//}

//func TestAddUserToGroup(t *testing.T) {
//	once.Do(startServer)

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	m := simplejson.New()
//	m.Set(JSON_GROUPID, 1)
//	m.Set(JSON_USERID, 1)
//	uids := []int64{4, 5}
//	m.Set(JSON_USERIDS, uids)

//	b, _ := m.Encode()
//	mmsg := ByteJoin([]byte(addUsersToGroup), b)

//	websocket.Message.Send(ws, mmsg)

//	remsg := make([]byte, 1024)
//	websocket.Message.Receive(ws, &remsg)
//	Println(string(remsg[7:]))
//}

//func TestSearchUserByEmail(t *testing.T) {
//	once.Do(startServer)

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}

//	m := simplejson.New()
//	m.Set(JSON_EMAIL, "zhao@163.com")

//	b, _ := m.Encode()
//	mmsg := ByteJoin([]byte(searchUserByEmail), b)

//	websocket.Message.Send(ws, mmsg)

//	remsg := make([]byte, 1024)
//	websocket.Message.Receive(ws, &remsg)
//	Println(string(remsg[7:]))
//}

func TestRTMPURL(t *testing.T) {
	once.Do(startServer)

	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}

	m := simplejson.New()
	m.Set(JSON_USERID, 1)

	b, _ := m.Encode()
	mmsg := ByteJoin([]byte(getRTMPURL), b)

	websocket.Message.Send(ws, mmsg)

	remsg := make([]byte, 1024)
	websocket.Message.Receive(ws, &remsg)
	Println(string(remsg[7:]))
}
