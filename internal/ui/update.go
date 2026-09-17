package ui

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

func update(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case splashTimeoutMsg:
		return updateSplashTimeout(m)
	case playbackFinishedMsg:
		return updatePlaybackFinished(m, msg)
	case playbackStepFinishedMsg:
		return updatePlaybackStepFinished(m, msg)
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.syncViewport()
		return m, nil
	case tea.KeyPressMsg:
		return updateKey(m, msg)
	default:
		return m, nil
	}
}

func updateSplashTimeout(m model) (model, tea.Cmd) {
	if m.screen == screenSplash {
		m.screen = screenMainMenu
	}

	return m, nil
}

func updatePlaybackFinished(m model, msg playbackFinishedMsg) (model, tea.Cmd) {
	m.isPlaying = false
	m.clearPlaybackHighlight()
	if msg.err != nil {
		m.status = "Playback failed: " + msg.err.Error()
	} else {
		m.status = ""
	}
	if msg.source == playbackSplash && m.screen == screenSplash {
		m.screen = screenMainMenu
	}

	return m, nil
}

func updatePlaybackStepFinished(m model, msg playbackStepFinishedMsg) (model, tea.Cmd) {
	// Playback commands can finish after the user has cancelled or started a new
	// session; ignore those stale completions so they cannot advance highlights.
	if msg.session != m.playbackSession || msg.kind != m.playbackKind {
		return m, nil
	}
	if msg.err != nil {
		m.finishPlayback(msg.err)
		return m, nil
	}

	switch msg.kind {
	case playbackKindScaleNotes:
		return m.advanceScaleNotePlayback(msg.index + 1)
	case playbackKindScaleChords:
		return m.advanceScaleChordPlayback(msg.index + 1)
	default:
		return m, nil
	}
}

func updateKey(m model, msg tea.KeyPressMsg) (model, tea.Cmd) {
	if key.Matches(msg, m.keys.Quit) {
		return m, tea.Quit
	}

	if m.showHelp {
		return updateHelpKey(m, msg)
	}

	switch {
	case key.Matches(msg, m.keys.Help):
		m.showHelp = true
	case key.Matches(msg, m.keys.ToggleChordTabs):
		m.toggleChordTabMode()
	case key.Matches(msg, m.keys.Back):
		return updateBackKey(m)
	case key.Matches(msg, m.keys.Select):
		m.openSelectedScreen()
	case key.Matches(msg, m.keys.MoveUp):
		m.moveCursor(-1)
	case key.Matches(msg, m.keys.MoveDown):
		m.moveCursor(1)
	case key.Matches(msg, m.keys.MoveLeft):
		m.moveVariation(-1)
	case key.Matches(msg, m.keys.MoveRight):
		m.moveVariation(1)
	case key.Matches(msg, m.keys.ProgressionRootL):
		m.moveProgressionRoot(-1)
	case key.Matches(msg, m.keys.ProgressionRootR):
		m.moveProgressionRoot(1)
	case key.Matches(msg, m.keys.RhythmLeft):
		m.moveRhythm(-1)
	case key.Matches(msg, m.keys.RhythmRight):
		m.moveRhythm(1)
	case key.Matches(msg, m.keys.Play):
		return updatePlayKey(m)
	case key.Matches(msg, m.keys.PlayScaleChords):
		return updateScaleChordPlaybackKey(m)
	case key.Matches(msg, m.keys.PageUp):
		m.mainViewport.PageUp()
	case key.Matches(msg, m.keys.PageDown):
		m.mainViewport.PageDown()
	case key.Matches(msg, m.keys.Home):
		m.mainViewport.GotoTop()
	case key.Matches(msg, m.keys.End):
		m.mainViewport.GotoBottom()
	}

	m.syncViewport()
	return m, nil
}

func updateHelpKey(m model, msg tea.KeyPressMsg) (model, tea.Cmd) {
	if key.Matches(msg, m.keys.Help, m.keys.Back, m.keys.Select) {
		m.showHelp = false
	}

	return m, nil
}

func updateBackKey(m model) (model, tea.Cmd) {
	if m.screen == screenChordBrowser || m.screen == screenScaleBrowser || m.screen == screenChordProgressions || m.screen == screenRiffs {
		m.screen = screenMainMenu
		m.syncViewport()
		return m, nil
	}

	return m, tea.Quit
}

func updatePlayKey(m model) (model, tea.Cmd) {
	if m.isPlaying {
		m.cancelPlayback()
		m.syncViewport()
		return m, nil
	}
	if m.screen == screenScaleBrowser {
		return m.startScaleNotePlayback()
	}

	cmd := m.playSelection()
	if cmd != nil {
		m.isPlaying = true
	}

	return m, cmd
}

func updateScaleChordPlaybackKey(m model) (model, tea.Cmd) {
	if m.screen != screenScaleBrowser {
		return m, nil
	}
	if m.isPlaying {
		m.cancelPlayback()
		m.syncViewport()
		return m, nil
	}

	return m.startScaleChordPlayback()
}
