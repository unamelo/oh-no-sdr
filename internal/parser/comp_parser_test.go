package parser

import (
	"strings"
	"testing"
)

func TestCOMPParser_Parse(t *testing.T) {
	parser := NewCOMPParser()
	
	// Test with sample data based on COMP9170.txt (65 characters per line)
	content := `9170917000047 2102-530            028092023 12033171106062024    
9170917000047 2102-510            028092023 12033171106062024    
9170917000047 2102-520            028092023 12033171106062024    `

	records, err := parser.Parse(content)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(records) != 3 {
		t.Errorf("Expected 3 records, got %d", len(records))
	}

	// Test first record
	if records[0]["INSTIT"] != "9170" {
		t.Errorf("Expected INSTIT '9170', got '%s'", records[0]["INSTIT"])
	}

	if records[0]["ID"] != "917000047" {
		t.Errorf("Expected ID '917000047', got '%s'", records[0]["ID"])
	}

	if records[0]["COURSE"] != "2102-530" {
		t.Errorf("Expected COURSE '2102-530', got '%s'", records[0]["COURSE"])
	}

	if records[0]["CRS_SRT"] != "28092023" {
		t.Errorf("Expected CRS_SRT '28092023', got '%s'", records[0]["CRS_SRT"])
	}

	if records[0]["NSN"] != "120331711" {
		t.Errorf("Expected NSN '120331711', got '%s'", records[0]["NSN"])
	}

	if records[0]["CRS_END"] != "06062024" {
		t.Errorf("Expected CRS_END '06062024', got '%s'", records[0]["CRS_END"])
	}

	if records[0]["PBRF_CRS_COMP_YR"] != "" {
		t.Errorf("Expected PBRF_CRS_COMP_YR '', got '%s'", records[0]["PBRF_CRS_COMP_YR"])
	}
}

