package main

import (
	"log"
	"net/http"

	"github.com/bhanuprasad1999/server/connection"
)


func main() {
	http.HandleFunc("/ws", connection.HandleWebSocket)

	log.Println("Websocket server running on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Listen and serve: ", err)
	}

}
