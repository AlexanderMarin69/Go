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

// writeSuccessWithRequest writes a success response with message and cache status injected by middleware
func writeSuccessWithRequest(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	// The caching middleware will inject cache status automatically
	writeResponse(w, statusCode, `{"message":"`+message+`"}`)
}

// writeSuccessWithID writes a success response with a message and ID
func writeSuccessWithID(w http.ResponseWriter, statusCode int, message string, id string) {
	writeResponse(w, statusCode, `{"message":"`+message+`","id":"`+id+`"}`)
}

// writeSuccessWithIDAndRequest writes a success response with message, ID, and cache status
func writeSuccessWithIDAndRequest(w http.ResponseWriter, r *http.Request, statusCode int, message string, id string) {
	// The caching middleware will inject cache status automatically
	writeResponse(w, statusCode, `{"message":"`+message+`","id":"`+id+`"}`)
}

// writeError writes an error response with a message and status code
func writeError(w http.ResponseWriter, statusCode int, message string) {
	writeResponse(w, statusCode, `{"error":"`+message+`","code":`+strconv.Itoa(statusCode)+`}`)
}
