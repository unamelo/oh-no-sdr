package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/oh-no-sdr/internal/parser"
)

func main() {
	// Test data from COMP9170.txt
	testData := `9170917000047 2102-530            028092023 12033171
9170917000047 2102-510            028092023 12033171
9170917000047 2102-520            028092023 12033171
9170917000047 2102-240            028092023 12033171
9170917000047 2102-516            028092023 12033171
9170917000047 2102-515            028092023 12033171
9170917000047 2102-511            028092023 12033171
9170917000047 2102-513            028092023 12033171
9170917000047 2102-514            228092023 12033171
9170917000047 2102-512            028092023 12033171`

	fmt.Println("Testing COMP Parser...")
	fmt.Println(strings.Repeat("=", 50))

	// Test 1: Basic parsing
	fmt.Println("\n1. Testing Basic Parsing:")
	compParser := parser.NewCOMPParser()
	
	records, err := compParser.Parse(testData)
	if err != nil {
		log.Fatalf("Parsing failed: %v", err)
	}

	fmt.Printf("   ✓ Successfully parsed %d records\n", len(records))
	fmt.Printf("   ✓ File type: %s\n", compParser.GetFileType())

	// Test 2: Headers
	fmt.Println("\n2. Testing Headers:")
	headers := compParser.GetHeaders()
	fmt.Printf("   Headers (%d): %v\n", len(headers), headers)

	// Test 3: Sample record
	fmt.Println("\n3. Testing Sample Record:")
	if len(records) > 0 {
		fmt.Printf("   First record:\n")
		for i, header := range headers {
			fieldName := compParser.GetSpec().Fields[i].Name
			fmt.Printf("     %s: '%s'\n", header, records[0][fieldName])
		}
	}

	// Test 4: Field extraction verification
	fmt.Println("\n4. Testing Field Extraction:")
	line := "9170917000047 2102-530            028092023 12033171"
	fmt.Printf("   Test line: '%s'\n", line)
	fmt.Printf("   Line length: %d\n", len(line))
	
	// Test parsing of this specific line
	testParser := parser.NewCOMPParser()
	testRecords, err := testParser.Parse(line)
	if err != nil {
		log.Fatalf("Single line parsing failed: %v", err)
	}
	
	if len(testRecords) > 0 {
		record := testRecords[0]
		fmt.Printf("   ✓ ID: '%s'\n", record["ID"])
		fmt.Printf("   ✓ COURSE: '%s'\n", record["COURSE"])
		fmt.Printf("   ✓ CRS_SRT: '%s'\n", record["CRS_SRT"])
		fmt.Printf("   ✓ CRS_END: '%s'\n", record["CRS_END"])
	}

	// Test 5: File processing
	fmt.Println("\n5. Testing File Processing:")
	
	// Create a temporary test file
	tempDir := "temp_test"
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		log.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	testFile := filepath.Join(tempDir, "COMP9170.txt")
	if err := os.WriteFile(testFile, []byte(testData), 0644); err != nil {
		log.Fatalf("Failed to create test file: %v", err)
	}
	
	// Process the file
	result := parser.ProcessFile(testFile, tempDir)
	if !result.Success {
		log.Fatalf("File processing failed: %v", result.Error)
	}
	
	fmt.Printf("   ✓ Input file: %s\n", result.InputFile)
	fmt.Printf("   ✓ Output file: %s\n", result.OutputFile)
	fmt.Printf("   ✓ Record count: %d\n", result.RecordCount)
	fmt.Printf("   ✓ File type: %s\n", result.FileType)
	
	// Read and display CSV output
	if csvContent, err := os.ReadFile(result.OutputFile); err == nil {
		fmt.Printf("   ✓ CSV output preview:\n")
		lines := strings.Split(string(csvContent), "\n")
		for i, line := range lines {
			if i < 3 && line != "" { // Show first 3 lines
				fmt.Printf("     %s\n", line)
			}
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("✓ All COMP Parser tests completed successfully!")
}
