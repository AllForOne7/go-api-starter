package main

// Message defines the core data model and its DB structure.
type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// Response is a standard wrapper for all JSON API responses (both success and error).
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
