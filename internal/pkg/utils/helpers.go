package utils

import (
	"fmt"
	"time"
)

func ConvertStringToTime(dateString string) (*time.Time, error) {
	layout := "02/01/2006"                      // Adjust the layout to match "day/month/year"
	date, err := time.Parse(layout, dateString) // Parse the string into a time.Time object.
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil, err
	}
	return &date, nil
}
