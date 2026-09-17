package main

import (
	"fmt"
	"os"

	"chordmaster/internal/synth"
	"chordmaster/internal/ui"

	tea "charm.land/bubbletea/v2"
)

func main() {
	audio, err := synth.NewEngine()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize audio: %v\n", err)
		os.Exit(1)
	}

	program := tea.NewProgram(ui.NewModel(audio))
	if _, err := program.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run chordmaster: %v\n", err)
		os.Exit(1)
	}
}
