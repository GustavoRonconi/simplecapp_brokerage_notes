package main

import (
	"fmt"
	"simplecapp_brokerage_notes/brokerage_note"
)

func main() {
	note := brokerage_note.GetBrokerageNoteInfo("nota2.pdf")
	fmt.Println(note)

}
