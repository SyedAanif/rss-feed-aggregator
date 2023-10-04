package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
	Function to convert and send a payload via a JSON format
*/
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload) // Marsahl a payload to a byte JSON structure

	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500) // Return server error
		return
	}
	w.Header().Add("Content-Type","application/json") // Content-Type of JSON
	w.WriteHeader(code) // Success HTTP code
	w.Write(data) // Response body
}

/*
	Function to convert and send a Error message via a JSON format
*/
func respondWithError(w http.ResponseWriter, code int, msg string) {
	// Server side errors
	if code > 499 {
		log.Println("Responding with 5XX error:",msg)
	}

	// To keep error message structured
	/*
		{
			"error": "something wrong"
		}
	*/
	type errResponse struct{
		Error string `json:"error"` // Reflect tag(@Json Property) this to tell in JSON use property error
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}