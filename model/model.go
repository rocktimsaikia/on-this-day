package model

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

type Type string

const (
	History Type = "history"
	Birth   Type = "birth"
	Death   Type = "death"
)

type Event struct {
	gorm.Model
	Type        Type `gorm:"type:enum('history', 'birth', 'death')"`
	Title       string
	Description string
	YearID      uint
	MonthID     uint
	DayID       uint
}
