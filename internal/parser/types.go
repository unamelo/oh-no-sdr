package parser

// FieldSpec defines a field in an SDR file
type FieldSpec struct {
	Name     string // Field name (e.g., "INSTIT")
	Title    string // Descriptive title (e.g., "Provider Code")
	Start    int    // 1-based starting position
	Length   int    // Field length in characters
	Required bool   // Whether field is required
}

// FileSpec defines the structure of an SDR file type
type FileSpec struct {
	FileType    string      // STUD, COUR, CREG, etc.
	Description string      // Human-readable description
	LineLength  int         // Expected line length
	Fields      []FieldSpec // Field definitions
}

// Parser interface for different file types
type Parser interface {
	Parse(content string) ([]map[string]string, error)
	GetHeaders() []string
	GetFileType() string
}
