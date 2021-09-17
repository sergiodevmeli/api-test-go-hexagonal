package utils

import (
	"encoding/json"
	"net/http"
)

const (
	Error = "error"
	Message = "message"
)

type response struct {
	MessageType string `json:"message_type"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func NewResponse(messageType string, message string,  data interface{}) response {
	return response {
		messageType,
		message,
		data,
	}
}

func ResponseJSON(writer http.ResponseWriter, statusCode int, resp response ) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	err := json.NewEncoder(writer).Encode(&resp)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
