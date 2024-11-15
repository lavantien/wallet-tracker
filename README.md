# Wallet Tracker

(**coverage**: 72%)

generates a history of all deposits and withdrawals for a specified month

## Project Structure

```
wallet-tracker/
├── main.go     # Main application
├── parser.go   # CSV handling
├── process.go  # Business logic
├── README.md   # Documentation
└── data.csv    # Dataset
```

## Diagram and Flowchart

<details>
    <summary>... expand</summary>

```txt
+-------------------+
|    User Input     |
| (Year/Month, CSV) |
+---------+---------+
          |
          v
+-------------------+
|  Parse Arguments  |
+---------+---------+
          |
          v
+-------------------+
|  Validate Input   |
| (Period, FilePath)|
+---------+---------+
          |
          v
+--------------------+
|   Process CSV      |
| (Read, Parse,      |
| Filter Transactions|
| Calculate Totals)  |
+---------+----------+
          |
          v
+-------------------+
|  Generate Output  |
| (JSON Format)     |
+-------------------+
          |
          v
+-------------------+
|  Standard Output  |
+-------------------+
```

```txt
[Start]
    |
[Get User Input]
    |
[Parse Arguments]
    |
[Validate Period]
    |----------------------|
    |                      |
[Valid]                [Invalid]
    |                      |
[Validate File Path]  [Output Error]
    |                      |
[Open CSV File]         [End]
    |
[Check if File is Empty]
    |----------------------|
    |                      |
[Yes]                    [No]
    |                      |
[Return Empty Summary] [Read Header]
                           |
                       [Read Transactions]
                           |
                     [Calculate Totals]
                           |
                   [Generate JSON Output]
                           |
               [Display Output to Standard Output]
                           |
                         [End]
```

</details>

## MVP Tasks

Check `Projects` tab.

1. **Basic Setup**

   - Init Go project
   - Create README

2. **Core Features**

   - ft3 - Parse CLI args (YYYYMM, filepath)
   - ft4 - Read & validate CSV & Filter & Fake Output
   - ft5 - Total Income & Expenditure & Sort Output
   - ft6 - Generate JSON output

3. **Testing & Docs**

   - enh7 - Add test cases & Handle edge cases
   - enh8 - Write solution explanation

## Technical Details

<details>
    <summary>... expand</summary>

### Thought Process

1. **Understanding Requirements**:

   - The first step involved thoroughly reviewing the problem statement to understand the requirements for generating a history of deposits and withdrawals for a specified month from a CSV file. Key aspects included input validation, data processing, and output formatting.

2. **Defining Inputs and Outputs**:

   - I identified the inputs: a year/month in `YYYYMM` format and a file path to the CSV containing transaction records. The expected output was a JSON object summarizing transactions for the specified period.

3. **Breaking Down the Problem**:

   - I broke down the problem into smaller tasks:
     - Parsing command-line arguments.
     - Validating inputs (period format and file existence).
     - Reading and processing the CSV file.
     - Filtering transactions based on the specified month/year.
     - Calculating totals (income and expenditures).
     - Formatting the output as JSON.

4. **Implementation Planning**:

   - I outlined a plan for implementation, deciding to use Go due to its strong support for concurrency, performance, and built-in libraries for handling CSV and JSON data.

5. **Iterative Development**:

   - I adopted an iterative approach, implementing one feature at a time (e.g., input parsing, CSV processing) and testing each component thoroughly before moving on to the next.

### Technology Choices

1. **Programming Language**:

   - **Go (Golang)**: I chose Go for its simplicity, efficiency, and robust standard library support for file handling, CSV parsing, and JSON encoding/decoding. Go's strong typing helps catch errors at compile time, which is beneficial for maintaining code quality.

2. **Libraries Used**:

   - **encoding/csv**: This standard library package simplifies reading from CSV files.
   - **encoding/json**: This package is used to marshal data structures into JSON format easily.
   - **os**: For file operations (reading/writing files).
   - **regexp**: For validating input formats using regular expressions.

### Design Decisions

1. **Modular Architecture**:

   - The program was designed with a modular architecture where each function has a single responsibility (e.g., parsing arguments, processing CSV data). This enhances maintainability and readability.

2. **Error Handling**:

   - Comprehensive error handling was implemented throughout the program to ensure robustness. Each function checks for potential errors and returns meaningful messages if something goes wrong.

3. **Separation of Concerns**:

   - I separated input validation, data processing, and output generation into distinct functions. This design pattern makes it easier to modify or extend individual components without affecting others.

4. **Output Flexibility**:

   - An interface-based approach was considered for output generation (e.g., console vs. file output). This allows easy expansion in the future if new output formats or destinations are required.

### Requirement Fulfillment

1. **Input Validation**:

   - The program validates both the year/month format using regular expressions and checks if the specified file exists and is in CSV format.

2. **Data Processing**:

   - It reads transactions from the CSV file, filters them based on the specified month/year, calculates total income and expenditures, and formats this data into a structured JSON output.

3. **Output Format**:

   - The final output meets the specified JSON structure requirements with properties like `period`, `total_income`, `total_expenditure`, and an array of transactions.

4. **Challenges Overcome**:

   - One challenge was ensuring robust error handling when reading from potentially malformed CSV files or when no transactions matched the specified period. This was addressed by implementing thorough validation checks and returning appropriate error messages.

