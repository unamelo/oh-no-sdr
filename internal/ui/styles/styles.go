package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colours (more vibrant like OH MY ZSH)
	Primary   = lipgloss.Color("#FF6B6B") // Bright red
	Secondary = lipgloss.Color("#4ECDC4") // Bright teal
	Accent    = lipgloss.Color("#45B7D1") // Bright blue
	Success   = lipgloss.Color("#96CEB4") // Bright green
	Warning   = lipgloss.Color("#FFEAA7") // Bright yellow
	Error     = lipgloss.Color("#FD79A8") // Bright pink
	Muted     = lipgloss.Color("#A29BFE") // Bright purple

	// ASCII Art Style
	ASCIIStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Primary).
			MarginLeft(1).
			MarginRight(1).
			MarginTop(0).
			MarginBottom(0)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(Secondary).
			Bold(true).
			MarginLeft(1).
			MarginBottom(0)

	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Accent).
			Padding(1, 2).
			Margin(0, 1)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(Success).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(Error).
			Bold(true)

	HighlightStyle = lipgloss.NewStyle().
			Foreground(Warning).
			Bold(true)
)
