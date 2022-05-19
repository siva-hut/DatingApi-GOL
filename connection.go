package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	name     = "postgres"
	password = "siva"
	dbname   = "dating"
)

type User struct {
	Id       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`
	Location float32 `json:"location"`
	Gender   string  `json:"gender"`
	Email    string  `json:"email"`
}
type like struct {
	Id           uint `json:"id" gorm:"primaryKey"`
	Who_likes    int  `json:"who_likes"`
	Who_is_liked int  `json:"who_is_liked"`
}

var DB *gorm.DB

func con() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, name, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	DB = db
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(User{})
	db.AutoMigrate(like{})
}
