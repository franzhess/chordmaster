package piano

import "chordmaster/internal/music"

type Key struct {
	Note music.Note
}

type Piano struct {
	Lowest  music.Note
	Highest music.Note
}

func StandardPiano() Piano {
	return Piano{
		Lowest:  music.Note{PitchClass: music.A, Octave: 0},
		Highest: music.Note{PitchClass: music.C, Octave: 8},
	}
}
