package main

import (
	"fmt"
	"os"
)

func main() {
	config, err := ParseArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Processing period: %s\n", config.Period)
	fmt.Printf("Using file: %s\n", config.FilePath)
}
