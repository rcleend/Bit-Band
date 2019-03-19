package sockhandler

import (
	"github.com/gorilla/websocket"
)

type subscription struct {
	connection *websocket.Conn
	bandID     int
}

type hub struct {
	bands      map[int]*band
	register   chan *subscription
	unregister chan *subscription
}

func (hub *hub) registerSubscription(subscription *subscription) {
	for _, band := range hub.bands {
		if len(band.connections) < len(possibleInstruments) {
			subscription.bandID = band.id
			band.addConnection(subscription)
			sendNewInstrumentMessage(hub, subscription)
			return
		}
	}
	newBand := createNewBand()
	subscription.bandID = newBand.id
	hub.bands[newBand.id] = &newBand
	newBand.connections[subscription.connection] = possibleInstruments[0]
	sendNewInstrumentMessage(hub, subscription)
}

func (hub *hub) unregisterSubscription(subscription *subscription) {
	band := hub.bands[subscription.bandID]

	delete(band.connections, subscription.connection)
	if len(band.connections) == 0 {
		delete(hub.bands, subscription.bandID)
	}

	sendRemoveInstrumentMessage(hub, subscription)
}

func (hub *hub) run() {
	for {
		select {
		case subscription := <-hub.register:
			hub.registerSubscription(subscription)

		case subscription := <-hub.unregister:
			hub.unregisterSubscription(subscription)
		}
	}
}
