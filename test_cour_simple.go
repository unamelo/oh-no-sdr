package main

import (
	"fmt"
	"log"
	"os"

	"github.com/unamelo/oh-no-sdr/internal/parser"
)

func main() {
	fmt.Println("Testing COUR parser...")
	
	// Test parsing COUR file
	inputFile := "COUR9170.txt"
	
	// Check if file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		log.Fatalf("File %s not found", inputFile)
	}
	
	// Process the file
	result := parser.ProcessFile(inputFile, ".")
	
	if result.Error != nil {
		log.Fatalf("Error processing file: %v", result.Error)
	}
	
	fmt.Printf("✅ Successfully processed %s\n", inputFile)
	fmt.Printf("📄 File type: %s\n", result.FileType)
	fmt.Printf("📊 Records processed: %d\n", result.RecordCount)
	fmt.Printf("💾 Output file: %s\n", result.OutputFile)
	
	// Verify output file exists
	if _, err := os.Stat(result.OutputFile); err != nil {
		log.Fatalf("Output file not created: %v", err)
	}
	
	fmt.Println("🎉 COUR parser test completed successfully!")
}
