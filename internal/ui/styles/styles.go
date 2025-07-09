package styles

import (
	"strings"

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

// Animation color palettes for the mascot
var animationPalettes = [][]lipgloss.Color{
	// Warm sunset
	{
		lipgloss.Color("#FF6B6B"), // Red
		lipgloss.Color("#FF8E53"), // Orange
		lipgloss.Color("#FF6B9D"), // Pink
		lipgloss.Color("#C44569"), // Dark pink
		lipgloss.Color("#F8B500"), // Yellow
	},
	// Cool ocean
	{
		lipgloss.Color("#4ECDC4"), // Teal
		lipgloss.Color("#45B7D1"), // Blue
		lipgloss.Color("#96CEB4"), // Green
		lipgloss.Color("#54C6EB"), // Light blue
		lipgloss.Color("#81C784"), // Light green
	},
	// Purple magic
	{
		lipgloss.Color("#A29BFE"), // Purple
		lipgloss.Color("#FD79A8"), // Pink
		lipgloss.Color("#E17055"), // Coral
		lipgloss.Color("#81ECEC"), // Cyan
		lipgloss.Color("#FDCB6E"), // Yellow
	},
	// Rainbow gradient
	{
		lipgloss.Color("#FF0080"), // Hot pink
		lipgloss.Color("#FF8000"), // Orange
		lipgloss.Color("#80FF00"), // Lime
		lipgloss.Color("#0080FF"), // Blue
		lipgloss.Color("#8000FF"), // Purple
	},
}

// CreateAnimatedMascot applies constantly changing colors to the mascot
func CreateAnimatedMascot(asciiArt string, animationFrame int) string {
	lines := strings.Split(asciiArt, "\n")

	// Select palette based on animation frame
	paletteIndex := (animationFrame / 10) % len(animationPalettes)
	palette := animationPalettes[paletteIndex]

	// Add some offset for color shifting within the palette
	colorOffset := animationFrame % len(palette)

	var styledLines []string
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			styledLines = append(styledLines, line)
			continue
		}

		// Calculate color index with animation offset
		colorIndex := (i + colorOffset) % len(palette)
		style := lipgloss.NewStyle().
			Bold(true).
			Foreground(palette[colorIndex])

		styledLines = append(styledLines, style.Render(line))
	}

	return strings.Join(styledLines, "\n")
}
