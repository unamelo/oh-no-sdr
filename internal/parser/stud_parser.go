package parser

import (
	"fmt"
	"strings"
)

// STUDParser parses STUD (Student) files
type STUDParser struct {
	spec FileSpec
}

// NewSTUDParser creates a new STUD parser
func NewSTUDParser() *STUDParser {
	return &STUDParser{
		spec: GetSTUDSpec(),
	}
}

// Parse parses the content and returns records as maps
func (p *STUDParser) Parse(content string) ([]map[string]string, error) {
	// Handle both Windows (\r\n) and Unix (\n) line endings
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")
	lines := strings.Split(strings.TrimSpace(content), "\n")
	var records []map[string]string

	for lineNum, line := range lines {
		// Skip empty lines and trim any extra whitespace
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Parse the line
		record, err := p.parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", lineNum+1, err)
		}

		records = append(records, record)
	}

	return records, nil
}

// parseLine extracts fields from a single line
func (p *STUDParser) parseLine(line string) (map[string]string, error) {
	// Pad line to expected length if it's shorter (common with trailing spaces missing)
	if len(line) < p.spec.LineLength {
		line = line + strings.Repeat(" ", p.spec.LineLength-len(line))
	}
	// Truncate line if it's longer (handle data quality issues)
	if len(line) > p.spec.LineLength {
		line = line[:p.spec.LineLength]
	}

	record := make(map[string]string)

	for _, field := range p.spec.Fields {
		// Convert 1-based position to 0-based for Go
		start := field.Start - 1
		end := start + field.Length

		// Bounds checking
		if start < 0 || end > len(line) {
			return nil, fmt.Errorf("field %s: position out of bounds (start: %d, end: %d, line length: %d)",
				field.Name, start, end, len(line))
		}

		// Extract field value
		value := line[start:end]

		// Trim both leading and trailing spaces
		value = strings.TrimSpace(value)

		// Check required fields
		if field.Required && value == "" {
			return nil, fmt.Errorf("required field %s is empty", field.Name)
		}

		record[field.Name] = value
	}

	return record, nil
}

// GetHeaders returns the field titles for CSV headers
func (p *STUDParser) GetHeaders() []string {
	headers := make([]string, len(p.spec.Fields))
	for i, field := range p.spec.Fields {
		headers[i] = field.Title
	}
	return headers
}

// GetFileType returns the file type
func (p *STUDParser) GetFileType() string {
	return p.spec.FileType
}
