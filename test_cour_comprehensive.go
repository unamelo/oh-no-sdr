package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/unamelo/oh-no-sdr/internal/parser"
)

func main() {
	fmt.Println("🔍 Testing COUR Parser Implementation")
	fmt.Println("====================================")
	
	// Create parser
	courParser := parser.NewCourseEnrolmentParser()
	
	// Test 1: Basic parser info
	fmt.Printf("📋 File Type: %s\n", courParser.GetFileType())
	fmt.Printf("📄 Description: %s\n", courParser.GetDescription())
	fmt.Printf("📏 Expected Line Length: %d\n", courParser.GetExpectedLineLength())
	fmt.Printf("🏷️  Number of Headers: %d\n", len(courParser.GetHeaders()))
	
	// Test 2: File type detection
	fmt.Println("\n🔍 Testing File Type Detection:")
	testFiles := []string{
		"COUR9170.txt",
		"cour9170.txt",
		"data.txt",
		"STUD9170.txt",
	}
	
	sampleLine := "9170917000047 NZ21022102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711"
	
	for _, filename := range testFiles {
		isMatch := courParser.IsMatchingFileType(filename, sampleLine)
		fmt.Printf("   %s: %v\n", filename, isMatch)
	}
	
	// Test 3: Line validation
	fmt.Println("\n✅ Testing Line Validation:")
	testLines := []string{
		sampleLine,
		"9170917000047 NZ2102", // Too short
		"    917000047 NZ21022102-530            2809202306062024        0029837NNNP122  1101090.11670.0117 0.0117 0.0117 0.0117 0.0117 0.0114 0.0000 0.0000 0.0000 0.0000 0.0000 0.0000  120331711", // Empty provider
	}
	
	for i, line := range testLines {
		err := courParser.ValidateLine(line)
		status := "✅ Valid"
		if err != nil {
			status = fmt.Sprintf("❌ Invalid: %v", err)
		}
		fmt.Printf("   Line %d: %s\n", i+1, status)
	}
	
	// Test 4: Field parsing
	fmt.Println("\n🔧 Testing Field Parsing:")
	values, err := courParser.ParseLine(sampleLine, 1)
	if err != nil {
		log.Fatalf("Failed to parse line: %v", err)
	}
	
	headers := courParser.GetHeaders()
	fmt.Printf("   Parsed %d fields:\n", len(values))
	for i, header := range headers {
		if i < len(values) {
			fmt.Printf("     %s: \"%s\"\n", header, values[i])
		}
	}
	
	// Test 5: Full file parsing
	fmt.Println("\n📁 Testing Full File Parsing:")
	
	// Read the COUR file
	inputFile := "COUR9170.txt"
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Printf("   ❌ File %s not found\n", inputFile)
		return
	}
	
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("   ❌ Error reading file: %v\n", err)
		return
	}
	
	// Parse the content
	records, err := courParser.Parse(string(content))
	if err != nil {
		fmt.Printf("   ❌ Error parsing file: %v\n", err)
		return
	}
	
	fmt.Printf("   ✅ Successfully parsed %d records\n", len(records))
	
	// Show sample record
	if len(records) > 0 {
		fmt.Println("\n📄 Sample Record (first record):")
		for field, value := range records[0] {
			fmt.Printf("     %s: \"%s\"\n", field, value)
		}
	}
	
	// Test 6: End-to-end processing
	fmt.Println("\n🎯 Testing End-to-End Processing:")
	result := parser.ProcessFile(inputFile, ".")
	
	if result.Error != nil {
		fmt.Printf("   ❌ Processing failed: %v\n", result.Error)
		return
	}
	
	fmt.Printf("   ✅ Processing successful!\n")
	fmt.Printf("   📊 Records processed: %d\n", result.RecordCount)
	fmt.Printf("   💾 Output file: %s\n", result.OutputFile)
	
	// Verify output file exists
	if _, err := os.Stat(result.OutputFile); err != nil {
		fmt.Printf("   ❌ Output file not created: %v\n", err)
		return
	}
	
	// Show first few lines of output
	outputContent, err := os.ReadFile(result.OutputFile)
	if err != nil {
		fmt.Printf("   ❌ Error reading output file: %v\n", err)
		return
	}
	
	lines := strings.Split(string(outputContent), "\n")
	fmt.Printf("\n📄 Output CSV Preview (first 3 lines):\n")
	for i, line := range lines {
		if i >= 3 || line == "" {
			break
		}
		fmt.Printf("   %d: %s\n", i+1, line)
	}
	
	fmt.Println("\n🎉 All tests completed successfully!")
	fmt.Println("\n💡 You can now use the TUI to parse COUR files:")
	fmt.Println("   go run main.go")
	fmt.Println("   Then select option 3: 'Parse COUR File'")
}
