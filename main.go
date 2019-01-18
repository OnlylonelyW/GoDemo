package main

import (
	_ "searchq-operationsys/routers"
	"github.com/astaxie/beego"
	_ "searchq-operationsys/controllers"
)

func main() {
	beego.Run()
}

