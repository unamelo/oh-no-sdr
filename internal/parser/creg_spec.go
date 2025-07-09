package parser

// GetCREGSpec returns the specification for CREG (Course Register) files
func GetCREGSpec() FileSpec {
	return FileSpec{
		FileType:    "CREG",
		Description: "Course Register File",
		LineLength:  148,
		Fields: []FieldSpec{
			{Name: "INSTIT", Title: "Provider Code", Start: 1, Length: 4, Required: true},
			{Name: "COURSE", Title: "Course Code", Start: 5, Length: 20, Required: true},
			{Name: "CTITLE", Title: "Course Title", Start: 25, Length: 75, Required: true},
			{Name: "QUAL", Title: "Qualification Code", Start: 100, Length: 6, Required: true},
			{Name: "CLASS", Title: "Course Classification", Start: 106, Length: 4, Required: true},
			{Name: "NZSCED", Title: "NZSCED Field of Study", Start: 110, Length: 6, Required: true},
			{Name: "NZQCFLEVEL", Title: "Level on the NZ Qualifications and Credentials Framework", Start: 116, Length: 1, Required: true},
			{Name: "CREDIT", Title: "Credit", Start: 117, Length: 3, Required: false},
			{Name: "CATEGORY", Title: "Funding Category", Start: 120, Length: 2, Required: true},
			{Name: "FACTOR", Title: "Course EFTS Factor", Start: 122, Length: 6, Required: true},
			{Name: "STAGE", Title: "Stage of Pre-Service Teacher Education Qualification", Start: 128, Length: 2, Required: false},
			{Name: "FEE", Title: "Course Tuition Fee", Start: 132, Length: 4, Required: false},
			{Name: "INTERNET", Title: "Internet Based Learning Indicator", Start: 136, Length: 1, Required: false},
			{Name: "PBRF_ELIGIBLE", Title: "PBRF Eligible Course Indicator", Start: 137, Length: 9, Required: false},
			{Name: "CCCOSTS_FEE", Title: "Compulsory Course Costs Fee", Start: 146, Length: 1, Required: false},
			{Name: "EXEMPT_INDICATOR", Title: "Course Exemption from AMFM", Start: 147, Length: 1, Required: false},
			{Name: "EMB_LIT_NUM", Title: "Embedded Literacy and Numeracy Flag", Start: 148, Length: 1, Required: false},
		},
	}
}
