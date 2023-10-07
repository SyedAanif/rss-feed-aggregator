package main

import (
	"fmt"
	"net/http"

	"github.com/SyedAanif/rss-feed-aggregator/internal/auth"
	"github.com/SyedAanif/rss-feed-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

/*
	This function takes in the authenticated handler with pointer tom DB config, and returns
	an HTTP handler that can be hooked on the CHI router. This is a way to prevent duplication of code.
	DRY(Don't Repeat Yourself)
*/
func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc{

	// Return an anonymous function of HTTp handler.
	// This is a CLOSURE --> access to types outside it's body
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Authorization Error: %v",err))
			return
		}
		
		// Context in GO has track of multiple routines running across. We can track, cancel etc for a context
		// using current context
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil{
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v",err))
			return
		}

		// Run the incoming handler that this function was presented with
		handler(w, r, user)
	}
}