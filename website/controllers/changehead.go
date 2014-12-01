package controllers

import (
	"code.google.com/p/graphics-go/graphics"
	"commonPackage"
	"commonPackage/viewModel"
	"fmt"
	"github.com/qiniu/api/io"
	"image"
	"image/png"
	"os"
	"service/db"
	"strconv"
	"website/utils"
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
		//rotation := h.GetString("rotation")
		fmt.Println(top)
		fmt.Println(left)
		fmt.Println(right)
		fmt.Println(bottom)

		now := utils.GetRandFileName()

		file, _, _ := h.GetFile("file") //获取上传的文件
		path := "./static/img/head/" + now + ".png"
		file.Close() // 关闭上传的文件，不然的话会出现临时文件不能清除的情况

		h.SaveToFile("file", path) //存文件

		//缩放
		src, err := LoadImage(path)
		if err != nil {
			panic(err)
		}
		dst := image.NewRGBA(image.Rect(0, 0, 280, 280)) //1200x800 tumbnail
		err = graphics.Scale(dst, src)
		graphics.Scale(dst, src)

		if err != nil {
			panic(err)
		}
		// 需要保存的文件
		savepath := "./static/img/head/" + now + ".png"
		saveImage(fmt.Sprintf(savepath), dst)

		/*//裁剪
		src, err = LoadImage(savepath)
		if err != nil {
			panic(err)
		}
		dst = image.NewRGBA(image.Rect(int(top), int(left), int(right)-int(left), int(bottom)-int(top)))
		rotationf, _ := strconv.ParseFloat(rotation, 64)
		err = graphics.Rotate(dst, src, &graphics.RotateOptions{rotationf})
		if err != nil {
			panic(err)
		}
		// 需要保存的文件
		//saveImage(fmt.Sprintf(savepath), dst)
		*/
		ret := new(io.PutRet)
		err = io.PutFile(nil, ret, commonPackage.Uptoken("meetinwareimgs", now), now, path, nil)
		if err != nil {
			panic(err)
		}

		durl := "http://meetinwareimgs.qiniudn.com/" + now + "?imageMogr2/crop/!" + strconv.Itoa(int(right)-int(left)) + "x" + strconv.Itoa(int(bottom)-int(top)) + "a" + strconv.Itoa(int(left)) + "a" + strconv.Itoa(int(top))
		h.userinfo.Lc_photoFile = durl
		var vm2 viewModel.EditUserInfoModel
		vm2.Lc_photoFile = durl
		err = db.UpdateAccount(h.userinfo.Id, &vm2)
		if err != nil {
			panic(err)
		}

		h.Redirect("/changehead", 302)
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
