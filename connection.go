package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)
// Database credentials
const (
	host     = "localhost"
	port     = 5432
	name     = "postgres"
	password = "siva"
	dbname   = "dating"
)
// DB is used to connect to database
var DB *gorm.DB

func con1() {
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
