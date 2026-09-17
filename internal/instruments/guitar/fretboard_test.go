package guitar

import (
	"strings"
	"testing"

	"chordmaster/internal/music"
)

func TestStandardGuitarFretboardRender(t *testing.T) {
	fretboard := StandardGuitarFretboard()
	rendered := fretboard.Render()

	if !strings.Contains(rendered, "e|--E--|--F--|--F#-|--G--|--G#-|--A--|--A#-|--B--|--C--|--C#-|--D--|--D#-|--E--|") {
		t.Fatalf("rendered fretboard missing high e string:\n%s", rendered)
	}

	if !strings.Contains(rendered, "B|--B--|--C--|--C#-|--D--|--D#-|--E--|--F--|--F#-|--G--|--G#-|--A--|--A#-|--B--|") {
		t.Fatalf("rendered fretboard missing B string:\n%s", rendered)
	}

	if !strings.Contains(rendered, " |0    |1    |2    |3    |4    |5    |6    |7    |8    |9    |10   |11   |12   |") {
		t.Fatalf("rendered fretboard missing fret numbers:\n%s", rendered)
	}
}

func TestFormatFretNoteUsesLeadingDashes(t *testing.T) {
	if got := FormatFretNote(music.E); got != "--E--" {
		t.Fatalf("FormatFretNote(E) = %q, want %q", got, "--E--")
	}

	if got := FormatFretNote(music.Fs); got != "--F#-" {
		t.Fatalf("FormatFretNote(F#) = %q, want %q", got, "--F#-")
	}
}

func TestFretPositionIncludesOctave(t *testing.T) {
	fretboard := StandardGuitarFretboard()

	highEOpen := fretboard.noteAt(fretboard.Strings[0], 0)
	if highEOpen.Note.PitchClass != music.E || highEOpen.Note.Octave != 4 {
		t.Fatalf("high e open = %+v, want E4", highEOpen)
	}

	highE12 := fretboard.noteAt(fretboard.Strings[0], 12)
	if highE12.Note.PitchClass != music.E || highE12.Note.Octave != 5 {
		t.Fatalf("high e fret 12 = %+v, want E5", highE12)
	}

	bStringFret1 := fretboard.noteAt(fretboard.Strings[1], 1)
	if bStringFret1.Note.PitchClass != music.C || bStringFret1.Note.Octave != 4 {
		t.Fatalf("B string fret 1 = %+v, want C4", bStringFret1)
	}
}
