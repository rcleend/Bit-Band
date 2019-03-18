package sockhandler

import (
	"fmt"
)

type Subscription struct {
	connection *Connection
	band       string
	instrument string
}

type hub struct {
	bands      map[string]*Band
	register   chan *Subscription
	unregister chan *Subscription
	broadcast  chan *Message
}

func sendUpdateInstrumentMessage(band *Band) {
	updateInstrument := Message{
		Type: "updateInstrument",
		Data: make(map[string]interface{}),
		Band: band.name,
	}

	updateInstrument.Data["newInstrument"] = band.instruments[len(band.instruments)-1]
	updateInstrument.Data["usedInstruments"] = band.instruments

	for connection, _ := range band.connections {
		connection.connection.WriteJSON(updateInstrument)
	}
}

func (hub *hub) run() {
	for {
		select {
		case subscription := <-hub.register:
			register(subscription, hub)

		case subscription := <-hub.unregister:
			unregister(subscription, hub)

		case message := <-hub.broadcast:
			fmt.Println(message)
			// 1. Get all connections of that specific room
			// 2. loop through all connections of that specific room
			// 3. Send message to all connections
		}
	}
}

func register(subscription *Subscription, hub *hub) {
	for bandName, band := range hub.bands {
		if len(band.connections) < len(possibleInstruments) {
			subscription.band = bandName
			band.addConnection(subscription)
			return
		}
	}
	newBand := CreateNewBand()
	subscription.band = newBand.name
	hub.bands[newBand.name] = &newBand
	newBand.connections[subscription.connection] = possibleInstruments[0]
	sendUpdateInstrumentMessage(&newBand)
}

func unregister(subscription *Subscription, hub *hub) {
	band := hub.bands[subscription.band]
	oldInstrument := band.connections[subscription.connection]
	for index, instrument := range band.instruments {
		if instrument == oldInstrument {
			band.instruments = append(band.instruments[:index], band.instruments[index+1:]...)
			break
		}
	}

	delete(band.connections, subscription.connection)
	if len(band.connections) == 0 {
		delete(hub.bands, subscription.band)
	}
}
