package music

import "testing"

func TestDefaultChordCatalogIncludesSupportedVariations(t *testing.T) {
	catalog := DefaultChordCatalog()

	if len(catalog.Variations) < 24 {
		t.Fatalf("variation count = %d, want at least 24", len(catalog.Variations))
	}

	if !catalogHasVariation(catalog, "Dominant7Sus4", "7sus4") {
		t.Fatal("catalog does not include Dominant7Sus4")
	}
}

func catalogHasVariation(catalog ChordCatalog, name string, suffix string) bool {
	for _, variation := range catalog.Variations {
		if variation.Name == name && variation.Suffix == suffix {
			return true
		}
	}

	return false
}

func catalogVariation(t *testing.T, catalog ChordCatalog, name string) ChordVariation {
	t.Helper()
	return catalog.Variations[catalogVariationIndex(t, catalog, name)]
}

func catalogVariationIndex(t *testing.T, catalog ChordCatalog, name string) int {
	t.Helper()
	for i, variation := range catalog.Variations {
		if variation.Name == name {
			return i
		}
	}
	t.Fatalf("catalog variation %q not found", name)
	return 0
}

func TestChordNameUsesVariationSuffix(t *testing.T) {
	catalog := DefaultChordCatalog()
	minorIndex := catalogVariationIndex(t, catalog, "Minor")

	if got := catalog.ChordName(0, minorIndex); got != "Cm" {
		t.Fatalf("ChordName(0, Minor) = %q, want %q", got, "Cm")
	}
}

func TestScaleDegreeSeparatesDegreeAndModifier(t *testing.T) {
	catalog := DefaultChordCatalog()
	minor := catalogVariation(t, catalog, "Minor")

	if got := minor.Formula[1]; got.Degree != 3 || got.Modifier != ModifierFlat {
		t.Fatalf("minor third = %+v, want flat third", got)
	}
}

func TestChordVariationFormulaString(t *testing.T) {
	catalog := DefaultChordCatalog()
	diminished7 := catalogVariation(t, catalog, "Diminished7")

	if got := diminished7.FormulaString(); got != "1,b3,b5,bb7" {
		t.Fatalf("FormulaString() = %q, want %q", got, "1,b3,b5,bb7")
	}
}

func TestNewChordAppliesFormulaToScale(t *testing.T) {
	catalog := DefaultChordCatalog()
	scale := NewScale(ChromaticPitchClasses()[0], SupportedScalePatterns()[0])
	minor := catalogVariation(t, catalog, "Minor")

	chord := NewChord(scale, minor)

	if got := chord.NoteNamesString(); got != "C, Eb, G" {
		t.Fatalf("NoteNamesString() = %q, want %q", got, "C, Eb, G")
	}
}

func TestNewChordUsesScaleLettersBeforeAlterations(t *testing.T) {
	catalog := DefaultChordCatalog()
	scale := NewScale(PitchClass{FlatName: "Db"}, SupportedScalePatterns()[0])
	augmented := catalogVariation(t, catalog, "Augmented")

	chord := NewChord(scale, augmented)

	if got := chord.NoteNamesString(); got != "Db, F, A" {
		t.Fatalf("NoteNamesString() = %q, want %q", got, "Db, F, A")
	}
}

func TestNewChordSupportsExtendedDegrees(t *testing.T) {
	catalog := DefaultChordCatalog()
	scale := NewScale(ChromaticPitchClasses()[0], SupportedScalePatterns()[0])
	dominant9 := catalogVariation(t, catalog, "Dominant9")

	chord := NewChord(scale, dominant9)

	if got := chord.NoteNamesString(); got != "C, E, G, Bb, D" {
		t.Fatalf("NoteNamesString() = %q, want %q", got, "C, E, G, Bb, D")
	}
}

func TestSupportedScaleChordSchemas(t *testing.T) {
	schemas := SupportedScaleChordSchemas()

	if len(schemas) != 5 {
		t.Fatalf("schema count = %d, want 5", len(schemas))
	}

	if got := schemas[0].Name; got != "Triad" {
		t.Fatalf("first schema = %q, want %q", got, "Triad")
	}

	if got := schemas[4].FormulaString(); got != "1,3,5,7,9,11,13" {
		t.Fatalf("13th schema formula = %q, want %q", got, "1,3,5,7,9,11,13")
	}
}

func TestNewScaleChordBuildsDiatonicTriadFromScaleDegree(t *testing.T) {
	scale := NewScale(ChromaticPitchClasses()[0], SupportedScalePatterns()[0])
	triad := SupportedScaleChordSchemas()[0]

	chord := NewScaleChord(scale, 2, triad)

	if got := chord.NoteNamesString(); got != "D, F, A" {
		t.Fatalf("NoteNamesString() = %q, want %q", got, "D, F, A")
	}
}

func TestNewScaleChordBuildsExtendedChordFromScaleDegree(t *testing.T) {
	scale := NewScale(ChromaticPitchClasses()[0], SupportedScalePatterns()[0])
	ninth := SupportedScaleChordSchemas()[2]

	chord := NewScaleChord(scale, 5, ninth)

	if got := chord.NoteNamesString(); got != "G, B, D, F, A" {
		t.Fatalf("NoteNamesString() = %q, want %q", got, "G, B, D, F, A")
	}
}

func TestScaleChordsBuildsOneChordPerScaleDegree(t *testing.T) {
	scale := NewScale(ChromaticPitchClasses()[0], SupportedScalePatterns()[0])
	triad := SupportedScaleChordSchemas()[0]

	chords := ScaleChords(scale, triad)

	if len(chords) != 7 {
		t.Fatalf("chord count = %d, want 7", len(chords))
	}

	if got := chords[6].NoteNamesString(); got != "B, D, F" {
		t.Fatalf("seventh scale chord = %q, want %q", got, "B, D, F")
	}
}

func TestChordCatalogLooksUpMinorChordByNotes(t *testing.T) {
	catalog := DefaultChordCatalog()
	scale := NewScale(ChromaticPitchClasses()[0], SupportedScalePatterns()[1])
	triad := SupportedScaleChordSchemas()[0]
	chord := NewScaleChord(scale, 1, triad)

	name, ok := chord.ChordName(catalog)
	if !ok {
		t.Fatal("expected C minor triad to match a known chord")
	}

	if name != "Cm" {
		t.Fatalf("ChordName() = %q, want %q", name, "Cm")
	}
}

func TestChordCatalogLooksUpVariationByUnorderedNotes(t *testing.T) {
	catalog := DefaultChordCatalog()

	variation, ok := catalog.ChordVariationByNotes([]string{"C", "G", "Eb"})
	if !ok {
		t.Fatal("expected C Eb G to match a known chord")
	}

	if variation.Name != "Minor" {
		t.Fatalf("variation = %q, want %q", variation.Name, "Minor")
	}
}
