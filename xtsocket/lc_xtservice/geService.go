package main

//图源对象处理
import (
	. "commonPackage"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
)

//同步图元对象
func SynGeModel(ws *websocket.Conn, body []byte) {
	var gm geModel
	JSONUnmarshal(body, &gm)
	m, has := G[gm.MEETINGID]
	if has {
		m.SynGeModels(gm)
		bf := ByteJoin([]byte(boardGeModel), body)
		m.BoardMsg(ws, bf)
	}
	rm := simplejson.New()
	rm.Set(JSON_OK, true)
	rmb, _ := rm.MarshalJSON()
	remsg := ByteJoin([]byte(reSynGeModel), rmb)
	ws.WriteMessage(websocket.BinaryMessage, remsg)
}
