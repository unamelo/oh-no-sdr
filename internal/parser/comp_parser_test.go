package parser

import (
	"strings"
	"testing"
)

func TestCOMPParser_Parse(t *testing.T) {
	parser := NewCOMPParser()
	
	// Test with sample data based on COMP9170.txt
	content := `9170917000047 2102-530            028092023 12033171
9170917000047 2102-510            028092023 12033171
9170917000047 2102-520            028092023 12033171`

	records, err := parser.Parse(content)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(records) != 3 {
		t.Errorf("Expected 3 records, got %d", len(records))
	}

	// Test first record
	if records[0]["ID"] != "9170917000047" {
		t.Errorf("Expected ID '9170917000047', got '%s'", records[0]["ID"])
	}

	if records[0]["COURSE"] != "2102-530" {
		t.Errorf("Expected COURSE '2102-530', got '%s'", records[0]["COURSE"])
	}

	if records[0]["CRS_SRT"] != "028092023" {
		t.Errorf("Expected CRS_SRT '028092023', got '%s'", records[0]["CRS_SRT"])
	}

	if records[0]["CRS_END"] != "12033171" {
		t.Errorf("Expected CRS_END '12033171', got '%s'", records[0]["CRS_END"])
	}
}

func TestCOMPParser_GetHeaders(t *testing.T) {
	parser := NewCOMPParser()
	headers := parser.GetHeaders()

	expected := []string{
		"Student Identification Code",
		"Course Code", 
		"Course Start Date",
		"Course End Date",
	}

	if len(headers) != len(expected) {
		t.Errorf("Expected %d headers, got %d", len(expected), len(headers))
	}

	for i, header := range headers {
		if header != expected[i] {
			t.Errorf("Header %d: expected '%s', got '%s'", i, expected[i], header)
		}
	}
}

func TestCOMPParser_GetFileType(t *testing.T) {
	parser := NewCOMPParser()
	fileType := parser.GetFileType()

	if fileType != "COMP" {
		t.Errorf("Expected file type 'COMP', got '%s'", fileType)
	}
}

func TestCOMPParser_ValidateLine(t *testing.T) {
	parser := NewCOMPParser()

	// Test valid line length (52 characters)
	validLine := "9170917000047 2102-530            028092023 12033171"
	if len(validLine) != 52 {
		t.Errorf("Test line should be 52 characters, got %d", len(validLine))
	}

	err := parser.ValidateLine(validLine)
	if err != nil {
		t.Errorf("Valid line should pass validation: %v", err)
	}

	// Test invalid line length
	invalidLine := "short"
	err = parser.ValidateLine(invalidLine)
	if err == nil {
		t.Error("Invalid line should fail validation")
	}
}

func TestCOMPParser_ParseLine(t *testing.T) {
	parser := NewCOMPParser()

	// Test normal line
	line := "9170917000047 2102-530            028092023 12033171"
	record, err := parser.parseLine(line)
	if err != nil {
		t.Fatalf("parseLine failed: %v", err)
	}

	// Verify all fields are extracted correctly
	if record["ID"] != "9170917000047" {
		t.Errorf("ID: expected '9170917000047', got '%s'", record["ID"])
	}

	if record["COURSE"] != "2102-530" {
		t.Errorf("COURSE: expected '2102-530', got '%s'", record["COURSE"])
	}

	if record["CRS_SRT"] != "028092023" {
		t.Errorf("CRS_SRT: expected '028092023', got '%s'", record["CRS_SRT"])
	}

	if record["CRS_END"] != "12033171" {
		t.Errorf("CRS_END: expected '12033171', got '%s'", record["CRS_END"])
	}
}

func TestCOMPParser_ParseWithSpaces(t *testing.T) {
	parser := NewCOMPParser()

	// Test line with extra spaces (should be trimmed)
	line := "9170917000047 2102-530            028092023 12033171"
	record, err := parser.parseLine(line)
	if err != nil {
		t.Fatalf("parseLine failed: %v", err)
	}

	// Fields should be trimmed
	if record["COURSE"] != "2102-530" {
		t.Errorf("COURSE should be trimmed: got '%s'", record["COURSE"])
	}
}

func TestCOMPParser_EmptyContent(t *testing.T) {
	parser := NewCOMPParser()
	
	records, err := parser.Parse("")
	if err != nil {
		t.Fatalf("Parse empty content failed: %v", err)
	}

	if len(records) != 0 {
		t.Errorf("Expected 0 records for empty content, got %d", len(records))
	}
}

func TestCOMPParser_WithEmptyLines(t *testing.T) {
	parser := NewCOMPParser()
	
	content := `9170917000047 2102-530            028092023 12033171

9170917000047 2102-510            028092023 12033171
   
9170917000047 2102-520            028092023 12033171`

	records, err := parser.Parse(content)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(records) != 3 {
		t.Errorf("Expected 3 records (empty lines should be skipped), got %d", len(records))
	}
}

func TestCOMPParser_InvalidLineLength(t *testing.T) {
	parser := NewCOMPParser()
	
	// Line too short
	content := "9170917000047 2102-530"
	
	_, err := parser.Parse(content)
	if err == nil {
		t.Error("Expected error for invalid line length")
	}

	if !strings.Contains(err.Error(), "invalid line length") {
		t.Errorf("Expected line length error, got: %v", err)
	}
}

