package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colours
	Primary   = lipgloss.Color("#7C3AED")
	Secondary = lipgloss.Color("#10B981")
	Success   = lipgloss.Color("#059669")
	Warning   = lipgloss.Color("#D97706")
	Error     = lipgloss.Color("#DC2626")
	Muted     = lipgloss.Color("#6B7280")

	// Common styles
	TitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(Primary).
		MarginLeft(2)

	SubtitleStyle = lipgloss.NewStyle().
		Foreground(Muted).
		MarginLeft(2)

	BoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Primary).
		Padding(1, 2)

	SuccessStyle = lipgloss.NewStyle().
		Foreground(Success).
		Bold(true)

	ErrorStyle = lipgloss.NewStyle().
		Foreground(Error).
		Bold(true)
)
