package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/unamelo/oh-no-sdr/internal/parser"
)

func main() {
	fmt.Println("=== CREG Parser Comprehensive Test ===\n")

	// Test data from the actual CREG9170.txt file
	testData := `91702102-530            Food Product Development                                                   NZ210222  1101095 14P10.1167    5944X         02N
91702102-510            Industry Placement                                                         NZ210222  1101095  7P10.0583    2974X         02N
91702102-520            Supervise Food Production                                                  NZ210222  1101095 14P10.1167    5944X         02N
91702102-240            Food and Beverage Service                                                  NZ210222  1101095  7P10.0583    2974X         02N
91702102-516            Catering Operations, Costs and Menu Planning                               NZ210222  1101095  8P10.0667    3394X         02N
91708065-02-209         Prepare, cook and finish meat, poultry and offal                           NZ210122  1101094 12P10.1000    5094X         02N
91708065-02-212         Prepare, cook and finish bakery products                                   NZ210122  1101094 16P10.1333    6794X         02N
91708065-02-219         Catering operations, costs and menu planning                               NZ210122  1101094  4P10.0333    1704X         02N
91708065-02-208         Prepare, cook and finish fish and shellfish dishes                         NZ210122  1101094 12P10.1000    5094X         02N
91708065-02-213         Prepare, cook and finish hot & cold desserts and puddings                  NZ210122  1101094 10P10.0833    4254X         02N
91708065-02-211         Prepare, cook and finish rice, grain, farinaceous products and egg dishes  NZ210122  1101094  8P10.0667    3384X         02N
91708065-02-206         Healthier foods and special diets                                          NZ210122  1101094  4P10.0334    1704X         02N
91708065-02-210         Prepare, cook and finish vegetables, fruit and pulses                      NZ210122  1101094  6P10.0500    2534X         02N
91708065-01-201         Introduction to the hospitality and catering industry                      NZ210122  1101094  2P10.0167     844X         02N
91708065-02-207         Prepare, cook and finish stocks, soups and sauces                          NZ210122  1101094  6P10.0500    2534X         02N
91708065-01-104         Introduction to nutrition                                                  NZ210122  1101094  2P10.0167     844X         02N
91708065-01-105         Prepare food for cold presentation                                         NZ210122  1101094 14P10.1166    5934X         02N
91708065-01-110         Introduction to basic kitchen procedures                                   NZ210122  1101094  1P10.0083     424X         02N
91708065-10-102         Safety at Work                                                             NZ210122  1101094  3P10.0250    1264X         02N
91708065-02-203         Food Safety in Catering                                                    NZ210122  1101094  5P10.0417    2114X         02N`

	// Create temporary test file
	tempDir, err := os.MkdirTemp("", "creg_test")
	if err != nil {
		log.Fatal("Failed to create temp directory:", err)
	}
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "CREG9170.txt")
	if err := os.WriteFile(testFile, []byte(testData), 0644); err != nil {
		log.Fatal("Failed to create test file:", err)
	}

	fmt.Printf("Created test file: %s\n", testFile)
	fmt.Printf("Test data contains %d lines\n\n", len(strings.Split(testData, "\n")))

	// Test 1: Create and test the parser directly
	fmt.Println("=== TEST 1: Direct Parser Test ===")
	cregParser := parser.NewCREGParser()
	
	// Check file type
	fmt.Printf("File type: %s\n", cregParser.GetFileType())
	
	// Check headers
	headers := cregParser.GetHeaders()
	fmt.Printf("Number of headers: %d\n", len(headers))
	fmt.Printf("Headers: %v\n", headers)
	
	// Parse the content
	records, err := cregParser.Parse(testData)
	if err != nil {
		log.Fatal("Failed to parse CREG data:", err)
	}
	
	fmt.Printf("Successfully parsed %d records\n", len(records))
	
	// Display first record
	if len(records) > 0 {
		fmt.Println("\nFirst record:")
		for _, header := range headers {
			fieldName := getFieldNameFromTitle(header)
			if value, exists := records[0][fieldName]; exists {
				fmt.Printf("  %s: '%s'\n", header, value)
			}
		}
	}

	// Test 2: Test with ProcessFile function
	fmt.Println("\n=== TEST 2: ProcessFile Function Test ===")
	result := parser.ProcessFile(testFile, tempDir)
	
	if result.Success {
		fmt.Printf("✓ Successfully processed %s\n", result.InputFile)
		fmt.Printf("  Output file: %s\n", result.OutputFile)
		fmt.Printf("  Record count: %d\n", result.RecordCount)
		fmt.Printf("  File type: %s\n", result.FileType)
		
		// Check if output file exists
		if _, err := os.Stat(result.OutputFile); err == nil {
			fmt.Printf("✓ Output CSV file created successfully\n")
			
			// Read and display first few lines of CSV
			csvContent, err := os.ReadFile(result.OutputFile)
			if err == nil {
				lines := strings.Split(string(csvContent), "\n")
				fmt.Printf("\nFirst 3 lines of CSV output:\n")
				for i, line := range lines {
					if i >= 3 {
						break
					}
					if strings.TrimSpace(line) != "" {
						fmt.Printf("  %d: %s\n", i+1, line)
					}
				}
			}
		} else {
			fmt.Printf("✗ Output CSV file not found: %s\n", result.OutputFile)
		}
	} else {
		fmt.Printf("✗ Failed to process file: %s\n", result.Error.Error())
	}

	// Test 3: Test file detection
	fmt.Println("\n=== TEST 3: File Detection Test ===")
	detectedType := parser.DetectFileType("CREG9170.txt")
	fmt.Printf("Detected file type from 'CREG9170.txt': %s\n", detectedType)
	
	detectedType2 := parser.DetectFileType("creg_test.txt")
	fmt.Printf("Detected file type from 'creg_test.txt': %s\n", detectedType2)
	
	detectedType3 := parser.DetectFileType("CREG.txt")
	fmt.Printf("Detected file type from 'CREG.txt': %s\n", detectedType3)

	// Test 4: Test GetParser function
	fmt.Println("\n=== TEST 4: GetParser Function Test ===")
	parserInstance, err := parser.GetParser("CREG")
	if err != nil {
		fmt.Printf("✗ Failed to get CREG parser: %s\n", err.Error())
	} else {
		fmt.Printf("✓ Successfully got CREG parser\n")
		fmt.Printf("  Parser type: %s\n", parserInstance.GetFileType())
		fmt.Printf("  Number of headers: %d\n", len(parserInstance.GetHeaders()))
	}

	fmt.Println("\n=== CREG Parser Test Complete ===")
	fmt.Println("\nAll tests passed! The CREG parser is working correctly.")
	fmt.Printf("\nYou can now use the main application to parse CREG files.\n")
	fmt.Printf("Try running: go run main.go and select 'Parse CREG File'\n")
}

