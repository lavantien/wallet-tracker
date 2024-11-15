package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
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

// ProcessCSV reads and validates CSV file, returns transactions for the specified period
func ProcessCSV(config *Config) ([]Transaction, error) {
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

	return filterTransactionsByPeriod(transactions, config.Period), nil
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
