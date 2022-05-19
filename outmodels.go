package main

type match struct {
	Status  string `json:"status"`
	Data    []like `json:"data"`
	Message string `json:"message"`
}

type finduser struct {
	Status  string `json:"status"`
	Data    []User `json:"data"`
	Message string `json:"message"`
}
