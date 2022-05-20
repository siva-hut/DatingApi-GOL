package main

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func con2() {

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	DB = db
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(User{})
	db.AutoMigrate(like{})
}
