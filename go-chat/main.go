package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// a simple chat server
// commands:
//   login:name
//   send:recepient:message
//   logout

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "home.html")
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err.Error())
		}

		for {
			log.Print("reading message....")
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				return
			}
			if messageType == websocket.TextMessage {
				log.Printf("%v", string(p))
			}
		}

	})

	quit := make(chan int)
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("failed to start server, Err=%v", err.Error())
			quit <- 0
			return
		}
	}()
	log.Print("Server started... press ctrl+c to quit")
	<-quit
}
