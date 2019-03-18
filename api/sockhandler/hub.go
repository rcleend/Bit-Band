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

func (hub *hub) run() {
	for {
		select {
		case subscription := <-hub.register:
			hub.registerSubscription(subscription)

		case subscription := <-hub.unregister:
			hub.unregisterSubscription(subscription)

		case message := <-hub.broadcast:
			fmt.Println(message)
			// 1. Get all connections of that specific room
			// 2. loop through all connections of that specific room
			// 3. Send message to all connections
		}
	}
}

func (hub *hub) registerSubscription(subscription *Subscription) {
	for _, band := range hub.bands {
		if len(band.connections) < len(possibleInstruments) {
			subscription.band = band.name
			band.addConnection(subscription)
			SendUpdateInstrumentMessage(band, subscription)
			return
		}
	}
	newBand := CreateNewBand()
	subscription.band = newBand.name
	hub.bands[newBand.name] = &newBand
	newBand.connections[subscription.connection] = Instrument{Type: possibleInstruments[0]}
	SendUpdateInstrumentMessage(&newBand, subscription)
}

func (hub *hub) unregisterSubscription(subscription *Subscription) {
	band := hub.bands[subscription.band]

	delete(band.connections, subscription.connection)
	if len(band.connections) == 0 {
		delete(hub.bands, subscription.band)
	}
}
