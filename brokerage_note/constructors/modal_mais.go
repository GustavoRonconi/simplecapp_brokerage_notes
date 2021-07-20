package constructors

import (
	"fmt"
	"simplecapp_brokerage_notes/brokerage_note/model"
	"simplecapp_brokerage_notes/utils"
	"strings"
)

func NewModalMaisNote(body string) *model.BrokerageNote {
	note := new(model.BrokerageNote)
	splitedBody := strings.Split(body, "\n\n")

	note.NumberNote = splitedBody[2]
	note.TradingDate = utils.ParseDefaultDate(splitedBody[5])
	note.CustomerCPF = splitedBody[8]

	fmt.Println(splitedBody)
	return note
}
