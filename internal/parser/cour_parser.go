package parser

import (
	"fmt"
	"strings"
)

// CourseEnrolmentParser handles parsing of COUR files
type CourseEnrolmentParser struct {
	spec              FileSpec
	comparisonService *ComparisonService
	comparisonEnabled bool
}

// NewCourseEnrolmentParser creates a new COUR parser
func NewCourseEnrolmentParser() *CourseEnrolmentParser {
	return &CourseEnrolmentParser{
		spec:              CourseEnrolmentSpec,
		comparisonService: NewComparisonService(),
		comparisonEnabled: false,
	}
}

// EnableComparison enables comparison mode and loads COMP data
func (p *CourseEnrolmentParser) EnableComparison(filePath string) error {
	p.comparisonEnabled = true
	return p.comparisonService.LoadCompData(filePath)
}

// GetComparisonWarnings returns any warnings from comparison loading
func (p *CourseEnrolmentParser) GetComparisonWarnings() []string {
	return p.comparisonService.GetWarnings()
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

	// Add comparison data headers if enabled
	if p.comparisonEnabled {
		// Add empty columns for spacing (2 columns separation)
		headers = append(headers, "", "") // Empty columns
		// Add comparison data column
		headers = append(headers, "Student Course Completion indicator")
	}

	return headers
}

// Parse parses the entire file content and returns records
func (p *CourseEnrolmentParser) Parse(content string) ([]map[string]string, error) {
	// Handle both Windows (\r\n) and Unix (\n) line endings
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")
	lines := strings.Split(strings.TrimSpace(content), "\n")
	var records []map[string]string

	for i, line := range lines {
		lineNum := i + 1

		// Skip empty lines and trim any extra whitespace
		line = strings.TrimSpace(line)
		if line == "" {
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

		// Add comparison data if enabled
		if p.comparisonEnabled {
			completionStatus := p.comparisonService.LookupCompletion(
				record["ID"],
				record["COURSE"],
				record["CRS_SRT"],
			)
			record["COMPLETE"] = completionStatus
		}

		records = append(records, record)
	}

	return records, nil
}

// ParseLine parses a single line and returns field values
func (p *CourseEnrolmentParser) ParseLine(line string, lineNum int) ([]string, error) {
	// Pad line to expected length if it's shorter (common with trailing spaces missing)
	if len(line) < p.spec.LineLength {
		line = line + strings.Repeat(" ", p.spec.LineLength-len(line))
	}
	// Truncate line if it's longer (handle data quality issues)
	if len(line) > p.spec.LineLength {
		line = line[:p.spec.LineLength]
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
		// Trim both leading and trailing spaces
		value = strings.TrimSpace(value)

		// Check required fields
		if field.Required && value == "" {
			return nil, fmt.Errorf("line %d: required field %s is empty", lineNum, field.Name)
		}

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
