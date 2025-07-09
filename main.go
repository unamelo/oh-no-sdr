package main

import (
	"log"
	// "os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/unamelo/oh-no-sdr/internal/ui/models"
)

func main() {
	p := tea.NewProgram(
		models.NewMainModel(),
		tea.WithAltScreen(), // Use full screen
	)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
