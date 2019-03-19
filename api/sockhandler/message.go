package sockhandler

import (
	"github.com/gorilla/websocket"
)

type Message struct {
	Type   string                 `json:"type"`
	BandID int                    `json:"bandID"`
	Data   map[string]interface{} `json:"data"`
}

func sendNewInstrumentMessage(hub *hub, subscription *subscription) {
	band := hub.bands[subscription.bandID]

	usedInstruments := make(map[string]instrument)
	for _, instrument := range band.connections {
		usedInstruments[instrument.name] = instrument
	}

	newInstrumentMessage := Message{"newInstrument", band.id, make(map[string]interface{})}
	newInstrumentMessage.Data["instrumentName"] = band.connections[subscription.connection].name
	newInstrumentMessage.Data["allInstruments"] = usedInstruments

	broadcastMessage(band.connections, &newInstrumentMessage)
}

func sendRemoveInstrumentMessage(hub *hub, subscription *subscription) {
	band := hub.bands[subscription.bandID]
	if band != nil {
		removeInstrumentMessage := Message{"removeInstrument", band.id, make(map[string]interface{})}
		removeInstrumentMessage.Data["removedInstrument"] = band.connections[subscription.connection].name
		broadcastMessage(band.connections, &removeInstrumentMessage)
	}
}

func broadcastMessage(connections map[*websocket.Conn]instrument, message *Message) {
	for connection, _ := range connections {
		connection.WriteJSON(message)
	}
}
