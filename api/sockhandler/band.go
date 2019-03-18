package sockhandler

import (
	"errors"
	"log"
	"math/rand"
)

type Band struct {
	name        string
	connections map[*Connection]string
	instruments []string
}

var possibleInstruments = []string{"Drum", "Bass", "Rhythm", "Lead"}

func CreateNewBand() Band {
	band := Band{
		name:        getNewBandName(),
		connections: make(map[*Connection]string),
		instruments: possibleInstruments[0:1],
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
	if instrument, err := getAvailableInstrument(band); err != nil {
		log.Fatal(err)
		return
	} else {
		band.connections[subscription.connection] = instrument
		band.instruments = append(band.instruments, instrument)
		sendUpdateInstrumentMessage(band)
	}
}

func getAvailableInstrument(band *Band) (string, error) {
	for _, instrument := range possibleInstruments {
		index := 0
		for _, usedInstrument := range band.instruments {
			if usedInstrument == instrument {
				break
			}
			if index == len(band.instruments)-1 && usedInstrument != instrument {
				return instrument, nil
			}
			index++
		}
	}
	return "No available instrument", errors.New("Something went wrong. No available instruments")
}
