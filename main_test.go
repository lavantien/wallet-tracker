// main_test.go
package main

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestValidatePeriod(t *testing.T) {
	tests := []struct {
		period string
		valid  bool
	}{
		{"202201", true},
		{"202213", false}, // Invalid month
		{"2022", false},   // Invalid format
	}

	for _, test := range tests {
		err := validatePeriod(test.period)
		if test.valid && err != nil {
			t.Errorf("Expected valid period for %s, got error: %v", test.period, err)
		}
		if !test.valid && err == nil {
			t.Errorf("Expected error for invalid period %s, got none", test.period)
		}
	}
}

func TestValidateFilePath(t *testing.T) {
	tests := []struct {
		path     string
		expected bool
	}{
		{"valid.csv", true}, // Assume this file will be created in the setup
	}

	for _, test := range tests {
		if test.expected {
			file, err := os.Create(test.path)
			if err != nil {
				t.Fatalf("Failed to create file %s: %v", test.path, err)
			}
			file.Close()
			defer os.Remove(test.path) // Clean up after test

			err = validateFilePath(test.path)
			if err != nil {
				t.Errorf("Expected valid file path for %s, got error: %v", test.path, err)
			}
		}

		// Test an invalid path
		invalidPath := "invalid.csv"
		err := validateFilePath(invalidPath)
		if err == nil {
			t.Errorf("Expected error for invalid file path %s, got none", invalidPath)
		}
	}
}

func TestParseTransaction(t *testing.T) {
	tests := []struct {
		record  []string
		lineNum int
		valid   bool
	}{
		{[]string{"2022/01/01", "1000", "income"}, 1, true},
		{[]string{"2022/01/01", "invalid", "income"}, 2, false}, // Invalid amount
		{[]string{"2022/01/01", "1000", ""}, 3, false},          // Empty content
	}

	for _, test := range tests {
		transaction, err := parseTransaction(test.record, test.lineNum)

		if test.valid && err != nil {
			t.Errorf("Expected valid transaction for record %v, got error: %v", test.record, err)
		}

		if !test.valid && err == nil {
			t.Errorf("Expected error for invalid transaction record %v, got none", test.record)
		}

		if test.valid && transaction.Content != "" && transaction.Amount <= 0 {
			t.Errorf("Expected positive amount for valid transaction %v, got: %d", test.record, transaction.Amount)
		}
	}
}

func TestFilterTransactionsByPeriod(t *testing.T) {
	testTransactions := []Transaction{
		{Date: parseDate(t, "2022/01/05"), Amount: -1000, Content: "eating out"},
		{Date: parseDate(t, "2022/01/06"), Amount: -10000, Content: "debit"},
		{Date: parseDate(t, "2022/02/03"), Amount: -1500, Content: "dining out"},
	}

	filteredTransactions := filterTransactionsByPeriod(testTransactions, "202201")

	if len(filteredTransactions) != 2 {
		t.Errorf("Expected 2 transactions for period 202201, got %d", len(filteredTransactions))
	}

	// Check if the correct transactions are filtered
	expectedContents := map[string]bool{
		"eating out": true,
		"debit":      true,
	}

	for _, trans := range filteredTransactions {
		if !expectedContents[trans.Content] {
			t.Errorf("Unexpected transaction content found: %s", trans.Content)
		}
	}
}

func parseDate(t *testing.T, date string) time.Time {
	parsedDate, err := time.Parse("2006/01/02", date)
	if err != nil {
		t.Fatalf("Failed to parse date %s: %v", date, err)
	}
	return parsedDate
}

func TestProcessCSV_ValidData(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testdata_*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up after test

	validCSV := "date,amount,content\n2022/01/01,1000,income\n2022/01/02,-500,eating out\n"
	if _, err := tempFile.WriteString(validCSV); err != nil {
		t.Fatalf("Failed to write valid CSV data: %v", err)
	}
	tempFile.Close() // Close file before processing

	config := &Config{Period: "202201", FilePath: tempFile.Name()}
	summary, err := ProcessCSV(config)
	if err != nil {
		t.Fatalf("Expected no error processing valid CSV, got: %v", err)
	}

	if summary.TotalIncome != 1000 || summary.TotalExpenditure != -500 {
		t.Errorf("Unexpected totals: %+v", summary)
	}
}

func TestProcessCSV_EmptyFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "emptydata_*.csv")
	if err != nil {
		t.Fatalf("Failed to create empty temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up after test

	config := &Config{Period: "202201", FilePath: tempFile.Name()}
	summary, err := ProcessCSV(config)
	if err != nil {
		t.Fatalf("Expected no error processing empty CSV, got: %v", err)
	}

	if len(summary.Transactions) != 0 || summary.TotalIncome != 0 || summary.TotalExpenditure != 0 {
		t.Errorf("Expected zero transactions and totals for empty CSV, got: %+v", summary)
	}
}

func TestProcessCSV_InvalidFormat(t *testing.T) {
	tempFile, err := os.CreateTemp("", "invaliddata_*.csv")
	if err != nil {
		t.Fatalf("Failed to create invalid temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up after test

	malformedCSV := "date|amount|content\n2022/01/01|1000|income\n" // Invalid delimiter
	if _, err := tempFile.WriteString(malformedCSV); err != nil {
		t.Fatalf("Failed to write invalid CSV data: %v", err)
	}
	tempFile.Close() // Close file before processing

	config := &Config{Period: "202201", FilePath: tempFile.Name()}
	_, err = ProcessCSV(config) // We only care about the error here
	if err == nil {
		t.Error("Expected error processing invalid CSV format, got none")
	}
}

func TestProcessCSV_NoTransactionsInPeriod(t *testing.T) {
	tempFile, err := os.CreateTemp("", "notransactions_*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up after test

	noTransactionsCSV := "date,amount,content\n2022/02/01,-1000,rent\n"
	if _, err := tempFile.WriteString(noTransactionsCSV); err != nil {
		t.Fatalf("Failed to write no transactions CSV data: %v", err)
	}
	tempFile.Close() // Close file before processing

	config := &Config{Period: "202201", FilePath: tempFile.Name()}
	summary, err := ProcessCSV(config)
	if err != nil {
		t.Fatalf("Expected no error processing CSV with no transactions in period but got: %v", err)
	}

	if len(summary.Transactions) != 0 || summary.TotalIncome != 0 || summary.TotalExpenditure != 0 {
		t.Errorf("Expected zero transactions and totals for CSV with no transactions in period; got summary %+v", summary)
	}
}

func TestJSONOutput(t *testing.T) {
	tempFile, _ := os.CreateTemp("", "jsondata_*.csv")
	defer os.Remove(tempFile.Name()) // Clean up after test

	validCSV := "date,amount,content\n2022/01/01,-500,eating out\n2022/01/02,-1000,rent\n"
	tempFile.WriteString(validCSV)

	config := &Config{Period: "202201", FilePath: tempFile.Name()}
	summary, _ := ProcessCSV(config)

	jsonOutputBytes, _ := json.Marshal(summary)

	var jsonOutput map[string]interface{}
	err := json.Unmarshal(jsonOutputBytes, &jsonOutput)

	if err != nil || jsonOutput["period"] == "" || jsonOutput["total_income"] == nil || jsonOutput["total_expenditure"] == nil || jsonOutput["transactions"] == nil {
		t.Error("JSON output is not structured correctly or is missing fields")
	}
}
