package utils

import (
	"code.sajari.com/docconv"
	"log"
)

func PdfToString() string {
	res, err := docconv.ConvertPath("nota2.pdf")
	if err != nil {
		log.Fatal(err)
	}

	return res.Body
}
