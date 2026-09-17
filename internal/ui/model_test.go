package ui

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
)

func keyPress(code rune) tea.KeyPressMsg {
	return tea.KeyPressMsg(tea.Key{Code: code})
}

func runePress(r rune) tea.KeyPressMsg {
	return tea.KeyPressMsg(tea.Key{Code: r, Text: string(r)})
}

func viewString(m model) string {
	return ansi.Strip(m.View().Content)
}

func styledViewString(m model) string {
	return m.View().Content
}

func TestCursorMovesWithinBounds(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenChordBrowser

	updated, _ := m.Update(keyPress(tea.KeyUp))
	m = updated.(model)
	if m.rootCursor != 0 {
		t.Fatalf("cursor moved above first item: %d", m.rootCursor)
	}

	for range m.chordCatalog.Roots {
		updated, _ = m.Update(keyPress(tea.KeyDown))
		m = updated.(model)
	}

	want := len(m.chordCatalog.Roots) - 1
	if m.rootCursor != want {
		t.Fatalf("cursor = %d, want %d", m.rootCursor, want)
	}
}

func TestVariationCursorMovesWithinBounds(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenChordBrowser

	updated, _ := m.Update(keyPress(tea.KeyLeft))
	m = updated.(model)
	if m.variationCursor != 0 {
		t.Fatalf("variation cursor moved before first item: %d", m.variationCursor)
	}

	for range m.chordCatalog.Variations {
		updated, _ = m.Update(keyPress(tea.KeyRight))
		m = updated.(model)
	}

	want := len(m.chordCatalog.Variations) - 1
	if m.variationCursor != want {
		t.Fatalf("variation cursor = %d, want %d", m.variationCursor, want)
	}
}

func TestScaleBrowserCursorsMoveWithinBounds(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenScaleBrowser

	updated, _ := m.Update(keyPress(tea.KeyUp))
	m = updated.(model)
	if m.scaleRootCursor != 0 {
		t.Fatalf("scale root cursor moved above first item: %d", m.scaleRootCursor)
	}

	for range m.pitchClasses {
		updated, _ = m.Update(keyPress(tea.KeyDown))
		m = updated.(model)
	}

	if want := len(m.pitchClasses) - 1; m.scaleRootCursor != want {
		t.Fatalf("scale root cursor = %d, want %d", m.scaleRootCursor, want)
	}

	updated, _ = m.Update(keyPress(tea.KeyLeft))
	m = updated.(model)
	if m.scalePatternCursor != 0 {
		t.Fatalf("scale pattern cursor moved before first item: %d", m.scalePatternCursor)
	}

	for range m.scalePatterns {
		updated, _ = m.Update(keyPress(tea.KeyRight))
		m = updated.(model)
	}

	if want := len(m.scalePatterns) - 1; m.scalePatternCursor != want {
		t.Fatalf("scale pattern cursor = %d, want %d", m.scalePatternCursor, want)
	}
}

func TestSplashAdvancesToMainMenu(t *testing.T) {
	m := NewModel().(model)

	updated, _ := m.Update(splashTimeoutMsg{})
	m = updated.(model)

	if m.screen != screenMainMenu {
		t.Fatalf("screen = %v, want %v", m.screen, screenMainMenu)
	}
}

func TestSplashCanBeSkipped(t *testing.T) {
	m := NewModel().(model)

	updated, _ := m.Update(keyPress(tea.KeyEnter))
	m = updated.(model)

	if m.screen != screenMainMenu {
		t.Fatalf("screen = %v, want %v", m.screen, screenMainMenu)
	}
}

func TestMainMenuOpensChordBrowser(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenMainMenu
	m.menuCursor = 1

	updated, _ := m.Update(keyPress(tea.KeyEnter))
	m = updated.(model)

	if m.screen != screenChordBrowser {
		t.Fatalf("screen = %v, want %v", m.screen, screenChordBrowser)
	}
}

func TestMainMenuOpensChordProgressions(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenMainMenu
	m.menuCursor = 2

	updated, _ := m.Update(keyPress(tea.KeyEnter))
	m = updated.(model)

	if m.screen != screenChordProgressions {
		t.Fatalf("screen = %v, want %v", m.screen, screenChordProgressions)
	}
}

func TestMainMenuOpensRiffs(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenMainMenu
	m.menuCursor = 3

	updated, _ := m.Update(keyPress(tea.KeyEnter))
	m = updated.(model)

	if m.screen != screenRiffs {
		t.Fatalf("screen = %v, want %v", m.screen, screenRiffs)
	}
}

