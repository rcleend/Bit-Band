package sockhandler

import (
	"errors"
	"log"
	"math/rand"
)

type Band struct {
	name        string
	connections map[*Connection]Instrument
}

var possibleInstruments = []string{"Drum", "Bass", "Rhythm", "Lead"}

func CreateNewBand() Band {
	band := Band{
		name:        getNewBandName(),
		connections: make(map[*Connection]Instrument),
	}
	return band
}

//TODO: implement real band name generator API
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func getNewBandName() string {
	n := 5
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (band *Band) addConnection(subscription *Subscription) {
	if instrument, err := band.getAvailableInstrument(); err != nil {
		log.Fatal(err)
		return
	} else {
		band.connections[subscription.connection] = Instrument{Type: instrument}
	}
}

func (band *Band) getAvailableInstrument() (string, error) {
	for _, instrument := range possibleInstruments {
		index := 0
		for _, usedInstrument := range band.connections {
			if usedInstrument.Type == instrument {
				break
			}
			if index == len(band.connections)-1 {
				return instrument, nil
			}
			index++
		}
	}
	return "No available instrument", errors.New("Something went wrong. No available instruments")
}
