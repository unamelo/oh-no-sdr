package models

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/unamelo/oh-no-sdr/internal/ui/styles"
)

type MenuModel struct {
	choices       []string
	cursor        int
	selectedIndex int
	width         int
	height        int
}

func NewMenuModel() MenuModel {
	return MenuModel{
		choices: []string{
			"Parse All Files",
			"Parse STUD File",
			"Parse COUR File", 
			"Parse CREG File",
			"Parse COMP File",
			"Parse QUAL File",
		},
		selectedIndex: -1,
	}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.selectedIndex = m.cursor
			return m, nil
		}
	}
	return m, nil
}

func (m MenuModel) View() string {
	var s strings.Builder
	
	// Header
	header := styles.HighlightStyle.Render(">> SELECT PARSING OPTION <<")
	s.WriteString(header + "\n\n")

	// Menu options
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
			cursor = lipgloss.NewStyle().Foreground(styles.Primary).Render(cursor)
		}
		
		// Style the selected option
		if m.cursor == i {
			choice = styles.HighlightStyle.Render(choice)
		}
		
		s.WriteString(fmt.Sprintf("%s %s\n", cursor, choice))
	}
	
	s.WriteString("\n")
	
	// Instructions
	instructions := styles.SubtitleStyle.Render("CONTROLS: [↑/↓] Navigate • [Enter] Select • [q] Quit")
	s.WriteString(instructions)

	return styles.BoxStyle.Render(s.String())
}

// GetSelectedOption returns the selected option type and any auto-detected files
func (m MenuModel) GetSelectedOption() (string, []string, error) {
	if m.selectedIndex < 0 {
		return "", nil, nil
	}
	
	currentDir, err := os.Getwd()
	if err != nil {
		return "", nil, fmt.Errorf("failed to get current directory: %w", err)
	}
	
	switch m.selectedIndex {
	case 0: // Parse All Files
		files, err := findAllSDRFiles(currentDir)
		return "all", files, err
	case 1: // Parse STUD File
		files, err := findFilesByType(currentDir, "STUD")
		return "stud", files, err
	case 2: // Parse COUR File
		files, err := findFilesByType(currentDir, "COUR")
		return "cour", files, err
	case 3: // Parse CREG File
		files, err := findFilesByType(currentDir, "CREG")
		return "creg", files, err
	case 4: // Parse COMP File
		files, err := findFilesByType(currentDir, "COMP")
		return "comp", files, err
	case 5: // Parse QUAL File
		files, err := findFilesByType(currentDir, "QUAL")
		return "qual", files, err
	}
	
	return "", nil, nil
}

// findAllSDRFiles finds all SDR files in the directory
func findAllSDRFiles(dir string) ([]string, error) {
	var files []string
	
	fileTypes := []string{"STUD", "COUR", "CREG", "COMP", "QUAL"}
	
	for _, fileType := range fileTypes {
		typeFiles, err := findFilesByType(dir, fileType)
		if err != nil {
			return nil, err
		}
		files = append(files, typeFiles...)
	}
	
	return files, nil
}

// findFilesByType finds files of a specific type (e.g., STUD, COUR, etc.)
func findFilesByType(dir, fileType string) ([]string, error) {
	var files []string
	
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}
	
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		
		name := entry.Name()
		// Check if filename contains the file type and has .txt extension
		if strings.Contains(strings.ToUpper(name), strings.ToUpper(fileType)) && 
		   strings.HasSuffix(strings.ToLower(name), ".txt") {
			files = append(files, filepath.Join(dir, name))
		}
	}
	
	return files, nil
}
