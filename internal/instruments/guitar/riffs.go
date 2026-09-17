package guitar

import "chordmaster/internal/music"

// SupportedRiffs contains original exercises composed for Chordmaster.
func SupportedRiffs() []music.Riff {
	return []music.Riff{
		{Artist: "Chordmaster", Song: "Midnight Climb", BPM: 108, Events: []music.NoteEvent{ev("E2", "e"), ev("G2", "e"), ev("A2", "q"), ev("B2", "e"), ev("A2", "e"), ev("G2", "q"), ev("E2", "h")}},
		{Artist: "Chordmaster", Song: "Copper Line", BPM: 124, Events: []music.NoteEvent{ev("A2", "q"), ev("C3", "e"), ev("D3", "e"), ev("E3", "q"), ev("D3", "q"), ev("A2", "h")}},
		{Artist: "Chordmaster", Song: "Open Road", BPM: 96, Events: []music.NoteEvent{ev("D3", "q"), ev("F#3", "e"), ev("A3", "e"), ev("B3", "q"), ev("A3", "q"), ev("F#3", "h")}},
		{Artist: "Chordmaster", Song: "Lantern Glow", BPM: 88, Events: []music.NoteEvent{ev("E3", "e"), ev("G3", "e"), ev("B3", "q"), ev("D4", "q"), ev("B3", "e"), ev("G3", "e"), ev("E3", "h")}},
		{Artist: "Chordmaster", Song: "Ridge Runner", BPM: 142, Events: []music.NoteEvent{ev("G2", "s"), ev("A2", "s"), ev("B2", "e"), ev("D3", "e"), ev("B2", "e"), ev("A2", "q"), ev("G2", "q")}},
		{Artist: "Chordmaster", Song: "Low Tide", BPM: 76, Events: []music.NoteEvent{ev("E2", "q"), ev("R", "e"), ev("G2", "e"), ev("A2", "q"), ev("G2", "q"), ev("E2", "h")}},
		{Artist: "Chordmaster", Song: "Glass Horizon", BPM: 116, Events: []music.NoteEvent{ev("C3", "e"), ev("D3", "e"), ev("G3", "q"), ev("A3", "e"), ev("G3", "e"), ev("D3", "q"), ev("C3", "h")}},
		{Artist: "Chordmaster", Song: "Signal Fire", BPM: 132, Events: []music.NoteEvent{ev("B2", "e"), ev("D3", "e"), ev("E3", "q"), ev("F#3", "e"), ev("E3", "e"), ev("D3", "q"), ev("B2", "h")}},
	}
}

func ev(note string, duration string) music.NoteEvent {
	return music.NoteEvent{Note: note, Duration: duration}
}
