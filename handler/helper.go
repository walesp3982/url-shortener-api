package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type bodyError struct {
	Message string
}

func newBodyError(message string) bodyError {
	return bodyError{
		Message: message,
	}
}

func generateError(w http.ResponseWriter, status_code int, message string) {
	w.WriteHeader(status_code)

	err := json.NewEncoder(w).Encode(newBodyError(message))
	if err != nil {
		panic(err)
	}
}

func setHeaderToJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func isValidUrl(url string) bool {
	re := regexp.MustCompile(`^(http|https):\/\/*`)
	return re.MatchString(url)
}
