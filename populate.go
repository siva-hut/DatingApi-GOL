package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func populate() {

	content, err := ioutil.ReadFile("users.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	var userdata []User
	json.Unmarshal(content, &userdata)
	for _, x := range userdata {
		DB.Create(&x)
	}
	c, err := ioutil.ReadFile("likes.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	var data []like
	json.Unmarshal(c, &data)
	for _, x := range data {
		DB.Create(&x)
	}
}
