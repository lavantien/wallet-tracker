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

	transactions, err := ProcessCSV(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing CSV: %v\n", err)
		os.Exit(1)
	}

	// Temporary output for testing
	fmt.Printf("Found %d transactions for period %s\n", len(transactions), config.Period)
	for _, t := range transactions {
		fmt.Printf("%v: %d - %s\n", t.Date.Format("2006/01/02"), t.Amount, t.Content)
	}
}
