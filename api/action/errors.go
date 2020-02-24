package action

import (
	"encoding/json"
	"net/http"
)

var (
	ErrInternalServer = Error{statusCode: http.StatusInternalServerError, Message: "Internal server Error"}
)

type Error struct {
	statusCode int
	Message    string `json:"message,omitempty"`
}

func (e Error) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.statusCode)
	return json.NewEncoder(w).Encode(e)
}