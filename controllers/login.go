package controllers

import (
	"app/models"
	"fmt"
	"net"
	"time"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	var ip string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				ip = ipnet.IP.String()
			}
		}
	}

	msg := models.Message{Message: time.Now().Format("00-01-01 15:04:05") + " : user login ip is " + ip}
	broadcast <- msg
	c.TplName = "login.html"
}
