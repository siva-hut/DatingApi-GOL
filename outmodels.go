package main
// Return Type for Matches API
type match struct {
	Status  string `json:"status"`
	Data    []like `json:"data"`
	Message string `json:"message"`
}
// Return Type for Getusers and find API
type finduser struct {
	Status  string `json:"status"`
	Data    []User `json:"data"`
	Message string `json:"message"`
}
