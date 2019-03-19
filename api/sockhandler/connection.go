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
	bands:      make(map[int]*band),
	register:   make(chan *subscription),
	unregister: make(chan *subscription),
}

func (subscription *subscription) handleIncomingMessage(message *Message) {
	message.BandID = subscription.bandID

	switch connType := message.Type; connType {
	case "toggleNote":

		break
	default:
		fmt.Println("default")
	}
}

func upgradeConnection(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	subscription := subscription{connection: socket}
	mHub.register <- &subscription

	defer func() {
		mHub.unregister <- &subscription
	}()
	for {
		var message Message
		if err := subscription.connection.ReadJSON(&message); err != nil {
			log.Println(err)
			return
		}
		subscription.handleIncomingMessage(&message)
	}
}

func StartConnection() {
	go mHub.run()
	http.HandleFunc("/", upgradeConnection)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
