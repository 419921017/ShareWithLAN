package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var wsupgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wshandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte)}
	client.hub.register <- client

	go client.writeDump()
	go client.readDump()
}

func HttpController(ctx *gin.Context, hub *Hub) {
	wshandler(hub, ctx.Writer, ctx.Request)
}
