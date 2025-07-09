package main

import (
	"fmt"
	"log"
	"os"

	"github.com/unamelo/oh-no-sdr/internal/parser"
)

func main() {
	fmt.Println("🔍 COUR Parser Demo")
	fmt.Println("===================")
	
	// Test the COUR parser
	inputFile := "COUR9170.txt"
	
	fmt.Printf("📁 Processing file: %s\n", inputFile)
	
	// Use the processor to handle the file
	result := parser.ProcessFile(inputFile, ".")
	
	if result.Error != nil {
		log.Fatalf("❌ Error processing file: %v", result.Error)
	}
	
	fmt.Printf("✅ Successfully processed %s\n", result.InputFile)
	fmt.Printf("📄 File type detected: %s\n", result.FileType)
	fmt.Printf("📊 Records processed: %d\n", result.RecordCount)
	fmt.Printf("💾 Output file created: %s\n", result.OutputFile)
	
	// Show file size
	if info, err := os.Stat(result.OutputFile); err == nil {
		fmt.Printf("📏 Output file size: %d bytes\n", info.Size())
	}
	
	fmt.Println("\n🎉 COUR parser is working correctly!")
	fmt.Println("You can now:")
	fmt.Println("  1. Run 'go run main.go' to use the TUI")
	fmt.Println("  2. Select option 3 'Parse COUR File'")
	fmt.Println("  3. The parser will automatically detect and process COUR files")
}
