package apiresponse

import (
	"encoding/json"
	"net/http"

	"ritsrnjn/dummy-cc/constants"
)

// ApiResponse is the common response object for all API calls
type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// send 200 OK response
func SendOK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ApiResponse{
		Success: true,
		Data:    data,
	})
}

// send 400 Bad Request response
func SendBadRequest(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	if message == constants.EmptyString {
		message = "Bad Request"
	}

	json.NewEncoder(w).Encode(ApiResponse{
		Success: false,
		Message: message,
	})
}

// send 401 Unauthorized response
func SendUnauthorized(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	if message == constants.EmptyString {
		message = "Unauthorized"
	}
	json.NewEncoder(w).Encode(ApiResponse{
		Success: false,
		Message: message,
	})
}

// send 500 Internal Server Error response
func SendInternalServerError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	if message == constants.EmptyString {
		message = "Internal Server Error"
	}
	json.NewEncoder(w).Encode(ApiResponse{
		Success: false,
		Message: message,
	})
}

// send 404 Not Found response
func SendNotFound(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	if message == constants.EmptyString {
		message = "Not Found"
	}
	json.NewEncoder(w).Encode(ApiResponse{
		Success: false,
		Message: message,
	})
}
