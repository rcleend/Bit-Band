package sockhandler

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

var possibleInstruments = []string{"Drum", "Bass", "Rhythm", "Lead"}

type Subscription struct {
	connection *Connection
	room       string
}

type hub struct {
	rooms      map[string]map[*Connection]string
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

func getAvailableInstrument(connections map[*Connection]string) (string, error) {
	for _, possibleInstrument := range possibleInstruments {
		index := 0
		for _, usedInstrument := range connections {
			if usedInstrument == possibleInstrument {
				fmt.Printf("%v is already used. breaking..\n", possibleInstrument)
				break
			}

			if index == len(connections)-1 && usedInstrument != possibleInstrument {
				fmt.Printf("Instrument not yet used. New instrument: %v\n", possibleInstrument)
				return possibleInstrument, nil
			}
			index++
		}
	}
	return "No available instrument", errors.New("Something went wrong. No available instruments")
}

func sendUpdateInstrumentMessage(subscription Subscription) {
	updateInstrument := Message{
		Type: "updateInstrument",
		Data: make(map[string]interface{}),
		Room: subscription.room,
	}
	updateInstrument.Data["instrument"] = mHub.rooms[subscription.room][subscription.connection]
	subscription.connection.connection.WriteJSON(updateInstrument)
}

func (hub *hub) run() {
	for {
	OuterLoop:
		select {
		case subscription := <-hub.register:
			// Add connection if a room has a place left
			for room, connections := range hub.rooms {
				if len(connections) < len(possibleInstruments) {
					subscription.room = room
					if instrument, err := getAvailableInstrument(connections); err != nil {
						log.Fatal(err)
						return
					} else {
						connections[subscription.connection] = instrument
						sendUpdateInstrumentMessage(*subscription)
					}
					goto OuterLoop
				}
			}
			// Add a new room with a new connection if no other room is available
			room := getNewRoom()
			subscription.room = room
			connections := hub.rooms
			connections[room] = make(map[*Connection]string)
			connections[room][subscription.connection] = possibleInstruments[0]
			sendUpdateInstrumentMessage(*subscription)

		case subscription := <-hub.unregister:
			connections := hub.rooms[subscription.room]
			delete(connections, subscription.connection)
			// Delete room if the room is emtpy
			if len(connections) == 0 {
				delete(hub.rooms, subscription.room)
			}

		case message := <-hub.broadcast:
			fmt.Println(message)
			// 1. Get all connections of that specific room
			// 2. loop through all connections of that specific room
			// 3. Send message to all connections
		}
	}
}
