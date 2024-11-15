// process.go
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Transaction represents a single financial transaction
type Transaction struct {
	Date    time.Time
	Amount  int
	Content string
}

// TransactionSummary represents the processed transaction data
type TransactionSummary struct {
	Period           string        `json:"period"` // YYYY/MM format
	TotalIncome      int           `json:"total_income"`
	TotalExpenditure int           `json:"total_expenditure"`
	Transactions     []Transaction `json:"transactions"`
}

func ProcessCSV(config *Config) (*TransactionSummary, error) {
	file, err := os.Open(config.FilePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read and validate header
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading header: %v", err)
	}
	if err := validateHeader(header); err != nil {
		return nil, err
	}

	// Read all transactions
	var transactions []Transaction
	lineNum := 1 // Start after header

	for {
		lineNum++
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading line %d: %v", lineNum, err)
		}

		// Parse and validate transaction
		transaction, err := parseTransaction(record, lineNum)
		if err != nil {
			return nil, fmt.Errorf("error on line %d: %v", lineNum, err)
		}

		transactions = append(transactions, transaction)
	}

	// Filter transactions for the specified period
	filtered := filterTransactionsByPeriod(transactions, config.Period)

	// Sort transactions by date in descending order
	sortTransactions(filtered)

	// Calculate totals
	summary := calculateSummary(filtered, config.Period)

	// Marshal summary to JSON
	jsonOutput, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling summary to JSON: %v", err)
	}

	// Output the JSON
	fmt.Println(string(jsonOutput))

	return summary, nil
}

// sortTransactions sorts transactions by date in descending order
func sortTransactions(transactions []Transaction) {
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Date.After(transactions[j].Date)
	})
}

// calculateSummary processes transactions and returns a summary
func calculateSummary(transactions []Transaction, period string) *TransactionSummary {
	summary := &TransactionSummary{
		Period:       formatPeriod(period), // Convert YYYYMM to YYYY/MM
		Transactions: transactions,
	}

	for _, t := range transactions {
		if t.Amount > 0 {
			summary.TotalIncome += t.Amount
		} else {
			summary.TotalExpenditure += t.Amount
		}
	}

	return summary
}

// formatPeriod converts YYYYMM to YYYY/MM format
func formatPeriod(period string) string {
	return period[:4] + "/" + period[4:]
}

func validateHeader(header []string) error {
	expectedHeader := []string{"date", "amount", "content"}
	if len(header) != len(expectedHeader) {
		return fmt.Errorf("invalid header: expected %v, got %v", expectedHeader, header)
	}
	for i, field := range header {
		if strings.ToLower(field) != expectedHeader[i] {
			return fmt.Errorf("invalid header field: expected %s, got %s", expectedHeader[i], field)
		}
	}
	return nil
}

func parseTransaction(record []string, lineNum int) (Transaction, error) {
	if len(record) != 3 {
		return Transaction{}, fmt.Errorf("invalid number of fields")
	}

	// Parse date
	date, err := time.Parse("2006/01/02", record[0])
	if err != nil {
		return Transaction{}, fmt.Errorf("invalid date format: %s", record[0])
	}

	// Parse amount
	amount, err := strconv.Atoi(record[1])
	if err != nil {
		return Transaction{}, fmt.Errorf("invalid amount: %s", record[1])
	}

	// Validate content
	content := strings.TrimSpace(record[2])
	if content == "" {
		return Transaction{}, fmt.Errorf("empty content")
	}

	return Transaction{
		Date:    date,
		Amount:  amount,
		Content: content,
	}, nil
}

func filterTransactionsByPeriod(transactions []Transaction, period string) []Transaction {
	targetYear := period[:4]
	targetMonth := period[4:]

	var filtered []Transaction
	for _, t := range transactions {
		if fmt.Sprintf("%04d%02d", t.Date.Year(), t.Date.Month()) == targetYear+targetMonth {
			filtered = append(filtered, t)
		}
	}
	return filtered
}
