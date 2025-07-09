package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/unamelo/oh-no-sdr/internal/ui/styles"
)

type sessionState int

const (
	menuView sessionState = iota
	filePickerView
	processingView
	resultsView
)

type MainModel struct {
	state      sessionState
	menu       MenuModel
	filePicker FilePickerModel
	progress   ProgressModel
	results    ResultsModel
	err        error
	width      int
	height     int
	currentFileType string
	filesToProcess  []string
}

func NewMainModel() MainModel {
	return MainModel{
		state:      menuView,
		menu:       NewMenuModel(),
		filePicker: NewFilePickerModel(),
		progress:   NewProgressModel(),
		results:    NewResultsModel(),
	}
}

func (m MainModel) Init() tea.Cmd {
	return m.menu.Init()
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// Update child models with new size
		m.filePicker.width = msg.Width
		m.filePicker.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	// Delegate to current view
	switch m.state {
	case menuView:
		newMenu, cmd := m.menu.Update(msg)
		m.menu = newMenu.(MenuModel)

		// Check if option was selected
		if m.menu.selectedIndex >= 0 {
			fileType, files, err := m.menu.GetSelectedOption()
			if err != nil {
				m.err = err
				return m, nil
			}
			
			m.currentFileType = fileType
			m.filesToProcess = files
			
			// If files were found, go directly to processing
			if len(files) > 0 {
				m.state = processingView
				comparisonEnabled := m.menu.GetGenerateComparison()
				return m, tea.Batch(cmd, m.progress.StartProcessingMultipleWithComparison(files, comparisonEnabled))
			} else {
				// No files found, show file picker
				m.state = filePickerView
				m.filePicker = NewFilePickerModel()
				// Set filter based on file type
				m.filePicker.SetFilter(fileType)
				return m, tea.Batch(cmd, m.filePicker.Init())
			}
		}
		return m, cmd

	case filePickerView:
		newFilePicker, cmd := m.filePicker.Update(msg)
		m.filePicker = newFilePicker.(FilePickerModel)

		// Check if file was selected
		if m.filePicker.selectedFile != "" {
			m.state = processingView
			comparisonEnabled := m.menu.GetGenerateComparison()
			return m, tea.Batch(cmd, m.progress.StartProcessingWithComparison(m.filePicker.selectedFile, comparisonEnabled))
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
		
		// Check if user wants to go back to menu
		if m.results.backToMenu {
			m.state = menuView
			m.menu = NewMenuModel()
			m.results.backToMenu = false
			return m, m.menu.Init()
		}
		return m, cmd
	}

	return m, nil
}

func (m MainModel) View() string {
	// ASCII Art Header
	asciiArt := styles.ASCIIStyle.Render(`
 ██████╗ ██╗  ██╗    ███╗   ██╗ ██████╗     ███████╗██████╗ ██████╗ 
██╔═══██╗██║  ██║    ████╗  ██║██╔═══██╗    ██╔════╝██╔══██╗██╔══██╗
██║   ██║███████║    ██╔██╗ ██║██║   ██║    ███████╗██║  ██║██████╔╝
██║   ██║██╔══██║    ██║╚██╗██║██║   ██║    ╚════██║██║  ██║██╔══██╗
╚██████╔╝██║  ██║    ██║ ╚████║╚██████╔╝    ███████║██████╔╝██║  ██║
 ╚═════╝ ╚═╝  ╚═╝    ╚═╝  ╚═══╝ ╚═════╝     ╚══════╝╚═════╝ ╚═╝  ╚═╝`)

	subtitle := styles.SubtitleStyle.Render("Parse Student Data Return files with style")

	var content string
	switch m.state {
	case menuView:
		content = m.menu.View()
	case filePickerView:
		content = m.filePicker.View()
	case processingView:
		content = m.progress.View()
	case resultsView:
		content = m.results.View()
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		asciiArt,
		subtitle,
		"",
		content,
	)
}
