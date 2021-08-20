package model

import (
	"gopkg.in/validator.v2"
	"log"
	"reflect"
	"strings"
	"time"
)

// Structs
type FinancialOperationPrimitive struct {
	Tickers        []string  `validate:"min=1"`
	OperationTypes []string  `validate:"min=1"`
	UnitaryValues  []float64 `validate:"min=1"`
	AmountValues   []float64 `validate:"min=1"`
}

type FinancialOperation struct {
	Ticker        string    `validate:"nonzero"`
	TradingDate   time.Time `validate:"nonzero"`
	OperationType string    `validate:"nonzero"`
	UnitaryValues float64   `validate:"nonzero"`
	AmountValues  float64   `validate:"nonzero"`
}

type BrokerageNote struct {
	FileName            string               `validate:"nonzero"`
	CustomerCPF         string               `validate:"nonzero"`
	NumberNote          string               `validate:"nonzero"`
	FinancialOperations []FinancialOperation `validate:"min=1"`
}

// Private Methods
func (financialOperation *FinancialOperation) setOperationType(operationTypeDefault string) {
	var mapOperationTypes = map[string]string{
		"C": "purchase",
		"V": "sale",
	}
	financialOperation.OperationType = mapOperationTypes[strings.Split(operationTypeDefault, " ")[1]]

}

func (financialOperationPrimitive FinancialOperationPrimitive) validFinancialOperationPrimitive(fileName string) {
	v := reflect.ValueOf(financialOperationPrimitive)

	numberOperations := 0
	for i := 0; i < v.NumField(); i++ {
		errs := validator.Validate(financialOperationPrimitive)
		if (i > 0 && v.Field(i).Len() != numberOperations) || (errs != nil) {
			log.Fatal("Error to parse financial operations of brokerage note ", fileName)
		}
		numberOperations = v.Field(i).Len()
	}
}

// External Methods
func (note *BrokerageNote) SetFinancialOperations(tradingDate time.Time, fileName string, financialOperationPrimitive *FinancialOperationPrimitive) {
	note.FinancialOperations = []FinancialOperation{}
	financialOperationPrimitive.validFinancialOperationPrimitive(fileName)

	for i := range financialOperationPrimitive.Tickers {
		financialOperation := FinancialOperation{
			Ticker:        strings.Split(financialOperationPrimitive.Tickers[i], " ")[0],
			TradingDate:   tradingDate,
			UnitaryValues: financialOperationPrimitive.UnitaryValues[i],
			AmountValues:  financialOperationPrimitive.AmountValues[i],
		}
		financialOperation.setOperationType(financialOperationPrimitive.OperationTypes[i])
		if errs := validator.Validate(financialOperation); errs != nil {
			log.Fatal("Error to parse financial operations of brokerage note", fileName)
		}
		note.FinancialOperations = append(note.FinancialOperations, financialOperation)
	}

	if errs := validator.Validate(note); errs != nil {
		log.Fatal("Error to parse financial operations of brokerage note: ", fileName)
	}
}
