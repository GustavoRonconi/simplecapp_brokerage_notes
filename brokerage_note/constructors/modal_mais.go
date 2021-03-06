package constructors

import (
	"fmt"
	"simplecapp_brokerage_notes/brokerage_note/model"
	"strings"
)

func NewModalMaisNote(body string, fileName string) *model.BrokerageNote {
	note := new(model.BrokerageNote)
	splitedBody := strings.Split(body, "\n\n")

	note.FileName = fileName
	note.NumberNote = splitedBody[2]
	note.CustomerCPF = splitedBody[8]

	fmt.Println(splitedBody)
	return note
}
