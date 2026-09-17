package ui

import (
	"fmt"
	"strings"

	"chordmaster/internal/instruments/guitar"
	"chordmaster/internal/instruments/guitar/chordshapes"
	"chordmaster/internal/music"

	"charm.land/lipgloss/v2"
)

func (m model) splashView() string {
	title := titleStyle.Render(strings.Join([]string{
		`   ____ _                   _                     _            `,
		`  / ___| |__   ___  _ __ __| |_ __ ___   __ _ ___| |_ ___ _ __ `,
		` | |   | '_ \ / _ \| '__/ _` + "`" + ` | '_ ` + "`" + ` _ \ / _` + "`" + ` / __| __/ _ \ '__|`,
		` | |___| | | | (_) | | | (_| | | | | | | (_| \__ \ ||  __/ |   `,
		`  \____|_|_|_|\___/|_|__\__,_|_| |_| |_|\__,_|___/\__\___|_|   `,
		` |___ \ / _ \ / _ \ / _ \                                      `,
		`   __) | | | | | | | | | |                                     `,
		`  / __/| |_| | |_| | |_| |                                     `,
		` |_____|\___/ \___/ \___/                                      `,
	}, "\n"))

	lines := []string{
		title,
		"",
		mutedStyle.Render("Loading cassette riffs..."),
	}
	if m.splashRiff.Title() != "" {
		lines = append(lines,
			"",
			mutedStyle.Render("Now playing:"),
			selectedStyle.Render(m.splashRiff.Title()),
			mutedStyle.Render(fmt.Sprintf("BPM %d", m.splashRiff.BPM)),
		)
	}
	lines = append(lines, "", mutedStyle.Render("Press enter to skip."))

	return strings.Join(lines, "\n")
}

func (m model) mainMenuView() string {
	var b strings.Builder

	for i, item := range m.menuItems {
		line := item
		if i == m.menuCursor {
			line = selectedStyle.Render(line)
		}
		b.WriteString(line)
		b.WriteString("\n")
	}

	width := dashboardWidth(m.width) - 4
	if width < 1 {
		width = 1
	}

	return lipgloss.Place(width, dashboardBodyHeight(m.height), lipgloss.Center, lipgloss.Center, b.String())
}

func (m model) chordBrowserView() string {
	var b strings.Builder
	width := dashboardWidth(m.width) - 8
	if width < 80 {
		width = 80
	}

	context := m.chordContextPanel(width)
	b.WriteString(context)

	return b.String()
}

func (m model) chordContextPanel(width int) string {
	var b strings.Builder

	selectedChord := m.chordName(m.rootCursor, m.variationCursor)
	selectedVariation := m.chordCatalog.Variations[m.variationCursor]
	root := string(m.chordCatalog.Roots[m.rootCursor])
	scale := music.NewScale(pitchClassFromRoot(root), m.scalePatterns[0])
	chord := music.NewChord(scale, selectedVariation)

	b.WriteString(titleStyle.Render("Chord"))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Root: %s %s", selectedStyle.Render(root), mutedStyle.Render(fmt.Sprintf("(%d/%d)", m.rootCursor+1, len(m.chordCatalog.Roots)))))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Variation: %s %s", selectedStyle.Render(selectedVariation.Name), mutedStyle.Render(fmt.Sprintf("(%d/%d)", m.variationCursor+1, len(m.chordCatalog.Variations)))))
	b.WriteString("\n\n")
	b.WriteString(lipgloss.NewStyle().Width(width).Align(lipgloss.Center).Render(selectedChord))
	b.WriteString("\n")
	b.WriteString(mutedStyle.Width(width).Align(lipgloss.Center).Render(chord.NoteNamesString()))
	b.WriteString("\n\n")
	b.WriteString(mutedStyle.Width(width).Align(lipgloss.Center).Render(selectedVariation.FormulaString()))
	b.WriteString("\n\n")
	if voicings, ok := chordVoicings(root, selectedVariation); ok {
		b.WriteString(flowViews(m.chordShapeCards(selectedChord, chord.NoteNamesString(), voicings), width, 3))
	} else {
		b.WriteString(lipgloss.NewStyle().Width(width).Align(lipgloss.Center).Render(mutedStyle.Render("Tab: unavailable")))
	}

	return b.String()
}

func (m model) chordShapeCards(chordName string, notes string, voicings []chordshapes.Voicing) []string {
	cards := make([]string, 0, len(voicings))
	for _, voicing := range voicings {
		title := chordName
		if voicing.Source != "" {
			title = chordName + " " + mutedStyle.Render("("+voicing.Source+")")
		}
		cards = append(cards, chordCard(title, notes, renderChordTab(voicing, m.chordTabMode), false))
	}

	return cards
}

