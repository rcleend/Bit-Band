package sockhandler

var drum = instrument{
	name: "Drum",
	Sounds: []sound{
		sound{
			Name: "Kick",
		},
		sound{
			Name: "Snare",
		},
		sound{
			Name: "Hi-Hat",
		},
		sound{
			Name: "Tom A",
		},
		sound{
			Name: "Tom B",
		},
	},
}
var bass = instrument{
	name: "Bass",
}
var rhythm = instrument{
	name: "Rhythm",
}
var lead = instrument{
	name: "Lead",
}

var possibleInstruments = []instrument{drum, bass, rhythm, lead}

type sound struct {
	Name string
}

type instrument struct {
	name   string
	Sounds []sound
}
