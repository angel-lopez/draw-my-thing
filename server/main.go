package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/angel-lopez/draw-my-thing/game"
	"github.com/gorilla/websocket"
)

var g = game.Game{}
var users = []user{}

type user struct {
	conn   *websocket.Conn
	player *game.Player
}

func (u *user) handleIncomingMessages() {
	for {
		messageType, p, err := u.conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		message := string(p)
		log.Printf("incoming message: %s\n", message)
		if message == "start" {
			g.StartNewRound("cat", users[0].player)
		} else {
			if _, err = u.player.Guess(string(p)); err != nil {
				if err = u.conn.WriteMessage(messageType, []byte(err.Error())); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

		u := user{conn, g.Join()}
		users = append(users, u)
		go u.handleIncomingMessages()
	})

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
