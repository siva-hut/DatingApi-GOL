package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input getuser
	_ = json.NewDecoder(r.Body).Decode(&input)
	var person User
	var result = finduser{}
	DB.Where("Id = ?", input.Id).First(&person)
	if person == (User{}) {
		result.Status = "Failed"
		result.Message = "Enter a ID in database"
		json.NewEncoder(w).Encode(result)
	} else {
		result.Status = "Successful"
		var Users []User
		DB.Find(&Users)
		for _, x := range Users {
			if x.Id != input.Id {

				var dist = x.Location - person.Location
				if dist < 0 {
					dist = dist * -1
				}
				if dist <= input.Distance {
					result.Data = append(result.Data, x)
				}
			}
		}
		json.NewEncoder(w).Encode(result)
	}
}

func getmatches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	map1 := make(map[helper]int)
	var likes []like
	var result = match{}
	result.Status = "Successful"
	DB.Find(&likes)
	for _, x := range likes {
		var y1 = x.Who_likes
		var y2 = x.Who_is_liked
		var find1 = helper{
			Id1: uint(y1), Id2: uint(y2),
		}
		var find2 = helper{
			Id1: uint(y2), Id2: uint(y1),
		}
		num, ok := map1[find2]
		_ = num
		if ok {
			result.Data = append(result.Data, x)
		} else {
			map1[find1] = 1
		}
	}
	json.NewEncoder(w).Encode(result)
}

func findusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var s = params["q"]
	var Users []User
	var result = finduser{}
	result.Status = "Successful"
	DB.Find(&Users)
	for _, x := range Users {
		if strings.Contains(x.Name, s) {
			result.Data = append(result.Data, x)
		}
	}
	json.NewEncoder(w).Encode(result)
}
