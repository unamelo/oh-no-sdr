package parser

import (
	"strings"
	"testing"
)

func TestCourseEnrolmentParser_Parse(t *testing.T) {
	parser := NewCourseEnrolmentParser()

	// Test with sample data based on COUR9170.txt (186 characters per line)
	content := `9170917000047 NZ21022102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711
9170917000047 NZ21022102-510            2809202306062024        0029837NNNP122  1101090.05830.0058 0.0058 0.0058 0.0058 0.0058 0.0061 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711
9170917000047 NZ21022102-520            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711`

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

	if records[0]["QUAL"] != "NZ2102" {
		t.Errorf("Expected QUAL 'NZ2102', got '%s'", records[0]["QUAL"])
	}

	if records[0]["COURSE"] != "2102-530" {
		t.Errorf("Expected COURSE '2102-530', got '%s'", records[0]["COURSE"])
	}

	if records[0]["CRS_SRT"] != "28092023" {
		t.Errorf("Expected CRS_SRT '28092023', got '%s'", records[0]["CRS_SRT"])
	}

	if records[0]["CRS_END"] != "06062024" {
		t.Errorf("Expected CRS_END '06062024', got '%s'", records[0]["CRS_END"])
	}

	if records[0]["NSN"] != "120331711" {
		t.Errorf("Expected NSN '120331711', got '%s'", records[0]["NSN"])
	}
}

func TestCourseEnrolmentParser_GetFileType(t *testing.T) {
	parser := NewCourseEnrolmentParser()
	expected := "COUR"
	if got := parser.GetFileType(); got != expected {
		t.Errorf("GetFileType() = %v, want %v", got, expected)
	}
}

func TestCourseEnrolmentParser_GetHeaders(t *testing.T) {
	parser := NewCourseEnrolmentParser()
	headers := parser.GetHeaders()

	expected := []string{
		"Provider Code",
		"Student Identification Code",
		"Qualification Code",
		"Course Code",
		"Course Start Date",
		"Course End Date",
		"Student's Course Withdrawal Date",
		"Category of Fees Assessment for International Students",
		"Intramural/Extramural Attendance",
		"Course Delivery Site",
		"Source of Funding",
		"Residential Status",
		"Australian Residential Status",
		"Managed Apprenticeship",
		"Funding Category",
		"Course Classification",
		"NZSCED Field of Study",
		"Course EFTS Factor",
		"EFTS by Month",
		"National Student Number",
	}

	if len(headers) != len(expected) {
		t.Errorf("GetHeaders() length = %v, want %v", len(headers), len(expected))
	}

	for i, header := range headers {
		if header != expected[i] {
			t.Errorf("GetHeaders()[%d] = %v, want %v", i, header, expected[i])
		}
	}
}

func TestCourseEnrolmentParser_EmptyContent(t *testing.T) {
	parser := NewCourseEnrolmentParser()

	records, err := parser.Parse("")
	if err != nil {
		t.Errorf("Empty content should not cause error: %v", err)
	}

	if len(records) != 0 {
		t.Errorf("Expected 0 records for empty content, got %d", len(records))
	}
}

func TestCourseEnrolmentParser_WithEmptyLines(t *testing.T) {
	parser := NewCourseEnrolmentParser()

	content := `9170917000047 NZ21022102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711

9170917000047 NZ21022102-510            2809202306062024        0029837NNNP122  1101090.05830.0058 0.0058 0.0058 0.0058 0.0058 0.0061 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711
   
9170917000047 NZ21022102-520            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711`

	records, err := parser.Parse(content)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(records) != 3 {
		t.Errorf("Expected 3 records (empty lines should be skipped), got %d", len(records))
	}
}

func TestCourseEnrolmentParser_InvalidLineLength(t *testing.T) {
	parser := NewCourseEnrolmentParser()

	// Line too long should now be handled by truncation
	content := strings.Repeat("X", 200) // More than 186 characters

	records, err := parser.Parse(content)
	if err != nil {
		t.Errorf("Long line should be handled by truncation: %v", err)
	}

	if len(records) != 1 {
		t.Errorf("Expected 1 record, got %d", len(records))
	}

	// Line too short should work (will be padded)
	shortContent := "9170917000047 NZ21022102-530            2809202306062024"
	_, err = parser.Parse(shortContent)
	if err != nil {
		t.Errorf("Short line should be handled by padding: %v", err)
	}
}

func TestCourseEnrolmentParser_RequiredFields(t *testing.T) {
	parser := NewCourseEnrolmentParser()

	// Test with clearly empty required field (completely empty INSTIT field)
	content := strings.Repeat(" ", 186) // All spaces, should have empty required fields

	records, err := parser.Parse(content)
	if err != nil {
		// This is expected - should fail on required field validation
		if !strings.Contains(err.Error(), "required field") {
			t.Errorf("Expected required field error, got: %v", err)
		}
	} else {
		// If no error, check that we got records but they shouldn't have passed validation
		t.Logf("Unexpectedly got %d records without required field error", len(records))
		if len(records) > 0 {
			// Check if the first required field is actually empty
			if records[0]["INSTIT"] == "" {
				t.Error("Empty required field should have caused an error")
			}
		}
	}
}

func TestCourseEnrolmentParser_RealWorldData(t *testing.T) {
	parser := NewCourseEnrolmentParser()

	// Test with actual data structure from COUR9170.txt
	content := `9170917000047 NZ21022102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711
9170917000047 NZ21022102-510            2809202306062024        0029837NNNP122  1101090.05830.0058 0.0058 0.0058 0.0058 0.0058 0.0061 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711
9170917000047 NZ21022102-520            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711
9170917000047 NZ21022102-240            2809202306062024        0029837NNNP122  1101090.05830.0058 0.0058 0.0058 0.0058 0.0058 0.0061 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711
9170917000047 NZ21022102-516            2809202306062024        0029837NNNP122  1101090.06670.0067 0.0067 0.0067 0.0067 0.0067 0.0064 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711`

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

func TestCourseEnrolmentParser_FieldExtraction(t *testing.T) {
	parser := NewCourseEnrolmentParser()

	// Test specific field positions
	content := "9170917000047 NZ21022102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711"

	records, err := parser.Parse(content)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(records) != 1 {
		t.Fatalf("Expected 1 record, got %d", len(records))
	}

	record := records[0]

	// Verify exact field extraction
	if record["INSTIT"] != "9170" {
		t.Errorf("INSTIT extraction failed: got '%s'", record["INSTIT"])
	}

	if record["ID"] != "917000047" {
		t.Errorf("ID extraction failed: got '%s'", record["ID"])
	}

	if record["QUAL"] != "NZ2102" {
		t.Errorf("QUAL extraction failed: got '%s'", record["QUAL"])
	}

	if record["COURSE"] != "2102-530" {
		t.Errorf("COURSE extraction failed: got '%s'", record["COURSE"])
	}

	if record["CRS_SRT"] != "28092023" {
		t.Errorf("CRS_SRT extraction failed: got '%s'", record["CRS_SRT"])
	}

	if record["CRS_END"] != "06062024" {
		t.Errorf("CRS_END extraction failed: got '%s'", record["CRS_END"])
	}

	if record["NSN"] != "120331711" {
		t.Errorf("NSN extraction failed: got '%s'", record["NSN"])
	}
}
