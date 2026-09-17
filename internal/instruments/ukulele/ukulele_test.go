package ukulele

import (
	"testing"

	"chordmaster/internal/music"
)

func TestStandardUkuleleUsesReentrantGCEA(t *testing.T) {
	ukulele := StandardUkulele()

	if len(ukulele.Strings) != 4 {
		t.Fatalf("string count = %d, want 4", len(ukulele.Strings))
	}

	if ukulele.Strings[0].OpenNote != (music.Note{PitchClass: music.A, Octave: 4}) {
		t.Fatalf("first string = %+v, want A4", ukulele.Strings[0].OpenNote)
	}

	if ukulele.Strings[3].OpenNote != (music.Note{PitchClass: music.G, Octave: 4}) {
		t.Fatalf("fourth string = %+v, want G4", ukulele.Strings[3].OpenNote)
	}
}
