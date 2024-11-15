// main.go
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

	// summary, err := ProcessCSV(config)
	summary, err := ProcessCSV(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing CSV: %v\n", err)
		os.Exit(1)
	}

	// Dump summary to null
	_ = summary

	// // Display summary (temporary output before JSON implementation)
	// fmt.Printf("Period: %s\n", summary.Period)
	// fmt.Printf("Total Income: %d\n", summary.TotalIncome)
	// fmt.Printf("Total Expenditure: %d\n", summary.TotalExpenditure)
	// fmt.Println("\nTransactions (sorted by date):")
	// for _, t := range summary.Transactions {
	// 	fmt.Printf("%v: %d - %s\n", t.Date.Format("2006/01/02"), t.Amount, t.Content)
	// }
}
