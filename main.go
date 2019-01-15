package main

import (
	_ "coscms.com/routers"
	"github.com/astaxie/beego"
	_ "coscms.com/controllers"
)

func main() {
	beego.Run()
}

