package guitar

import "testing"

func TestSupportedRiffs(t *testing.T) {
	riffs := SupportedRiffs()
	if len(riffs) != 8 {
		t.Fatalf("len(riffs) = %d, want 8", len(riffs))
	}

	if riffs[0].Artist != "Chordmaster" || riffs[0].Song != "Midnight Climb" {
		t.Fatalf("first riff = %+v, want Chordmaster - Midnight Climb", riffs[0])
	}

	if riffs[5].Events[1].Note != "R" {
		t.Fatalf("Low Tide second event = %+v, want rest", riffs[5].Events[1])
	}
}
