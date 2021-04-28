package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

func StatusCode(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

func JSON(w http.ResponseWriter, payload interface{}, statusCode int) {
	StatusCode(w, statusCode)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Panicln("Error encoding http return:", err.Error())
	}
}

func Error(w http.ResponseWriter, payload interface{}, statusCode int) {
	JSON(w, ResponseError{payload}, statusCode)
}

type ResponseError struct {
	Message interface{} `json:"message"`
}
