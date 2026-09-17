package music

import "testing"

func TestChromaticPitchClassesContainsTwelvePitchClasses(t *testing.T) {
	pitchClasses := ChromaticPitchClasses()

	if len(pitchClasses) != 12 {
		t.Fatalf("pitch class count = %d, want 12", len(pitchClasses))
	}
}

func TestPitchClassNameIncludesEnharmonicSpelling(t *testing.T) {
	pitchClasses := ChromaticPitchClasses()

	if got := pitchClasses[1].Name(); got != "C# / Db" {
		t.Fatalf("pitchClasses[1].Name() = %q, want %q", got, "C# / Db")
	}

	if got := pitchClasses[4].Name(); got != "E" {
		t.Fatalf("pitchClasses[4].Name() = %q, want %q", got, "E")
	}
}

func TestPitchClassNamesString(t *testing.T) {
	got := PitchClassNamesString(ChromaticPitchClasses()[:3])
	want := "C, C# / Db, D"

	if got != want {
		t.Fatalf("PitchClassNamesString() = %q, want %q", got, want)
	}
}

func TestParsePitchClassNameAcceptsEnharmonicAliases(t *testing.T) {
	bb, err := ParsePitchClassName("Bb")
	if err != nil {
		t.Fatal(err)
	}
	if bb != As {
		t.Fatalf("ParsePitchClassName(Bb) = %v, want %v", bb, As)
	}

	bSharp, err := ParsePitchClassName("B#")
	if err != nil {
		t.Fatal(err)
	}
	if bSharp != C {
		t.Fatalf("ParsePitchClassName(B#) = %v, want %v", bSharp, C)
	}
}

func TestParsePitchClassNameAcceptsDoubleAccidentals(t *testing.T) {
	tests := map[string]PitchClass{
		"C##": D,
		"E##": Fs,
		"B##": Cs,
		"Dbb": C,
		"Abb": G,
		"Bbb": A,
	}

	for name, want := range tests {
		got, err := ParsePitchClassName(name)
		if err != nil {
			t.Fatalf("ParsePitchClassName(%s) returned error: %v", name, err)
		}
		if got != want {
			t.Fatalf("ParsePitchClassName(%s) = %v, want %v", name, got, want)
		}
	}
}

func TestPitchClassBySemitoneWraps(t *testing.T) {
	if got := PitchClassBySemitone(13); got != Cs {
		t.Fatalf("PitchClassBySemitone(13) = %v, want %v", got, Cs)
	}

	if got := PitchClassBySemitone(-1); got != B {
		t.Fatalf("PitchClassBySemitone(-1) = %v, want %v", got, B)
	}
}

func TestTransposePitchClassWrapsChromaticScale(t *testing.T) {
	if got := TransposePitchClass(B, 1); got != C {
		t.Fatalf("TransposePitchClass(B, 1) = %v, want %v", got, C)
	}

	if got := TransposePitchClass(C, -1); got != B {
		t.Fatalf("TransposePitchClass(C, -1) = %v, want %v", got, B)
	}
}

func TestNoteStringIncludesPitchClassAndOctave(t *testing.T) {
	note := Note{PitchClass: Fs, Octave: 4}
	if got := note.String(); got != "F#4" {
		t.Fatalf("Note.String() = %q, want %q", got, "F#4")
	}
}
