package main

import (
	// "encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	// "os"
	"strconv"
	"strings"
	// "time"
)

type HistoryEvent struct {
	EventTitle string `json:"event_title"`
	EventDesc  string `json:"event_desc"`
	Year       int    `json:"year"`
}

type Person struct {
	Name  string `json:"name"`
	About string `json:"about"`
	Year  int    `json:"year"`
}

type AllEvents struct {
	HistoryEvents []HistoryEvent `json:"history_events"`
	Birthdays     []Person       `json:"birthdays"`
	Deaths        []Person       `json:"deaths"`
}

func Scrape(month string, day int) AllEvents {
	c := colly.NewCollector(
		colly.AllowedDomains("www.timeanddate.com"),
	)

	historyEvents := []HistoryEvent{}
	birthdays := []Person{}
	deaths := []Person{}

	c.OnHTML("article.otd-row.otd-detail > ul", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			h3Text := el.ChildText("h3")
			h3TextWords := strings.Fields(h3Text)
			eventTitle := strings.Join(h3TextWords[1:], " ")

			yearStr := h3TextWords[0]
			yearInt, _ := strconv.Atoi(yearStr)

			eventDesc := el.ChildText("p")

			event := HistoryEvent{
				EventTitle: eventTitle,
				EventDesc:  eventDesc,
				Year:       yearInt,
			}
			historyEvents = append(historyEvents, event)
		})
	})

	c.OnHTML("article.otd-row.otd-life", func(e *colly.HTMLElement) {
		// This temp idx is being used to differentiate between
		// the two li tags under the same article tag container.
		idx := 0
		e.ForEach("ul", func(_ int, ul *colly.HTMLElement) {
			idx += 1
			ul.ForEach("li", func(_ int, li *colly.HTMLElement) {
				h3Text := li.ChildText("h3")
				h3TextWords := strings.Fields(h3Text)
				name := strings.Join(h3TextWords[1:], " ")

				yearStr := h3TextWords[0]
				yearInt, _ := strconv.Atoi(yearStr)

				about := li.ChildText("p")
				person := Person{
					Name:  name,
					About: about,
					Year:  yearInt,
				}
				if idx == 1 {
					birthdays = append(birthdays, person)
				} else {
					deaths = append(deaths, person)
				}
			})
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(err)
	})

	c.Visit("https://www.timeanddate.com/on-this-day/" + month + "/" + strconv.Itoa(day))

	fmt.Println(historyEvents)

	return AllEvents{
		HistoryEvents: historyEvents,
		Birthdays:     birthdays,
		Deaths:        deaths,
	}
}

func main() {
	// DATA_DIR := "../data"
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

	// Create the current data-source sub directory if it doesn't exist
	// currentYear := time.Now().Year()
	// yearDir := fmt.Sprintf("%s/%d", DATA_DIR, currentYear)
	// if _, err := os.Stat(yearDir); os.IsNotExist(err) {
	// 	os.Mkdir(yearDir, 0755)
	// }

	for month, days := range monthsAndDays {
		// monthDir := fmt.Sprintf("%s/%s", yearDir, month)
		// if _, err := os.Stat(monthDir); os.IsNotExist(err) {
		// 	os.Mkdir(monthDir, 0755)
		// }
		for i := 1; i <= days; i++ {
			Scrape(month, i)
			// eventDataJson, err := json.Marshal(eventData)

			// if err != nil {
			// 	fmt.Println("Error marshalling event data", err)
			// 	return
			// }

			// fmt.Println("Event data for", month, i, "is:")
			// fmt.Println(string(eventDataJson))

			// filePath := fmt.Sprintf("%s/%d.json", monthDir, i)
			// err = os.WriteFile(filePath, eventDataJson, 0644)

			// if err != nil {
			// 	fmt.Println("Error writing event data to file", err)
			// 	return
			// }

			// fmt.Printf("Successfully wrote event data to file %s\n", filePath)
		}
	}
}
