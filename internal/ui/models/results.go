package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/unamelo/oh-no-sdr/internal/ui/styles"
)

type ResultsModel struct {
	results string
}

func NewResultsModel() ResultsModel {
	return ResultsModel{}
}

func (m ResultsModel) Init() tea.Cmd {
	return nil
}

func (m ResultsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// TODO: Implement results navigation
	return m, nil
}

func (m ResultsModel) View() string {
	header := styles.SuccessStyle.Render(">> PARSING COMPLETE <<")
	content := styles.SubtitleStyle.Render("\nSUCCESS: File parsed successfully!\n\nOutput: " + m.results + "\n\nPress [q] to quit or [r] to restart")
	
	return styles.BoxStyle.Render(header + "\n" + content)
}

func (m *ResultsModel) SetResults(results string) {
	m.results = results
}
