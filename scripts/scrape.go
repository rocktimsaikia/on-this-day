package scripts

import (
  "fmt"
  "github.com/gocolly/colly"
  "encoding/json"
  "strings"
  "strconv"
)

type HistoryEvent struct {
  EventTitle string `json:"event_title"`
  EventDesc string `json:"event_desc"`
  Year int `json:"year"`
}

type Person struct {
  Name string `json:"name"`
  About string `json:"about"`
  Year int `json:"year"`
}

type AllEvents struct {
  HistoryEvents []HistoryEvent `json:"history_events"`
  Birthdays []Person `json:"birthdays"`
  Deaths []Person `json:"deaths"`
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
	EventDesc: eventDesc,
	Year: yearInt,
      }
      historyEvents = append(historyEvents, event)
    })
  })

  c.OnHTML("article.otd-row.otd-life", func(e *colly.HTMLElement) {
    // This temp idx is being used to differentiate between
    // the two li tags under the same article tag container.
    idx := 0
    e.ForEach("ul", func(_ int, ul *colly.HTMLElement){
      idx += 1
      ul.ForEach("li", func(_ int, li *colly.HTMLElement) {
	h3Text := li.ChildText("h3")
	h3TextWords := strings.Fields(h3Text)
	name := strings.Join(h3TextWords[1:], " ")

	yearStr := h3TextWords[0]
	yearInt, _ := strconv.Atoi(yearStr)

	about := li.ChildText("p")
	person := Person{
	  Name: name,
	  About: about,
	  Year: yearInt,
	}
	if idx == 1 {
	  birthdays = append(birthdays, person)
	} else  {
	  deaths = append(deaths, person)
	}
      })
    })
  })

  c.OnRequest(func(r *colly.Request) {
    fmt.Println("Visiting", r.URL.String())
  })

  c.OnError(func(r *colly.Response, err error){
    fmt.Println(err)
  })

  c.OnScraped(func(r *colly.Response) {
    fmt.Println("Finished", r.Request.URL)
    historyEventsJson, _ := json.MarshalIndent(historyEvents, "", "  ")
    birtdaysJson, _ := json.MarshalIndent(birthdays, "", "  ")
    deathsJson, _ := json.MarshalIndent(deaths, "", "  ")

    fmt.Println(string(historyEventsJson))
    fmt.Println(string(birtdaysJson))
    fmt.Println(string(deathsJson))
  })

  c.Visit("https://www.timeanddate.com/on-this-day/" + month + "/" + strconv.Itoa(day))

  return AllEvents{
    HistoryEvents: historyEvents,
    Birthdays: birthdays,
    Deaths: deaths,
  }
}
