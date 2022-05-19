package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func getmatches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	map1 := make(map[int]string)
	var likes []like
	DB.Find(&likes)
	for _, x := range likes {
		var y1 = x.Who_likes
		var y2 = x.Who_is_liked
		num, ok := map1[y1+y2]
		_ = num
		if ok {
			json.NewEncoder(w).Encode(x)
		} else {
			map1[y1+y2] = "present"
		}
	}

}

type getuser struct {
	Id       uint    `json:"id"`
	Location float32 `json:"location"`
}

func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input getuser
	_ = json.NewDecoder(r.Body).Decode(&input)
	fmt.Println(input.Id)
	var person User
	DB.Where("Id = ?", input.Id).First(&person)
	if person == (User{}) {
		json.NewEncoder(w).Encode("punda")
	} else {
		var Users []User
		DB.Find(&Users)
		for _, x := range Users {
			if x.Id != input.Id {

				var dist = x.Location - person.Location
				if dist < 0 {
					dist = dist * -1
				}
				if dist <= input.Location {
					json.NewEncoder(w).Encode(x)
				}
			}
		}
	}
}

func findusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var s = params["q"]
	var Users []User
	DB.Find(&Users)
	for _, x := range Users {
		if strings.Contains(x.Name, s) {
			json.NewEncoder(w).Encode(x)
		}
	}
}
func route() {
	r := mux.NewRouter()
	r.HandleFunc("/matches", getmatches).Methods("GET")
	r.HandleFunc("/getusers", getusers).Methods("POST")
	r.HandleFunc("/find/{q}", findusers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
