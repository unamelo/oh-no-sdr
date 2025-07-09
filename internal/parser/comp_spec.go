package parser

// GetCOMPSpec returns the specification for COMP (Course Completion) files
func GetCOMPSpec() FileSpec {
	return FileSpec{
		FileType:    "COMP",
		Description: "Course Completion records",
		LineLength:  65, // Based on official specification ending at position 65
		Fields: []FieldSpec{
			{
				Name:     "INSTIT",
				Title:    "Provider Code",
				Start:    1,
				Length:   4,
				Required: true,
			},
			{
				Name:     "ID",
				Title:    "Student Identification Code",
				Start:    5,
				Length:   10,
				Required: true,
			},
			{
				Name:     "COURSE",
				Title:    "Course Code",
				Start:    15,
				Length:   20,
				Required: true,
			},
			{
				Name:     "COMPLETE",
				Title:    "Student Course Completion indicator",
				Start:    35,
				Length:   1,
				Required: false,
			},
			{
				Name:     "CRS_SRT",
				Title:    "Course Start Date",
				Start:    36,
				Length:   8,
				Required: true,
			},
			{
				Name:     "NSN",
				Title:    "National Student Number",
				Start:    44,
				Length:   10,
				Required: false,
			},
			{
				Name:     "CRS_END",
				Title:    "Course End Date",
				Start:    54,
				Length:   8,
				Required: true,
			},
			{
				Name:     "PBRF_CRS_COMP_YR",
				Title:    "PBRF Course Completion Year",
				Start:    62,
				Length:   4,
				Required: false,
			},
		},
	}
}
