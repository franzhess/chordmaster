package chordshapes

import "testing"

func TestEshapeMajorBarreTransposesFToG(t *testing.T) {
	shape := findTemplate(t, "E-shape major barre")
	voicing, err := TransposeShape(shape, G)
	if err != nil {
		t.Fatal(err)
	}

	want := [6]int{3, 5, 5, 4, 3, 3}
	if voicing.Frets != want {
		t.Fatalf("frets = %v, want %v", voicing.Frets, want)
	}
}

func TestEshapeMajorBarreTransposesFToA(t *testing.T) {
	shape := findTemplate(t, "E-shape major barre")
	voicing, err := TransposeShape(shape, A)
	if err != nil {
		t.Fatal(err)
	}

	want := [6]int{5, 7, 7, 6, 5, 5}
	if voicing.Frets != want {
		t.Fatalf("frets = %v, want %v", voicing.Frets, want)
	}
}

func TestAshapeMajorBarreTransposesBbToC(t *testing.T) {
	shape := findTemplate(t, "A-shape major barre")
	voicing, err := TransposeShape(shape, C)
	if err != nil {
		t.Fatal(err)
	}

	want := [6]int{-1, 3, 5, 5, 5, 3}
	if voicing.Frets != want {
		t.Fatalf("frets = %v, want %v", voicing.Frets, want)
	}
}

func TestOpenCIsReturnedForCMajor(t *testing.T) {
	voicings := GetVoicings(C, Major)

	if !hasSource(voicings, "C open") {
		t.Fatalf("C open not found in C major voicings: %+v", voicings)
	}
}

func TestOpenFIsPreferredForFMajor(t *testing.T) {
	voicings := GetVoicings(F, Major)

	if len(voicings) == 0 || voicings[0].Source != "F open" {
		t.Fatalf("first F major voicing = %+v, want F open", voicings)
	}
}

func TestOpenB7IsReturnedForBDominant7(t *testing.T) {
	voicings := GetVoicings(B, Dominant7)

	if !hasSource(voicings, "B7 open") {
		t.Fatalf("B7 open not found in B7 voicings: %+v", voicings)
	}
}

func TestOpenBdimIsReturnedForBDiminished(t *testing.T) {
	voicings := GetVoicings(B, Diminished)

	if !hasSource(voicings, "Bdim open") {
		t.Fatalf("Bdim open not found in B diminished voicings: %+v", voicings)
	}
}

func TestOpenCIsNotReturnedForDMajor(t *testing.T) {
	voicings := GetVoicings(D, Major)

	if hasSource(voicings, "C open") {
		t.Fatal("C open was returned for D major")
	}
}

func TestNonMovableOpenShapesAreNotTransposed(t *testing.T) {
	shape := findTemplate(t, "C open")

	if _, err := TransposeShape(shape, D); err == nil {
		t.Fatal("expected non-movable C open shape to reject D target")
	}
}

func TestMovableTemplatesRejectTargetShapesAboveFret24(t *testing.T) {
	shape := ShapeTemplate{
		Name:     "too high",
		BaseRoot: C,
		Quality:  Major,
		Frets:    [6]int{24, 24, 24, 24, 24, 24},
		Movable:  true,
	}

	if _, err := TransposeShape(shape, Cs); err == nil {
		t.Fatal("expected shape above fret 24 to be rejected")
	}
}

func TestValidateCMajorOpenVoicing(t *testing.T) {
	voicing := Voicing{Root: C, Quality: Major, Frets: [6]int{-1, 3, 2, 0, 1, 0}}

	if !ValidateVoicing(voicing, []Note{C, E, G}, StandardTuning) {
		t.Fatal("expected x32010 to validate as C major")
	}
}

func TestValidateGMajorBarreVoicing(t *testing.T) {
	voicing := Voicing{Root: G, Quality: Major, Frets: [6]int{3, 5, 5, 4, 3, 3}}

	if !ValidateVoicing(voicing, []Note{G, B, D}, StandardTuning) {
		t.Fatal("expected 355433 to validate as G major")
	}
}

func TestValidateBdimOpenVoicing(t *testing.T) {
	voicing := Voicing{Root: B, Quality: Diminished, Frets: [6]int{-1, 2, 3, 4, 3, -1}}

	if !ValidateVoicing(voicing, []Note{B, D, F}, StandardTuning) {
		t.Fatal("expected x2343x to validate as B diminished")
	}
}

func TestParseAndRenderFrets(t *testing.T) {
	frets, err := ParseFrets("x32010")
	if err != nil {
		t.Fatal(err)
	}

	want := [6]int{-1, 3, 2, 0, 1, 0}
	if frets != want {
		t.Fatalf("frets = %v, want %v", frets, want)
	}

	if got := RenderFrets(frets); got != "x 3 2 0 1 0" {
		t.Fatalf("RenderFrets() = %q", got)
	}
}

func TestParseNoteNameAndRenderFingers(t *testing.T) {
	note, err := ParseNoteName("Bb")
	if err != nil {
		t.Fatal(err)
	}
	if note != As {
		t.Fatalf("ParseNoteName(Bb) = %v, want %v", note, As)
	}

	fingers := [6]int{0, 3, 2, 0, 1, 0}
	if got := RenderFingers(fingers); got != "- 3 2 - 1 -" {
		t.Fatalf("RenderFingers() = %q", got)
	}
}

func findTemplate(t *testing.T, name string) ShapeTemplate {
	t.Helper()
	for _, shape := range Templates() {
		if shape.Name == name {
			return shape
		}
	}

	t.Fatalf("template %q not found", name)
	return ShapeTemplate{}
}

func hasSource(voicings []Voicing, source string) bool {
	for _, voicing := range voicings {
		if voicing.Source == source {
			return true
		}
	}

	return false
}
