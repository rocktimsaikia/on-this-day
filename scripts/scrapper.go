package scripts

import (
    "encoding/json"
    "fmt"
)

func Main() {
  monthsAndDays := map[string]int{
    "January": 31,
    "February": 28,
    "March": 31,
    "April": 30,
    "May": 31,
    "June": 30,
    "July": 31,
    "August": 31,
    "September": 30,
    "October": 31,
    "November": 30,
    "December": 31,
  }
  for month, days := range monthsAndDays {
    for i:=1; i<=days; i++ {
      event_data := Scrape(month, i)
      fmt.Println(json.MarshalIndent(event_data, "", "  "))
      // TODO: Process and write the data to files in the `data` directory
    }
  }
}
