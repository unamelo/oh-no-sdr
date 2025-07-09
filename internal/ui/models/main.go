package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/unamelo/oh-no-sdr/internal/ui/styles"
)

type sessionState int

const (
	filePickerView sessionState = iota
	processingView
	resultsView
)

type MainModel struct {
	state      sessionState
	filePicker FilePickerModel
	progress   ProgressModel
	results    ResultsModel
	err        error
	width      int
	height     int
}

func NewMainModel() MainModel {
	return MainModel{
		state:      filePickerView,
		filePicker: NewFilePickerModel(),
		progress:   NewProgressModel(),
		results:    NewResultsModel(),
	}
}

func (m MainModel) Init() tea.Cmd {
	return m.filePicker.Init()
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	// Delegate to current view
	switch m.state {
	case filePickerView:
		newFilePicker, cmd := m.filePicker.Update(msg)
		m.filePicker = newFilePicker.(FilePickerModel)

		// Check if file was selected
		if m.filePicker.selectedFile != "" {
			m.state = processingView
			return m, tea.Batch(cmd, m.progress.StartProcessing(m.filePicker.selectedFile))
		}
		return m, cmd

	case processingView:
		newProgress, cmd := m.progress.Update(msg)
		m.progress = newProgress.(ProgressModel)

		// Check if processing is complete
		if m.progress.isComplete {
			m.state = resultsView
			m.results.SetResults(m.progress.results)
		}
		return m, cmd

	case resultsView:
		newResults, cmd := m.results.Update(msg)
		m.results = newResults.(ResultsModel)
		return m, cmd
	}

	return m, nil
}

func (m MainModel) View() string {
	header := styles.TitleStyle.Render("🚀 SDR Parser")
	subtitle := styles.SubtitleStyle.Render("Parse Student Data Return files with ease")

	var content string
	switch m.state {
	case filePickerView:
		content = m.filePicker.View()
	case processingView:
		content = m.progress.View()
	case resultsView:
		content = m.results.View()
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		subtitle,
		"",
		content,
	)
}
