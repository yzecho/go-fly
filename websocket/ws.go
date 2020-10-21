package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type User struct {
	conn   *websocket.Conn
	name   string
	id     string
	avator string
	to_id  string
}
type Message struct {
	conn        *websocket.Conn
	context     *gin.Context
	content     []byte
	messageType int
}
type TypeMessage struct {
	Type interface{} `json:"type"`
	Data interface{} `json:"data"`
}
type ClientMessage struct {
	Name      string `json:"name"`
	Avator    string `json:"avator"`
	Id        string `json:"id"`
	VisitorId string `json:"visitor_id"`
	Group     string `json:"group"`
	Time      string `json:"time"`
	ToId      string `json:"to_id"`
	Content   string `json:"content"`
	City      string `json:"city"`
	ClientIp  string `json:"client_ip"`
	Refer     string `json:"refer"`
}

var clientList = make(map[string]*User)
var kefuList = make(map[string][]*User)
var message = make(chan *Message)
var upgrader = websocket.Upgrader{}

func init() {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}