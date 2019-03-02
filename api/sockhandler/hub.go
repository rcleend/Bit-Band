package sockhandler

import (
	"fmt"
	"math/rand"
)

var roomsize = 4

type Subscription struct {
	connection *Connection
	room       string
}

type hub struct {
	rooms      map[string]map[*Connection]bool
	register   chan *Subscription
	unregister chan *Subscription
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
			for room, connections := range hub.rooms {
				if len(connections) <= roomsize {
					subscription.room = room
					connections[subscription.connection] = true
					fmt.Printf("Amount of connections: %v\n", len(connections))
					// Return to the outer loop when a new connection has been added
					goto OuterLoop
				}
			}
			// Add a new room with a new connection if no other room is available
			room := getNewRoom()
			subscription.room = room
			connections := hub.rooms
			connections[room] = make(map[*Connection]bool)
			connections[room][subscription.connection] = true
			fmt.Printf("new room created: %v\n", room)
			fmt.Printf("Amount of connections: %v\n", len(connections))

		case subscription := <-hub.unregister:
			// 1. if the connection is the last in the room: delete room
			// 2. else just delete the collection from the room
			connections := hub.rooms[subscription.room]
			delete(connections, subscription.connection)
			fmt.Printf("Amount of connections: %v\n", len(connections))

			if len(connections) == 0 {
				delete(hub.rooms, subscription.room)
				fmt.Printf("delete room: %v", subscription.room)
			}
		}
	}
}
