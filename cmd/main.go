package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var allClients map[*Client]int

type Client struct {
	conn websocket.Conn
}

func newClient(connection websocket.Conn) *Client {
	client := &Client{
		conn: connection,
	}
	return client
}

func handle(ws *websocket.Conn) {
	var err error
	client := newClient(*ws)
	allClients[client] = 1

	for {
		var reply string

		if err = websocket.Message.Receive(&client.conn, &reply); err != nil {
			fmt.Println("Can't receive")
			delete(allClients, client)
			break
		}

		// Handle reply JSON
		var replyMap map[string]interface{}
		json.Unmarshal([]byte(reply), &replyMap)

		switch replyMap["type"] {
		case "update":
			updatePlayer(replyMap)
		case "delete":
		default:
			fmt.Println("default")
		}

		// Send reply to all clients
		for clientList, _ := range allClients {
			if err = websocket.Message.Send(&clientList.conn, reply); err != nil {
				fmt.Println("Can't send")
				break
			}
		}
		fmt.Println(len(allClients))

	}
}

func updatePlayer(replyMap map[string]interface{}) {
	fmt.Println(replyMap["data"])
}

func main() {
	allClients = make(map[*Client]int)
	http.Handle("/", websocket.Handler(handle))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
