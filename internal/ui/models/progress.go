package models

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/unamelo/oh-no-sdr/internal/parser"
	"github.com/unamelo/oh-no-sdr/internal/ui/styles"
)

type ProgressModel struct {
	isComplete    bool
	results       string
	filesToProcess []string
	currentFile   string
	processedFiles int
	totalFiles    int
	error         error
}

func NewProgressModel() ProgressModel {
	return ProgressModel{}
}

func (m ProgressModel) Init() tea.Cmd {
	return nil
}

func (m ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ProcessCompleteMsg:
		m.isComplete = true
		m.results = strings.Join(msg.Results, "\n")
		m.error = msg.Error
		return m, nil
	case ProcessProgressMsg:
		m.processedFiles = msg.ProcessedFiles
		m.currentFile = msg.CurrentFile
		return m, nil
	}
	return m, nil
}

func (m ProgressModel) View() string {
	header := styles.HighlightStyle.Render(">> PROCESSING SDR FILES <<")
	
	if m.error != nil {
		error := styles.ErrorStyle.Render("\nERROR: " + m.error.Error())
		return styles.BoxStyle.Render(header + "\n" + error)
	}
	
	var content string
	if m.totalFiles > 1 {
		content = styles.SubtitleStyle.Render(
			"\nProcessing multiple files...\n" +
			"Current: " + filepath.Base(m.currentFile) + "\n" +
			"Progress: " + fmt.Sprintf("%d/%d files", m.processedFiles, m.totalFiles) + "\n\n" +
			"[████████████████████████████████████████] Processing...")
	} else {
		content = styles.SubtitleStyle.Render(
			"\nAnalyzing file structure...\n" +
			"Current: " + filepath.Base(m.currentFile) + "\n" +
			"Extracting fields...\n" +
			"Generating CSV...\n\n" +
			"[████████████████████████████████████████] Processing...")
	}
	
	return styles.BoxStyle.Render(header + "\n" + content)
}

func (m ProgressModel) StartProcessing(filename string) tea.Cmd {
	return m.StartProcessingWithComparison(filename, false)
}

func (m ProgressModel) StartProcessingWithComparison(filename string, enableComparison bool) tea.Cmd {
	m.filesToProcess = []string{filename}
	m.currentFile = filename
	m.totalFiles = 1
	m.processedFiles = 0
	m.error = nil
	return processFilesWithComparison([]string{filename}, enableComparison)
}

func (m ProgressModel) StartProcessingMultiple(files []string) tea.Cmd {
	return m.StartProcessingMultipleWithComparison(files, false)
}

func (m ProgressModel) StartProcessingMultipleWithComparison(files []string, enableComparison bool) tea.Cmd {
	m.filesToProcess = files
	m.totalFiles = len(files)
	m.processedFiles = 0
	m.error = nil
	if len(files) > 0 {
		m.currentFile = files[0]
	}
	return processFilesWithComparison(files, enableComparison)
}

// ProcessCompleteMsg is sent when processing is complete
type ProcessCompleteMsg struct {
	Results []string
	Error   error
}

// ProcessProgressMsg is sent for progress updates
type ProcessProgressMsg struct {
	ProcessedFiles int
	CurrentFile    string
}

// processFiles processes one or more files
func processFiles(files []string) tea.Cmd {
	return processFilesWithComparison(files, false)
}

// processFilesWithComparison processes one or more files with optional comparison mode
func processFilesWithComparison(files []string, enableComparison bool) tea.Cmd {
	return func() tea.Msg {
		var results []string
		var processingError error
		
		// Get current directory for output
		currentDir, err := os.Getwd()
		if err != nil {
			return ProcessCompleteMsg{Error: fmt.Errorf("failed to get current directory: %w", err)}
		}
		
		for i, file := range files {
			// Send progress update
			if i > 0 {
				// This is a bit of a hack - we can't send multiple messages from one command
				// In a real implementation, you'd want to use a proper progress system
			}
			
			result := parser.ProcessFileWithComparison(file, currentDir, enableComparison)
			if result.Success {
				results = append(results, fmt.Sprintf("✓ %s → %s (%d records)", 
					filepath.Base(result.InputFile), 
					filepath.Base(result.OutputFile), 
					result.RecordCount))
			} else {
				results = append(results, fmt.Sprintf("✗ %s - ERROR: %s", 
					filepath.Base(result.InputFile), 
					result.Error.Error()))
				processingError = result.Error
			}
		}
		
		return ProcessCompleteMsg{
			Results: results,
			Error:   processingError,
		}
	}
}
