package parser

// GetQUALSpec returns the specification for QUAL (Qualification Completion) files
func GetQUALSpec() FileSpec {
	return FileSpec{
		FileType:    "QUAL",
		Description: "Qualification Completion records",
		LineLength:  50, // Based on the file specification
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
				Name:     "NSN",
				Title:    "National Student Number",
				Start:    15,
				Length:   10,
				Required: true,
			},
			{
				Name:     "QUAL",
				Title:    "Qualification Code",
				Start:    25,
				Length:   6,
				Required: true,
			},
			{
				Name:     "MAIN_1",
				Title:    "Main Subject 1",
				Start:    31,
				Length:   4,
				Required: false,
			},
			{
				Name:     "MAIN_2",
				Title:    "Main Subject 2",
				Start:    35,
				Length:   4,
				Required: false,
			},
			{
				Name:     "MAIN_3",
				Title:    "Main Subject 3",
				Start:    39,
				Length:   4,
				Required: false,
			},
			{
				Name:     "YR_REQ_MET",
				Title:    "Year Requirements Met",
				Start:    43,
				Length:   4,
				Required: true,
			},
			{
				Name:     "PADDING",
				Title:    "Padding",
				Start:    47,
				Length:   4,
				Required: false,
			},
		},
	}
}
