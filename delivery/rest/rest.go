package rest

import (
	"encoding/json"
	"net/http"
)

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	Type string `json:"type"`
}

func StatusCode(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

func JSON(w http.ResponseWriter, payload interface{}, statusCode int) {
	StatusCode(w, statusCode)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(payload)
}

func Error(w http.ResponseWriter, payload interface{}, statusCode int) {
	JSON(w, payload, statusCode)
}
