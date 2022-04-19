package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	greet := "<h1> Welcome to my go websocket echo server </h1>"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, greet)
}

var upgrader = websocket.Upgrader{}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } //never blindly trust any origin. But since test ok
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Ws connection error", err)
	}
	defer ws.Close()
	for { //make this a goroutine...
		_, bytes, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Client has disconnected") //lifecycle event
			return
			//handleDisconnection()
		}
		//msg:= string(bytes)
		ws.WriteMessage(websocket.TextMessage, bytes) //handleIncomingMessage(websocket.TextMessage,msg)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/v1/ws", socketHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
