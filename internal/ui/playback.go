package ui

import (
	"errors"
	"time"

	"chordmaster/internal/music"
	"chordmaster/internal/synth"

	tea "charm.land/bubbletea/v2"
)

var errAudioUnavailable = errors.New("audio is not available")

func (m *model) cancelPlayback() {
	if m.playbackCancel != nil {
		close(m.playbackCancel)
		m.playbackCancel = nil
	}
	m.isPlaying = false
	m.playbackKind = playbackKindNone
	m.playbackSession++
	m.clearPlaybackHighlight()
	m.status = ""
}

func (m *model) finishPlayback(err error) {
	m.isPlaying = false
	m.playbackKind = playbackKindNone
	m.playbackCancel = nil
	m.clearPlaybackHighlight()
	if err != nil && !errors.Is(err, synth.ErrPlaybackCanceled) {
		m.status = "Playback failed: " + err.Error()
		return
	}
	m.status = ""
}

func (m *model) clearPlaybackHighlight() {
	m.activeScaleNote = -1
	m.activeScaleChord = -1
}

func (m model) beginCancellablePlayback(kind playbackKind) model {
	m.playbackSession++
	m.playbackKind = kind
	m.playbackCancel = make(chan struct{})
	m.isPlaying = true
	m.clearPlaybackHighlight()
	return m
}

func audioUnavailableCmd() tea.Cmd {
	return func() tea.Msg {
		return playbackFinishedMsg{err: errAudioUnavailable, source: playbackGeneral}
	}
}

func (m model) playSelection() tea.Cmd {
	if m.audio == nil {
		return audioUnavailableCmd()
	}

	switch m.screen {
	case screenScaleBrowser:
		return nil
	case screenChordBrowser:
		root := pitchClassFromRoot(string(m.chordCatalog.Roots[m.rootCursor]))
		scale := music.NewScale(root, m.scalePatterns[0])
		chord := music.NewChord(scale, m.chordCatalog.Variations[m.variationCursor])
		return func() tea.Msg {
			return playbackFinishedMsg{err: m.audio.PlayChord(chord.NoteNames), source: playbackGeneral}
		}
	case screenChordProgressions:
		scale := music.NewScale(m.pitchClasses[m.scaleRootCursor], m.scalePatterns[m.scalePatternCursor])
		progression := m.progressions[m.progressionCursor]
		rhythm := m.rhythmPatterns[m.rhythmCursor]
		resolved := music.ResolveChordProgression(scale, progression)
		chords := make([][]string, 0, len(resolved))
		durations := make([]time.Duration, 0, len(resolved))
		for _, chord := range resolved {
			chords = append(chords, chord.NoteNames)
		}
		for i := range resolved {
			unit := rhythm.Units[i%len(rhythm.Units)]
			durations = append(durations, time.Duration(unit)*synth.ProgressionBeatTime)
		}
		return func() tea.Msg {
			return playbackFinishedMsg{err: m.audio.PlayTimedChordSequence(chords, durations), source: playbackGeneral}
		}
	case screenRiffs:
		if len(m.riffs) == 0 {
			return nil
		}
		return m.playRiff(m.riffs[m.riffCursor], playbackGeneral)
	default:
		return nil
	}
}

func (m model) startScaleNotePlayback() (model, tea.Cmd) {
	if m.audio == nil {
		return m, audioUnavailableCmd()
	}

	m = m.beginCancellablePlayback(playbackKindScaleNotes)
	return m.advanceScaleNotePlayback(0)
}

func (m model) advanceScaleNotePlayback(index int) (model, tea.Cmd) {
	scale := music.NewScale(m.pitchClasses[m.scaleRootCursor], m.scalePatterns[m.scalePatternCursor])
	if index >= len(scale.NoteNames) {
		m.finishPlayback(nil)
		return m, nil
	}

	m.activeScaleNote = index
	notes := synth.NoteNamesWithOctaves(scale.NoteNames, 4)
	note := notes[index]
	session := m.playbackSession
	cancel := m.playbackCancel
	return m, func() tea.Msg {
		freq, err := synth.NoteFrequency(note)
		if err == nil {
			err = m.audio.PlayToneCancellable(freq, synth.ScaleNoteTime, cancel)
		}
		return playbackStepFinishedMsg{err: err, kind: playbackKindScaleNotes, session: session, index: index}
	}
}

func (m model) startScaleChordPlayback() (model, tea.Cmd) {
	if m.audio == nil {
		return m, audioUnavailableCmd()
	}

	m = m.beginCancellablePlayback(playbackKindScaleChords)
	return m.advanceScaleChordPlayback(0)
}

func (m model) advanceScaleChordPlayback(index int) (model, tea.Cmd) {
	scale := music.NewScale(m.pitchClasses[m.scaleRootCursor], m.scalePatterns[m.scalePatternCursor])
	chords := music.ScaleChords(scale, music.SupportedScaleChordSchemas()[0])
	if index >= len(chords) {
		m.finishPlayback(nil)
		return m, nil
	}

	m.activeScaleChord = index
	session := m.playbackSession
	cancel := m.playbackCancel
	notes := chords[index].NoteNames
	return m, func() tea.Msg {
		err := m.audio.PlayChordCancellable(notes, synth.ChordNoteTime, cancel)
		return playbackStepFinishedMsg{err: err, kind: playbackKindScaleChords, session: session, index: index}
	}
}

func (m model) playRiff(riff music.Riff, source playbackSource) tea.Cmd {
	notes := riff.TimedNotes()
	timedNotes := make([]synth.TimedNote, 0, len(notes))
	for _, note := range notes {
		timedNotes = append(timedNotes, synth.TimedNote{Note: note.Note, Duration: note.Duration})
	}

	return func() tea.Msg { return playbackFinishedMsg{err: m.audio.PlayTimedNotes(timedNotes), source: source} }
}
