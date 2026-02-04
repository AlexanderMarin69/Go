package controllers

import (
	"net/http"
	"strconv"
)

// writeResponse writes a JSON response to the http.ResponseWriter
func writeResponse(w http.ResponseWriter, statusCode int, body string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(body))
}

// writeSuccess writes a success response with a message
func writeSuccess(w http.ResponseWriter, statusCode int, message string) {
	writeResponse(w, statusCode, `{"message":"`+message+`"}`)
}

// writeSuccessWithCacheStatus writes a success response with cache status
func writeSuccessWithCacheStatus(w http.ResponseWriter, statusCode int, message string, cached bool) {
	cacheVal := "false"
	if cached {
		cacheVal = "true"
	}
	writeResponse(w, statusCode, `{"message":"`+message+`","cached":`+cacheVal+`}`)
}

// writeSuccessWithID writes a success response with a message and ID
func writeSuccessWithID(w http.ResponseWriter, statusCode int, message string, id string) {
	writeResponse(w, statusCode, `{"message":"`+message+`","id":"`+id+`"}`)
}

// writeSuccessWithIDAndCacheStatus writes a success response with message, ID, and cache status
func writeSuccessWithIDAndCacheStatus(w http.ResponseWriter, statusCode int, message string, id string, cached bool) {
	cacheVal := "false"
	if cached {
		cacheVal = "true"
	}
	writeResponse(w, statusCode, `{"message":"`+message+`","id":"`+id+`","cached":`+cacheVal+`}`)
}

// writeError writes an error response with a message and status code
func writeError(w http.ResponseWriter, statusCode int, message string) {
	writeResponse(w, statusCode, `{"error":"`+message+`","code":`+strconv.Itoa(statusCode)+`}`)
}
