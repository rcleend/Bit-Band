package main

import (
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

		for clientList, _ := range allClients {
			if err = websocket.Message.Send(&clientList.conn, reply); err != nil {
				fmt.Println("Can't send")
				break
			}
		}
		fmt.Println(len(allClients))

	}
}

func main() {
	allClients = make(map[*Client]int)
	http.Handle("/", websocket.Handler(handle))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
