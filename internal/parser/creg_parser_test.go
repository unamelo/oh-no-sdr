package parser

import (
	"testing"
)

// Sample CREG data for testing (based on the actual CREG9170.txt file)
const sampleCREGData = `91702102-530            Food Product Development                                                   NZ210222  1101095 14P10.1167    5944X         02N
91702102-510            Industry Placement                                                         NZ210222  1101095  7P10.0583    2974X         02N
91702102-520            Supervise Food Production                                                  NZ210222  1101095 14P10.1167    5944X         02N
91708065-02-209         Prepare, cook and finish meat, poultry and offal                           NZ210122  1101094 12P10.1000    5094X         02N
91708065-01-201         Introduction to the hospitality and catering industry                      NZ210122  1101094  2P10.0167     844X         02N`

func TestCREGParser_Parse(t *testing.T) {
	parser := NewCREGParser()

	records, err := parser.Parse(sampleCREGData)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	// Should have 5 records
	if len(records) != 5 {
		t.Fatalf("Expected 5 records, got %d", len(records))
	}

	// Test first record (Food Product Development)
	firstRecord := records[0]

	// Test some key fields
	expectedFields := map[string]string{
		"INSTIT":           "9170",
		"COURSE":           "2102-530",
		"CTITLE":           "Food Product Development",
		"QUAL":             "NZ2102",
		"CLASS":            "22",
		"NZSCED":           "110109",
		"NZQCFLEVEL":       "5",
		"CREDIT":           "14",
		"CATEGORY":         "P1",
		"FACTOR":           "0.1167",
		"FEE":              "5944",
		"INTERNET":         "X",
		"EXEMPT_INDICATOR": "2", // Fixed: actual value is "2"
	}

	for fieldName, expectedValue := range expectedFields {
		if actualValue, exists := firstRecord[fieldName]; !exists {
			t.Errorf("Field %s missing from record", fieldName)
		} else if actualValue != expectedValue {
			t.Errorf("Field %s: expected '%s', got '%s'", fieldName, expectedValue, actualValue)
		}
	}
}

func TestCREGParser_ParseSingleLine(t *testing.T) {
	parser := NewCREGParser()

	line := "91702102-530            Food Product Development                                                   NZ210222  1101095 14P10.1167    5944X         02N"

	record, err := parser.parseLine(line)
	if err != nil {
		t.Fatalf("parseLine failed: %v", err)
	}

	// Test specific field extractions
	tests := []struct {
		fieldName string
		expected  string
	}{
		{"INSTIT", "9170"},
		{"COURSE", "2102-530"},
		{"CTITLE", "Food Product Development"},
		{"QUAL", "NZ2102"},
		{"CLASS", "22"},
		{"NZSCED", "110109"},
		{"NZQCFLEVEL", "5"},
		{"CREDIT", "14"},
		{"CATEGORY", "P1"},
		{"FACTOR", "0.1167"},
		{"FEE", "5944"},
		{"INTERNET", "X"},
		{"CCCOSTS_FEE", "0"},
		{"EXEMPT_INDICATOR", "2"},
		{"EMB_LIT_NUM", "N"},
	}

	for _, test := range tests {
		if actual, exists := record[test.fieldName]; !exists {
			t.Errorf("Field %s missing from record", test.fieldName)
		} else if actual != test.expected {
			t.Errorf("Field %s: expected '%s', got '%s'", test.fieldName, test.expected, actual)
		}
	}
}

func TestCREGParser_GetHeaders(t *testing.T) {
	parser := NewCREGParser()
	headers := parser.GetHeaders()

	// Should have 17 headers (all fields in spec)
	if len(headers) != 17 {
		t.Errorf("Expected 17 headers, got %d", len(headers))
	}

	// Test some specific headers
	expectedHeaders := []string{
		"Provider Code",
		"Course Code",
		"Course Title",
		"Qualification Code",
		"Course Classification",
		"NZSCED Field of Study",
		"Level on the NZ Qualifications and Credentials Framework",
		"Credit",
		"Funding Category",
		"Course EFTS Factor",
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

func TestCREGParser_GetFileType(t *testing.T) {
	parser := NewCREGParser()
	if fileType := parser.GetFileType(); fileType != "CREG" {
		t.Errorf("Expected file type 'CREG', got '%s'", fileType)
	}
}

func TestCREGParser_EmptyLines(t *testing.T) {
	parser := NewCREGParser()

	// Content with empty lines
	content := `91702102-530            Food Product Development                                                   NZ210222  1101095 14P10.1167    5944X         02N

91702102-510            Industry Placement                                                         NZ210222  1101095  7P10.0583    2974X         02N
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

func TestCREGParser_InvalidData(t *testing.T) {
	parser := NewCREGParser()

	// Line that's too short
	invalidContent := "91702102-530            Food Product"

	_, err := parser.Parse(invalidContent)
	if err == nil {
		t.Error("Expected error for invalid line length")
	}
}

func TestCREGParser_LongCourseTitles(t *testing.T) {
	parser := NewCREGParser()

	// Test with course that has a long title (should be trimmed to 75 chars)
	line := "91708065-02-211         Prepare, cook and finish rice, grain, farinaceous products and egg dishes  NZ210122  1101094  8P10.0667    3384X         02N"

	record, err := parser.parseLine(line)
	if err != nil {
		t.Fatalf("parseLine failed: %v", err)
	}

	expectedTitle := "Prepare, cook and finish rice, grain, farinaceous products and egg dishes"
	if actualTitle := record["CTITLE"]; actualTitle != expectedTitle {
		t.Errorf("Course title: expected '%s', got '%s'", expectedTitle, actualTitle)
	}
}

// Benchmark test for performance
func BenchmarkCREGParser_Parse(b *testing.B) {
	parser := NewCREGParser()

	for i := 0; i < b.N; i++ {
		_, err := parser.Parse(sampleCREGData)
		if err != nil {
			b.Fatalf("Parse failed: %v", err)
		}
	}
}