func TestMainMenuOpensScaleBrowser(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenMainMenu

	updated, _ := m.Update(keyPress(tea.KeyEnter))
	m = updated.(model)

	if m.screen != screenScaleBrowser {
		t.Fatalf("screen = %v, want %v", m.screen, screenScaleBrowser)
	}
}

func TestViewIncludesSelectedChord(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenChordBrowser
	view := viewString(m)

	if strings.Contains(view, "Chord Browser") || strings.Contains(view, "Select a root and chord variation") {
		t.Fatal("view includes removed chord browser heading")
	}

	if !strings.Contains(view, "Root") || !strings.Contains(view, "Variation:") || !strings.Contains(view, "Major") {
		t.Fatal("view does not include selected chord context")
	}

	if strings.Contains(view, "C6") || strings.Contains(view, "Major6") {
		t.Fatal("view includes unselected chord variations")
	}

	if strings.Contains(view, "RootMajor") || strings.Contains(view, "----------") {
		t.Fatal("view includes old chord table")
	}

	if !strings.Contains(view, "Variation:") || !strings.Contains(view, "(1/12)") || !strings.Contains(view, "(1/") || !strings.Contains(view, "1,3,5") {
		t.Fatal("view does not include selected chord context")
	}

	if !strings.Contains(view, "Tabs: notes") || !strings.Contains(view, "e E||---|---|---|") || !strings.Contains(view, "E X||---|---|---|") {
		t.Fatal("view does not include selected chord tab")
	}
	if strings.Contains(view, "Tab: notes") {
		t.Fatal("view includes inline tab mode label")
	}

	if !strings.Contains(view, m.chordName(m.rootCursor, m.variationCursor)) {
		t.Fatal("view does not include selected chord")
	}

	if !strings.Contains(view, "Guitar") {
		t.Fatal("view does not include selected instrument")
	}

	if strings.Contains(view, "Instrument: Guitar") {
		t.Fatal("view includes verbose instrument label")
	}
}

func TestChordTabFingeringToggle(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenChordBrowser

	updated, _ := m.Update(runePress('f'))
	m = updated.(model)
	view := viewString(m)

	if !strings.Contains(view, "Tabs: fingers") || !strings.Contains(view, "B  ||-1-|---|---|") || !strings.Contains(view, "A  ||---|---|-3-|") {
		t.Fatal("view does not include fingering tab after toggle")
	}

	updated, _ = m.Update(runePress('f'))
	m = updated.(model)
	if m.chordTabMode != chordTabNotes {
		t.Fatal("second toggle did not return to note tab mode")
	}
}

func TestViewIncludesSelectedScale(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenScaleBrowser
	view := viewString(m)

	if strings.Contains(view, "Scale Browser") || strings.Contains(view, "Select a root and scale pattern") {
		t.Fatal("view includes removed scale browser heading")
	}

	if !strings.Contains(view, "Root:") || !strings.Contains(view, "(1/12)") || !strings.Contains(view, "Pattern:") || !strings.Contains(view, "(1/16)") {
		t.Fatal("view does not include scale selection counts")
	}

	if !strings.Contains(view, "Major (Ionian)") {
		t.Fatal("view does not include selected scale pattern")
	}

	if !strings.Contains(view, "C, D, E, F, G, A, B") {
		t.Fatal("view does not include concrete scale")
	}

	if !strings.Contains(view, "Triads") {
		t.Fatal("view does not include triads section")
	}

	if !strings.Contains(view, "C") || !strings.Contains(view, "D, F, A") || !strings.Contains(view, "Bdim") {
		t.Fatal("view does not include expected scale triads")
	}

	if !strings.Contains(view, "e E||---|---|---|") || !strings.Contains(view, "E X||---|---|---|") {
		t.Fatal("view does not include triad tabs")
	}

	if !strings.Contains(view, "Fretboard") || !strings.Contains(view, "e") || !strings.Contains(view, "E") {
		t.Fatal("view does not include guitar fretboard")
	}
}

func TestWideScaleLayoutFlowsTriadsWithinRightPanel(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenScaleBrowser
	m.width = 160
	view := viewString(m)

	if strings.Contains(view, "C        C, E, G    Dm       D, F, A    Em       E, G, B    F") {
		t.Fatal("wide scale layout packs too many triad cards into the right panel")
	}
}

func TestScaleBrowserOmitsTriadsForNonHeptatonicScales(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenScaleBrowser
	m.scalePatternCursor = 9
	view := styledViewString(m)

	if strings.Contains(view, "Triads") {
		t.Fatal("view includes triads for non-heptatonic scale")
	}
}

