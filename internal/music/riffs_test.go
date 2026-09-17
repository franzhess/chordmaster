package music

import (
	"testing"
	"time"
)

func TestParseRiffCatalog(t *testing.T) {
	content := `# Original Riffs

1. Chordmaster — Midnight Climb — BPM: 108
	   E2(e) G2(e) A2(q) B2(e)
`

	riffs, err := ParseRiffCatalog(content)
	if err != nil {
		t.Fatal(err)
	}

	if len(riffs) != 1 {
		t.Fatalf("len(riffs) = %d, want 1", len(riffs))
	}

	if riffs[0].Artist != "Chordmaster" || riffs[0].Song != "Midnight Climb" || riffs[0].BPM != 108 {
		t.Fatalf("riff = %+v", riffs[0])
	}

	if len(riffs[0].Events) != 4 || riffs[0].Events[1].Note != "G2" || riffs[0].Events[1].Duration != "e" || riffs[0].Events[3].Duration != "e" {
		t.Fatalf("events = %+v", riffs[0].Events)
	}
}

func TestRiffTimedNotesPreserveRests(t *testing.T) {
	riff := Riff{
		Artist: "Test Artist",
		Song:   "Rest Test",
		BPM:    120,
		Events: []NoteEvent{
			{Note: "G3", Duration: "q"},
			{Note: "R", Duration: "e"},
		},
	}

	notes := riff.TimedNotes()
	if len(notes) != 2 {
		t.Fatalf("len(notes) = %d, want 2", len(notes))
	}

	if notes[1].Note != "R" || notes[1].Duration != 250*time.Millisecond {
		t.Fatalf("rest note = %+v", notes[1])
	}
}

func TestRiffDuration(t *testing.T) {
	quarter := 500 * time.Millisecond
	if got := RiffDuration("q", 120); got != quarter {
		t.Fatalf("quarter = %s, want %s", got, quarter)
	}

	if got := RiffDuration("e", 120); got != quarter/2 {
		t.Fatalf("eighth = %s, want %s", got, quarter/2)
	}

	if got := RiffDuration("s", 120); got != quarter/4 {
		t.Fatalf("sixteenth = %s, want %s", got, quarter/4)
	}

	if got := RiffDuration("h", 120); got != quarter*2 {
		t.Fatalf("half = %s, want %s", got, quarter*2)
	}

	if got := RiffDuration("w", 120); got != quarter*4 {
		t.Fatalf("whole = %s, want %s", got, quarter*4)
	}
}
