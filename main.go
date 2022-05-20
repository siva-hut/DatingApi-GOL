package main

import (
	_ "github.com/lib/pq"
)

// Main function
func main() {
	//Create database connection
	con1()
	//populate the database with data from like and users JSON
	populate()
	//Start server and route the API's
	route()
}