func (m model) scaleBrowserView() string {
	var b strings.Builder

	pattern := m.scalePatterns[m.scalePatternCursor]
	scale := music.NewScale(m.pitchClasses[m.scaleRootCursor], pattern)
	width := dashboardWidth(m.width) - 8
	if width < 80 {
		width = 80
	}

	context := m.scaleContextPanel(scale, pattern, width)
	if width >= 144 && len(scale.NoteNames) == 7 {
		leftWidth := 82
		rightWidth := width - leftWidth - 4
		context = m.scaleContextPanel(scale, pattern, leftWidth)
		triads := m.scaleTriadsPanel(scale, rightWidth)
		b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top,
			lipgloss.NewStyle().Width(leftWidth).Render(context),
			"    ",
			lipgloss.NewStyle().Width(rightWidth).Render(triads),
		))
	} else {
		b.WriteString(context)
		if len(scale.NoteNames) == 7 {
			b.WriteString("\n\n")
			triads := m.scaleTriadsPanel(scale, width)
			b.WriteString(triads)
		}
	}

	return b.String()
}

func (m model) scaleContextPanel(scale music.Scale, pattern music.ScalePattern, width int) string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("Scale"))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Root: %s %s", selectedStyle.Render(scale.Root.PreferredName(music.PreferSharps)), mutedStyle.Render(fmt.Sprintf("(%d/%d)", m.scaleRootCursor+1, len(m.pitchClasses)))))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Pattern: %s %s", selectedStyle.Render(pattern.Name), mutedStyle.Render(fmt.Sprintf("(%d/%d)", m.scalePatternCursor+1, len(m.scalePatterns)))))
	b.WriteString("\n\n")
	b.WriteString(lipgloss.NewStyle().Width(width).Align(lipgloss.Center).Render(m.scaleNoteNamesView(scale)))
	b.WriteString("\n")
	b.WriteString(mutedStyle.Width(width).Align(lipgloss.Center).Render(pattern.StepString()))
	b.WriteString("\n\n")
	b.WriteString(titleStyle.Render("Fretboard"))
	b.WriteString("\n")
	b.WriteString(m.scaleFretboardView(scale))

	return b.String()
}

func (m model) scaleTriadsPanel(scale music.Scale, width int) string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("Triads"))
	b.WriteString("\n")
	b.WriteString(flowViews(m.scaleTriadCards(scale), width, 3))

	return b.String()
}

func (m model) scaleTriadCards(scale music.Scale) []string {
	triads := music.ScaleChords(scale, music.SupportedScaleChordSchemas()[0])
	cards := make([]string, 0, len(triads))
	for i, triad := range triads {
		chordName, ok := triad.ChordName(m.chordCatalog)
		if !ok {
			chordName = "Unknown"
		}

		tab := "Tab: unavailable"
		if variation, ok := m.chordCatalog.ChordVariationByNotes(triad.NoteNames); ok {
			if rendered, ok := chordTab(triad.NoteNames[0], variation, m.chordTabMode); ok {
				tab = rendered
			}
		}

		cards = append(cards, chordCard(chordName, triad.NoteNamesString(), tab, i == m.activeScaleChord))
	}

	return cards
}

func (m model) scaleNoteNamesView(scale music.Scale) string {
	parts := make([]string, 0, len(scale.NoteNames))
	for i, noteName := range scale.NoteNames {
		if i == m.activeScaleNote {
			parts = append(parts, selectedStyle.Render(noteName))
			continue
		}
		parts = append(parts, noteName)
	}

	return strings.Join(parts, ", ")
}

func chordCard(name string, notes string, tab string, active bool) string {
	header := name
	if active {
		header = selectedStyle.Render(name)
	}

	return strings.Join([]string{
		header,
		mutedStyle.Render(notes),
		"",
		tab,
	}, "\n")
}

func (m model) scaleFretboardView(scale music.Scale) string {
	scalePitchClasses := make(map[string]bool, len(scale.PitchClasses))
	for _, pitchClass := range scale.PitchClasses {
		scalePitchClasses[pitchClass.PreferredName(music.PreferSharps)] = true
	}

	return guitar.StandardGuitarFretboard().RenderWithFormatter(func(position guitar.FretPosition) string {
		formatted := guitar.FormatFretNote(position.Note.PitchClass)
		if scalePitchClasses[position.Note.PitchClass.String()] {
			return octaveStyle(position.Note.Octave).Render(formatted)
		}

		return mutedStyle.Render(formatted)
	})
}

func (m model) visibleProgressionRange() (int, int) {
	const visibleProgressions = 10

	count := len(m.progressions)
	if count <= visibleProgressions {
		return 0, count
	}

	start := m.progressionCursor - (visibleProgressions / 2)
	start = clamp(start, 0, count-visibleProgressions)

	return start, start + visibleProgressions
}

