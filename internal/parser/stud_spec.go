package parser

// GetSTUDSpec returns the specification for STUD (Student) files
func GetSTUDSpec() FileSpec {
	return FileSpec{
		FileType:    "STUD",
		Description: "Student File",
		LineLength:  116,
		Fields: []FieldSpec{
			{Name: "INSTIT", Title: "Provider Code", Start: 1, Length: 4, Required: true},
			{Name: "ID", Title: "Student Identification Code", Start: 5, Length: 10, Required: true},
			{Name: "GENDER", Title: "Gender", Start: 15, Length: 1, Required: true},
			{Name: "DOB", Title: "Date of Birth", Start: 16, Length: 8, Required: true},
			{Name: "TOTAL_FEE", Title: "Total fee for domestic student", Start: 24, Length: 6, Required: false},
			{Name: "NAMEID", Title: "Name ID Code", Start: 30, Length: 5, Required: true},
			{Name: "PRIOR_A", Title: "Main Activity at 1 October in Year Prior to Formal Enrolment", Start: 35, Length: 2, Required: false},
			{Name: "FIRST_YR", Title: "First Year of Tertiary Education", Start: 37, Length: 4, Required: false},
			{Name: "DIS_ACCESS", Title: "Disability Services Accessed Indicator", Start: 41, Length: 1, Required: false},
			{Name: "S_SCHOOL", Title: "Last Secondary School Attended", Start: 42, Length: 4, Required: false},
			{Name: "Y_SCHOOL", Title: "Last Year at Secondary School", Start: 46, Length: 4, Required: false},
			{Name: "SEC_QUAL", Title: "Highest Secondary School Qualification", Start: 50, Length: 2, Required: false},
			{Name: "CITIZEN", Title: "Country of Citizenship", Start: 52, Length: 3, Required: false},
			{Name: "FEES_FREE_ELIGIBLE", Title: "Fees Free Eligibility indicator", Start: 55, Length: 1, Required: false},
			{Name: "REMOVED_FIELD", Title: "Removed field (padded blanks)", Start: 56, Length: 1, Required: false},
			{Name: "DISABILITY", Title: "Disability Indicator", Start: 57, Length: 1, Required: false},
			{Name: "FINISH", Title: "Expectation to Complete a Qualification this year", Start: 58, Length: 1, Required: false},
			{Name: "IWI", Title: "Iwi Affiliation", Start: 59, Length: 12, Required: false},
			{Name: "IRDNOS", Title: "Padded Blanks (previously IRD Number)", Start: 71, Length: 9, Required: false},
			{Name: "NSN", Title: "National Student Number", Start: 80, Length: 10, Required: false},
			{Name: "FOREIGN_FEE", Title: "Tuition fee paid by international fee-paying student", Start: 90, Length: 5, Required: false},
			{Name: "MAX_EXEMPT_FEE", Title: "Maxima Exempt Fees", Start: 95, Length: 5, Required: false},
			{Name: "ETHNIC", Title: "Ethnicity", Start: 100, Length: 9, Required: false},
			{Name: "PERM_POST_CODE", Title: "Permanent Post Code", Start: 109, Length: 4, Required: false},
			{Name: "TERM_POST_CODE", Title: "Term Post Code", Start: 113, Length: 4, Required: false},
		},
	}
}
