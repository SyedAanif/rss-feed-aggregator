package main

import "net/http"

/*
	Handler for errors
*/
func handleError(w http.ResponseWriter, r *http.Request){
	respondWithError(w,400,"Something went wrong")
}