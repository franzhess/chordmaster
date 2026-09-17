package synth

import (
	"encoding/binary"
	"math"
	"testing"
	"time"
)

func TestNoteFrequency(t *testing.T) {
	freq, err := noteFrequency("A4")
	if err != nil {
		t.Fatal(err)
	}

	if math.Abs(freq-440) > 0.001 {
		t.Fatalf("A4 frequency = %f, want 440", freq)
	}

	cSharp, err := noteFrequency("C#4")
	if err != nil {
		t.Fatal(err)
	}
	dFlat, err := noteFrequency("Db4")
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(cSharp-dFlat) > 0.001 {
		t.Fatalf("C#4 frequency = %f, Db4 frequency = %f", cSharp, dFlat)
	}

	eDoubleSharp, err := noteFrequency("E##4")
	if err != nil {
		t.Fatal(err)
	}
	fSharp, err := noteFrequency("F#4")
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(eDoubleSharp-fSharp) > 0.001 {
		t.Fatalf("E##4 frequency = %f, F#4 frequency = %f", eDoubleSharp, fSharp)
	}

	bDoubleFlat, err := noteFrequency("Bbb4")
	if err != nil {
		t.Fatal(err)
	}
	a, err := noteFrequency("A4")
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(bDoubleFlat-a) > 0.001 {
		t.Fatalf("Bbb4 frequency = %f, A4 frequency = %f", bDoubleFlat, a)
	}
}

func TestGenerateSineProduces16BitMonoPCM(t *testing.T) {
	data := generateSine(440, 100*time.Millisecond, 0.5)
	want := SampleRate / 10 * 2

	if len(data) != want {
		t.Fatalf("PCM byte count = %d, want %d", len(data), want)
	}

	if first := int16(binary.LittleEndian.Uint16(data[:2])); first != 0 {
		t.Fatalf("first sample = %d, want 0 due to fade-in", first)
	}
}

func TestGenerateChordAvoidsClipping(t *testing.T) {
	data := generateChord([]float64{261.63, 329.63, 392.00}, 100*time.Millisecond, 1)

	for i := 0; i < len(data); i += 2 {
		sample := int16(binary.LittleEndian.Uint16(data[i:]))
		if sample == math.MaxInt16 || sample == math.MinInt16 {
			t.Fatalf("sample %d clipped to %d", i/2, sample)
		}
	}
}

func TestNoteNamesWithOctavesAscendsAcrossOctaves(t *testing.T) {
	notes := noteNamesWithOctaves([]string{"A", "B", "C"}, 4)
	want := []string{"A4", "B4", "C5"}

	for i := range want {
		if notes[i] != want[i] {
			t.Fatalf("notes[%d] = %q, want %q", i, notes[i], want[i])
		}
	}
}
