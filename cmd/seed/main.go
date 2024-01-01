package main

import (
	"encoding/json"
	"fmt"
	// "on-this-day/lib"
	"on-this-day/types"
	"os"
)

func main() {
	// db := lib.ConnectToDB()

	DATA_DIR := "data/2023"
	monthsAndDays := map[string]int{
		"January":   31,
		"February":  28,
		"March":     31,
		"April":     30,
		"May":       31,
		"June":      30,
		"July":      31,
		"August":    31,
		"September": 30,
		"October":   31,
		"November":  30,
		"December":  31,
	}

	for month, days := range monthsAndDays {
		for i := 1; i <= days; i++ {
			jsonPath := fmt.Sprintf("%s/%s/%d.json", DATA_DIR, month, i)
			fmt.Println("jsonPath:", jsonPath)
			jsonFile, err := os.Open(jsonPath)

			if err != nil {
				fmt.Println(err)
				return
			}

			defer jsonFile.Close()

			byteValue, err := os.ReadFile(jsonPath)
			if err != nil {
				fmt.Println(err)
				return
			}

			var allEvents types.AllEvents
			err = json.Unmarshal(byteValue, &allEvents)
			if err != nil {
				return
			}

			historyEvents := allEvents.HistoryEvents
			birthdays := allEvents.Birthdays
			deaths := allEvents.Deaths

			fmt.Println("history:", len(historyEvents))
			fmt.Println("births:", len(birthdays))
			fmt.Println("deaths:", len(deaths))

			for _, event := range historyEvents {
				eventTitle := event.EventTitle
				eventDesc := event.EventDesc
				eventYear := event.Year
				eventMonth := month
				eventDay := i
				fmt.Println("eventTitle:", eventTitle)
				fmt.Println("eventDesc:", eventDesc)
				fmt.Println("eventYear:", eventYear)
				fmt.Println("eventMonth:", eventMonth)
				fmt.Println("eventDay:", eventDay)
			}
		}
	}
}
