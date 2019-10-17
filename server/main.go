package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/angel-lopez/draw-my-thing/game"
	"github.com/gorilla/websocket"
)

type user struct {
	Conn   *websocket.Conn
	Player *game.Player
}

func main() {
	g := game.Game{}
	users := []user{}

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

		player := g.Join()
		u := user{conn, player}
		users = append(users, u)

		if len(users) >= 2 {
			g.StartNewRound("cat", users[0].Player)
		}

		go func() {
			for {
				messageType, p, err := conn.ReadMessage()
				if err != nil {
					log.Println(err)
					return
				}
				fmt.Println(string(p))
				if _, err = player.Guess(string(p)); err != nil {
					if err = conn.WriteMessage(messageType, []byte(err.Error())); err != nil {
						fmt.Println(err)
					}
				}
				fmt.Println(err)
			}
		}()

		conn.SetCloseHandler(func(code int, text string) error {
			fmt.Println("connection closed")
			return nil
		})
	})

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
