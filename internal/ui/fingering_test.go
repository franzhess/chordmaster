package ui

import (
	"strings"
	"testing"

	"chordmaster/internal/instruments/guitar/chordshapes"
	"chordmaster/internal/music"

	"github.com/charmbracelet/x/ansi"
)

func TestChordTabOpenPositionUsesNutAndFourFrets(t *testing.T) {
	variation := music.DefaultChordCatalog().Variations[0]
	tab, ok := chordTab("C", variation, chordTabNotes)
	if !ok {
		t.Fatal("expected C major tab")
	}
	tab = ansi.Strip(tab)

	if !strings.Contains(tab, "e E||---|---|---|---|") {
		t.Fatalf("open tab did not render four frets with nut: %q", tab)
	}
	if strings.Contains(tab, "\n    1") {
		t.Fatalf("open tab includes fret guide: %q", tab)
	}
}

func TestChordTabShiftedPositionUsesFretGuideAndFourFrets(t *testing.T) {
	tab := renderChordTab(chordshapes.Voicing{Frets: [6]int{9, 11, 11, 10, 9, 9}}, chordTabNotes)
	tab = ansi.Strip(tab)

	if strings.Contains(tab, "||") {
		t.Fatalf("shifted tab uses nut marker: %q", tab)
	}
	if !strings.Contains(tab, "\n    9") {
		t.Fatalf("shifted tab does not include starting fret guide: %q", tab)
	}
	if !strings.Contains(tab, "e  |C#-|---|---|---|") {
		t.Fatalf("shifted tab did not render four fret cells from starting fret: %q", tab)
	}
}
