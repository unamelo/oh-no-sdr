package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colours (more vibrant like OH MY ZSH)
	Primary   = lipgloss.Color("#FF6B6B")  // Bright red
	Secondary = lipgloss.Color("#4ECDC4")  // Bright teal
	Accent    = lipgloss.Color("#45B7D1")  // Bright blue
	Success   = lipgloss.Color("#96CEB4")  // Bright green
	Warning   = lipgloss.Color("#FFEAA7")  // Bright yellow
	Error     = lipgloss.Color("#FD79A8")  // Bright pink
	Muted     = lipgloss.Color("#A29BFE")  // Bright purple

	// ASCII Art Style
	ASCIIStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(Primary).
		MarginLeft(4).
		MarginRight(4).
		MarginTop(1).
		MarginBottom(1)

	SubtitleStyle = lipgloss.NewStyle().
		Foreground(Secondary).
		Bold(true).
		MarginLeft(4).
		MarginBottom(2)

	BoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Accent).
		Padding(3, 6).
		Margin(2, 4).
		Width(100)

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
