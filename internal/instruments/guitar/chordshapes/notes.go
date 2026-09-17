package chordshapes

import "chordmaster/internal/music"

type Note = music.PitchClass

var (
	C  = music.C
	Cs = music.Cs
	D  = music.D
	Ds = music.Ds
	E  = music.E
	F  = music.F
	Fs = music.Fs
	G  = music.G
	Gs = music.Gs
	A  = music.A
	As = music.As
	B  = music.B
)

func SemitoneDistance(from Note, to Note) int {
	return music.SemitoneDistance(from, to)
}

func TransposeNote(note Note, semitones int) Note {
	return music.TransposePitchClass(note, semitones)
}

func normalizeNote(note Note) Note {
	return music.TransposePitchClass(note, 0)
}
