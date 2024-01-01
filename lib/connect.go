package lib

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	fmt.Println("Connecting to DB...")

	db, err := gorm.Open(sqlite.Open("on-this-day.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	// Migrate the model schema
	db.AutoMigrate(&Year{}, &Month{}, &Day{}, &Event{})

	fmt.Println("Connected to DB!")

	return db
}