func TestScaleBrowserHighlightsActivePlaybackNote(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenScaleBrowser
	m.activeScaleNote = 1
	view := styledViewString(m)

	if !strings.Contains(view, selectedStyle.Render("D")) {
		t.Fatal("view does not highlight active scale note")
	}
}

func TestScaleBrowserHighlightsActivePlaybackChord(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenScaleBrowser
	m.activeScaleChord = 1
	view := styledViewString(m)

	if !strings.Contains(ansi.Strip(view), "Dm") || !strings.Contains(view, "\x1b[") {
		t.Fatal("view does not highlight active scale chord")
	}
}

func TestPressingPCancelsPlayback(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenScaleBrowser
	m.isPlaying = true
	m.playbackKind = playbackKindScaleNotes
	m.playbackSession = 1
	m.playbackCancel = make(chan struct{})
	m.activeScaleNote = 2

	updated, _ := m.Update(runePress('p'))
	m = updated.(model)

	if m.isPlaying {
		t.Fatal("playback did not stop")
	}
	if m.activeScaleNote != -1 || m.activeScaleChord != -1 {
		t.Fatal("playback highlights were not cleared")
	}
	if m.playbackKind != playbackKindNone {
		t.Fatal("playback kind was not reset")
	}
}

func TestFlowViewsWrapsVariableWidthItems(t *testing.T) {
	rendered := flowViews([]string{"short", "much-wider", "tiny"}, 16, 2)
	lines := strings.Split(rendered, "\n")

	if len(lines) != 3 {
		t.Fatalf("flow layout line count = %d, want 3", len(lines))
	}
	if strings.TrimRight(lines[0], " ") != "short" {
		t.Fatalf("first row = %q, want %q", lines[0], "short")
	}
	if lines[1] != "" {
		t.Fatalf("second row = %q, want blank padding row", lines[1])
	}
	if strings.TrimRight(lines[2], " ") != "much-wider  tiny" {
		t.Fatalf("third row = %q, want %q", lines[2], "much-wider  tiny")
	}
}

func TestProgressionBrowserCursorsMoveWithinBounds(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenChordProgressions

	updated, _ := m.Update(keyPress(tea.KeyUp))
	m = updated.(model)
	if m.progressionCursor != 0 {
		t.Fatalf("progression cursor moved above first item: %d", m.progressionCursor)
	}

	for range m.progressions {
		updated, _ = m.Update(keyPress(tea.KeyDown))
		m = updated.(model)
	}

	if want := len(m.progressions) - 1; m.progressionCursor != want {
		t.Fatalf("progression cursor = %d, want %d", m.progressionCursor, want)
	}

	updated, _ = m.Update(keyPress(tea.KeyLeft))
	m = updated.(model)
	if m.scalePatternCursor != 0 {
		t.Fatalf("scale pattern cursor moved before first item: %d", m.scalePatternCursor)
	}

	for range m.scalePatterns {
		updated, _ = m.Update(keyPress(tea.KeyRight))
		m = updated.(model)
	}

	if want := len(m.scalePatterns) - 1; m.scalePatternCursor != want {
		t.Fatalf("scale pattern cursor = %d, want %d", m.scalePatternCursor, want)
	}

	updated, _ = m.Update(runePress(','))
	m = updated.(model)
	if m.rhythmCursor != 0 {
		t.Fatalf("rhythm cursor moved before first item: %d", m.rhythmCursor)
	}

	for range m.rhythmPatterns {
		updated, _ = m.Update(runePress('.'))
		m = updated.(model)
	}

	if want := len(m.rhythmPatterns) - 1; m.rhythmCursor != want {
		t.Fatalf("rhythm cursor = %d, want %d", m.rhythmCursor, want)
	}
}

func TestRiffBrowserCursorMovesWithinBounds(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenRiffs

	updated, _ := m.Update(keyPress(tea.KeyUp))
	m = updated.(model)
	if m.riffCursor != 0 {
		t.Fatalf("riff cursor moved above first item: %d", m.riffCursor)
	}

	for range m.riffs {
		updated, _ = m.Update(keyPress(tea.KeyDown))
		m = updated.(model)
	}

	if want := len(m.riffs) - 1; m.riffCursor != want {
		t.Fatalf("riff cursor = %d, want %d", m.riffCursor, want)
	}
}

