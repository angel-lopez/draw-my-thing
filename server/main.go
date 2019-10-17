package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func loopMessage(connections *[]*websocket.Conn) {
	for i, conn := range *connections {
		conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprint("Hello connection:", i)))
	}

	time.Sleep(2 * time.Second)
	loopMessage(connections)
}

func main() {
	connections := []*websocket.Conn{}

	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		conn.SetCloseHandler(func(code int, text string) error {
			fmt.Println("connection closed")
			return nil
		})

		// for {
		// 	messageType, p, err := conn.ReadMessage()
		// 	if err != nil {
		// 		log.Println(err)
		// 		return
		// 	}
		// 	if err := conn.WriteMessage(messageType, p); err != nil {
		// 		log.Println(err)
		// 		return
		// 	}
		// }

		connections = append(connections, conn)
	})

	go loopMessage(&connections)

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
