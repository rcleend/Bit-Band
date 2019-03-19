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

	usedInstruments := make(map[string]instrument)
	for _, instrument := range band.connections {
		usedInstruments[instrument.Type] = instrument
	}

	newInstrumentMessage := Message{"newInstrument", band.name, make(map[string]interface{})}
	newInstrumentMessage.Data["newInstrumentType"] = band.connections[subscription.connection].Type
	newInstrumentMessage.Data["usedInstruments"] = usedInstruments

	broadcastMessage(band.connections, &newInstrumentMessage)
}

func sendRemoveInstrumentMessage(hub *hub, subscription *subscription) {
	band := hub.bands[subscription.band]
	if band != nil {
		removeInstrumentMessage := Message{"removeInstrument", band.name, make(map[string]interface{})}
		removeInstrumentMessage.Data["removedInstrument"] = band.connections[subscription.connection].Type
		broadcastMessage(band.connections, &removeInstrumentMessage)
	}
}

func broadcastMessage(connections map[*websocket.Conn]instrument, message *Message) {
	for connection, _ := range connections {
		connection.WriteJSON(message)
	}
}
