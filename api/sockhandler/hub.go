package sockhandler

import (
	"fmt"
	"math/rand"
)

var bandSize = 4

type Subscription struct {
	connection *Connection
	band       string
}

type hub struct {
	bands      map[string]map[*Connection]bool
	register   chan *Subscription
	unregister chan *Subscription
	broadcast  chan *Message
}

//TODO: implement real band name generator API
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func getNewRoom() string {
	return randSeq(5)
}

func (hub *hub) run() {
	for {
	OuterLoop:
		select {
		case subscription := <-hub.register:
			// Add connection if a room has a place left
			for room, connections := range hub.bands {
				if len(connections) < bandSize {
					subscription.room = room
					connections[subscription.connection] = true
					fmt.Printf("adding in bands: %v\n", hub.bands)
					// Return to the outer loop when a new connection has been added
					goto OuterLoop
				}
			}
			// Add a new room with a new connection if no other room is available
			room := getNewRoom()
			subscription.room = room
			connections := hub.bands
			connections[room] = make(map[*Connection]bool)
			connections[room][subscription.connection] = true
			fmt.Printf("new room created: %v\n", room)
			fmt.Printf("adding in bands: %v\n", hub.bands)

		case subscription := <-hub.unregister:
			connections := hub.bands[subscription.room]
			delete(connections, subscription.connection)
			fmt.Printf("Amount of connections: %v\n", len(connections))

			if len(connections) == 0 {
				delete(hub.bands, subscription.room)
				fmt.Printf("delete room: %v\n", subscription.room)
			}
			fmt.Printf("Deleting in bands: %v\n", hub.bands)
		case message := <-hub.broadcast:
			fmt.Println(message)
			// 1. Get all connections of that specific room
			// 2. loop through all connections of that specific room
			// 3. Send message to all connections
		}
	}
}
