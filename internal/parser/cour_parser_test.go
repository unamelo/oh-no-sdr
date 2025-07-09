package parser

import (
	"testing"
)

func TestCourseEnrolmentParser_GetFileType(t *testing.T) {
	parser := NewCourseEnrolmentParser()
	expected := "COUR"
	if got := parser.GetFileType(); got != expected {
		t.Errorf("GetFileType() = %v, want %v", got, expected)
	}
}

func TestCourseEnrolmentParser_GetDescription(t *testing.T) {
	parser := NewCourseEnrolmentParser()
	expected := "Course Enrolment File"
	if got := parser.GetDescription(); got != expected {
		t.Errorf("GetDescription() = %v, want %v", got, expected)
	}
}

func TestCourseEnrolmentParser_GetExpectedLineLength(t *testing.T) {
	parser := NewCourseEnrolmentParser()
	expected := 186
	if got := parser.GetExpectedLineLength(); got != expected {
		t.Errorf("GetExpectedLineLength() = %v, want %v", got, expected)
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

func TestCourseEnrolmentParser_ValidateLine(t *testing.T) {
	parser := NewCourseEnrolmentParser()
	
	tests := []struct {
		name    string
		line    string
		wantErr bool
	}{
		{
			name:    "Valid line",
			line:    "9170917000047 NZ21022102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711",
			wantErr: false,
		},
		{
			name:    "Line too short",
			line:    "9170917000047 NZ2102",
			wantErr: true,
		},
		{
			name:    "Empty required field - Provider Code",
			line:    "    917000047 NZ21022102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711",
			wantErr: true,
		},
		{
			name:    "Empty required field - Student ID",
			line:    "9170          NZ21022102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711",
			wantErr: true,
		},
		{
			name:    "Empty required field - Qualification Code",
			line:    "9170917000047       2102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711",
			wantErr: true,
		},
		{
			name:    "Empty required field - Course Code",
			line:    "9170917000047 NZ2102                    2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711",
			wantErr: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := parser.ValidateLine(tt.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateLine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCourseEnrolmentParser_ParseLine(t *testing.T) {
	parser := NewCourseEnrolmentParser()
	
	tests := []struct {
		name    string
		line    string
		lineNum int
		want    []string
		wantErr bool
	}{
		{
			name:    "Valid line - complete record",
			line:    "9170917000047 NZ21022102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711",
			lineNum: 1,
			want: []string{
				"9170",                                   // Provider Code
				"917000047",                              // Student Identification Code
				"NZ2102",                                 // Qualification Code
				"2102-530",                               // Course Code
				"28092023",                               // Course Start Date
				"06062024",                               // Course End Date
				"",                                       // Student's Course Withdrawal Date
				"00",                                     // Category of Fees Assessment for International Students
				"2",                                      // Intramural/Extramural Attendance
				"98",                                     // Course Delivery Site
				"37",                                     // Source of Funding
				"N",                                      // Residential Status
				"N",                                      // Australian Residential Status
				"N",                                      // Managed Apprenticeship
				"P1",                                     // Funding Category
				"2211",                                   // Course Classification (Need to check actual position)
				"010109",                                 // NZSCED Field of Study
				"0.1167",                                 // Course EFTS Factor
				"0.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000", // EFTS by Month
				"120331711",                              // National Student Number
			},
			wantErr: false,
		},
		{
			name:    "Invalid line - too short",
			line:    "9170917000047 NZ2102",
			lineNum: 1,
			want:    nil,
			wantErr: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.ParseLine(tt.line, tt.lineNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != len(tt.want) {
				t.Errorf("ParseLine() got length = %v, want length %v", len(got), len(tt.want))
				return
			}
			for i := 0; i < len(tt.want) && i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("ParseLine() got[%d] = %v, want[%d] = %v", i, got[i], i, tt.want[i])
				}
			}
		})
	}
}

func TestCourseEnrolmentParser_IsMatchingFileType(t *testing.T) {
	parser := NewCourseEnrolmentParser()
	
	tests := []struct {
		name      string
		filename  string
		firstLine string
		want      bool
	}{
		{
			name:      "Matching filename - COUR",
			filename:  "COUR9170.txt",
			firstLine: "9170917000047 NZ21022102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711",
			want:      true,
		},
		{
			name:      "Matching filename - lowercase",
			filename:  "cour9170.txt",
			firstLine: "9170917000047 NZ21022102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711",
			want:      true,
		},
		{
			name:      "Matching content - correct length and format",
			filename:  "data.txt",
			firstLine: "9170917000047 NZ21022102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711",
			want:      true,
		},
		{
			name:      "Non-matching filename",
			filename:  "STUD9170.txt",
			firstLine: "9170917000047F05081991  4955YUILB04202020223200913NZL0 1Y                      120331711    0    0111      41204120",
			want:      false,
		},
		{
			name:      "Non-matching content - wrong length",
			filename:  "data.txt",
			firstLine: "9170917000047F05081991  4955YUILB04202020223200913NZL0 1Y                      120331711    0    0111      41204120",
			want:      false,
		},
		{
			name:      "Non-matching content - wrong format",
			filename:  "data.txt",
			firstLine: "abcd" + strings.Repeat(" ", 182),
			want:      false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parser.IsMatchingFileType(tt.filename, tt.firstLine); got != tt.want {
				t.Errorf("IsMatchingFileType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCourseEnrolmentParser_FieldPositions(t *testing.T) {
	parser := NewCourseEnrolmentParser()
	
	// Test with actual sample data
	sampleLine := "9170917000047 NZ21022102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711"
	
	values, err := parser.ParseLine(sampleLine, 1)
	if err != nil {
		t.Fatalf("ParseLine() error = %v", err)
	}
	
	// Test specific field extractions
	tests := []struct {
		fieldName string
		index     int
		expected  string
	}{
		{"Provider Code", 0, "9170"},
		{"Student Identification Code", 1, "917000047"},
		{"Qualification Code", 2, "NZ2102"},
		{"Course Code", 3, "2102-530"},
		{"Course Start Date", 4, "28092023"},
		{"Course End Date", 5, "06062024"},
		{"National Student Number", 19, "120331711"},
	}
	
	for _, tt := range tests {
		t.Run(tt.fieldName, func(t *testing.T) {
			if values[tt.index] != tt.expected {
				t.Errorf("Field %s: got %v, want %v", tt.fieldName, values[tt.index], tt.expected)
			}
		})
	}
}
