package model

import (
	"time"
)

type BrokerageNote struct {
	CustomerCPF string
	NumberNote  string
	TradingDate time.Time
}
