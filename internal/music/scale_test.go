package music

import "testing"

func TestSupportedScalePatternsIncludesAllPatterns(t *testing.T) {
	patterns := SupportedScalePatterns()

	if len(patterns) != 16 {
		t.Fatalf("pattern count = %d, want 16", len(patterns))
	}

	if patterns[0].Name != "Major (Ionian)" {
		t.Fatalf("first pattern = %q, want Major (Ionian)", patterns[0].Name)
	}

	if patterns[len(patterns)-1].Name != "Diminished (Whole-Half)" {
		t.Fatalf("last pattern = %q, want Diminished (Whole-Half)", patterns[len(patterns)-1].Name)
	}
}

func TestScalePatternStepString(t *testing.T) {
	major := SupportedScalePatterns()[0]

	if got := major.StepString(); got != "W W H W W W H" {
		t.Fatalf("StepString() = %q, want %q", got, "W W H W W W H")
	}
}

func TestScalePatternsSpanOctave(t *testing.T) {
	for _, pattern := range SupportedScalePatterns() {
		if got := pattern.SemitoneCount(); got != 12 {
			t.Fatalf("%s semitone count = %d, want 12", pattern.Name, got)
		}
	}
}

func TestHarmonicMinorUsesAugmentedSecond(t *testing.T) {
	harmonicMinor := SupportedScalePatterns()[2]

	if got := harmonicMinor.StepString(); got != "W H W W H WH H" {
		t.Fatalf("StepString() = %q, want %q", got, "W H W W H WH H")
	}

	if harmonicMinor.Steps[5].Semitones != 3 {
		t.Fatalf("harmonic minor WH semitones = %d, want 3", harmonicMinor.Steps[5].Semitones)
	}
}

func TestNewScaleAppliesPatternToPitchClasses(t *testing.T) {
	root := ChromaticPitchClasses()[0]
	major := SupportedScalePatterns()[0]

	scale := NewScale(root, major)
	want := "C, D, E, F, G, A, B"

	if got := scale.PitchClassNamesString(); got != want {
		t.Fatalf("PitchClassNamesString() = %q, want %q", got, want)
	}
}

func TestNewScaleWrapsAroundChromaticPitchClasses(t *testing.T) {
	root := ChromaticPitchClasses()[9]
	naturalMinor := SupportedScalePatterns()[1]

	scale := NewScale(root, naturalMinor)
	want := "A, B, C, D, E, F, G"

	if got := scale.PitchClassNamesString(); got != want {
		t.Fatalf("PitchClassNamesString() = %q, want %q", got, want)
	}
}

func TestNewScaleAcceptsFlatEnharmonicRoot(t *testing.T) {
	root := PitchClass{FlatName: "Db"}
	major := SupportedScalePatterns()[0]

	scale := NewScale(root, major)
	want := "Db, Eb, F, Gb, Ab, Bb, C"

	if got := scale.PitchClassNamesString(); got != want {
		t.Fatalf("PitchClassNamesString() = %q, want %q", got, want)
	}
}

func TestNewScaleSpellsSharpKeysWithEachLetterOnce(t *testing.T) {
	root := ChromaticPitchClasses()[1]
	major := SupportedScalePatterns()[0]

	scale := NewScale(root, major)
	want := "C#, D#, E#, F#, G#, A#, B#"

	if got := scale.PitchClassNamesString(); got != want {
		t.Fatalf("PitchClassNamesString() = %q, want %q", got, want)
	}
}

func TestNewChromaticScaleContainsTwelvePitchClasses(t *testing.T) {
	root := ChromaticPitchClasses()[0]
	chromatic := SupportedScalePatterns()[12]

	scale := NewScale(root, chromatic)

	if got := len(scale.PitchClasses); got != 12 {
		t.Fatalf("scale pitch class count = %d, want 12", got)
	}
}