func TestCOMPParser_FieldBounds(t *testing.T) {
	parser := NewCOMPParser()
	
	// Test with line exactly at boundary
	line := strings.Repeat("X", 52)
	_, err := parser.parseLine(line)
	if err != nil {
		t.Errorf("Valid length line should parse: %v", err)
	}
}

func TestCOMPParser_RequiredFields(t *testing.T) {
	parser := NewCOMPParser()
	
	// Test with missing required field (empty ID)
	line := "             2102-530            028092023 12033171"
	_, err := parser.parseLine(line)
	if err == nil {
		t.Error("Expected error for empty required field")
	}

	if !strings.Contains(err.Error(), "required field") {
		t.Errorf("Expected required field error, got: %v", err)
	}
}

func TestCOMPParser_MultipleRecords(t *testing.T) {
	parser := NewCOMPParser()
	
	content := `9170917000047 2102-530            028092023 12033171
9170917000440 2102-530            027022024 16264822
9170917000116 2102-530            028092023 11450702`

	records, err := parser.Parse(content)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(records) != 3 {
		t.Errorf("Expected 3 records, got %d", len(records))
	}

	// Verify different student IDs
	expectedIDs := []string{"9170917000047", "9170917000440", "9170917000116"}
	for i, expectedID := range expectedIDs {
		if records[i]["ID"] != expectedID {
			t.Errorf("Record %d: expected ID '%s', got '%s'", i, expectedID, records[i]["ID"])
		}
	}
}

func TestCOMPParser_RealWorldData(t *testing.T) {
	parser := NewCOMPParser()
	
	// Test with actual data from COMP9170.txt
	content := `9170917000047 2102-530            028092023 12033171
9170917000047 2102-510            028092023 12033171
9170917000047 2102-520            028092023 12033171
9170917000047 2102-240            028092023 12033171
9170917000047 2102-516            028092023 12033171`

	records, err := parser.Parse(content)
	if err != nil {
		t.Fatalf("Parse real data failed: %v", err)
	}

	if len(records) != 5 {
		t.Errorf("Expected 5 records, got %d", len(records))
	}

	// Verify all records have the same student ID
	for i, record := range records {
		if record["ID"] != "9170917000047" {
			t.Errorf("Record %d: expected ID '9170917000047', got '%s'", i, record["ID"])
		}
	}

	// Verify different course codes
	expectedCourses := []string{"2102-530", "2102-510", "2102-520", "2102-240", "2102-516"}
	for i, expectedCourse := range expectedCourses {
		if records[i]["COURSE"] != expectedCourse {
			t.Errorf("Record %d: expected COURSE '%s', got '%s'", i, expectedCourse, records[i]["COURSE"])
		}
	}
}

func TestCOMPParser_FieldExtraction(t *testing.T) {
	parser := NewCOMPParser()
	
	// Test specific field positions
	line := "9170917000047 2102-530            028092023 12033171"
	//        ^            ^                   ^         ^
	//        1            15                  36        45
	//        |<-- 13 -->| |<----- 20 ----->| |<- 8 ->| |<-8->|
	
	record, err := parser.parseLine(line)
	if err != nil {
		t.Fatalf("parseLine failed: %v", err)
	}

	// Verify exact field extraction
	if record["ID"] != "9170917000047" {
		t.Errorf("ID extraction failed: got '%s'", record["ID"])
	}

	if record["COURSE"] != "2102-530" {
		t.Errorf("COURSE extraction failed: got '%s'", record["COURSE"])
	}

	if record["CRS_SRT"] != "028092023" {
		t.Errorf("CRS_SRT extraction failed: got '%s'", record["CRS_SRT"])
	}

	if record["CRS_END"] != "12033171" {
		t.Errorf("CRS_END extraction failed: got '%s'", record["CRS_END"])
	}
}

func TestCOMPParser_Specification(t *testing.T) {
	spec := GetCOMPSpec()
	
	// Test specification details
	if spec.FileType != "COMP" {
		t.Errorf("Expected FileType 'COMP', got '%s'", spec.FileType)
	}

	if spec.LineLength != 52 {
		t.Errorf("Expected LineLength 52, got %d", spec.LineLength)
	}

	if len(spec.Fields) != 4 {
		t.Errorf("Expected 4 fields, got %d", len(spec.Fields))
	}

	// Test field specifications
	expectedFields := []struct {
		name   string
		start  int
		length int
	}{
		{"ID", 1, 13},
		{"COURSE", 15, 20},
		{"CRS_SRT", 36, 8},
		{"CRS_END", 45, 8},
	}

	for i, expected := range expectedFields {
		field := spec.Fields[i]
		if field.Name != expected.name {
			t.Errorf("Field %d: expected name '%s', got '%s'", i, expected.name, field.Name)
		}
		if field.Start != expected.start {
			t.Errorf("Field %d: expected start %d, got %d", i, expected.start, field.Start)
		}
		if field.Length != expected.length {
			t.Errorf("Field %d: expected length %d, got %d", i, expected.length, field.Length)
		}
	}
}
