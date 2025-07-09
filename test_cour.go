package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/unamelo/oh-no-sdr/internal/parser"
)

func main() {
	// Test COUR parser
	fmt.Println("Testing COUR parser...")
	
	// Create parser
	courParser := parser.NewCourseEnrolmentParser()
	fmt.Printf("File type: %s\n", courParser.GetFileType())
	fmt.Printf("Description: %s\n", courParser.GetDescription())
	fmt.Printf("Expected line length: %d\n", courParser.GetExpectedLineLength())
	
	// Test with actual file
	inputFile := "COUR9170.txt"
	if _, err := os.Stat(inputFile); err != nil {
		fmt.Printf("Error: Cannot find file %s: %v\n", inputFile, err)
		return
	}
	
	// Process file
	result := parser.ProcessFile(inputFile, ".")
	if result.Error != nil {
		fmt.Printf("Error processing file: %v\n", result.Error)
		return
	}
	
	fmt.Printf("Successfully processed %s\n", inputFile)
	fmt.Printf("File type: %s\n", result.FileType)
	fmt.Printf("Records processed: %d\n", result.RecordCount)
	fmt.Printf("Output file: %s\n", result.OutputFile)
	
	// Check if output file was created
	if _, err := os.Stat(result.OutputFile); err != nil {
		fmt.Printf("Warning: Output file not found: %v\n", err)
	} else {
		fmt.Printf("Output file created successfully!\n")
	}
}
