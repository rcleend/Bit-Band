package sockhandler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var mHub = hub{
	rooms:      make(map[string]map[*Connection]string),
	register:   make(chan *Subscription),
	unregister: make(chan *Subscription),
}

type Message struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
	Room string
}

type Connection struct {
	connection *websocket.Conn
}

func createSubscription(socket *websocket.Conn) Subscription {
	connection := Connection{
		connection: socket,
	}
	subscription := Subscription{
		connection: &connection,
	}
	return subscription
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	// Upgrade http connection to a websocket connection
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	subscription := createSubscription(socket)
	mHub.register <- &subscription
	// Unregister connection when the connection is closed
	defer func() {
		mHub.unregister <- &subscription
	}()
	for {
		var message Message
		// Assign the JSON data to the message variable if there is no error
		if err := socket.ReadJSON(&message); err != nil {
			log.Println(err)
			return
		}
		message.Room = subscription.room
		fmt.Printf("message room: %v\n", message.Room)

		switch connType := message.Type; connType {
		case "toggleNote":
			fmt.Println(connType)
		default:
			fmt.Println("default")
		}

		fmt.Println(message.Data["note"])
	}
}

func StartConnection() {
	go mHub.run()
	http.HandleFunc("/", handleConnection)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
