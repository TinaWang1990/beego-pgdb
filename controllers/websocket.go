package controllers

import (
	"fmt"
	"log"
	"time"

	"app/models"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type WebsocketController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{}

func (c *WebsocketController) Get() {
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	clients[ws] = true

	for {
		time.Sleep(time.Second * 3)

		var msg models.Message // Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("page is broke ws.ReadJSON error: %v", err)
			delete(clients, ws)
			break
		} else {
			fmt.Println("received: ", msg.Message)
		}

		broadcast <- msg
	}
}
