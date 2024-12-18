package connection

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/bhanuprasad1999/server/services"
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to websocket", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message", err)
			break
		}
		var data map[string]interface{}
		json.Unmarshal(message, &data)

		fmt.Println(data)
		if data["type"] == "channel"{
			services.EditChannel(data["channelName"].(string), data["username"].(string))
		}else if data["type"] == "message"{
			services.AddMessage(data["message"].(string), data["username"].(string), data["channelName"].(string))
		}

		fmt.Printf("Recevied: %s\n", message)

	}
}
