package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func route() {
	r := mux.NewRouter()
	r.HandleFunc("/matches", getmatches).Methods("GET")
	r.HandleFunc("/getusers", getusers).Methods("POST")
	r.HandleFunc("/find/{q}", findusers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
