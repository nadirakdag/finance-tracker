package models

type Summary struct {
	TotalIncome       float64                 `json:"totalIncome"`
	TotalExpenses     float64                 `json:"totalExpenses"`
	RemainingAmount   float64                 `json:"remainingAmount"`
	MonthlyBreakdown  map[string]MonthData    `json:"monthlyBreakdown"`
	CategoryBreakdown map[string]CategoryData `json:"categoryBreakdown"`
}

type MonthData struct {
	Income   float64 `json:"income"`
	Expenses float64 `json:"expenses"`
	Balance  float64 `json:"balance"`
}

type CategoryData struct {
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
}
