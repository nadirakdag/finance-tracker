# Finance Tracker

A modern finance tracking application with Go backend and Vue.js frontend for managing personal income and expenses.

## Features

- Track income and expenses
- Categorize transactions
- View monthly summaries
- Category-based reporting
- SQLite database storage
- RESTful API
- Responsive design

## Tech Stack

**Backend:**
- Go
- SQLite
- Gorilla Mux
- Prometheus metrics

**Frontend:**
- Vue.js 3
- Tailwind CSS
- Composition API

## Getting Started

### Prerequisites
- Go 1.21+
- Node.js 18+
- SQLite 3

### Backend Setup
```bash
cd backend
go mod download
go run cmd/server/main.go
```

### Frontend Setup
```bash
cd frontend
npm install
npm run dev
```

### Docker Setup
```bash
docker-compose up --build
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

### Example Request
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

## Project Structure

```
.
├── backend/
│   ├── cmd/
│   │   └── server/
│   ├── internal/
│   │   ├── api/
│   │   ├── domain/
│   │   └── storage/
│   └── pkg/
└── frontend/
    ├── src/
    │   ├── components/
    │   ├── services/
    │   └── types/
    └── public/
```

## Development

### Running Tests
```bash
# Backend tests
cd backend
go test ./...

# Frontend tests
cd frontend
npm run test
```

### Database Migrations
Database schema is automatically created on first run.

## Docker Support

Build and run with Docker Compose:
```bash
docker-compose up --build
```

Accessing services:
- Frontend: http://localhost:80
- Backend: http://localhost:8080
- API Documentation: http://localhost:8080/api/v1
