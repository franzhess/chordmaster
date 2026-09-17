package music

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type NoteEvent struct {
	Note     string
	Duration string
}

type Riff struct {
	Artist string
	Song   string
	BPM    int
	Events []NoteEvent
}

type TimedNote struct {
	Note     string
	Duration time.Duration
}

var riffHeaderPattern = regexp.MustCompile(`^\s*\d+\.\s+(.+?)\s+—\s+(.+?)\s+—\s+BPM:\s+(\d+)\s*$`)
var noteEventPattern = regexp.MustCompile(`(R|[A-G](?:#|b)?\d+)\((s|e|q|h|w)\)`)

func ParseRiffCatalog(content string) ([]Riff, error) {
	lines := strings.Split(content, "\n")
	riffs := make([]Riff, 0, 100)
	for i := 0; i < len(lines); i++ {
		matches := riffHeaderPattern.FindStringSubmatch(lines[i])
		if matches == nil {
			continue
		}

		bpm, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, err
		}

		if i+1 >= len(lines) {
			return nil, fmt.Errorf("missing notation for riff %q", matches[2])
		}

		events := parseRiffEvents(lines[i+1])
		if len(events) == 0 {
			return nil, fmt.Errorf("missing note events for riff %q", matches[2])
		}

		riffs = append(riffs, Riff{Artist: matches[1], Song: matches[2], BPM: bpm, Events: events})
	}

	if len(riffs) == 0 {
		return nil, fmt.Errorf("no riffs found")
	}

	return riffs, nil
}

func (r Riff) TimedNotes() []TimedNote {
	notes := make([]TimedNote, 0, len(r.Events))
	for _, event := range r.Events {
		notes = append(notes, TimedNote{Note: event.Note, Duration: RiffDuration(event.Duration, r.BPM)})
	}

	return notes
}

func (r Riff) Notation() string {
	parts := make([]string, 0, len(r.Events))
	for _, event := range r.Events {
		parts = append(parts, fmt.Sprintf("%s(%s)", event.Note, event.Duration))
	}

	return strings.Join(parts, " ")
}

func (r Riff) Title() string {
	if r.Artist == "" {
		return r.Song
	}

	if r.Song == "" {
		return r.Artist
	}

	return r.Artist + " - " + r.Song
}

func RiffDuration(symbol string, bpm int) time.Duration {
	if bpm <= 0 {
		return 0
	}

	quarter := time.Minute / time.Duration(bpm)
	switch symbol {
	case "s":
		return quarter / 4
	case "e":
		return quarter / 2
	case "q":
		return quarter
	case "h":
		return quarter * 2
	case "w":
		return quarter * 4
	default:
		return 0
	}
}

func parseRiffEvents(line string) []NoteEvent {
	matches := noteEventPattern.FindAllStringSubmatch(line, -1)
	events := make([]NoteEvent, 0, len(matches))
	for _, match := range matches {
		events = append(events, NoteEvent{Note: match[1], Duration: match[2]})
	}

	return events
}
