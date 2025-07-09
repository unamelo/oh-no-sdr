package parser

import (
	"testing"
)

// Sample STUD data for testing (FAKE DATA)
const sampleSTUDData = `1234567890123 M01011990  9999AAAAA01199920999199009NZL1 1Y                      999999999    0    0111      12341234
1234567890456 F15061985  8888BBBBB02200520999198509AUS9 9N                      888888888    0    0222      56785678
1234567890789 M22121992  7777CCCCC03199820999199209GBR0 1N                      777777777    0    0333      90129012`

func TestSTUDParser_Parse(t *testing.T) {
	parser := NewSTUDParser()
	
	records, err := parser.Parse(sampleSTUDData)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	
	// Should have 3 records
	if len(records) != 3 {
		t.Fatalf("Expected 3 records, got %d", len(records))
	}
	
	// Test first record
	firstRecord := records[0]
	
	// Test some key fields
	expectedFields := map[string]string{
		"INSTIT":  "1234",
		"ID":      "567890123",
		"GENDER":  "M",
		"DOB":     "01011990",
		"NAMEID":  "AAAAA",
		"CITIZEN": "NZL",
		"NSN":     "999999999",
	}
	
	for fieldName, expectedValue := range expectedFields {
		if actualValue, exists := firstRecord[fieldName]; !exists {
			t.Errorf("Field %s missing from record", fieldName)
		} else if actualValue != expectedValue {
			t.Errorf("Field %s: expected '%s', got '%s'", fieldName, expectedValue, actualValue)
		}
	}
}

func TestSTUDParser_ParseSingleLine(t *testing.T) {
	parser := NewSTUDParser()
	
	line := "1234567890123 M01011990  9999AAAAA01199920999199009NZL1 1Y                      999999999    0    0111      12341234"
	
	record, err := parser.parseLine(line)
	if err != nil {
		t.Fatalf("parseLine failed: %v", err)
	}
	
	// Test specific field extractions
	tests := []struct {
		fieldName string
		expected  string
	}{
		{"INSTIT", "1234"},
		{"ID", "567890123"},
		{"GENDER", "M"},
		{"DOB", "01011990"},
		{"TOTAL_FEE", "9999"},  // Should be trimmed
		{"NAMEID", "AAAAA"},
		{"CITIZEN", "NZL"},
		{"FEES_FREE_ELIGIBLE", "1"},
		{"DISABILITY", "1"},
		{"NSN", "999999999"},
		{"PERM_POST_CODE", "1234"},
		{"TERM_POST_CODE", "1234"},
	}
	
	for _, test := range tests {
		if actual, exists := record[test.fieldName]; !exists {
			t.Errorf("Field %s missing from record", test.fieldName)
		} else if actual != test.expected {
			t.Errorf("Field %s: expected '%s', got '%s'", test.fieldName, test.expected, actual)
		}
	}
}

func TestSTUDParser_ValidateLine(t *testing.T) {
	parser := NewSTUDParser()
	
	// Valid line (116 characters)
	validLine := "1234567890123 F01011990  9999AAAAA01199920999199009NZL0 1Y                      999999999    0    0111      12341234"
	if err := parser.ValidateLine(validLine); err != nil {
		t.Errorf("Valid line failed validation: %v", err)
	}
	
	// Invalid line (too short)
	shortLine := "1234567890123 F01011990  9999Y"
	if err := parser.ValidateLine(shortLine); err == nil {
		t.Error("Short line should have failed validation")
	}
	
	// Invalid line (too long)
	longLine := validLine + "EXTRA"
	if err := parser.ValidateLine(longLine); err == nil {
		t.Error("Long line should have failed validation")
	}
}

func TestSTUDParser_GetHeaders(t *testing.T) {
	parser := NewSTUDParser()
	headers := parser.GetHeaders()
	
	// Should have 25 headers (all fields in spec)
	if len(headers) != 25 {
		t.Errorf("Expected 25 headers, got %d", len(headers))
	}
	
	// Test some specific headers
	expectedHeaders := []string{
		"Provider Code",
		"Student Identification Code", 
		"Gender",
		"Date of Birth",
		"National Student Number",
	}
	
	for _, expected := range expectedHeaders {
		found := false
		for _, header := range headers {
			if header == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected header '%s' not found", expected)
		}
	}
}

func TestSTUDParser_GetFileType(t *testing.T) {
	parser := NewSTUDParser()
	if fileType := parser.GetFileType(); fileType != "STUD" {
		t.Errorf("Expected file type 'STUD', got '%s'", fileType)
	}
}

func TestSTUDParser_EmptyLines(t *testing.T) {
	parser := NewSTUDParser()
	
	// Content with empty lines
	content := `1234567890123 F01011990  9999AAAAA01199920999199009NZL0 1Y                      999999999    0    0111      12341234

5678567890456 M15061985  8888BBBBB02200520999198509AUS9 9N                      888888888    0    0222      56785678
`
	
	records, err := parser.Parse(content)
	if err != nil {
		t.Fatalf("Parse with empty lines failed: %v", err)
	}
	
	// Should still have 2 records (empty line ignored)
	if len(records) != 2 {
		t.Errorf("Expected 2 records, got %d", len(records))
	}
}

func TestSTUDParser_InvalidData(t *testing.T) {
	parser := NewSTUDParser()
	
	// Line that's too short
	invalidContent := "1234567890123 F01011990  9999Y"
	
	_, err := parser.Parse(invalidContent)
	if err == nil {
		t.Error("Expected error for invalid line length")
	}
}

// Benchmark test for performance
func BenchmarkSTUDParser_Parse(b *testing.B) {
	parser := NewSTUDParser()
	
	for i := 0; i < b.N; i++ {
		_, err := parser.Parse(sampleSTUDData)
		if err != nil {
			b.Fatalf("Parse failed: %v", err)
		}
	}
}
