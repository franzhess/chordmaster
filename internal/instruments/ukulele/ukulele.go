package ukulele

import "chordmaster/internal/music"

type String struct {
	Label    string
	OpenNote music.Note
}

type Ukulele struct {
	Strings []String
	Frets   int
}

func StandardUkulele() Ukulele {
	return Ukulele{
		Strings: []String{
			{Label: "A", OpenNote: music.Note{PitchClass: music.A, Octave: 4}},
			{Label: "E", OpenNote: music.Note{PitchClass: music.E, Octave: 4}},
			{Label: "C", OpenNote: music.Note{PitchClass: music.C, Octave: 4}},
			{Label: "G", OpenNote: music.Note{PitchClass: music.G, Octave: 4}},
		},
		Frets: 12,
	}
}
