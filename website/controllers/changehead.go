package controllers

import (
	"code.google.com/p/graphics-go/graphics"
	"commonPackage"
	"fmt"
	"github.com/qiniu/api/io"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
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
		fmt.Println(top)
		fmt.Println(left)
		fmt.Println(right)
		fmt.Println(bottom)
		file, h1, _ := h.GetFile("file") //获取上传的文件
		path := "./static/img/" + h1.Filename
		file.Close() // 关闭上传的文件，不然的话会出现临时文件不能清除的情况

		h.SaveToFile("file", path) //存文件
		src, err := LoadImage(path)
		if err != nil {
			panic(err)
		}

		dst := image.NewRGBA(image.Rect(0, 0, 280, 280)) //1200x800 tumbnail
		err = graphics.Scale(dst, src)
		graphics.Scale(dst, src)

		//dst = image.NewRGBA(image.Rect(int(top), int(left), int(right), int(bottom)))
		//err = graphics.Rotate(dst, src, &graphics.RotateOptions{3.5})
		if err != nil {
			panic(err)
		}

		// 需要保存的文件
		imgcounter := 123
		saveImage(fmt.Sprintf("./static/img/IMG_2718tttt111.PNG", imgcounter), dst)

		h.Redirect("/changehead", 302)
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

		imgfile, err := os.Create("./static/img/IMG_2718111.png")
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

// LoadImage decodes an image from a file.
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

// 保存Png图片
func saveImage(path string, img image.Image) (err error) {
	// 需要保存的文件
	imgfile, err := os.Create(path)
	defer imgfile.Close()

	// 以PNG格式保存文件
	err = png.Encode(imgfile, img)
	if err != nil {
		panic(err)
	}
	return
}
