# Finance Tracker Frontend

Vue.js frontend for personal finance tracking application.

## Tech Stack

- Vue 3
- Composition API
- Tailwind CSS

## Project Structure
```
.
├── Dockerfile
├── README.md
├── env.d.ts
├── eslint.config.js
├── index.html
├── nginx.conf
├── package-lock.json
├── package.json
├── postcss.config.js
├── public
│   └── favicon.ico
├── src
│   ├── App.vue
│   ├── assets
│   │   ├── logo.svg
│   │   └── main.css
│   ├── components
│   │   ├── common
│   │   │   └── LoadingSpinner.vue
│   │   └── dashboard
│   │       ├── CategoryBreakdown.vue
│   │       ├── ExpenseForm.vue
│   │       ├── IncomeForm.vue
│   │       ├── MonthlyOverview.vue
│   │       ├── SummaryCards.vue
│   │       └── TransactionForm.vue
│   ├── main.ts
│   ├── router
│   │   └── index.ts
│   ├── services
│   │   └── api.ts
│   ├── stores
│   │   └── finance.ts
│   ├── types
│   │   ├── forms.ts
│   │   └── transaction.ts
│   └── views
│       └── DashboardView.vue
├── tailwind.config.js
├── tsconfig.app.json
├── tsconfig.json
├── tsconfig.node.json
└── vite.config.ts
```

## Setup

1. Install dependencies:
```bash
npm install
```

2. Run development server:
```bash
npm run dev
```

3. Build for production:
```bash
npm run build
```

## Features

### Dashboard Layout
- Summary cards showing total income, expenses, and balance
- Forms for adding income and expenses
- Monthly overview table
- Category breakdown table

### Transaction Forms
- Income form with source selection
- Expense form with category selection
- Amount validation
- Date selection
- Dynamic category loading from API

### Data Display
- Color-coded transactions (green for income, red for expenses)
- Sortable tables
- Loading states
- Error handling

## Docker Support

Build image:
```bash
docker build -t finance-tracker-frontend .
```

Run container:
```bash
docker run -p 80:80 finance-tracker-frontend
```

## API Integration

The frontend integrates with the backend API at `http://localhost:8080/api/v1/`:

### Endpoints Used:
- `/categories` - Get transaction categories
- `/expenses` - Manage expenses
- `/incomes` - Manage incomes
- `/summary` - Get financial overview

## Environment Variables

```env
VITE_API_URL=http://localhost:8080
```

## Browser Support

- Chrome (latest)
- Firefox (latest)
- Safari (latest)
- Edge (latest)
