package parser

// CourseEnrolmentSpec defines the field specifications for COUR files
// Based on FILE SPECIFICATIONS – COURSE ENROLMENT FILE [COUR]
var CourseEnrolmentSpec = FileSpec{
	FileType:    "COUR",
	Description: "Course Enrolment File",
	LineLength:  186,
	Fields: []FieldSpec{
		{Name: "INSTIT", Title: "Provider Code", Start: 1, Length: 4, Required: true},
		{Name: "ID", Title: "Student Identification Code", Start: 5, Length: 10, Required: true},
		{Name: "QUAL", Title: "Qualification Code", Start: 15, Length: 6, Required: true},
		{Name: "COURSE", Title: "Course Code", Start: 21, Length: 20, Required: true},
		{Name: "CRS_SRT", Title: "Course Start Date", Start: 41, Length: 8, Required: false},
		{Name: "CRS_END", Title: "Course End Date", Start: 49, Length: 8, Required: false},
		{Name: "CRS_WTD", Title: "Student's Course Withdrawal Date", Start: 57, Length: 8, Required: false},
		{Name: "ASSIST", Title: "Category of Fees Assessment for International Students", Start: 65, Length: 2, Required: false},
		{Name: "ATTEND", Title: "Intramural/Extramural Attendance", Start: 67, Length: 1, Required: false},
		{Name: "CRS_SITE", Title: "Course Delivery Site", Start: 68, Length: 2, Required: false},
		{Name: "FUNDING", Title: "Source of Funding", Start: 70, Length: 2, Required: false},
		{Name: "RESIDENCY", Title: "Residential Status", Start: 72, Length: 1, Required: false},
		{Name: "AUS_RESIDENCY", Title: "Australian Residential Status", Start: 73, Length: 1, Required: false},
		{Name: "MANAAPPR", Title: "Managed Apprenticeship", Start: 74, Length: 1, Required: false},
		{Name: "CATEGORY", Title: "Funding Category", Start: 75, Length: 2, Required: false},
		{Name: "CLASS", Title: "Course Classification", Start: 77, Length: 4, Required: false},
		{Name: "NZSCED", Title: "NZSCED Field of Study", Start: 81, Length: 6, Required: false},
		{Name: "FACTOR", Title: "Course EFTS Factor", Start: 87, Length: 6, Required: false},
		{Name: "EFTS_MTH", Title: "EFTS by Month", Start: 93, Length: 84, Required: false},
		{Name: "NSN", Title: "National Student Number", Start: 177, Length: 10, Required: false},
	},
}