### Future Work

1. **Enhanced Output Options**:

   - Future enhancements could include adding options for different output formats (e.g., XML or Excel) or allowing users to specify additional filters (e.g., filtering by transaction type).
   - Dependency Injection: Define an interface for outputting data. For example, I could create an Outputter interface with methods like Write(data interface{}) error. Pass the chosen Outputter implementation as a parameter to the function that generates the output. This allows me to implement different output strategies (e.g., console output, file output) without changing the core logic.
   - Use a configuration structure or command-line flags to allow users to specify the desired output method (e.g., console or file).

1. **Containerization and CI/CD Pipeline**:

   - Using Docker Composed and Github Actions.

1. **User Interface Improvements**:

   - Implementing a graphical user interface (GUI) could make it easier for non-technical users to interact with the program.

1. **Performance Optimization**:

   - For large datasets, performance optimizations could be explored, such as parallel processing of transactions or more efficient data structures for storing transaction summaries.

1. **Unit Testing Expansion**:

   - While unit tests were created for core functionality, additional tests could be added to cover edge cases more comprehensively.

1. **Documentation and Examples**:

   - Providing detailed documentation and usage examples would help users understand how to use the program effectively.

</details>

## Data Flow

### Input (data.csv)

```csv
date,amount,content
2023/06/15,-720,transportation
2022/01/05,-1000,eating out
2022/01/06,-10000,debit
2022/02/03,-1500,dining out
2022/01/25,-100000,rent
2023/03/01,200000,salary
2022/01/15,50000,freelance payment
2022/01/30,-2500,utilities
2023/05/20,-15000,shopping
2022/01/10,-3000,groceries
2023/04/01,180000,bonus
2022/01/08,-800,coffee and snacks
2022/02/15,-95000,rent
2023/01/20,-12000,phone bill
2022/01/18,-5000,internet bill
```

### Output (stdout JSON)

```bash
go build . && ./wallet-tracker 202201 data.csv
```

```json
{
  "period": "2022/01",
  "total_income": 50000,
  "total_expenditure": -122300,
  "transactions": [
    {
      "Date": "2022-01-30T00:00:00Z",
      "Amount": -2500,
      "Content": "utilities"
    },
    {
      "Date": "2022-01-25T00:00:00Z",
      "Amount": -100000,
      "Content": "rent"
    },
    {
      "Date": "2022-01-18T00:00:00Z",
      "Amount": -5000,
      "Content": "internet bill"
    },
    {
      "Date": "2022-01-15T00:00:00Z",
      "Amount": 50000,
      "Content": "freelance payment"
    },
    {
      "Date": "2022-01-10T00:00:00Z",
      "Amount": -3000,
      "Content": "groceries"
    },
    {
      "Date": "2022-01-08T00:00:00Z",
      "Amount": -800,
      "Content": "coffee and snacks"
    },
    {
      "Date": "2022-01-06T00:00:00Z",
      "Amount": -10000,
      "Content": "debit"
    },
    {
      "Date": "2022-01-05T00:00:00Z",
      "Amount": -1000,
      "Content": "eating out"
    }
  ]
}
```

## Test Coverage

```bash
go test -v -cover
```

```go
=== RUN   TestValidatePeriod
--- PASS: TestValidatePeriod (0.00s)
=== RUN   TestValidateFilePath
--- PASS: TestValidateFilePath (0.00s)
=== RUN   TestParseTransaction
--- PASS: TestParseTransaction (0.00s)
=== RUN   TestFilterTransactionsByPeriod
--- PASS: TestFilterTransactionsByPeriod (0.00s)
=== RUN   TestProcessCSV_ValidData
{
  "period": "2022/01",
  "total_income": 1000,
  "total_expenditure": -500,
  "transactions": [
    {
      "Date": "2022-01-02T00:00:00Z",
      "Amount": -500,
      "Content": "eating out"
    },
    {
      "Date": "2022-01-01T00:00:00Z",
      "Amount": 1000,
      "Content": "income"
    }
  ]
}
--- PASS: TestProcessCSV_ValidData (0.00s)
=== RUN   TestProcessCSV_EmptyFile
--- PASS: TestProcessCSV_EmptyFile (0.00s)
=== RUN   TestProcessCSV_InvalidFormat
--- PASS: TestProcessCSV_InvalidFormat (0.00s)
=== RUN   TestProcessCSV_NoTransactionsInPeriod
{
  "period": "2022/01",
  "total_income": 0,
  "total_expenditure": 0,
  "transactions": null
}
--- PASS: TestProcessCSV_NoTransactionsInPeriod (0.00s)
=== RUN   TestJSONOutput
{
  "period": "2022/01",
  "total_income": 0,
  "total_expenditure": -1500,
  "transactions": [
    {
      "Date": "2022-01-02T00:00:00Z",
      "Amount": -1000,
      "Content": "rent"
    },
    {
      "Date": "2022-01-01T00:00:00Z",
      "Amount": -500,
      "Content": "eating out"
    }
  ]
}
--- PASS: TestJSONOutput (0.00s)
PASS
coverage: 72.0% of statements
ok      wallet-tracker  0.131s
```

## Business Logic
