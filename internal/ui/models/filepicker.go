package models

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/unamelo/oh-no-sdr/internal/ui/styles"
)

type FilePickerModel struct {
	filepicker   filepicker.Model
	selectedFile string
	err          error
	width        int
	height       int
}

func NewFilePickerModel() FilePickerModel {
	fp := filepicker.New()
	fp.AllowedTypes = []string{".txt"}
	fp.CurrentDirectory, _ = os.Getwd()

	return FilePickerModel{
		filepicker: fp,
	}
}

func (m FilePickerModel) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m FilePickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)

	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		m.selectedFile = path
	}

	if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
		m.err = fmt.Errorf("selected file is not a .txt file: %s", path)
	}

	return m, cmd
}

func (m FilePickerModel) View() string {
	var s strings.Builder
	
	// Header
	header := styles.HighlightStyle.Render(">> SELECT AN SDR FILE TO PARSE <<")
	s.WriteString(header + "\n\n")

	if m.err != nil {
		s.WriteString(styles.ErrorStyle.Render(fmt.Sprintf("ERROR: %v\n\n", m.err)))
	}

	// File picker with more space
	s.WriteString(m.filepicker.View())
	s.WriteString("\n\n")
	
	// Instructions
	instructions := styles.SubtitleStyle.Render("CONTROLS: [↑/↓] Navigate • [Enter] Select • [q] Quit")
	s.WriteString(instructions)

	return styles.BoxStyle.Render(s.String())
}
