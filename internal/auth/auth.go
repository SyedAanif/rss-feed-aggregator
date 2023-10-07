package auth

import (
	"errors"
	"net/http"
	"strings"
)

/*
	Exported(capital) function that can be used as abstraction
	to extract API key from headers of HTTP request
	ex:
	Authorization: ApiKey {your-api-key}

	NOTE:: erros in go should be samle cased start(syntactic sugar)
*/
func GetAPIKey(headers http.Header) (string, error){
	val := headers.Get("Authorization") // Get Authorization header
	if val == ""{
		return "",errors.New("no authentication info found")
	}

	vals := strings.Split(val," ") // Split on space
	if len(vals) != 2{ // ApiKey, value
		return "", errors.New("malformed Authorization Header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed Authorization Header Key")
	}
	return vals[1], nil // API key
}