func (m model) visibleRhythmRange() (int, int) {
	const visibleRhythms = 6

	count := len(m.rhythmPatterns)
	if count <= visibleRhythms {
		return 0, count
	}

	start := m.rhythmCursor - (visibleRhythms / 2)
	start = clamp(start, 0, count-visibleRhythms)

	return start, start + visibleRhythms
}

func (m model) chordName(rootIndex, variationIndex int) string {
	return m.chordCatalog.ChordName(rootIndex, variationIndex)
}

func (m model) chordTabModeLabel() string {
	if m.chordTabMode == chordTabFingers {
		return "fingers"
	}

	return "notes"
}

func (m model) chordProgressionsView() string {
	var b strings.Builder

	scale := music.NewScale(m.pitchClasses[m.scaleRootCursor], m.scalePatterns[m.scalePatternCursor])
	progression := m.progressions[m.progressionCursor]
	chords := music.ResolveChordProgression(scale, progression)
	width := dashboardWidth(m.width) - 8
	if width < 80 {
		width = 80
	}

	rhythm := m.rhythmPatterns[m.rhythmCursor]
	context := m.progressionContextPanel(scale, progression, rhythm, width)
	if width >= 120 {
		leftWidth := 58
		rightWidth := width - leftWidth - 4
		context = m.progressionContextPanel(scale, progression, rhythm, leftWidth)
		chordsPanel := m.progressionChordsPanel(chords, rightWidth)
		b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top,
			lipgloss.NewStyle().Width(leftWidth).Render(context),
			"    ",
			lipgloss.NewStyle().Width(rightWidth).Render(chordsPanel),
		))
	} else {
		b.WriteString(context)
		b.WriteString("\n\n")
		b.WriteString(m.progressionChordsPanel(chords, width))
	}

	return b.String()
}

func (m model) progressionContextPanel(scale music.Scale, progression music.ChordProgression, rhythm music.RhythmPattern, width int) string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("Selection"))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Root: %s %s", selectedStyle.Render(scale.Root.PreferredName(music.PreferSharps)), mutedStyle.Render(fmt.Sprintf("(%d/%d)", m.scaleRootCursor+1, len(m.pitchClasses)))))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Pattern: %s %s", selectedStyle.Render(scale.Pattern.Name), mutedStyle.Render(fmt.Sprintf("(%d/%d)", m.scalePatternCursor+1, len(m.scalePatterns)))))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Progression: %s %s", selectedStyle.Render(progression.Name), mutedStyle.Render(fmt.Sprintf("(%d/%d)", m.progressionCursor+1, len(m.progressions)))))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Rhythm: %s %s", selectedStyle.Render(rhythm.Name), mutedStyle.Render(fmt.Sprintf("(%d/%d)", m.rhythmCursor+1, len(m.rhythmPatterns)))))
	b.WriteString("\n\n")
	b.WriteString(lipgloss.NewStyle().Width(width).Align(lipgloss.Center).Render(strings.Join(progression.Symbols, " - ")))
	b.WriteString("\n")
	b.WriteString(mutedStyle.Width(width).Align(lipgloss.Center).Render(music.RhythmPatternLabel(rhythm)))
	b.WriteString("\n\n")

	b.WriteString(titleStyle.Render("Progressions"))
	b.WriteString("\n")
	start, end := m.visibleProgressionRange()
	for i := start; i < end; i++ {
		marker := "  "
		name := m.progressions[i].Name
		if i == m.progressionCursor {
			marker = "> "
			name = selectedStyle.Render(name)
		}

		b.WriteString(fmt.Sprintf("%s%-30s %s", marker, name, mutedStyle.Render(strings.Join(m.progressions[i].Symbols, " - "))))
		b.WriteString("\n")
	}
	b.WriteString("\n")

	b.WriteString(titleStyle.Render("Rhythms"))
	b.WriteString("\n")
	rhythmStart, rhythmEnd := m.visibleRhythmRange()
	for i := rhythmStart; i < rhythmEnd; i++ {
		marker := "  "
		name := m.rhythmPatterns[i].Name
		if i == m.rhythmCursor {
			marker = "> "
			name = selectedStyle.Render(name)
		}

		b.WriteString(fmt.Sprintf("%s%-28s %s", marker, name, mutedStyle.Render(music.RhythmPatternLabel(m.rhythmPatterns[i]))))
		b.WriteString("\n")
	}

	return b.String()
}

