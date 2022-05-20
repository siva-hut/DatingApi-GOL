package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func populate() {
	//Read JSON file
	content, err := ioutil.ReadFile("users.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	var userdata []User
	json.Unmarshal(content, &userdata)
	for _, x := range userdata {
		//Insert User into DB
		DB.Create(&x)
	}
	c, err := ioutil.ReadFile("likes.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	var data []like
	json.Unmarshal(c, &data)
	for _, x := range data {
		//Insert Like into DB
		DB.Create(&x)
	}
}
