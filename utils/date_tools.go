package utils

import (
	"log"
	"time"
)

// Function to parse simple brasilian date
func ParseDefaultDate(defaultDate string) time.Time {
	defaultDateTime, err := time.Parse("02/01/2006", defaultDate)
	if err != nil {
		log.Fatal(err)
	}

	return defaultDateTime

}
