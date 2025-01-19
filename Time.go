package main

import (
	"fmt"
	"time"
)

func Time() {
	currentTime := time.Now()
	formattedTime_24Hr := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formattedTime_24Hr)

	// Use the "3:04 PM" format for a 12-hour clock with AM/PM
	formattedTime_12Hr := currentTime.Format("2006-01-02 3:04 PM")
	fmt.Println("Formatted Time:", formattedTime_12Hr)

	/* The time.Parse function in Go is used to parse a formatted string representing a time
	and convert it into a time.Time value. */
	dateStr := "2023-11-25 12:04:15"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", dateStr)

	if err == nil {
		fmt.Println("Parsed Time:", parsedTime)
	} else {
		fmt.Println("Error parsing time:", err)
	}

	// We can also add or subtract Time
	newTime := currentTime.Add(24 * time.Hour)
	fmt.Println("current time :", currentTime)
	fmt.Println("New Time (after adding 1 day) :", newTime)

	timeWithDay := currentTime.Format("02.01.2006 Monday")
	fmt.Println("current time (with day format) :", timeWithDay)

}
