package parser

import (
	"fmt"
	"strings"
)

// CourseEnrolmentParser handles parsing of COUR files
type CourseEnrolmentParser struct {
	spec FileSpec
}

// NewCourseEnrolmentParser creates a new COUR parser
func NewCourseEnrolmentParser() *CourseEnrolmentParser {
	return &CourseEnrolmentParser{
		spec: CourseEnrolmentSpec,
	}
}

// GetFileType returns the file type identifier
func (p *CourseEnrolmentParser) GetFileType() string {
	return p.spec.FileType
}

// GetDescription returns the file description
func (p *CourseEnrolmentParser) GetDescription() string {
	return p.spec.Description
}

// GetExpectedLineLength returns the expected line length
func (p *CourseEnrolmentParser) GetExpectedLineLength() int {
	return p.spec.LineLength
}

// GetHeaders returns the CSV headers
func (p *CourseEnrolmentParser) GetHeaders() []string {
	headers := make([]string, len(p.spec.Fields))
	for i, field := range p.spec.Fields {
		headers[i] = field.Title
	}
	return headers
}

// Parse parses the entire file content and returns records
func (p *CourseEnrolmentParser) Parse(content string) ([]map[string]string, error) {
	lines := strings.Split(content, "\n")
	var records []map[string]string
	
	for i, line := range lines {
		lineNum := i + 1
		
		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}
		
		// Parse line
		values, err := p.ParseLine(line, lineNum)
		if err != nil {
			return nil, fmt.Errorf("error parsing line %d: %w", lineNum, err)
		}
		
		// Create record map
		record := make(map[string]string)
		for j, field := range p.spec.Fields {
			record[field.Name] = values[j]
		}
		
		records = append(records, record)
	}
	
	return records, nil
}

// ValidateLine validates a single line format (updated signature for Parser interface)
func (p *CourseEnrolmentParser) ValidateLine(line string) error {
	return p.validateLineWithNumber(line, 1)
}

// validateLineWithNumber validates a single line format with line number
func (p *CourseEnrolmentParser) validateLineWithNumber(line string, lineNum int) error {
	if len(line) < p.spec.LineLength {
		return fmt.Errorf("line %d: insufficient length %d, expected %d", lineNum, len(line), p.spec.LineLength)
	}
	
	// Check required fields
	for _, field := range p.spec.Fields {
		if field.Required {
			start := field.Start - 1
			end := start + field.Length
			if start >= len(line) || end > len(line) {
				return fmt.Errorf("line %d: cannot extract required field %s (positions %d-%d)", lineNum, field.Name, field.Start, field.Start+field.Length-1)
			}
			fieldValue := strings.TrimSpace(line[start:end])
			if fieldValue == "" {
				return fmt.Errorf("line %d: required field %s is empty", lineNum, field.Name)
			}
		}
	}
	
	return nil
}

// ParseLine parses a single line and returns field values
func (p *CourseEnrolmentParser) ParseLine(line string, lineNum int) ([]string, error) {
	if err := p.validateLineWithNumber(line, lineNum); err != nil {
		return nil, err
	}
	
	values := make([]string, len(p.spec.Fields))
	for i, field := range p.spec.Fields {
		start := field.Start - 1
		end := start + field.Length
		
		if start >= len(line) || end > len(line) {
			values[i] = ""
			continue
		}
		
		value := line[start:end]
		// Only trim trailing spaces for string fields to preserve data integrity
		value = strings.TrimRight(value, " ")
		values[i] = value
	}
	
	return values, nil
}

// IsMatchingFileType checks if the file matches this parser type
func (p *CourseEnrolmentParser) IsMatchingFileType(filename string, firstLine string) bool {
	// Check filename pattern
	upperFilename := strings.ToUpper(filename)
	if strings.Contains(upperFilename, "COUR") {
		return true
	}
	
	// Check line length
	if len(firstLine) == p.spec.LineLength {
		// Additional validation: check if it looks like a COUR record
		if len(firstLine) >= 20 {
			// Check if first 4 characters look like institution code (numeric)
			instit := strings.TrimSpace(firstLine[:4])
			if len(instit) == 4 && isNumeric(instit) {
				// Check if positions 5-14 look like student ID
				studentID := strings.TrimSpace(firstLine[4:14])
				if len(studentID) > 0 {
					// Check if positions 15-20 look like qualification code
					qualCode := strings.TrimSpace(firstLine[14:20])
					if len(qualCode) > 0 {
						return true
					}
				}
			}
		}
	}
	
	return false
}

// isNumeric checks if a string contains only digits
func isNumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
