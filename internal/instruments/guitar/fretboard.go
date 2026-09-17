package guitar

import (
	"fmt"
	"strings"

	"chordmaster/internal/music"
)

type GuitarString struct {
	Label    string
	OpenNote music.Note
}

type FretPosition struct {
	Note music.Note
}

type GuitarFretboard struct {
	Strings []GuitarString
	Frets   int
}

func StandardGuitarFretboard() GuitarFretboard {
	return GuitarFretboard{
		Strings: []GuitarString{
			{Label: "e", OpenNote: music.Note{PitchClass: music.E, Octave: 4}},
			{Label: "B", OpenNote: music.Note{PitchClass: music.B, Octave: 3}},
			{Label: "G", OpenNote: music.Note{PitchClass: music.G, Octave: 3}},
			{Label: "D", OpenNote: music.Note{PitchClass: music.D, Octave: 3}},
			{Label: "A", OpenNote: music.Note{PitchClass: music.A, Octave: 2}},
			{Label: "E", OpenNote: music.Note{PitchClass: music.E, Octave: 2}},
		},
		Frets: 12,
	}
}

func (f GuitarFretboard) Render() string {
	return f.RenderWithFormatter(func(position FretPosition) string {
		return FormatFretNote(position.Note.PitchClass)
	})
}

func (f GuitarFretboard) RenderWithFormatter(formatter func(FretPosition) string) string {
	var b strings.Builder
	for _, guitarString := range f.Strings {
		b.WriteString(guitarString.Label)
		for fret := 0; fret <= f.Frets; fret++ {
			b.WriteString("|")
			b.WriteString(formatter(f.noteAt(guitarString, fret)))
		}
		b.WriteString("|")
		b.WriteString("\n")
	}

	b.WriteString(" ")
	for fret := 0; fret <= f.Frets; fret++ {
		b.WriteString(fmt.Sprintf("|%-5d", fret))
	}
	b.WriteString("|")

	return b.String()
}

func (f GuitarFretboard) noteAt(guitarString GuitarString, fret int) FretPosition {
	openSemitone, _ := guitarString.OpenNote.PitchClass.Semitone()
	return FretPosition{
		Note: music.Note{
			PitchClass: music.TransposePitchClass(guitarString.OpenNote.PitchClass, fret),
			Octave:     guitarString.OpenNote.Octave + (openSemitone+fret)/12,
		},
	}
}

func FormatFretNote(note music.PitchClass) string {
	name := note.String()
	if strings.Contains(name, "#") {
		return "--" + name + "-"
	}

	return "--" + name + "--"
}
