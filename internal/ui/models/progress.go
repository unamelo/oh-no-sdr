package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ProgressModel struct {
	isComplete bool
	results    string
}

func NewProgressModel() ProgressModel {
	return ProgressModel{}
}

func (m ProgressModel) Init() tea.Cmd {
	return nil
}

func (m ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// TODO: Implement progress logic
	return m, nil
}

func (m ProgressModel) View() string {
	return "Processing... (TODO: Implement progress bar)"
}

func (m ProgressModel) StartProcessing(filename string) tea.Cmd {
	// TODO: Return command to start processing
	return nil
}
