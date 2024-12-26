package models

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"` // "income" or "expense"
}
