package utils

import (
	"encoding/json"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/types"
	"log"
	"net/http"
)

type responseWriter struct {
}

func NewResponseWriter() *responseWriter {
	return &responseWriter{}
}

func (r *responseWriter) GenerateError(w http.ResponseWriter, status int, msg string) {
	err := json.NewEncoder(w).Encode(types.CustomError{
		Status: status,
		Msg:    msg,
	})

	if err != nil {
		log.Println("Failed to encode err:" + err.Error())
	}
}
