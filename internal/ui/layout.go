package ui

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	v := tea.NewView(m.viewString())
	v.AltScreen = true
	return v
}

func (m model) viewString() string {
	if m.screen == screenSplash {
		content := m.splashView()
		if m.width <= 0 || m.height <= 0 {
			return content
		}

		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
	}

	return m.dashboardView()
}

func (m model) screenView() string {
	switch m.screen {
	case screenSplash:
		return m.splashView()
	case screenMainMenu:
		return m.mainMenuView()
	case screenChordBrowser:
		return m.chordBrowserView()
	case screenScaleBrowser:
		return m.scaleBrowserView()
	case screenChordProgressions:
		return m.chordProgressionsView()
	case screenRiffs:
		return m.riffsView()
	default:
		return ""
	}
}

func (m model) dashboardView() string {
	width := dashboardWidth(m.width)
	innerWidth := dashboardContentWidth(width)
	if innerWidth < 1 {
		innerWidth = width
	}
	bodyHeight := dashboardBodyHeight(m.height)
	m.syncViewport()
	body := m.mainViewport.View()
	if m.showHelp {
		body = lipgloss.Place(innerWidth, bodyHeight, lipgloss.Center, lipgloss.Center, m.helpPopupView(innerWidth))
	}

	content := strings.Join([]string{
		m.topBar(innerWidth),
		separator(innerWidth),
		lipgloss.NewStyle().Width(innerWidth).Height(bodyHeight).Render(body),
		separator(innerWidth),
		m.bottomBar(innerWidth),
	}, "\n")

	return dashboardStyle.Width(width).Render(content)
}

func (m model) topBar(width int) string {
	right := m.screenTitle()
	left := titleStyle.Render("Chordmaster")
	padding := width - lipgloss.Width(left) - lipgloss.Width(right)
	if padding < 1 {
		padding = 1
	}

	return barStyle.Width(width).Render(left + strings.Repeat(" ", padding) + mutedStyle.Render(right))
}

func (m model) bottomBar(width int) string {
	left := instrumentStyle(m.selectedInstrument).Render(m.selectedInstrument) + " " + mutedStyle.Render("Tabs: "+m.chordTabModeLabel())
	middle := " "
	if m.isPlaying {
		middle = audioPlayingStyle.Render("♪")
	}
	right := mutedStyle.Render("? help")

	sideWidth := width / 3
	centerWidth := width - (sideWidth * 2)
	line := lipgloss.JoinHorizontal(lipgloss.Top,
		lipgloss.NewStyle().Width(sideWidth).Align(lipgloss.Left).Render(left),
		lipgloss.NewStyle().Width(centerWidth).Align(lipgloss.Center).Render(middle),
		lipgloss.NewStyle().Width(sideWidth).Align(lipgloss.Right).Render(right),
	)

	return barStyle.Width(width).Render(line)
}

func (m model) screenTitle() string {
	switch m.screen {
	case screenMainMenu:
		return ""
	case screenChordBrowser:
		return "Chords"
	case screenScaleBrowser:
		return "Scales"
	case screenChordProgressions:
		return "Progressions"
	case screenRiffs:
		return "Riffs"
	default:
		return ""
	}
}

func (m model) helpPopupView(maxWidth int) string {
	popupWidth := maxWidth
	if popupWidth < 1 {
		popupWidth = 1
	}
	contentWidth := popupWidth - helpPopupStyle.GetHorizontalFrameSize()
	if contentWidth < 1 {
		contentWidth = 1
	}

	columns := []string{
		helpSection("Global", []helpLine{
			{"?", "toggle this help"},
			{"esc", "go back or close help"},
			{"q / ctrl+c", "quit"},
		}),
		helpSection("Navigation", []helpLine{
			{"j / k", "move primary selection"},
			{"h / l", "move secondary selection"},
			{"enter", "open selected screen"},
			{"pgup / pgdn", "scroll content"},
			{"home / end", "jump to top or bottom"},
		}),
		helpSection("Playback", []helpLine{
			{"p", "play selection or cancel"},
			{"c", "play scale chords"},
			{"f", "toggle chord notes/fingers"},
		}),
		helpSection("Progressions", []helpLine{
			{"[ / ]", "change root"},
			{", / .", "change rhythm"},
		}),
	}

	content := flowViews(columns, contentWidth, 4)

	return helpPopupStyle.Width(popupWidth).Render(titleStyle.Render("Help") + "\n\n" + content)
}

type helpLine struct {
	keys string
	desc string
}

func helpSection(title string, lines []helpLine) string {
	rows := []string{titleStyle.Render(title)}
	for _, line := range lines {
		rows = append(rows, mutedStyle.Render(line.keys)+"  "+line.desc)
	}

	return strings.Join(rows, "\n")
}

func (m *model) syncViewport() {
	if m.screen == screenSplash {
		return
	}

	width := dashboardContentWidth(dashboardWidth(m.width))
	if width < 1 {
		width = 1
	}
	height := dashboardBodyHeight(m.height)
	if height < 1 {
		height = 1
	}

	m.mainViewport.SetWidth(width)
	m.mainViewport.SetHeight(height)
	m.mainViewport.SoftWrap = false
	m.mainViewport.SetContent(m.screenView())
}

func instrumentStyle(instrument string) lipgloss.Style {
	switch instrument {
	case "Guitar":
		return instrumentGuitarStyle
	default:
		return titleStyle
	}
}

func separator(width int) string {
	return mutedStyle.Render(strings.Repeat("─", width))
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	}

	if value > max {
		return max
	}

	return value
}

func dashboardWidth(width int) int {
	const defaultWidth = 120
	const minWidth = 88

	if width <= 0 {
		return defaultWidth
	}

	if width < minWidth {
		return minWidth
	}

	return width
}

func dashboardContentWidth(width int) int {
	return width - dashboardStyle.GetHorizontalFrameSize()
}

func dashboardBodyHeight(height int) int {
	const defaultHeight = 120
	const minHeight = 20

	if height <= 0 {
		return defaultHeight
	}

	available := height - 8
	if available < minHeight {
		return minHeight
	}

	return available
}
