package model

import (
	"time"
)

type FinancialOperation struct {
	Ticker      string
	TradingDate time.Time
	OperationType string
}

type BrokerageNote struct {
	CustomerCPF         string
	NumberNote          string
	FinancialOperations FinancialOperation
}
