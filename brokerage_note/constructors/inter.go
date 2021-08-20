package constructors

import (
	"regexp"
	"simplecapp_brokerage_notes/brokerage_note/model"
	"simplecapp_brokerage_notes/utils"
	"strings"
	"time"
)

var Corretora string = "INTER"

func NewInterNote(body string, fileName string) *model.BrokerageNote {
	note := new(model.BrokerageNote)
	note.FileName = fileName
	splitedBody := strings.Split(body, "\n")
	var tradingDate time.Time

	//Financial Operations Primitive
	financialOperationPrimitive := new(model.FinancialOperationPrimitive)

	
	//Aux vars
	var indexesSubtotal []int

	for i1, e1 := range splitedBody {

		// cpf
		if utils.CheckOperation("C.P.F./C.N.P.J./C.V.M./C.O.B.", e1, "=") {
			cpfRegex := regexp.MustCompile(`[^\w]`)
			note.CustomerCPF = cpfRegex.ReplaceAllString(splitedBody[i1+1], "")
		}

		// number note
		if utils.CheckOperation("Nº Nota: ", e1, "in") {
			note.NumberNote = strings.Split(splitedBody[i1], ": ")[1]
		}

		// operations date
		if utils.CheckOperation("Data pregão: ", e1, "in") {
			tradingDate = utils.ParseDefaultDate(strings.Split(splitedBody[i1], ": ")[1])
		}

		// tickers & operations_types
		if utils.CheckOperation("SubTotal :", e1, "=") {
			// tickers = append(tickers, splitedBody[i1-1])
			financialOperationPrimitive.Tickers = append(financialOperationPrimitive.Tickers, splitedBody[i1-1])
			// operationTypes = append(operationTypes, splitedBody[i1-2])
			financialOperationPrimitive.OperationTypes = append(financialOperationPrimitive.OperationTypes, splitedBody[i1-2])

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
		if utils.CheckOperation("Preço Liquidação (R$)", e1, "=") || utils.CheckOperation("Quantidade", e1, "=") {
			referIndex := i1
			for _, indexSubtotal := range indexesSubtotal {
				referIndex = referIndex + indexSubtotal
				e3 := utils.ConvertToFloat(splitedBody[referIndex])

				if e1 == "Quantidade" {
					// amountValues = append(amountValues, e3)
					financialOperationPrimitive.AmountValues = append(financialOperationPrimitive.AmountValues, e3)
				} else {
					// unitaryValues = append(unitaryValues, e3)
					financialOperationPrimitive.UnitaryValues = append(financialOperationPrimitive.UnitaryValues, e3)
				}
			}
		}
	}
	note.SetFinancialOperations(tradingDate, fileName, financialOperationPrimitive)

	return note
}
