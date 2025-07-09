package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ProcessorResult contains the results of file processing
type ProcessorResult struct {
	InputFile   string
	OutputFile  string
	RecordCount int
	FileType    string
	Success     bool
	Error       error
}

// CSVWriter handles writing parsed data to CSV files
type CSVWriter struct{}

// NewCSVWriter creates a new CSV writer
func NewCSVWriter() *CSVWriter {
	return &CSVWriter{}
}

// WriteCSV writes parsed records to a CSV file
func (w *CSVWriter) WriteCSV(records []map[string]string, headers []string, outputPath string, parser Parser) error {
	// Create output file
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	// Create CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write headers: %w", err)
	}

	// Handle different parser types
	if studParser, ok := parser.(*STUDParser); ok {
		return w.writeSTUDRecords(writer, records, studParser)
	} else if courParser, ok := parser.(*CourseEnrolmentParser); ok {
		return w.writeCourseEnrolmentRecords(writer, records, courParser)
	}

	return fmt.Errorf("unsupported parser type")
}

// writeSTUDRecords writes STUD records using the proper field mapping
func (w *CSVWriter) writeSTUDRecords(writer *csv.Writer, records []map[string]string, parser *STUDParser) error {
	for i, record := range records {
		row := make([]string, len(parser.spec.Fields))
		for j, field := range parser.spec.Fields {
			row[j] = record[field.Name]
		}
		
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write record %d: %w", i+1, err)
		}
	}
	return nil
}

// writeCourseEnrolmentRecords writes COUR records using the proper field mapping
func (w *CSVWriter) writeCourseEnrolmentRecords(writer *csv.Writer, records []map[string]string, parser *CourseEnrolmentParser) error {
	for i, record := range records {
		row := make([]string, len(parser.spec.Fields))
		for j, field := range parser.spec.Fields {
			row[j] = record[field.Name]
		}
		
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write record %d: %w", i+1, err)
		}
	}
	return nil
}

// ProcessFile processes a single SDR file
func ProcessFile(inputPath string, outputDir string) ProcessorResult {
	result := ProcessorResult{
		InputFile: inputPath,
		Success:   false,
	}

	// Read input file
	content, err := os.ReadFile(inputPath)
	if err != nil {
		result.Error = fmt.Errorf("failed to read input file: %w", err)
		return result
	}

	// Determine file type from filename
	filename := filepath.Base(inputPath)
	fileType := DetectFileType(filename)
	if fileType == "" {
		result.Error = fmt.Errorf("unable to determine file type from filename: %s", filename)
		return result
	}

	result.FileType = fileType

	// Get appropriate parser
	parser, err := GetParser(fileType)
	if err != nil {
		result.Error = fmt.Errorf("failed to get parser for file type %s: %w", fileType, err)
		return result
	}

	// Parse content
	records, err := parser.Parse(string(content))
	if err != nil {
		result.Error = fmt.Errorf("failed to parse file: %w", err)
		return result
	}

	result.RecordCount = len(records)

	// Generate output filename
	baseFilename := strings.TrimSuffix(filename, ".txt")
	outputFilename := fmt.Sprintf("%s_parsed.csv", baseFilename)
	outputPath := filepath.Join(outputDir, outputFilename)
	result.OutputFile = outputPath

	// Write CSV
	csvWriter := NewCSVWriter()
	headers := parser.GetHeaders()
	
	if err := csvWriter.WriteCSV(records, headers, outputPath, parser); err != nil {
		result.Error = fmt.Errorf("failed to write CSV: %w", err)
		return result
	}

	result.Success = true
	return result
}

// DetectFileType attempts to determine file type from filename
func DetectFileType(filename string) string {
	upper := strings.ToUpper(filename)
	
	if strings.Contains(upper, "STUD") {
		return "STUD"
	} else if strings.Contains(upper, "COUR") {
		return "COUR"
	} else if strings.Contains(upper, "CREG") {
		return "CREG"
	} else if strings.Contains(upper, "COMP") {
		return "COMP"
	} else if strings.Contains(upper, "QUAL") {
		return "QUAL"
	}
	
	return ""
}

// GetParser returns the appropriate parser for a given file type
func GetParser(fileType string) (Parser, error) {
	switch strings.ToUpper(fileType) {
	case "STUD":
		return NewSTUDParser(), nil
	case "COUR":
		return NewCourseEnrolmentParser(), nil
	default:
		return nil, fmt.Errorf("unsupported file type: %s", fileType)
	}
}
