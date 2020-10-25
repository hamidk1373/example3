package helpers

import (
	"encoding/json"
	"net/http"
)

// JSONResponse the type for responses
type JSONResponse map[string]interface{}

// SendError returns the error using responseWriter
func SendError(w http.ResponseWriter, code int, textMessage string, content interface{}) {
	responseData := JSONResponse{"message": textMessage, "content": content}
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		SendError(w, http.StatusInternalServerError, "internal server Error", nil)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonData)
}

// SendSuccessResponse returns correct response using responseWriter
func SendSuccessResponse(w http.ResponseWriter, code int, content JSONResponse) {
	responseData := JSONResponse{"data": content}
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		SendError(w, http.StatusInternalServerError, "internal server Error", nil)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonData)
}
