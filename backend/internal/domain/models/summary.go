package models

type Summary struct {
	TotalIncome      float64              `json:"totalIncome"`
	TotalExpenses    float64              `json:"totalExpenses"`
	RemainingAmount  float64              `json:"remainingAmount"`
	CategoryTotals   map[string]float64   `json:"categoryTotals"`
	MonthlyBreakdown map[string]MonthData `json:"monthlyBreakdown"`
}

type MonthData struct {
	Income   float64 `json:"income"`
	Expenses float64 `json:"expenses"`
	Balance  float64 `json:"balance"`
}
