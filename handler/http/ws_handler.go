package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	dto "../../dto"
	"github.com/gorilla/websocket"

	redis "../../connection"
)

var upgrader = websocket.Upgrader{}

func InitWsHandler() *WsHandler {
	return &WsHandler{}
}

type WsHandler struct{}

func checkSameOrigin(r *http.Request) bool {
	return true
}

func (authHandler *WsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = checkSameOrigin
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		wsRequest := dto.WsRequest{}
		json.Unmarshal(message, &wsRequest)

		if wsRequest.Action == "subscribe" {
			pubsub := redis.GetRedis().Subscribe(wsRequest.Channel)

			_, err := pubsub.Receive()
			if err != nil {
				panic(err)
			}

			ch := pubsub.Channel()

			for msg := range ch {
				fmt.Println(msg.Channel, msg.Payload)
				c.WriteMessage(mt, []byte(msg.Payload))
			}
		}

		log.Printf("recv: %s", message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