func (m model) progressionChordsPanel(chords []music.ProgressionChord, width int) string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("Chords"))
	b.WriteString("\n")
	b.WriteString(flowViews(m.progressionChordCards(chords), width, 3))

	return lipgloss.NewStyle().Width(width).Align(lipgloss.Center).Render(b.String())
}

func (m model) progressionChordCards(chords []music.ProgressionChord) []string {
	cards := make([]string, 0, len(chords))
	for i, chord := range chords {
		tab := "Tab: unavailable"
		if variation, ok := m.chordCatalog.ChordVariationByNotes(chord.NoteNames); ok {
			if rendered, ok := chordTab(chord.NoteNames[0], variation, m.chordTabMode); ok {
				tab = rendered
			}
		}

		cards = append(cards, strings.Join([]string{
			mutedStyle.Render(fmt.Sprintf("%d.", i+1)),
			chordCard(chord.Name, strings.Join(chord.NoteNames, ", "), tab, false),
		}, "\n"))
	}

	return cards
}

func (m model) riffsView() string {
	var b strings.Builder

	if len(m.riffs) == 0 {
		return strings.Join([]string{
			titleStyle.Render("Riffs"),
			mutedStyle.Render("No riffs available."),
		}, "\n")
	}

	selected := m.riffs[m.riffCursor]
	leftWidth := 38
	rightWidth := 72
	start, end := m.visibleRiffRange()

	var left strings.Builder
	left.WriteString(titleStyle.Render("Riffs"))
	left.WriteString("\n")
	for i := start; i < end; i++ {
		marker := "  "
		name := m.riffs[i].Title()
		if i == m.riffCursor {
			marker = "> "
			name = selectedStyle.Render(name)
		}

		left.WriteString(fmt.Sprintf("%s%s", marker, name))
		left.WriteString("\n")
	}

	var right strings.Builder
	right.WriteString(titleStyle.Render("Selected"))
	right.WriteString("\n")
	right.WriteString(selectedStyle.Render(selected.Song))
	right.WriteString("\n")
	right.WriteString(mutedStyle.Render(selected.Artist))
	right.WriteString("\n")
	right.WriteString(mutedStyle.Render(fmt.Sprintf("BPM: %d", selected.BPM)))
	right.WriteString("\n\n")
	right.WriteString(titleStyle.Render("Notes"))
	right.WriteString("\n")
	right.WriteString(wrapText(selected.Notation(), rightWidth))
	right.WriteString("\n\n")
	right.WriteString(mutedStyle.Render(fmt.Sprintf("Riff %d/%d", m.riffCursor+1, len(m.riffs))))

	b.WriteString(lipglossJoin(left.String(), right.String(), leftWidth, rightWidth))

	return b.String()
}

func (m model) visibleRiffRange() (int, int) {
	const visibleRiffs = 16

	count := len(m.riffs)
	if count <= visibleRiffs {
		return 0, count
	}

	start := m.riffCursor - (visibleRiffs / 2)
	start = clamp(start, 0, count-visibleRiffs)

	return start, start + visibleRiffs
}

func wrapText(text string, width int) string {
	if width <= 0 {
		return text
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return ""
	}

	var lines []string
	line := words[0]
	for _, word := range words[1:] {
		if len(line)+1+len(word) > width {
			lines = append(lines, line)
			line = word
			continue
		}

		line += " " + word
	}
	lines = append(lines, line)

	return strings.Join(lines, "\n")
}

func lipglossJoin(left, right string, leftWidth, rightWidth int) string {
	leftPanel := mutedStyle.Width(leftWidth).Render(left)
	rightPanel := mutedStyle.Width(rightWidth).Render(right)
	return lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, "  ", rightPanel)
}

func flowViews(items []string, maxWidth int, gap int) string {
	if len(items) == 0 {
		return ""
	}
	if maxWidth <= 0 {
		return lipgloss.JoinVertical(lipgloss.Left, items...)
	}
	if gap < 0 {
		gap = 0
	}

	spacing := strings.Repeat(" ", gap)
	rows := make([]string, 0)
	row := make([]string, 0)
	rowWidth := 0
	for _, item := range items {
		// Measure rendered strings instead of assuming chord diagrams are fixed-width;
		// shifted shapes include fret guides and can be wider than open-position tabs.
		itemWidth := lipgloss.Width(item)
		nextWidth := itemWidth
		if len(row) > 0 {
			nextWidth = rowWidth + gap + itemWidth
		}

		if len(row) > 0 && nextWidth > maxWidth {
			rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, row...))
			row = row[:0]
			rowWidth = 0
		}

		if len(row) > 0 {
			row = append(row, spacing)
			rowWidth += gap
		}
		row = append(row, item)
		rowWidth += itemWidth
	}

	if len(row) > 0 {
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, row...))
	}

	return strings.Join(rows, "\n\n")
}
