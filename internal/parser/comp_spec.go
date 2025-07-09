package parser

// GetCOMPSpec returns the specification for COMP (Course Completion) files
func GetCOMPSpec() FileSpec {
	return FileSpec{
		FileType:    "COMP",
		Description: "Course Completion records",
		LineLength:  52, // Based on the sample data ending at position 52
		Fields: []FieldSpec{
			{
				Name:     "ID",
				Title:    "Student Identification Code",
				Start:    1,
				Length:   13,
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
				Name:     "CRS_SRT",
				Title:    "Course Start Date",
				Start:    36,
				Length:   8,
				Required: true,
			},
			{
				Name:     "CRS_END",
				Title:    "Course End Date",
				Start:    45,
				Length:   8,
				Required: true,
			},
		},
	}
}
