package parser

import (
	"strings"
	"testing"
)

func TestQUALParser_Parse(t *testing.T) {
	parser := NewQUALParser()
	
	// Test with sample data based on QUAL9170.txt
	content := `9170917000478  140261767NZ2101            2024    
9170917000441  171046090NZ2101            2024    
9170917000409  138474289NZ2101            2024    `

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

	if records[0]["ID"] != "917000478" {
		t.Errorf("Expected ID '917000478', got '%s'", records[0]["ID"])
	}

	if records[0]["NSN"] != "140261767" {
		t.Errorf("Expected NSN '140261767', got '%s'", records[0]["NSN"])
	}

	if records[0]["QUAL"] != "NZ2101" {
		t.Errorf("Expected QUAL 'NZ2101', got '%s'", records[0]["QUAL"])
	}

	if records[0]["YR_REQ_MET"] != "2024" {
		t.Errorf("Expected YR_REQ_MET '2024', got '%s'", records[0]["YR_REQ_MET"])
	}
}

func TestQUALParser_GetHeaders(t *testing.T) {
	parser := NewQUALParser()
	headers := parser.GetHeaders()

	expected := []string{
		"Provider Code",
		"Student Identification Code",
		"National Student Number",
		"Qualification Code",
		"Main Subject 1",
		"Main Subject 2", 
		"Main Subject 3",
		"Year Requirements Met",
		"Padding",
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

func TestQUALParser_GetFileType(t *testing.T) {
	parser := NewQUALParser()
	fileType := parser.GetFileType()

	if fileType != "QUAL" {
		t.Errorf("Expected file type 'QUAL', got '%s'", fileType)
	}
}

func TestQUALParser_ValidateLine(t *testing.T) {
	parser := NewQUALParser()

	// Test valid line length (50 characters)
	validLine := "9170917000478  140261767NZ2101            2024    "
	if len(validLine) != 50 {
		t.Errorf("Test line should be 50 characters, got %d", len(validLine))
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

func TestQUALParser_ParseLine(t *testing.T) {
	parser := NewQUALParser()

	// Test normal line
	line := "9170917000478  140261767NZ2101            2024    "
	record, err := parser.parseLine(line)
	if err != nil {
		t.Fatalf("parseLine failed: %v", err)
	}

	// Verify all fields are extracted correctly
	if record["INSTIT"] != "9170" {
		t.Errorf("INSTIT: expected '9170', got '%s'", record["INSTIT"])
	}

	if record["ID"] != "917000478" {
		t.Errorf("ID: expected '917000478', got '%s'", record["ID"])
	}

	if record["NSN"] != "140261767" {
		t.Errorf("NSN: expected '140261767', got '%s'", record["NSN"])
	}

	if record["QUAL"] != "NZ2101" {
		t.Errorf("QUAL: expected 'NZ2101', got '%s'", record["QUAL"])
	}

	if record["YR_REQ_MET"] != "2024" {
		t.Errorf("YR_REQ_MET: expected '2024', got '%s'", record["YR_REQ_MET"])
	}
}

func TestQUALParser_EmptyContent(t *testing.T) {
	parser := NewQUALParser()
	
	records, err := parser.Parse("")
	if err != nil {
		t.Fatalf("Parse empty content failed: %v", err)
	}

	if len(records) != 0 {
		t.Errorf("Expected 0 records for empty content, got %d", len(records))
	}
}

func TestQUALParser_WithEmptyLines(t *testing.T) {
	parser := NewQUALParser()
	
	content := `9170917000478  140261767NZ2101            2024    

9170917000441  171046090NZ2101            2024    
   
9170917000409  138474289NZ2101            2024    `

	records, err := parser.Parse(content)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(records) != 3 {
		t.Errorf("Expected 3 records (empty lines should be skipped), got %d", len(records))
	}
}

func TestQUALParser_InvalidLineLength(t *testing.T) {
	parser := NewQUALParser()
	
	// Line too short
	content := "9170917000478  140261767"
	
	_, err := parser.Parse(content)
	if err == nil {
		t.Error("Expected error for invalid line length")
	}

	if !strings.Contains(err.Error(), "invalid line length") {
		t.Errorf("Expected line length error, got: %v", err)
	}
}

func TestQUALParser_RequiredFields(t *testing.T) {
	parser := NewQUALParser()
	
	// Test with missing required field (empty INSTIT)
	line := "    917000478  140261767NZ2101            2024    "
	_, err := parser.parseLine(line)
	if err == nil {
		t.Error("Expected error for empty required field")
	}

	if !strings.Contains(err.Error(), "required field") {
		t.Errorf("Expected required field error, got: %v", err)
	}
}

func TestQUALParser_MultipleRecords(t *testing.T) {
	parser := NewQUALParser()
	
	content := `9170917000478  140261767NZ2101            2024    
9170917000441  171046090NZ2101            2024    
9170917000409  138474289NZ2101            2024    `

	records, err := parser.Parse(content)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(records) != 3 {
		t.Errorf("Expected 3 records, got %d", len(records))
	}

	// Verify different NSNs
	expectedNSNs := []string{"140261767", "171046090", "138474289"}
	for i, expectedNSN := range expectedNSNs {
		if records[i]["NSN"] != expectedNSN {
			t.Errorf("Record %d: expected NSN '%s', got '%s'", i, expectedNSN, records[i]["NSN"])
		}
	}
}

func TestQUALParser_RealWorldData(t *testing.T) {
	parser := NewQUALParser()
	
	// Test with actual data from QUAL9170.txt
	content := `9170917000478  140261767NZ2101            2024    
9170917000441  171046090NZ2101            2024    
9170917000409  138474289NZ2101            2024    
9170917000404  134424676NZ2101            2024    
9170917000281   98180910NZ2101            2024    `

	records, err := parser.Parse(content)
	if err != nil {
		t.Fatalf("Parse real data failed: %v", err)
	}

	if len(records) != 5 {
		t.Errorf("Expected 5 records, got %d", len(records))
	}

	// Verify all records have the same qualification
	for i, record := range records {
		if record["QUAL"] != "NZ2101" {
			t.Errorf("Record %d: expected QUAL 'NZ2101', got '%s'", i, record["QUAL"])
		}
		if record["YR_REQ_MET"] != "2024" {
			t.Errorf("Record %d: expected YR_REQ_MET '2024', got '%s'", i, record["YR_REQ_MET"])
		}
	}
}

func TestQUALParser_FieldExtraction(t *testing.T) {
	parser := NewQUALParser()
	
	// Test specific field positions
	line := "9170917000478  140261767NZ2101            2024    "
	//        ^   ^         ^        ^     ^   ^   ^   ^   ^
	//        1   5         15       25    31  35  39  43  47
	//        |<4>|<-- 10 ->|<- 10 ->|<-6->|<4>|<4>|<4>|<4>|<4>|
	
	record, err := parser.parseLine(line)
	if err != nil {
		t.Fatalf("parseLine failed: %v", err)
	}

	// Verify exact field extraction
	if record["INSTIT"] != "9170" {
		t.Errorf("INSTIT extraction failed: got '%s'", record["INSTIT"])
	}

	if record["ID"] != "917000478" {
		t.Errorf("ID extraction failed: got '%s'", record["ID"])
	}

	if record["NSN"] != "140261767" {
		t.Errorf("NSN extraction failed: got '%s'", record["NSN"])
	}

	if record["QUAL"] != "NZ2101" {
		t.Errorf("QUAL extraction failed: got '%s'", record["QUAL"])
	}

	if record["YR_REQ_MET"] != "2024" {
		t.Errorf("YR_REQ_MET extraction failed: got '%s'", record["YR_REQ_MET"])
	}
}

func TestQUALParser_Specification(t *testing.T) {
	spec := GetQUALSpec()
	
	// Test specification details
	if spec.FileType != "QUAL" {
		t.Errorf("Expected FileType 'QUAL', got '%s'", spec.FileType)
	}

	if spec.LineLength != 50 {
		t.Errorf("Expected LineLength 50, got %d", spec.LineLength)
	}

	if len(spec.Fields) != 9 {
		t.Errorf("Expected 9 fields, got %d", len(spec.Fields))
	}

	// Test field specifications
	expectedFields := []struct {
		name   string
		start  int
		length int
	}{
		{"INSTIT", 1, 4},
		{"ID", 5, 10},
		{"NSN", 15, 10},
		{"QUAL", 25, 6},
		{"MAIN_1", 31, 4},
		{"MAIN_2", 35, 4},
		{"MAIN_3", 39, 4},
		{"YR_REQ_MET", 43, 4},
		{"PADDING", 47, 4},
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

func TestQUALParser_OptionalFields(t *testing.T) {
	parser := NewQUALParser()
	
	// Test with optional fields that may be empty
	line := "9170917000478  140261767NZ2101            2024    "
	record, err := parser.parseLine(line)
	if err != nil {
		t.Fatalf("parseLine failed: %v", err)
	}

	// Optional fields should be empty or have values
	for _, fieldName := range []string{"MAIN_1", "MAIN_2", "MAIN_3", "PADDING"} {
		if _, exists := record[fieldName]; !exists {
			t.Errorf("Optional field %s should exist in record", fieldName)
		}
	}
}
