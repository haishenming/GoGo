package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/echo", echo)    // 设置路由
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, _ := upgrader.Upgrade(w, r, nil)
	defer c.Close()

	for {
		mt, message, _ := c.ReadMessage()
		c.WriteMessage(mt, append([]byte("hello "), message[:]...))
	}


}



