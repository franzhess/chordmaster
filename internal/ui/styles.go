package ui

import "charm.land/lipgloss/v2"

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	dashboardStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("238")).
			Padding(1, 1)

	barStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))

	instrumentGuitarStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("214"))

	audioPlayingStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("212"))

	helpPopupStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Padding(1, 3)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("212"))

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("230")).
			Background(lipgloss.Color("62")).
			Padding(0, 1)

	mutedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
)

var octaveStyles = map[int]lipgloss.Style{
	2: lipgloss.NewStyle().Foreground(lipgloss.Color("75")),
	3: lipgloss.NewStyle().Foreground(lipgloss.Color("112")),
	4: lipgloss.NewStyle().Foreground(lipgloss.Color("214")),
	5: lipgloss.NewStyle().Foreground(lipgloss.Color("212")),
}

func octaveStyle(octave int) lipgloss.Style {
	if style, ok := octaveStyles[octave]; ok {
		return style
	}

	return titleStyle
}
