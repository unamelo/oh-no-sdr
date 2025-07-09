package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ComparisonService handles lookup operations between COUR and COMP files
type ComparisonService struct {
	compData map[string]string // key: ID+COURSE+CRS_SRT, value: COMPLETE
	loaded   bool
	warnings []string
}

// NewComparisonService creates a new comparison service
func NewComparisonService() *ComparisonService {
	return &ComparisonService{
		compData: make(map[string]string),
		loaded:   false,
		warnings: []string{},
	}
}

// LoadCompData loads COMP file data for lookup operations
func (cs *ComparisonService) LoadCompData(courFilePath string) error {
	// Find COMP file in the same directory as COUR file
	compFilePath := cs.findCompFile(courFilePath)
	if compFilePath == "" {
		warning := "No COMP file found in the same directory - completion data will show as N/A"
		cs.warnings = append(cs.warnings, warning)
		cs.loaded = true // Mark as loaded even if no file found
		return nil
	}

	// Read COMP file
	content, err := os.ReadFile(compFilePath)
	if err != nil {
		warning := fmt.Sprintf("Failed to read COMP file (%s): %v - completion data will show as N/A", compFilePath, err)
		cs.warnings = append(cs.warnings, warning)
		cs.loaded = true
		return nil
	}

	// Parse COMP file
	compParser := NewCOMPParser()
	records, err := compParser.Parse(string(content))
	if err != nil {
		warning := fmt.Sprintf("Failed to parse COMP file (%s): %v - completion data will show as N/A", compFilePath, err)
		cs.warnings = append(cs.warnings, warning)
		cs.loaded = true
		return nil
	}

	// Build lookup map
	for _, record := range records {
		key := cs.buildCompositeKey(record["ID"], record["COURSE"], record["CRS_SRT"])
		cs.compData[key] = record["COMPLETE"]
	}

	cs.loaded = true
	return nil
}

// findCompFile looks for COMP file in the same directory as the COUR file
func (cs *ComparisonService) findCompFile(courFilePath string) string {
	dir := filepath.Dir(courFilePath)
	
	// Get the base name pattern (e.g., if COUR9170.txt, look for COMP9170.txt)
	courFileName := filepath.Base(courFilePath)
	courFileName = strings.ToUpper(courFileName)
	
	// Try to extract the number pattern from COUR file
	var pattern string
	if strings.HasPrefix(courFileName, "COUR") && strings.HasSuffix(courFileName, ".TXT") {
		// Extract number part (e.g., "9170" from "COUR9170.txt")
		pattern = strings.TrimPrefix(courFileName, "COUR")
		pattern = strings.TrimSuffix(pattern, ".TXT")
	}
	
	// Look for COMP file with same pattern
	possibleNames := []string{
		fmt.Sprintf("COMP%s.txt", pattern),
		fmt.Sprintf("comp%s.txt", pattern),
		fmt.Sprintf("Comp%s.txt", pattern),
		"COMP.txt",
		"comp.txt",
		"Comp.txt",
	}
	
	for _, name := range possibleNames {
		fullPath := filepath.Join(dir, name)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath
		}
	}
	
	return ""
}

// buildCompositeKey creates a composite key from ID, COURSE, and CRS_SRT
func (cs *ComparisonService) buildCompositeKey(id, course, crsStart string) string {
	// Clean and normalize the values
	id = strings.TrimSpace(id)
	course = strings.TrimSpace(course)
	crsStart = strings.TrimSpace(crsStart)
	
	// Combine with a separator that's unlikely to appear in the data
	return fmt.Sprintf("%s||%s||%s", id, course, crsStart)
}

// LookupCompletion looks up completion status for a given ID, COURSE, and CRS_SRT
func (cs *ComparisonService) LookupCompletion(id, course, crsStart string) string {
	if !cs.loaded {
		return "N/A"
	}
	
	key := cs.buildCompositeKey(id, course, crsStart)
	if completion, exists := cs.compData[key]; exists {
		return completion
	}
	
	return "N/A"
}

// GetWarnings returns any warnings encountered during loading
func (cs *ComparisonService) GetWarnings() []string {
	return cs.warnings
}

// GetStats returns statistics about the loaded data
func (cs *ComparisonService) GetStats() map[string]interface{} {
	return map[string]interface{}{
		"loaded":           cs.loaded,
		"total_records":    len(cs.compData),
		"warnings_count":   len(cs.warnings),
	}
}
