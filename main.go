package main

import (
	"fmt"
	"ledgerCo-loans/entity"
	"ledgerCo-loans/processor"
	"os"
)

func main() {
	// command line args validation
	if len(os.Args) != 2 {
		fmt.Println("Please provide input text file location")
		return
	}

	// create file processor object to process commands
	if proc := processor.New(entity.FileProcessor, os.Args); proc != nil {
		proc.Process()
	}
}
