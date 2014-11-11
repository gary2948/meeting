package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"github.com/gorilla/websocket"
	//	"unicode/utf8"
	"commonPackage"
	"service/db"
)

const (
	ConnectionMax = 1000 //cs max Connection
)

var (
	cmdLeg = 7               //命令长度
	bufLeg = 4 * 1024 * 1024 //字节长度
)
var G *Group

type Client struct {
	Id       int64
	Addr     string
	Conn     *websocket.Conn
	CurGroup *Group
}

type Group struct {
	sync.Mutex
	IdClients  map[int64]*Client
	IpClients  map[string]*Client
	MaxClients int
	CurrentId  int
}

func NewGroup(max int) *Group {
	group := new(Group)
	group.IdClients = make(map[int64]*Client, max)
	group.IpClients = make(map[string]*Client, max)
	group.MaxClients = max
	return group
}

func (g *Group) AddClient(userId int64, ws *websocket.Conn) *Client {
	if len(g.IdClients) > g.MaxClients {
		return nil
	}
	g.Lock()
	g.CurrentId++
	c := new(Client)
	c.Id = userId
	c.Conn = ws
	c.Addr = ws.RemoteAddr().String()
	c.CurGroup = g
	g.IdClients[c.Id] = c
	g.IpClients[c.Addr] = c
	g.Unlock()
	return c
}

func (g *Group) RemoveClient(clientId int64) bool {
	c := g.IdClients[clientId]
	if c == nil {
		return false
	}
	g.Lock()
	delete(g.IdClients, c.Id)
	delete(g.IpClients, c.Addr)
	g.Unlock()
	return true
}

func (g *Group) RemoveClientAddr(clientAddr string) bool {
	c := g.IpClients[clientAddr]
	if c == nil {
		return false
	}
	g.Lock()
	delete(g.IdClients, c.Id)
	delete(g.IpClients, c.Addr)
	g.Unlock()
	return true
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

func ping(msg string) error {
	commonPackage.Println("ping:" + msg)
	return nil
}

func pong(msg string) error {
	commonPackage.Println("pong:" + msg)
	return nil
}

func callback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	_, err := db.AddFileMap(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}

func main() {
	fmt.Printf("Welcome lcsoft server!")
	G = NewGroup(ConnectionMax)
	h := new(myServer)
	http.HandleFunc("/callback", callback)
	http.HandleFunc("/ws", h.ServeHTTP)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Printf("ListenAndServe: " + err.Error())
	}
}