func TestViewIncludesSelectedProgression(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenChordProgressions
	m.progressionCursor = 5
	view := viewString(m)

	if strings.Contains(view, "Chord Progressions") || strings.Contains(view, "Select a scale and progression") {
		t.Fatal("view includes removed progression heading")
	}

	if !strings.Contains(view, "Root:") || !strings.Contains(view, "Pattern:") || !strings.Contains(view, "Progression:") || !strings.Contains(view, "Rhythm:") {
		t.Fatal("view does not include progression selection context")
	}

	if !strings.Contains(view, "I-V-vi-IV") || !strings.Contains(view, "C") || !strings.Contains(view, "G") || !strings.Contains(view, "Am") || !strings.Contains(view, "F") {
		t.Fatal("view does not include selected progression chords")
	}

	if !strings.Contains(view, "Tabs: notes") || !strings.Contains(view, "e E||---|---|---|") {
		t.Fatal("view does not include progression chord tabs")
	}
	if strings.Contains(view, "Tab: notes") {
		t.Fatal("view includes inline tab mode label")
	}

	if !strings.Contains(view, "Whole Measure Chords") || !strings.Contains(view, "4") {
		t.Fatal("view does not include selected rhythm pattern")
	}
}

func TestWideProgressionLayoutUsesTwoColumns(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenChordProgressions
	m.progressionCursor = 5
	m.width = 200
	view := viewString(m)

	if !strings.Contains(view, "Selection") || !strings.Contains(view, "Chords") || !strings.Contains(view, "Progressions") || !strings.Contains(view, "Rhythms") {
		t.Fatal("wide progression layout does not include both columns")
	}
}

func TestViewIncludesSelectedRiff(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenRiffs
	view := viewString(m)

	if !strings.Contains(view, "Riffs") {
		t.Fatal("view does not include riffs title")
	}

	if !strings.Contains(view, "Midnight Climb") || !strings.Contains(view, "Chordmaster") || !strings.Contains(view, "BPM: 108") || !strings.Contains(view, "E2(e)") {
		t.Fatal("view does not include selected riff")
	}
}

func TestSplashPlaybackCompletionAdvancesToMainMenu(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenSplash
	m.isPlaying = true

	updated, _ := m.Update(playbackFinishedMsg{source: playbackSplash})
	m = updated.(model)
	if m.screen != screenMainMenu {
		t.Fatalf("screen = %v, want %v", m.screen, screenMainMenu)
	}
}

func TestGeneralPlaybackCompletionDoesNotNavigate(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenRiffs
	m.isPlaying = true

	updated, _ := m.Update(playbackFinishedMsg{source: playbackGeneral})
	m = updated.(model)
	if m.screen != screenRiffs {
		t.Fatalf("screen = %v, want %v", m.screen, screenRiffs)
	}
}

func TestBottomBarIncludesInstrumentAndPlaybackIndicator(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenMainMenu

	view := viewString(m)
	if !strings.Contains(view, "Guitar") {
		t.Fatal("bottom bar does not include selected instrument")
	}

	if strings.Contains(view, "♪") {
		t.Fatal("audio indicator is visible while idle")
	}

	if !strings.Contains(view, "? help") {
		t.Fatal("bottom bar does not include help hint")
	}

	if !strings.Contains(view, "Tabs: notes") {
		t.Fatal("bottom bar does not include tab mode")
	}

	m.isPlaying = true
	view = viewString(m)
	if !strings.Contains(view, "♪") {
		t.Fatal("audio indicator is hidden while playing")
	}
}

func TestPlaybackFinishedClearsStatus(t *testing.T) {
	m := NewModel().(model)
	m.isPlaying = true
	m.status = "Playback finished."

	updated, _ := m.Update(playbackFinishedMsg{})
	m = updated.(model)
	if m.isPlaying {
		t.Fatal("playback state did not clear")
	}

	if m.status != "" {
		t.Fatalf("status = %q, want empty", m.status)
	}
}

func TestHelpPopupTogglesWithQuestionMark(t *testing.T) {
	m := NewModel().(model)
	m.screen = screenChordProgressions

	updated, _ := m.Update(runePress('?'))
	m = updated.(model)
	if !m.showHelp {
		t.Fatal("help popup did not open")
	}

	view := viewString(m)
	if !strings.Contains(view, "Help") || !strings.Contains(view, "j / k") || !strings.Contains(view, "Playback") || !strings.Contains(view, "p") {
		t.Fatal("help popup does not include keybindings")
	}

	updated, _ = m.Update(runePress('?'))
	m = updated.(model)
	if m.showHelp {
		t.Fatal("help popup did not close")
	}
}
