package controllers

import (
	"commonPackage"
	"fmt"
	"github.com/qiniu/api/io"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

type ChangeheadController struct {
	BaseController
}

func (h *ChangeheadController) Get() {
	if h.userinfo != nil {
		h.Data["userinfo"] = h.userinfo

		h.Data["doctotal"] = 0

		h.TplNames = "pages/home/changehead.html"
	}

}

func (h *ChangeheadController) Post() {
	if h.userinfo != nil {
		top, _ := h.GetInt("top")
		left, _ := h.GetInt("left")
		right, _ := h.GetInt("right")
		bottom, _ := h.GetInt("bottom")
		file, h1, _ := h.GetFile("file") //获取上传的文件
		path := "./static/img/" + h1.Filename
		file.Close() // 关闭上传的文件，不然的话会出现临时文件不能清除的情况

		h.SaveToFile("file", path) //存文件
		f1, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f1.Close()
		m1, err := jpeg.Decode(f1)
		if err != nil {
			panic(err)
		}
		bounds := m1.Bounds()

		m := image.NewRGBA(bounds)
		draw.Draw(m, image.Rect(int(top), int(left), int(right), int(bottom)), m1, image.Pt(250, 60), draw.Src)
		imgfile, err := os.Create("./static/img/IMG_2718111.JPG")
		defer imgfile.Close()
		err = jpeg.Encode(imgfile, m, &jpeg.Options{90})
		fmt.Println("00000000")
		fmt.Println(err)
		fmt.Println("00000000")
		ret := new(io.PutRet)
		//err := io.PutFile(nil, ret, commonPackage.Uptoken(), h1.Filename, path, nil)

		durl := commonPackage.DownloadUrl(h1.Filename)
		fmt.Println(durl)
		fmt.Println(ret)
		h.Redirect("/changehead", 302)

		h.Data["userinfo"] = h.userinfo

		h.Data["doctotal"] = 0

	}

}