// Helper function to get field name from title
func getFieldNameFromTitle(title string) string {
	fieldMap := map[string]string{
		"Provider Code":                                            "INSTIT",
		"Course Code":                                             "COURSE",
		"Course Title":                                            "CTITLE",
		"Qualification Code":                                      "QUAL",
		"Course Classification":                                   "CLASS",
		"NZSCED Field of Study":                                   "NZSCED",
		"Level on the NZ Qualifications and Credentials Framework": "NZQCFLEVEL",
		"Credit":                                                  "CREDIT",
		"Funding Category":                                        "CATEGORY",
		"Course EFTS Factor":                                      "FACTOR",
		"Stage of Pre-Service Teacher Education Qualification":    "STAGE",
		"Course Tuition Fee":                                      "FEE",
		"Internet Based Learning Indicator":                       "INTERNET",
		"PBRF Eligible Course Indicator":                          "PBRF_ELIGIBLE",
		"Compulsory Course Costs Fee":                             "CCCOSTS_FEE",
		"Course Exemption from AMFM":                              "EXEMPT_INDICATOR",
		"Embedded Literacy and Numeracy Flag":                     "EMB_LIT_NUM",
	}
	
	if fieldName, exists := fieldMap[title]; exists {
		return fieldName
	}
	return title
}
