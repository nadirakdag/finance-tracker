# Finance Tracker Backend

REST API for personal finance tracking built with Go and SQLite.

## Tech Stack

- Go 1.21+
- SQLite3
- Gorilla Mux
- Prometheus metrics

## Project Structure
```
.
├── Dockerfile
├── README.md
├── cmd
│   └── server
│       └── main.go
├── configs
│   └── config.yaml
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── handlers
│   │   │   ├── category.go
│   │   │   ├── expense.go
│   │   │   ├── health.go
│   │   │   ├── income.go
│   │   │   ├── summary.go
│   │   │   └── utils.go
│   │   ├── middleware
│   │   │   ├── cors.go
│   │   │   └── logging.go
│   │   └── routes.go
│   ├── config
│   │   └── config.go
│   ├── domain
│   │   ├── errors
│   │   │   └── errors.go
│   │   ├── models
│   │   │   ├── category.go
│   │   │   ├── expense.go
│   │   │   ├── income.go
│   │   │   └── summary.go
│   │   └── services
│   │       ├── category.go
│   │       ├── expense.go
│   │       ├── income.go
│   │       └── summary.go
│   └── storage
│       ├── memory
│       │   └── store.go
│       ├── sqlite
│       │   └── store.go
│       └── storage.go
└── pkg
    ├── logger
    │   └── logger.go
    └── utils
        └── id.go
```

## Setup

1. Install dependencies:
```bash
go mod download
```

2. Run server:
```bash
go run cmd/server/main.go
```

## API Endpoints

### Categories
```
GET /api/v1/categories         # Get all categories
GET /api/v1/categories/income  # Get income sources
GET /api/v1/categories/expense # Get expense categories
```

### Transactions
```
POST /api/v1/expenses         # Create expense
GET  /api/v1/expenses         # List expenses
POST /api/v1/incomes          # Create income
GET  /api/v1/incomes          # List incomes
GET  /api/v1/summary          # Get financial summary
```

## Example Requests

Create Expense:
```bash
curl -X POST http://localhost:8080/api/v1/expenses \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 50.99,
    "category": "food",
    "description": "Grocery shopping",
    "date": "2024-02-24T00:00:00Z"
  }'
```

Create Income:
```bash
curl -X POST http://localhost:8080/api/v1/incomes \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 2500.00,
    "source": "salary",
    "description": "Monthly salary",
    "date": "2024-02-01T00:00:00Z"
  }'
```

Get Summary:
```bash
curl http://localhost:8080/api/v1/summary
```

## Database Schema

### Categories Table
```sql
CREATE TABLE categories (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL CHECK (type IN ('income', 'expense'))
);
```

### Expenses Table
```sql
CREATE TABLE expenses (
    id TEXT PRIMARY KEY,
    amount REAL NOT NULL,
    category TEXT NOT NULL,
    description TEXT NOT NULL,
    date DATE NOT NULL
);
```

### Incomes Table
```sql
CREATE TABLE incomes (
    id TEXT PRIMARY KEY,
    amount REAL NOT NULL,
    source TEXT NOT NULL,
    description TEXT NOT NULL,
    date DATE NOT NULL
);
```

## Metrics

Prometheus metrics available at `/metrics`:
- Total number of expenses
- Total number of incomes
- Transaction processing duration
- Active users count

## Docker

Build image:
```bash
docker build -t finance-tracker-backend .
```

Run container:
```bash
docker run -p 8080:8080 finance-tracker-backend
```
