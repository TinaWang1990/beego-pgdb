package main

import (
	_ "app/models"
	_ "app/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
