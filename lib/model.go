package lib

import (
	"gorm.io/gorm"
)

type Year struct {
	gorm.Model
	Year int
}

type Month struct {
	gorm.Model
	Month     int
	MonthName string
}

type Day struct {
	gorm.Model
	Day     int
	DayName string
}

type EventType string

const (
	History EventType = "history"
	Birth   EventType = "birth"
	Death   EventType = "death"
)

type Event struct {
	gorm.Model
	Type        EventType `gorm:"type:ENUM('history', 'birth', 'death')"`
	Title       string    `gorm:"unique"`
	Description string    `gorm:"unique"`
	YearID      uint
	MonthID     uint
	DayID       uint
}
