package parser

import (
	"fmt"
	"strings"
)

// CREGParser parses CREG (Course Register) files
type CREGParser struct {
	spec FileSpec
}

// NewCREGParser creates a new CREG parser
func NewCREGParser() *CREGParser {
	return &CREGParser{
		spec: GetCREGSpec(),
	}
}

// Parse parses the content and returns records as maps
func (p *CREGParser) Parse(content string) ([]map[string]string, error) {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	var records []map[string]string

	for lineNum, line := range lines {
		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Validate line length
		if err := p.ValidateLine(line); err != nil {
			return nil, fmt.Errorf("line %d: %w", lineNum+1, err)
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
func (p *CREGParser) parseLine(line string) (map[string]string, error) {
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
func (p *CREGParser) GetHeaders() []string {
	headers := make([]string, len(p.spec.Fields))
	for i, field := range p.spec.Fields {
		headers[i] = field.Title
	}
	return headers
}

// GetFileType returns the file type
func (p *CREGParser) GetFileType() string {
	return p.spec.FileType
}

// ValidateLine validates a line's basic structure
func (p *CREGParser) ValidateLine(line string) error {
	if len(line) != p.spec.LineLength {
		return fmt.Errorf("invalid line length: expected %d, got %d", p.spec.LineLength, len(line))
	}
	return nil
}