func TestCOMPParser_GetHeaders(t *testing.T) {
	parser := NewCOMPParser()
	headers := parser.GetHeaders()

	expected := []string{
		"Provider Code",
		"Student Identification Code",
		"Course Code", 
		"Student Course Completion indicator",
		"Course Start Date",
		"National Student Number",
		"Course End Date",
		"PBRF Course Completion Year",
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

	// Test valid line length (65 characters)
	validLine := "9170917000047 2102-530            028092023 12033171106062024    "
	if len(validLine) != 65 {
		t.Errorf("Test line should be 65 characters, got %d", len(validLine))
	}

	err := parser.ValidateLine(validLine)
	if err != nil {
		t.Errorf("Valid line should pass validation: %v", err)
	}

	// Test short line (should be allowed)
	shortLine := "9170917000047 2102-530            028092023 12033171106062024"
	err = parser.ValidateLine(shortLine)
	if err != nil {
		t.Errorf("Short line should pass validation (will be padded): %v", err)
	}
	
	// Test line too long
	longLine := strings.Repeat("X", 70)
	err = parser.ValidateLine(longLine)
	if err == nil {
		t.Error("Line too long should fail validation")
	}
}

func TestCOMPParser_ParseLine(t *testing.T) {
	parser := NewCOMPParser()

	// Test normal line (65 characters)
	line := "9170917000047 2102-530            028092023 12033171106062024    "
	record, err := parser.parseLine(line)
	if err != nil {
		t.Fatalf("parseLine failed: %v", err)
	}

	// Verify all fields are extracted correctly
	if record["INSTIT"] != "9170" {
		t.Errorf("INSTIT: expected '9170', got '%s'", record["INSTIT"])
	}

	if record["ID"] != "917000047" {
		t.Errorf("ID: expected '917000047', got '%s'", record["ID"])
	}

	if record["COURSE"] != "2102-530" {
		t.Errorf("COURSE: expected '2102-530', got '%s'", record["COURSE"])
	}

	if record["CRS_SRT"] != "28092023" {
		t.Errorf("CRS_SRT: expected '28092023', got '%s'", record["CRS_SRT"])
	}

	if record["NSN"] != "120331711" {
		t.Errorf("NSN: expected '120331711', got '%s'", record["NSN"])
	}

	if record["CRS_END"] != "06062024" {
		t.Errorf("CRS_END: expected '06062024', got '%s'", record["CRS_END"])
	}

	if record["PBRF_CRS_COMP_YR"] != "" {
		t.Errorf("PBRF_CRS_COMP_YR: expected '', got '%s'", record["PBRF_CRS_COMP_YR"])
	}
}

func TestCOMPParser_ParseWithSpaces(t *testing.T) {
	parser := NewCOMPParser()

	// Test line with extra spaces (should be trimmed)
	line := "9170917000047 2102-530            028092023 12033171106062024    "
	record, err := parser.parseLine(line)
	if err != nil {
		t.Fatalf("parseLine failed: %v", err)
	}

	// Fields should be trimmed
	if record["COURSE"] != "2102-530" {
		t.Errorf("COURSE should be trimmed: got '%s'", record["COURSE"])
	}

	// Test that trailing spaces are properly handled
	if record["PBRF_CRS_COMP_YR"] != "" {
		t.Errorf("PBRF_CRS_COMP_YR should be trimmed: got '%s'", record["PBRF_CRS_COMP_YR"])
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
	
	content := `9170917000047 2102-530            028092023 12033171106062024    

9170917000047 2102-510            028092023 12033171106062024    
   
9170917000047 2102-520            028092023 12033171106062024    `

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
	
	// Line too long should fail
	content := strings.Repeat("X", 70) // 70 characters, more than 65
	
	_, err := parser.Parse(content)
	if err == nil {
		t.Error("Expected error for line too long")
	}

	if !strings.Contains(err.Error(), "line too long") {
		t.Errorf("Expected line too long error, got: %v", err)
	}
	
	// Line too short should work (will be padded)
	// Use a line that has all required fields but missing trailing spaces
	shortContent := "9170917000047 2102-530            028092023 12033171106062024"
	_, err = parser.Parse(shortContent)
	if err != nil {
		t.Errorf("Short line should be handled by padding: %v", err)
	}
}

func TestCOMPParser_FieldBounds(t *testing.T) {
	parser := NewCOMPParser()
	
	// Test with line exactly at boundary
	line := strings.Repeat("X", 65)
	_, err := parser.parseLine(line)
	if err != nil {
		t.Errorf("Valid length line should parse: %v", err)
	}
}

func TestCOMPParser_RequiredFields(t *testing.T) {
	parser := NewCOMPParser()
	
	// Test with missing required field (empty INSTIT)
	line := "    917000047 2102-530            028092023 12033171106062024    "
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
	
	content := `9170917000047 2102-530            028092023 12033171106062024    
9170917000440 2102-530            027022024 16264822205112024    
9170917000116 2102-530            028092023 11450702406062024    `

	records, err := parser.Parse(content)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(records) != 3 {
		t.Errorf("Expected 3 records, got %d", len(records))
	}

	// Verify different student IDs
	expectedIDs := []string{"917000047", "917000440", "917000116"}
	for i, expectedID := range expectedIDs {
		if records[i]["ID"] != expectedID {
			t.Errorf("Record %d: expected ID '%s', got '%s'", i, expectedID, records[i]["ID"])
		}
	}
}

func TestCOMPParser_RealWorldData(t *testing.T) {
	parser := NewCOMPParser()
	
	// Test with actual data from COMP9170.txt
	content := `9170917000047 2102-530            028092023 12033171106062024    
9170917000047 2102-510            028092023 12033171106062024    
9170917000047 2102-520            028092023 12033171106062024    
9170917000047 2102-240            028092023 12033171106062024    
9170917000047 2102-516            028092023 12033171106062024    `

	records, err := parser.Parse(content)
	if err != nil {
		t.Fatalf("Parse real data failed: %v", err)
	}

	if len(records) != 5 {
		t.Errorf("Expected 5 records, got %d", len(records))
	}

	// Verify all records have the same student ID
	for i, record := range records {
		if record["ID"] != "917000047" {
			t.Errorf("Record %d: expected ID '917000047', got '%s'", i, record["ID"])
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
	line := "9170917000047 2102-530            028092023 12033171106062024    "
	//        ^   ^        ^                 ^^         ^        ^    ^
	//        1   5        15                3536       44       54   62
	//        |<4>|<--10-->|<------ 20 ----->||<--8-->|<--10-->|<-8->|<4>|
	
	record, err := parser.parseLine(line)
	if err != nil {
		t.Fatalf("parseLine failed: %v", err)
	}

	// Verify exact field extraction
	if record["INSTIT"] != "9170" {
		t.Errorf("INSTIT extraction failed: got '%s'", record["INSTIT"])
	}

	if record["ID"] != "917000047" {
		t.Errorf("ID extraction failed: got '%s'", record["ID"])
	}

	if record["COURSE"] != "2102-530" {
		t.Errorf("COURSE extraction failed: got '%s'", record["COURSE"])
	}

	if record["CRS_SRT"] != "28092023" {
		t.Errorf("CRS_SRT extraction failed: got '%s'", record["CRS_SRT"])
	}

	if record["NSN"] != "120331711" {
		t.Errorf("NSN extraction failed: got '%s'", record["NSN"])
	}

	if record["CRS_END"] != "06062024" {
		t.Errorf("CRS_END extraction failed: got '%s'", record["CRS_END"])
	}

	if record["PBRF_CRS_COMP_YR"] != "" {
		t.Errorf("PBRF_CRS_COMP_YR extraction failed: got '%s'", record["PBRF_CRS_COMP_YR"])
	}
}

func TestCOMPParser_Specification(t *testing.T) {
	spec := GetCOMPSpec()
	
	// Test specification details
	if spec.FileType != "COMP" {
		t.Errorf("Expected FileType 'COMP', got '%s'", spec.FileType)
	}

	if spec.LineLength != 65 {
		t.Errorf("Expected LineLength 65, got %d", spec.LineLength)
	}

	if len(spec.Fields) != 8 {
		t.Errorf("Expected 8 fields, got %d", len(spec.Fields))
	}

	// Test field specifications
	expectedFields := []struct {
		name   string
		start  int
		length int
	}{
		{"INSTIT", 1, 4},
		{"ID", 5, 10},
		{"COURSE", 15, 20},
		{"COMPLETE", 35, 1},
		{"CRS_SRT", 36, 8},
		{"NSN", 44, 10},
		{"CRS_END", 54, 8},
		{"PBRF_CRS_COMP_YR", 62, 4},
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

// Test with actual COMP9170.txt data structure
func TestCOMPParser_ActualDataStructure(t *testing.T) {
	parser := NewCOMPParser()
	
	// Sample line from the actual COMP9170.txt file
	line := "9170917000047 2102-530            028092023 12033171106062024    "
	
	// Verify the line is exactly 65 characters
	if len(line) != 65 {
		t.Errorf("Expected line length 65, got %d", len(line))
	}
	
	record, err := parser.parseLine(line)
	if err != nil {
		t.Fatalf("Failed to parse actual data: %v", err)
	}
	
	// Verify the parsed fields match the expected structure
	expectedValues := map[string]string{
		"INSTIT":           "9170",
		"ID":               "917000047",
		"COURSE":           "2102-530",
		"COMPLETE":         "0",         // This is '0' in the actual data
		"CRS_SRT":          "28092023",
		"NSN":              "120331711",
		"CRS_END":          "06062024",
		"PBRF_CRS_COMP_YR": "",         // This is empty (spaces) in the actual data
	}
	
	for fieldName, expectedValue := range expectedValues {
		if record[fieldName] != expectedValue {
			t.Errorf("Field %s: expected '%s', got '%s'", fieldName, expectedValue, record[fieldName])
		}
	}
	
	// Verify all fields are present
	for _, field := range parser.GetSpec().Fields {
		if _, exists := record[field.Name]; !exists {
			t.Errorf("Field %s missing from parsed record", field.Name)
		}
	}
}
