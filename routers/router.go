package routers

import (
	"APP/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/ws", &controllers.WebsocketController{})
	beego.Router("/login", &controllers.LoginController{})
}
