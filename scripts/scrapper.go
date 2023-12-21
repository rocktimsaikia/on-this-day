package scripts

import (
    "encoding/json"
    "fmt"
    "os"
    "time"
)

func main() {
  DATA_DIR := "data"
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

  // Create the current data-source sub directory if it doesn't exist
  currentYear := time.Now().Year()
  yearDir := fmt.Sprintf("%s/%d", DATA_DIR, currentYear)
  if _, err := os.Stat(yearDir); os.IsNotExist(err) {
    os.Mkdir(yearDir, 0755)
  }

  for month, days := range monthsAndDays {
    monthDir := fmt.Sprintf("%s/%s", yearDir, month)
    if _, err := os.Stat(monthDir); os.IsNotExist(err) {
      os.Mkdir(monthDir, 0755)
    }
    for i:=1; i<=days; i++ {
      eventData := Scrape(month, i)
      eventDataJson, err := json.Marshal(eventData)

      if err != nil {
	fmt.Println("Error marshalling event data", err)
	return
      }

      filePath := fmt.Sprintf("%s/%d.json", monthDir, i)
      err = os.WriteFile(filePath, eventDataJson, 0644)

      if err != nil {
	fmt.Println("Error writing event data to file", err)
	return
      }

      fmt.Printf("Successfully wrote event data to file %s\n", filePath)
    }
  }
}
