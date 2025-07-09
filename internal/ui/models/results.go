package models

import (
	tea "github.com/charmbracelet/bubbletea"
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
	return "Results: " + m.results
}

func (m *ResultsModel) SetResults(results string) {
	m.results = results
}
