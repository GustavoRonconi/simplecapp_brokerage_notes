package constructors

import (
	"fmt"
	"log"
	"regexp"
	"simplecapp_brokerage_notes/brokerage_note/model"
	"simplecapp_brokerage_notes/utils"
	"strings"
	"time"
)

var mapOperationTypes = map[string]string{
	"C": "purchase",
	"V": "sale",
}

func searchOperation(searchString string, element string, operation string) bool {
	switch operation {
	case "=":
		if element == searchString {
			return true
		}
	case "in":
		if strings.Contains(element, searchString) {
			return true
		}
	}
	return false
}

func setNoteAttrs(note *model.BrokerageNote, splitedBody []string) {
	// var financialOperations []model.FinancialOperation
	var tickers []string
	var operationTypes []string
	var indexesSubtotal []int
	var unitaryValues []float64
	var amountValues []float64

	var tradingDate time.Time
	// var operationTypes []string

	for i1, e1 := range splitedBody {

		// cpf
		if searchOperation("C.P.F./C.N.P.J./C.V.M./C.O.B.", e1, "=") {
			cpfRegex := regexp.MustCompile(`[^\w]`)
			note.CustomerCPF = cpfRegex.ReplaceAllString(splitedBody[i1+1], "")
		}

		// number note
		if searchOperation("Nº Nota: ", e1, "in") {
			note.NumberNote = strings.Split(splitedBody[i1], ": ")[1]
		}

		// data operacoes
		if searchOperation("Data pregão: ", e1, "in") {
			tradingDate = utils.ParseDefaultDate(strings.Split(splitedBody[i1], ": ")[1])
		}

		// tickers & operations_types
		if searchOperation("SubTotal :", e1, "=") {
			tickers = append(tickers, splitedBody[i1-1])
			operationTypes = append(operationTypes, mapOperationTypes[strings.Split(splitedBody[i1-2], " ")[1]])

			e2 := splitedBody[i1-1]
			i2 := 1
			indexSubtotal := 1
			for e2 != "SubTotal :" && e2 != "C/V Tipo Mercado Especificação do Título" {
				if i2%2 == 0 {
					indexSubtotal += 1
				}
				i2 += 1
				e2 = splitedBody[i1-i2]
			}

			indexesSubtotal = append(indexesSubtotal, indexSubtotal)
		}

		// unitary values & amount
		if searchOperation("Preço Liquidação (R$)", e1, "=") || searchOperation("Quantidade", e1, "=") {
			referIndex := i1
			for _, indexSubtotal := range indexesSubtotal {
				referIndex = referIndex + indexSubtotal
				e3 := utils.ConvertToFloat(splitedBody[referIndex])

				if e1 == "Quantidade" {
					amountValues = append(amountValues, e3)
				} else {
					unitaryValues = append(unitaryValues, e3)
				}
			}
		}
	}

	if len(tickers) != len(operationTypes) {
		log.Fatal("Não foi possível processar a nota nº: ", note.NumberNote, " da corretora INTER")
	}

	// for i, item := range tickers {

	// }

	fmt.Println(tradingDate, tickers)
}

func NewInterNote(body string) *model.BrokerageNote {
	note := new(model.BrokerageNote)
	fmt.Println(body)
	splitedBody := strings.Split(body, "\n")

	setNoteAttrs(note, splitedBody)

	return note
}
