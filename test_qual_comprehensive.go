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
	// Test data from QUAL9170.txt
	testData := `9170917000478  140261767NZ2101            2024    
9170917000441  171046090NZ2101            2024    
9170917000409  138474289NZ2101            2024    
9170917000404  134424676NZ2101            2024    
9170917000281   98180910NZ2101            2024    
9170917000394  107766870NZ2101            2024    
9170917000380  132387160NZ2101            2024    
9170917000352  169588026NZ2101            2024    
9170917000340  129213856NZ2101            2024    
9170917000339  117184867NZ2101            2024    
9170917000311  126701699NZ2101            2024    `

	fmt.Println("Testing QUAL Parser...")
	fmt.Println(strings.Repeat("=", 50))

	// Test 1: Basic parsing
	fmt.Println("\n1. Testing Basic Parsing:")
	qualParser := parser.NewQUALParser()
	
	records, err := qualParser.Parse(testData)
	if err != nil {
		log.Fatalf("Parsing failed: %v", err)
	}

	fmt.Printf("   ✓ Successfully parsed %d records\n", len(records))
	fmt.Printf("   ✓ File type: %s\n", qualParser.GetFileType())

	// Test 2: Headers
	fmt.Println("\n2. Testing Headers:")
	headers := qualParser.GetHeaders()
	fmt.Printf("   Headers (%d): %v\n", len(headers), headers)

	// Test 3: Sample record
	fmt.Println("\n3. Testing Sample Record:")
	if len(records) > 0 {
		fmt.Printf("   First record:\n")
		for i, header := range headers {
			fieldName := qualParser.GetSpec().Fields[i].Name
			fmt.Printf("     %s: '%s'\n", header, records[0][fieldName])
		}
	}

	// Test 4: Field extraction verification
	fmt.Println("\n4. Testing Field Extraction:")
	line := "9170917000478  140261767NZ2101            2024    "
	fmt.Printf("   Test line: '%s'\n", line)
	fmt.Printf("   Line length: %d\n", len(line))
	
	// Test parsing of this specific line
	testParser := parser.NewQUALParser()
	testRecords, err := testParser.Parse(line)
	if err != nil {
		log.Fatalf("Single line parsing failed: %v", err)
	}
	
	if len(testRecords) > 0 {
		record := testRecords[0]
		fmt.Printf("   ✓ INSTIT: '%s'\n", record["INSTIT"])
		fmt.Printf("   ✓ ID: '%s'\n", record["ID"])
		fmt.Printf("   ✓ NSN: '%s'\n", record["NSN"])
		fmt.Printf("   ✓ QUAL: '%s'\n", record["QUAL"])
		fmt.Printf("   ✓ YR_REQ_MET: '%s'\n", record["YR_REQ_MET"])
	}

	// Test 5: File processing
	fmt.Println("\n5. Testing File Processing:")
	
	// Create a temporary test file
	tempDir := "temp_test"
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		log.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	testFile := filepath.Join(tempDir, "QUAL9170.txt")
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

	// Test 6: Verify all qualifications are NZ2101
	fmt.Println("\n6. Testing Data Consistency:")
	allNZ2101 := true
	for i, record := range records {
		if record["QUAL"] != "NZ2101" {
			fmt.Printf("   ✗ Record %d has qualification '%s', expected 'NZ2101'\n", i+1, record["QUAL"])
			allNZ2101 = false
		}
	}
	if allNZ2101 {
		fmt.Printf("   ✓ All %d records have qualification 'NZ2101'\n", len(records))
	}

	// Test 7: Verify year requirements
	fmt.Println("\n7. Testing Year Requirements:")
	all2024 := true
	for i, record := range records {
		if record["YR_REQ_MET"] != "2024" {
			fmt.Printf("   ✗ Record %d has year '%s', expected '2024'\n", i+1, record["YR_REQ_MET"])
			all2024 = false
		}
	}
	if all2024 {
		fmt.Printf("   ✓ All %d records have year requirement '2024'\n", len(records))
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("✓ All QUAL Parser tests completed successfully!")
}
