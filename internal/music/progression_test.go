package music

import "testing"

func TestResolveChordProgressionUsesSelectedScale(t *testing.T) {
	scale := NewScale(C, SupportedScalePatterns()[0])
	progression := ChordProgression{Name: "I-V-vi-IV", Symbols: []string{"I", "V", "vi", "IV"}}

	chords := ResolveChordProgression(scale, progression)
	if len(chords) != 4 {
		t.Fatalf("len(chords) = %d, want 4", len(chords))
	}

	wantNames := []string{"C", "G", "Am", "F"}
	for i, want := range wantNames {
		if chords[i].Name != want {
			t.Fatalf("chords[%d].Name = %q, want %q", i, chords[i].Name, want)
		}
	}
}

func TestResolveChordProgressionHandlesBorrowedAndSeventhChords(t *testing.T) {
	scale := NewScale(C, SupportedScalePatterns()[0])
	progression := ChordProgression{Name: "I-bIII7-ii-bII7", Symbols: []string{"I", "bIII7", "ii", "bII7"}}

	chords := ResolveChordProgression(scale, progression)
	wantNames := []string{"C", "Eb7", "Dm", "Db7"}
	for i, want := range wantNames {
		if chords[i].Name != want {
			t.Fatalf("chords[%d].Name = %q, want %q", i, chords[i].Name, want)
		}
	}

	wantNotes := []string{"Db", "F", "Ab", "Cb"}
	for i, want := range wantNotes {
		if chords[3].NoteNames[i] != want {
			t.Fatalf("bII7 note %d = %q, want %q", i, chords[3].NoteNames[i], want)
		}
	}
}

func TestResolveChordProgressionHandlesSecondaryDominants(t *testing.T) {
	scale := NewScale(C, SupportedScalePatterns()[0])
	progression := ChordProgression{Name: "V/V-V-I", Symbols: []string{"V/V", "V", "I"}}

	chords := ResolveChordProgression(scale, progression)
	if chords[0].Name != "D" {
		t.Fatalf("secondary dominant = %q, want D", chords[0].Name)
	}
}

func TestSupportedRhythmPatternsIncludeCommonDurations(t *testing.T) {
	patterns := SupportedRhythmPatterns()
	if len(patterns) == 0 {
		t.Fatal("no rhythm patterns returned")
	}

	foundPopSplit := false
	foundModalJazz := false
	seen := map[string]string{}
	for _, pattern := range patterns {
		label := RhythmPatternLabel(pattern)
		if previous, ok := seen[label]; ok {
			t.Fatalf("rhythm pattern %q duplicates %q as %q", pattern.Name, previous, label)
		}
		seen[label] = pattern.Name

		switch pattern.Name {
		case "Pop Split Bar":
			foundPopSplit = label == "224"
		case "Modal Jazz":
			foundModalJazz = len(pattern.Units) == 1 && pattern.Units[0] == 16 && label == "16"
		}
	}

	if seen["4"] != "Whole Measure Chords" {
		t.Fatalf("looped 4-beat rhythm was not reduced and deduplicated: %q", seen["4"])
	}

	if !foundPopSplit {
		t.Fatal("Pop Split Bar rhythm missing or incorrect")
	}
	if !foundModalJazz {
		t.Fatal("Modal Jazz rhythm missing or incorrect")
	}
}
