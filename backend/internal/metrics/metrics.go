package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ExpenseTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "finance_tracker_expenses_total",
		Help: "The total number of expenses recorded",
	})

	IncomeTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "finance_tracker_incomes_total",
		Help: "The total number of incomes recorded",
	})

	TransactionDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "finance_tracker_transaction_duration_seconds",
		Help:    "Time taken to process transactions",
		Buckets: prometheus.DefBuckets,
	})

	ActiveUsers = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "finance_tracker_active_users",
		Help: "The current number of active users",
	})
)
