package sockhandler

import (
	"github.com/gorilla/websocket"
)

type subscription struct {
	connection *websocket.Conn
	band       string
}

type hub struct {
	bands      map[string]*band
	register   chan *subscription
	unregister chan *subscription
}

func (hub *hub) registerSubscription(subscription *subscription) {
	for _, band := range hub.bands {
		if len(band.connections) < len(possibleInstrumentTypes) {
			subscription.band = band.name
			band.addConnection(subscription)
			sendNewInstrumentMessage(hub, subscription)
			return
		}
	}
	newBand := createNewBand()
	subscription.band = newBand.name
	hub.bands[newBand.name] = &newBand
	newBand.connections[subscription.connection] = instrument{Type: possibleInstrumentTypes[0]}
	sendNewInstrumentMessage(hub, subscription)
}

func (hub *hub) unregisterSubscription(subscription *subscription) {
	band := hub.bands[subscription.band]

	delete(band.connections, subscription.connection)
	if len(band.connections) == 0 {
		delete(hub.bands, subscription.band)
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
