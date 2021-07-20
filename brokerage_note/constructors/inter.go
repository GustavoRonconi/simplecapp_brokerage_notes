package constructors

import (
	"fmt"
	"log"
	"regexp"
	"simplecapp_brokerage_notes/brokerage_note/model"
	"simplecapp_brokerage_notes/utils"
	"strings"
)

func NewInterNote(body string) *model.BrokerageNote {
	note := new(model.BrokerageNote)
	fmt.Println(body)
	splitedBody := strings.Split(body, "\n")
	cpfRegex, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}

	note.NumberNote = strings.Split(splitedBody[9], ": ")[1]
	note.TradingDate = utils.ParseDefaultDate(strings.Split(splitedBody[5], ": ")[1])
	note.CustomerCPF = cpfRegex.ReplaceAllString(splitedBody[45], "")

	return note
}
