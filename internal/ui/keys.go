package ui

import "charm.land/bubbles/v2/key"

type keyMap struct {
	Quit             key.Binding
	Back             key.Binding
	Help             key.Binding
	Select           key.Binding
	MoveUp           key.Binding
	MoveDown         key.Binding
	MoveLeft         key.Binding
	MoveRight        key.Binding
	ProgressionRootL key.Binding
	ProgressionRootR key.Binding
	RhythmLeft       key.Binding
	RhythmRight      key.Binding
	Play             key.Binding
	PlayScaleChords  key.Binding
	ToggleChordTabs  key.Binding
	PageUp           key.Binding
	PageDown         key.Binding
	Home             key.Binding
	End              key.Binding
}

func defaultKeyMap() keyMap {
	return keyMap{
		Quit:             key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
		Back:             key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "back")),
		Help:             key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
		Select:           key.NewBinding(key.WithKeys("enter", "space"), key.WithHelp("enter", "select")),
		MoveUp:           key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("j/k", "move")),
		MoveDown:         key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("j/down", "move down")),
		MoveLeft:         key.NewBinding(key.WithKeys("left", "h"), key.WithHelp("h/left", "previous")),
		MoveRight:        key.NewBinding(key.WithKeys("right", "l"), key.WithHelp("l/right", "next")),
		ProgressionRootL: key.NewBinding(key.WithKeys("["), key.WithHelp("[", "root down")),
		ProgressionRootR: key.NewBinding(key.WithKeys("]"), key.WithHelp("]", "root up")),
		RhythmLeft:       key.NewBinding(key.WithKeys(","), key.WithHelp(",", "rhythm down")),
		RhythmRight:      key.NewBinding(key.WithKeys("."), key.WithHelp(".", "rhythm up")),
		Play:             key.NewBinding(key.WithKeys("p"), key.WithHelp("p", "play/cancel")),
		PlayScaleChords:  key.NewBinding(key.WithKeys("c"), key.WithHelp("c", "scale chords")),
		ToggleChordTabs:  key.NewBinding(key.WithKeys("f"), key.WithHelp("f", "tabs")),
		PageUp:           key.NewBinding(key.WithKeys("pgup"), key.WithHelp("pgup", "scroll up")),
		PageDown:         key.NewBinding(key.WithKeys("pgdown"), key.WithHelp("pgdn", "scroll down")),
		Home:             key.NewBinding(key.WithKeys("home"), key.WithHelp("home", "top")),
		End:              key.NewBinding(key.WithKeys("end"), key.WithHelp("end", "bottom")),
	}
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Back, k.MoveUp, k.MoveDown, k.MoveLeft, k.MoveRight, k.Play}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Back, k.Quit},
		{k.MoveUp, k.MoveDown, k.MoveLeft, k.MoveRight, k.Select},
		{k.Play, k.PlayScaleChords, k.ToggleChordTabs},
		{k.ProgressionRootL, k.ProgressionRootR, k.RhythmLeft, k.RhythmRight},
		{k.PageUp, k.PageDown, k.Home, k.End},
	}
}
