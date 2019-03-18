package sockhandler

import ()

type Message struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
	Band string
}

func SendUpdateInstrumentMessage(band *Band, subscription *Subscription) {
	updateInstrument := Message{
		Type: "updateInstrument",
		Data: make(map[string]interface{}),
		Band: band.name,
	}

	usedInstruments := make(map[string]Instrument)
	for _, instrument := range band.connections {
		usedInstruments[instrument.Type] = instrument
	}

	updateInstrument.Data["newInstrument"] = band.connections[subscription.connection].Type
	updateInstrument.Data["usedInstruments"] = usedInstruments

	for connection, _ := range band.connections {
		connection.connection.WriteJSON(updateInstrument)
	}
}
