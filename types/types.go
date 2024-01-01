package types

// Description: This file contains the types used in the application
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
