package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/unamelo/oh-no-sdr/internal/ui/styles"
)

type ResultsModel struct {
	results    string
	backToMenu bool
}

func NewResultsModel() ResultsModel {
	return ResultsModel{}
}

func (m ResultsModel) Init() tea.Cmd {
	return nil
}

func (m ResultsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "r", "enter":
			m.backToMenu = true
			return m, nil
		}
	}
	return m, nil
}

func (m ResultsModel) View() string {
	header := styles.SuccessStyle.Render(">> PARSING COMPLETE <<")
	content := styles.SubtitleStyle.Render("\nSUCCESS: Files parsed successfully!\n\nOutput: " + m.results + "\n\nPress [r] to return to menu or [q] to quit")
	
	return styles.BoxStyle.Render(header + "\n" + content)
}

func (m *ResultsModel) SetResults(results string) {
	m.results = results
}
