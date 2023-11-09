package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Urls struct {
	ID           uint   `gorm:"primaryKey"`
	ShortURLCode string `gorm:"unique" gorm:"index"`
	OriginalURL  string `gorm:"index"`
	BaseURL      string
	ExpiryDate   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func InitDB() {

	dbURL := "postgres://urlpro:TheStillRiver@1@localhost:5432/urlpro"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DB Connected in ORM Style")
		fmt.Println(db)
	}

	e := db.AutoMigrate(&Urls{})
	if e != nil {
		fmt.Println(e)
	} else {
		url := Urls{OriginalURL: "https://www.cyware.com2", ShortURLCode: "ABCDEF3", BaseURL: "https://www.baseurl.com"}

		result := db.Create(&url) // pass pointer of data to Create

		if result.Error != nil {
			fmt.Println(result.Error)
		} else {
			fmt.Println(result.RowsAffected)
		}

	}

	//return db

}
