package brokerage_note

import (
	"simplecapp_brokerage_notes/brokerage_note/constructors"
	"simplecapp_brokerage_notes/brokerage_note/model"
	"simplecapp_brokerage_notes/utils"
	"strings"
)

func constructorByBroker(body string, fileName string) *model.BrokerageNote {
	m := map[bool]func(string, string) *model.BrokerageNote{
		strings.Contains(body, "MODAL DTVM LTDA"): constructors.NewModalMaisNote,
		strings.Contains(body, "Inter DTVM Ltda."): constructors.NewInterNote,
	}
	for k, v := range m {
		if k {
			return v(body, fileName)
		}
	}
	return nil
}

func GetBrokerageNoteInfo(fileName string) *model.BrokerageNote {
	body := utils.PdfToString(fileName)
	return constructorByBroker(body, fileName)
}
