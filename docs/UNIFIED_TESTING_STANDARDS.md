# Unified Parser Testing Standards

## Overview
This document defines the standardized testing approach for all SDR file parsers (STUD, COUR, CREG, COMP, QUAL) to ensure consistency and maintainability.

## Standard Testing Pattern

### 1. Required Test Functions

Every parser should implement these core test functions:

```go
func Test{Parser}Parser_Parse(t *testing.T)           // Test main parsing functionality
func Test{Parser}Parser_GetFileType(t *testing.T)     // Test file type identifier
func Test{Parser}Parser_GetHeaders(t *testing.T)      // Test CSV headers
func Test{Parser}Parser_ValidateLine(t *testing.T)    // Test line validation
func Test{Parser}Parser_EmptyContent(t *testing.T)    // Test empty content handling
func Test{Parser}Parser_WithEmptyLines(t *testing.T)  // Test handling of empty lines
func Test{Parser}Parser_InvalidLineLength(t *testing.T) // Test error handling
func Test{Parser}Parser_RequiredFields(t *testing.T)  // Test required field validation
func Test{Parser}Parser_RealWorldData(t *testing.T)   // Test with actual file data
func Test{Parser}Parser_FieldExtraction(t *testing.T) // Test field position accuracy
```

### 2. Standard Interface Usage

All parsers must implement the common `Parser` interface:

```go
type Parser interface {
    Parse(content string) ([]map[string]string, error)
    GetHeaders() []string
    GetFileType() string
    ValidateLine(line string) error
}
```

### 3. Test Data Format

#### Input Data
- Use realistic test data with correct line lengths
- Include multiple records for comprehensive testing
- Use actual file samples when available

#### Expected Output
- Parse should return `[]map[string]string`
- Each record should be a map with field names as keys
- Field names should match the specification (e.g., "INSTIT", "ID", "COURSE")

### 4. Common Test Patterns

#### Basic Parse Test
```go
func TestCOMPParser_Parse(t *testing.T) {
    parser := NewCOMPParser()
    
    // Test with sample data (correct line length)
    content := `actual_test_data_here`
    
    records, err := parser.Parse(content)
    if err != nil {
        t.Fatalf("Parse failed: %v", err)
    }
    
    if len(records) != expected_count {
        t.Errorf("Expected %d records, got %d", expected_count, len(records))
    }
    
    // Test specific field values
    if records[0]["FIELD_NAME"] != "expected_value" {
        t.Errorf("Expected FIELD_NAME 'expected_value', got '%s'", records[0]["FIELD_NAME"])
    }
}
```

#### Header Test
```go
func TestCOMPParser_GetHeaders(t *testing.T) {
    parser := NewCOMPParser()
    headers := parser.GetHeaders()
    
    expected := []string{
        "Provider Code",
        "Student Identification Code",
        // ... all expected headers
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
```

#### Validation Test
```go
func TestCOMPParser_ValidateLine(t *testing.T) {
    parser := NewCOMPParser()
    
    // Test valid line
    validLine := "line_with_correct_length"
    if len(validLine) != expected_length {
        t.Errorf("Test line should be %d characters, got %d", expected_length, len(validLine))
    }
    
    err := parser.ValidateLine(validLine)
    if err != nil {
        t.Errorf("Valid line should pass validation: %v", err)
    }
    
    // Test invalid line
    invalidLine := "short"
    err = parser.ValidateLine(invalidLine)
    if err == nil {
        t.Error("Invalid line should fail validation")
    }
}
```

### 5. File-Specific Requirements

Each parser should follow these specific requirements:

#### STUD Parser (116 characters)
- Line length: 116 characters
- Required fields: INSTIT, ID, GENDER, DOB

#### COUR Parser (186 characters)  
- Line length: 186 characters
- Required fields: INSTIT, ID, QUAL, COURSE

#### CREG Parser (148 characters)
- Line length: 148 characters
- Required fields: INSTIT, COURSE, CTITLE

#### COMP Parser (65 characters)
- Line length: 65 characters
- Required fields: INSTIT, ID, COURSE, CRS_SRT, CRS_END

#### QUAL Parser (50 characters)
- Line length: 50 characters
- Required fields: INSTIT, ID, NSN, QUAL

### 6. Error Testing Standards

All parsers should test these error conditions:

1. **Empty content** - should return empty slice, no error
2. **Invalid line length** - should return error with descriptive message
3. **Missing required fields** - should return error specifying which field
4. **Empty lines** - should be skipped without error
5. **Malformed data** - should return error with line number

### 7. Import Requirements

All test files should include these imports:

```go
import (
    "strings"
    "testing"
)
```

### 8. Testing Data Sources

- Use actual data from the provided sample files when possible
- Ensure test data matches the exact line length requirements
- Include edge cases and boundary conditions
- Test with multiple records to verify consistency

### 9. Common Anti-Patterns to Avoid

❌ **Don't do this:**
- Mixed testing approaches (some tests using ParseLine, others using Parse)
- Hardcoded magic numbers without explanation
- Tests that depend on specific ordering of map keys
- Missing error message validation
- Tests without proper line length verification

✅ **Do this instead:**
- Consistent use of the Parse() method for all tests
- Clear documentation of expected line lengths
- Proper error message validation
- Comprehensive field extraction testing
- Real-world data validation

### 10. Test Execution

To run tests for a specific parser:
```bash
go test -v ./internal/parser -run TestCOMP
go test -v ./internal/parser -run TestCOUR
go test -v ./internal/parser -run TestSTUD
go test -v ./internal/parser -run TestCREG
go test -v ./internal/parser -run TestQUAL
```

To run all parser tests:
```bash
go test -v ./internal/parser
```

## Current Status

✅ **COMP Parser** - Fully updated with unified approach
✅ **COUR Parser** - Fully updated with unified approach  
⚠️ **STUD Parser** - Needs review and standardization
⚠️ **CREG Parser** - Needs review and standardization
⚠️ **QUAL Parser** - Needs review and standardization

## Benefits of This Approach

1. **Consistency** - All parsers follow the same testing pattern
2. **Maintainability** - Easy to understand and modify tests
3. **Reliability** - Comprehensive error handling and validation
4. **Documentation** - Tests serve as usage examples
5. **Debugging** - Clear error messages and proper failure reporting

This unified approach ensures that all SDR parsers are tested consistently and thoroughly, making the codebase more maintainable and reliable.
