package main

import (
	. "commonPackage"
	"fmt"
	"log"
	"net/http"
	"service/db"
	"strconv"
	"sync"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
)

const (
	ConnectionMax = 100 //cs max Connection
)

var (
	cmdLeg = 7           //命令长度
	bufLeg = 1024 * 1024 //字节长度
)
var G map[int64]*GMeeting

type Client struct {
	Email      string
	Addr       string
	Conn       *websocket.Conn
	CurMeeting *GMeeting `json:"-"`
	RoomPem    []byte
	UserInfo   []byte
}

type GMeeting struct {
	sync.Mutex
	EmailClients map[string]*Client
	IpClients    map[string]*Client
	MaxClients   int
	CurrentId    int64
	MeetingId    int64
	GeModels     map[string]*simplejson.Json `json:"-"`
	UserPower    map[int64]*simplejson.Json  `json:"-"`
	MeetingFiles map[int64]*simplejson.Json  `json:"-"`
}

func NewGMeetings(id int64, max int) *GMeeting {
	if G == nil {
		G = make(map[int64]*GMeeting, ConnectionMax)
	}

	meeting := new(GMeeting)
	meeting.EmailClients = make(map[string]*Client, max)
	meeting.IpClients = make(map[string]*Client, max)
	meeting.GeModels = make(map[string]*simplejson.Json, max)
	meeting.MaxClients = max
	meeting.MeetingId = id
	db.RedisObjSet(strconv.FormatInt(id, 10), meeting)
	G[id] = meeting
	return meeting
}

func (g *GMeeting) AddClient(email string, userInfo []byte, ws *websocket.Conn) *Client {
	//if len(g.EmailClients) > g.MaxClients {
	//	return nil
	//}
	g.Lock()
	g.CurrentId++
	c := new(Client)
	c.Email = email
	c.Conn = ws
	c.Addr = ws.RemoteAddr().String()
	c.CurMeeting = g
	c.UserInfo = userInfo
	g.EmailClients[c.Email] = c
	g.IpClients[c.Addr] = c
	db.RedisObjSet(strconv.FormatInt(g.MeetingId, 10), g)
	g.Unlock()

	return c
}

func (g *GMeeting) RemoveClient(email string) bool {
	c := g.EmailClients[email]
	if c == nil {
		return false
	}
	g.Lock()
	delete(g.EmailClients, c.Email)
	delete(g.IpClients, c.Addr)

	g.Unlock()
	return true
}

func (g *GMeeting) RemoveClientAddr(clientAddr string) bool {
	c := g.IpClients[clientAddr]
	if c == nil {
		return false
	}
	g.Lock()
	delete(g.EmailClients, c.Email)
	delete(g.IpClients, c.Addr)
	g.Unlock()
	return true
}

func (g *GMeeting) BoardMsg(ws *websocket.Conn, msg []byte) {
	clientAddr := ws.RemoteAddr().String()
	c := g.IpClients[clientAddr]
	if c == nil {
		return
	}
	g.Lock()
	addStr := ws.RemoteAddr().String()
	fmt.Print("开始进行广播")
	fmt.Print(addStr)
	for _, v := range g.IpClients {
		fmt.Println(v.Addr)
		if v.Addr != addStr {
			Println(v.Addr)
			Println(v.Email)
			v.Conn.WriteMessage(websocket.BinaryMessage, msg)
		}
	}
	g.Unlock()
}

func (g *GMeeting) SynGeModels(gm geModel) {
	g.Lock()
	gem, has := g.GeModels[gm.GeID]

	if has { //如果存在做merge
		MergeJson(gm.CONTENT, gem)
		g.GeModels[gm.GeID] = gem
	} else { //如果不存在，就直接添加
		g.GeModels[gm.GeID], _ = simplejson.NewJson(gm.CONTENT)
	}
	g.Unlock()
}

func (g *GMeeting) SynUserPower(up userPower) {
	g.Lock()
	upm, has := g.UserPower[up.POWERID]

	if has { //如果存在做merge
		MergeJson(up.CONTENT, upm)
		g.UserPower[up.POWERID] = upm
	} else { //如果不存在，就直接添加
		g.UserPower[up.POWERID], _ = simplejson.NewJson(up.CONTENT)
	}
	g.Unlock()
}

func (g *GMeeting) SynMeetingFileInfo(mf meetingFileInfo) {
	g.Lock()
	mfm, has := g.MeetingFiles[mf.FILEID]

	if has { //如果存在做merge
		MergeJson(mf.CONTENT, mfm)
		g.MeetingFiles[mf.FILEID] = mfm
	} else { //如果不存在，就直接添加
		g.MeetingFiles[mf.FILEID], _ = simplejson.NewJson(mf.CONTENT)
	}
	g.Unlock()
}

type myServer struct {
}

func (ms myServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  bufLeg,
		WriteBufferSize: bufLeg,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}
	WsServer(conn)
}

func main() {
	fmt.Printf("Welcome lcsoft xt server!")
	h := new(myServer)

	err := http.ListenAndServe(":7000", h)
	if err != nil {
		fmt.Printf("ListenAndServe: " + err.Error())
	}

}
