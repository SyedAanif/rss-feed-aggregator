package main

import "net/http"

/*
	HTTP Handler in the format of GO
*/
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}