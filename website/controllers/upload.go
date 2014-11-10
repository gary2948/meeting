package controllers

import (
	"bytes"
	//"commonPackage/model/account"
	"encoding/binary"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"os"
	//"service/db"
	"website/models"
)

type WebSocketController struct {
	beego.Controller
	ReceiveLength int
}

func (this *WebSocketController) Get() {
	filename := this.GetString("FileName")
	folderID, _ := this.GetInt("Folderid")
	var content []byte
	beego.Info(folderID)
	beego.Info(filename)
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	buff := make([]byte, 1024)
	if err == nil {
		for {
			_, read, err := ws.NextReader()
			if err == nil {
				file, fileerr := os.OpenFile(filename, os.O_CREATE, 0660)
				if fileerr != nil {
					beego.Info(fileerr)
					break
				}
				for {
					count, _ := read.Read(buff)
					if count == 0 {

						err = file.Close()
						if err != nil {
							beego.Info(err)
						}
						beego.Info(filename + "closed")
						goto disposeFile
						break
					}
					file.Write(buff[0:count])
					this.ReceiveLength += count
					//beego.Info(this.ReceiveLength)
					content, _ = json.Marshal(models.UploadFileResponseModel{Code: models.UploadingPercent, Data: this.ReceiveLength})
					ws.WriteMessage(websocket.TextMessage, content)
				}
			} else {
				break
			}
		}
	} else {

	}
	defer ws.Close()
	//////////////////////向数据库中插入文件数据\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
disposeFile:
	//userinfo, _ := this.GetSession("Lctb_userInfo").(*account.Lctb_userInfo)
	//_, err = db.AddFile(userinfo.Id, folderID, filename, 0)
	////////////////////////////////////////////////////////////
	content, _ = json.Marshal(models.UploadFileResponseModel{Code: models.Finished, Data: "receive data insert into db"})
	err = ws.WriteMessage(websocket.TextMessage, content)
	beego.Info(err)

	this.EnableRender = false
	return
}

func (this *WebSocketController) Post() {

}

func int2bytes(value int32) []byte {
	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, value)
	return b_buf.Bytes()
}

func int2string(value int32) string {
	return string(value)
}
func string2bytes(value string) []byte {
	return []byte(value)
}
