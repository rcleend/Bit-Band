package sockhandler

type Message struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
	Band string
}

func SendNewInstrumentMessage(band *Band, subscription *Subscription) {
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

	for connection, _ := range band.connections {
		connection.connection.WriteJSON(newInstrumentMessage)
	}
}

func SendRemoveInstrumentMessage(band *Band, subscription *Subscription) {
	removeInstrumentMessage := Message{
		Type: "removeInstrument",
		Data: make(map[string]interface{}),
		Band: band.name,
	}

	removeInstrumentMessage.Data["removedInstrument"] = band.connections[subscription.connection].Type

	for connection, _ := range band.connections {
		connection.connection.WriteJSON(removeInstrumentMessage)
	}
}
