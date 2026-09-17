package ui

func (m *model) openSelectedScreen() {
	switch m.screen {
	case screenSplash:
		m.screen = screenMainMenu
	case screenMainMenu:
		switch m.menuCursor {
		case 0:
			m.screen = screenScaleBrowser
		case 1:
			m.screen = screenChordBrowser
		case 2:
			m.screen = screenChordProgressions
		default:
			m.screen = screenRiffs
		}
	}
}

func (m *model) toggleChordTabMode() {
	if m.chordTabMode == chordTabFingers {
		m.chordTabMode = chordTabNotes
		return
	}

	m.chordTabMode = chordTabFingers
}

func (m *model) moveCursor(direction int) {
	switch m.screen {
	case screenMainMenu:
		m.menuCursor = clamp(m.menuCursor+direction, 0, len(m.menuItems)-1)
	case screenChordBrowser:
		m.rootCursor = clamp(m.rootCursor+direction, 0, len(m.chordCatalog.Roots)-1)
	case screenScaleBrowser:
		m.scaleRootCursor = clamp(m.scaleRootCursor+direction, 0, len(m.pitchClasses)-1)
	case screenChordProgressions:
		m.progressionCursor = clamp(m.progressionCursor+direction, 0, len(m.progressions)-1)
	case screenRiffs:
		m.riffCursor = clamp(m.riffCursor+direction, 0, len(m.riffs)-1)
	}
}

func (m *model) moveVariation(direction int) {
	if m.screen == screenChordBrowser {
		m.variationCursor = clamp(m.variationCursor+direction, 0, len(m.chordCatalog.Variations)-1)
		return
	}

	if m.screen == screenScaleBrowser {
		m.scalePatternCursor = clamp(m.scalePatternCursor+direction, 0, len(m.scalePatterns)-1)
		return
	}

	if m.screen == screenChordProgressions {
		m.scalePatternCursor = clamp(m.scalePatternCursor+direction, 0, len(m.scalePatterns)-1)
	}
}

func (m *model) moveRhythm(direction int) {
	if m.screen == screenChordProgressions {
		m.rhythmCursor = clamp(m.rhythmCursor+direction, 0, len(m.rhythmPatterns)-1)
	}
}

func (m *model) moveProgressionRoot(direction int) {
	if m.screen == screenChordProgressions {
		m.scaleRootCursor = clamp(m.scaleRootCursor+direction, 0, len(m.pitchClasses)-1)
	}
}
