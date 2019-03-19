package sockhandler

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
)

type band struct {
	id          int
	connections map[*websocket.Conn]instrument
}

func createNewBand() band {
	band := band{
		id:          rand.Intn(1000000),
		connections: make(map[*websocket.Conn]instrument),
	}
	return band
}

func (band *band) addConnection(subscription *subscription) {
	if instrument, err := band.getAvailableInstrumentType(); err != nil {
		log.Fatal(err)
		return
	} else {
		band.connections[subscription.connection] = instrument
	}
}

func (band *band) getAvailableInstrumentType() (instrument, error) {
	for _, instrument := range possibleInstruments {
		index := 0
		for _, usedInstrument := range band.connections {
			if usedInstrument.name == instrument.name {
				break
			}
			if index == len(band.connections)-1 {
				return instrument, nil
			}
			index++
		}
	}
	return instrument{}, errors.New("Something went wrong. No available instruments")
}
