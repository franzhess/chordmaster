package piano

import (
	"testing"

	"chordmaster/internal/music"
)

func TestStandardPianoUsesA0ToC8Range(t *testing.T) {
	piano := StandardPiano()

	if piano.Lowest != (music.Note{PitchClass: music.A, Octave: 0}) {
		t.Fatalf("lowest = %+v, want A0", piano.Lowest)
	}

	if piano.Highest != (music.Note{PitchClass: music.C, Octave: 8}) {
		t.Fatalf("highest = %+v, want C8", piano.Highest)
	}
}
