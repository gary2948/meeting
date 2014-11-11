package main

import (
	"commonPackage"
	"fmt"
	"log"
	"net"
	"net/http/httptest"
	"service/uuid"
	"sync"
	"testing"

	"code.google.com/p/go.net/websocket"
)

var (
	serverAddr string
	once       sync.Once
)

//var origin = "http://localhost/"
//var url = "ws://localhost:8001/"

func startServer() {
	//G = NewGroup(ConnectionMax)
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

func TestStartMeeting(t *testing.T) {
	once.Do(startServer)

	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}
	msg := make([]byte, 512)
	cmd := "6001000" + `{"MEETINGID":123}`

	ws.Write([]byte(cmd))
	_, err = ws.Read(msg)
	if err != nil {
		t.Error(err)
	}

	commonPackage.Printb(msg)

}

func TestLoginMeeting(t *testing.T) {
	once.Do(startServer)

	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}
	msg := make([]byte, 512)
	cmd := "6002000" + `{"MEETINGID":123, "EMAIL":"zhao@163.com"}`

	ws.Write([]byte(cmd))
	_, err = ws.Read(msg)
	if err != nil {
		t.Error(err)
	}

	commonPackage.Printb(msg)
}

//func TestMouseMove(t *testing.T) {
//	once.Do(startServer)

//	ws, err := newConn(t)
//	if err != nil {
//		t.Errorf("WebSocket handshake error: %v", err)
//		return
//	}
//	msg := make([]byte, 512)
//	cmd := `6004000{"CONTENT":{"X":908,"Y":1008},"EMAIL":"zileyuan@QQ.com","MEETINGID":123}`

//	ws.Write([]byte(cmd))
//	_, err = ws.Read(msg)
//	if err != nil {
//		t.Error(err)
//	}

//	commonPackage.Printb(msg)
//}

func TestSynGe(t *testing.T) {
	ws, err := newConn(t)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}

	for i := 0; i < 10; i++ {
		gid, _ := uuid.GenUUID4()
		cmd := `6006000{"CONTENT":{"X":908,"Y":1008},"EMAIL":"zileyuan@QQ.com","MEETINGID":123,"GeID":"` + gid + `"}`
		ws.Write([]byte(cmd))
		msg := make([]byte, 512)
		_, err = ws.Read(msg)
		if err != nil {
			t.Error(err)
		}

		commonPackage.Printb(msg)
	}

}
