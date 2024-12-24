# Finance Tracker Backend

A Go-based backend service for tracking personal finances, including income and expenses.

## Features

- RESTful API for managing income and expenses
- In-memory storage with thread-safe operations
- Prometheus metrics for monitoring
- Structured logging
- CORS middleware
- Configuration management
- Docker support

## Project Structure

```
finance-tracker/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── expense.go         # Expense handler
│   │   │   ├── income.go          # Income handler
│   │   │   ├── summary.go         # Summary handler
│   │   │   └── utils.go           # Handler utilities (error handling, JSON writing)
│   │   ├── middleware/
│   │   │   ├── cors.go           # CORS middleware
│   │   │   ├── logging.go        # Logging middleware
│   │   │   └── auth.go           # Authentication middleware
│   │   └── routes.go             # Route definitions
│   ├── config/
│   │   └── config.go             # Configuration management
│   ├── domain/
│   │   ├── errors/
│   │   │   └── errors.go         # Custom error definitions
│   │   ├── models/
│   │   │   ├── expense.go        # Expense model
│   │   │   ├── income.go         # Income model
│   │   │   ├── summary.go        # Summary model
│   │   │   └── validation.go     # Model validation
│   │   └── services/
│   │       ├── expense.go        # Expense business logic
│   │       ├── income.go         # Income business logic
│   │       └── summary.go        # Summary business logic
│   ├── metrics/
│   │   └── metrics.go            # Prometheus metrics
│   └── storage/
│       ├── memory/
│       │   └── store.go          # In-memory storage implementation
│       └── storage.go            # Storage interface
├── pkg/
│   ├── logger/
│   │   └── logger.go             # Logger implementation
│   └── utils/
│       └── id.go                 # ID generation utility
├── configs/
│   └── config.yaml              # Configuration file
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.21 or later
- Docker (optional)

### Running Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/finance-tracker.git
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the server:
   ```bash
   go run cmd/server/main.go
   ```

### Using Docker

1. Build the image:
   ```bash
   docker build -t finance-tracker .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 finance-tracker
   ```

## API Endpoints

- `POST /api/v1/expenses` - Create expense
- `GET /api/v1/expenses` - List expenses
- `POST /api/v1/incomes` - Create income
- `GET /api/v1/incomes` - List incomes
- `GET /api/v1/summary` - Get financial summary

## Metrics

Prometheus metrics are available at `/metrics`

## Configuration

Configuration is managed through `configs/config.yaml`. Key configurations include:

- Server settings
- CORS configuration
- Logging settings
- Metrics configuration

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request