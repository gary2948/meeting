package main

import (
	"github.com/astaxie/beego"
	_ "website/routers"
)

func main() {
	beego.SessionOn = true
	beego.Run()
}
