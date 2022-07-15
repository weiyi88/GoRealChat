package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 定义一个upgrade 配置，需要定义读写内存
var upgrader = websocket.Upgrader{
	// 读写内存
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// 检查链来源
	// 允许react 服务向这里发出请求
	// 不需要检查并运行任务链接
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 定义一个reader 监听往ws发送的新消息
func reader(conn *websocket.Conn) {
	for {
		// 读消息
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))	//打印读到的消息
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// 定义一个Websocket 服务处理函数
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// 将链接更新为websocket 链接
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// 一直监听Websoket 链接上传来的新消息
	reader(ws)

}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "http 链接成功")
	})

	// 将/ws 端点交给 serveWs 函数处理
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":3100", nil)
}
