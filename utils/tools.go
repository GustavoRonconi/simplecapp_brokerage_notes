package utils

import (
	"code.sajari.com/docconv"
	"log"
	"strconv"
	"strings"
	"time"
)

// Find if val exist in array
func FindSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Convert pdf to string
func PdfToString() string {
	res, err := docconv.ConvertPath("nota2.pdf")
	if err != nil {
		log.Fatal(err)
	}

	return res.Body
}

// Function to parse simple brasilian date
func ParseDefaultDate(defaultDate string) time.Time {
	defaultDateTime, err := time.Parse("02/01/2006", defaultDate)
	if err != nil {
		log.Fatal(err)
	}

	return defaultDateTime

}

// Convert string to float
func ConvertToFloat(floatString string) float64 {
	floatString = strings.ReplaceAll(floatString, ".", "")
	floatString = strings.ReplaceAll(floatString, ",", ".")

	floatValue, err := strconv.ParseFloat(floatString, 64)
	if err != nil {
		log.Fatal(err)
	}
	return floatValue
}
