package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Echo(r http.ResponseWriter, req *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(r, req, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			var mainSend Main
			json.Unmarshal(msg, &mainSend)
			switch mainSend.Type {
			case "LOGIN":
				var LoginDetails Login

				json.Unmarshal(msg, &LoginDetails)

				fmt.Print(LoginDetails)
			}
		}
	}()
}

func main() {
	http.HandleFunc("/timer", Echo)
	http.ListenAndServe(":8080", nil)
}
