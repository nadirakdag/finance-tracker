package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nadirakdag/finance-tracker/internal/domain/models"
	"github.com/nadirakdag/finance-tracker/pkg/utils"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewStore(dbPath string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := initializeTables(db); err != nil {
		return nil, err
	}

	return &SQLiteStore{db: db}, nil
}

func initializeTables(db *sql.DB) error {
	expenseTable := `CREATE TABLE IF NOT EXISTS expenses (
        id TEXT PRIMARY KEY,
        amount REAL NOT NULL,
        category TEXT NOT NULL,
        description TEXT NOT NULL,
        date DATE NOT NULL
    )`

	incomeTable := `CREATE TABLE IF NOT EXISTS incomes (
        id TEXT PRIMARY KEY,
        amount REAL NOT NULL,
        source TEXT NOT NULL,
        description TEXT NOT NULL,
        date DATE NOT NULL
    )`

	categoryTable := `CREATE TABLE IF NOT EXISTS categories (
    	id TEXT PRIMARY KEY,
    	name TEXT NOT NULL,
    	type TEXT NOT NULL CHECK (type IN ('income', 'expense'))
	)`

	categoryTableData := `INSERT OR IGNORE INTO categories (id, name, type) VALUES
		('salary', 'Salary', 'income'),
		('other_income', 'Other', 'income'),
		('food', 'Food & Dining', 'expense'),
		('transport', 'Transportation', 'expense'),
		('utilities', 'Utilities', 'expense'),
		('entertainment', 'Entertainment', 'expense'),
		('health', 'Healthcare', 'expense'),
		('vacation', 'Vacation', 'expense'),
		('tech', 'Tech', 'expense'),
		('subscription', 'Subscription', 'expense'),
		('other_expense', 'Other', 'expense');
	`

	queries := []string{expenseTable, incomeTable, categoryTable, categoryTableData}

	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return err
		}
	}

	return nil
}

func (s *SQLiteStore) CreateExpense(expense *models.Expense) error {
	query := `INSERT INTO expenses (id, amount, category, description, date)
              VALUES (?, ?, ?, ?, ?)`

	_, err := s.db.Exec(query,
		utils.GenerateID(16),
		expense.Amount,
		expense.Category,
		expense.Description,
		expense.Date)

	return err
}

func (s *SQLiteStore) GetExpenses() ([]models.Expense, error) {
	query := `SELECT id, amount, category, description, date FROM expenses`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		if err := rows.Scan(
			&expense.ID,
			&expense.Amount,
			&expense.Category,
			&expense.Description,
			&expense.Date); err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (s *SQLiteStore) CreateIncome(income *models.Income) error {
	query := `INSERT INTO incomes (id, amount, source, description, date)
              VALUES (?, ?, ?, ?, ?)`

	_, err := s.db.Exec(query,
		utils.GenerateID(16),
		income.Amount,
		income.Source,
		income.Description,
		income.Date)

	return err
}

func (s *SQLiteStore) GetIncomes() ([]models.Income, error) {
	query := `SELECT id, amount, source, description, date FROM incomes`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incomes []models.Income
	for rows.Next() {
		var income models.Income
		if err := rows.Scan(
			&income.ID,
			&income.Amount,
			&income.Source,
			&income.Description,
			&income.Date); err != nil {
			return nil, err
		}
		incomes = append(incomes, income)
	}

	return incomes, nil
}

func (s *SQLiteStore) GetCategories(categoryType string) ([]models.Category, error) {
	query := `SELECT id, name, type FROM categories WHERE type = ?`

	rows, err := s.db.Query(query, categoryType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Type); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *SQLiteStore) CheckCategory(c string, cType string) error {
	query := `SELECT id, name, type FROM categories WHERE type = ? and id = ?`

	rows, err := s.db.Query(query, cType, c)
	if err != nil {
		return err
	}
	defer rows.Close()

	var category models.Category
	for rows.Next() {
		err = rows.Scan(&category.ID, &category.Name, &category.Type)
	}

	return err
}

func (s *SQLiteStore) Close() error {
	return s.db.Close()
}
