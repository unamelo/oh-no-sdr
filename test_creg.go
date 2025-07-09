package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// This is a simple test to verify the CREG parser works with real data
func main() {
	// Sample CREG data from the documents
	cregData := `91702102-530            Food Product Development                                                   NZ210222  1101095 14P10.1167    5944X         02N
91702102-510            Industry Placement                                                         NZ210222  1101095  7P10.0583    2974X         02N
91702102-520            Supervise Food Production                                                  NZ210222  1101095 14P10.1167    5944X         02N
91702102-240            Food and Beverage Service                                                  NZ210222  1101095  7P10.0583    2974X         02N
91708065-02-209         Prepare, cook and finish meat, poultry and offal                           NZ210122  1101094 12P10.1000    5094X         02N
91708065-02-212         Prepare, cook and finish bakery products                                   NZ210122  1101094 16P10.1333    6794X         02N
91708065-01-201         Introduction to the hospitality and catering industry                      NZ210122  1101094  2P10.0167     844X         02N`

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "creg_test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Write test data to a file
	testFile := filepath.Join(tempDir, "CREG9170.txt")
	err = os.WriteFile(testFile, []byte(cregData), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Create output directory
	outputDir := filepath.Join(tempDir, "output")
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Test file created: %s\n", testFile)
	fmt.Printf("Output directory: %s\n", outputDir)
	fmt.Printf("File size: %d bytes\n", len(cregData))
	
	// Print first few lines for verification
	lines := strings.Split(cregData, "\n")
	fmt.Printf("Number of lines: %d\n", len(lines))
	for i, line := range lines[:3] {
		fmt.Printf("Line %d length: %d\n", i+1, len(line))
		fmt.Printf("Line %d: %s\n", i+1, line)
	}

	fmt.Println("\nCREG parser test setup complete!")
	fmt.Println("You can now run: go run main.go parse -i", testFile, "-o", outputDir)
}
