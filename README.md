# wallet-tracker

generates a history of all deposits and withdrawals for a specified month

# Wallet Tracker - Core Tasks

## Project Structure

```
wallet-tracker/
├── main.go      # Main application
├── parser.go    # CSV handling
├── process.go   # Business logic
└── README.md    # Documentation
```

## Kanban Board Items

### 📌 MVP Tasks

1. **Basic Setup**

   - Init Go project
   - Create README

2. **Core Features**

   - Parse CLI args (YYYYMM, filepath)
   - Read & validate CSV
   - Filter transactions by month
   - Generate JSON output

3. **Testing & Docs**

   - Add test cases
   - Write solution explanation
   - Update README

## Dependency List

```go
// Minimal dependencies
encoding/csv   // CSV parsing
encoding/json  // JSON output
time          // Date handling
sort          // Transaction sorting
```

## Implementation Order

1. Setup project
2. Implement CSV reading
3. Add business logic
4. Add tests
5. Write documentation
