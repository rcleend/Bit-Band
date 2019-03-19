package sockhandler

import (
	"github.com/gorilla/websocket"
)

type Message struct {
	Type string `json:"type"`
	Band string
	Data map[string]interface{} `json:"data"`
}

func sendNewInstrumentMessage(hub *hub, subscription *subscription) {
	band := hub.bands[subscription.band]
	newInstrumentMessage := Message{
		Type: "newInstrument",
		Data: make(map[string]interface{}),
		Band: band.name,
	}

	usedInstruments := make(map[string]Instrument)
	for _, instrument := range band.connections {
		usedInstruments[instrument.Type] = instrument
	}

	newInstrumentMessage.Data["newInstrument"] = band.connections[subscription.connection].Type
	newInstrumentMessage.Data["usedInstruments"] = usedInstruments

	broadcastMessage(band.connections, &newInstrumentMessage)
}

func sendRemoveInstrumentMessage(hub *hub, subscription *subscription) {
	band := hub.bands[subscription.band]
	removeInstrumentMessage := Message{
		Type: "removeInstrument",
		Data: make(map[string]interface{}),
		Band: band.name,
	}

	removeInstrumentMessage.Data["removedInstrument"] = band.connections[subscription.connection].Type

	broadcastMessage(band.connections, &removeInstrumentMessage)
}

func broadcastMessage(connections map[*websocket.Conn]Instrument, message *Message) {
	for connection, _ := range connections {
		connection.WriteJSON(message)
	}
}
