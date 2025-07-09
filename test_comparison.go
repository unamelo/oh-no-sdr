package main

import (
	"fmt"
	"os"
	"github.com/unamelo/oh-no-sdr/internal/parser"
)

func main() {
	// Test the comparison service
	fmt.Println("Testing comparison feature...")
	
	// Create a COUR parser
	courParser := parser.NewCourseEnrolmentParser()
	
	// Test enabling comparison
	courFilePath := "E:\\Development\\Go Project\\oh-no-sdr\\COUR9170.txt"
	if err := courParser.EnableComparison(courFilePath); err != nil {
		fmt.Printf("Error enabling comparison: %v\n", err)
		return
	}
	
	// Get warnings
	warnings := courParser.GetComparisonWarnings()
	for _, warning := range warnings {
		fmt.Printf("Warning: %s\n", warning)
	}
	
	// Test headers
	headers := courParser.GetHeaders()
	fmt.Printf("Headers count: %d\n", len(headers))
	fmt.Printf("Last few headers: %v\n", headers[len(headers)-3:])
	
	fmt.Println("Test completed successfully!")
}
