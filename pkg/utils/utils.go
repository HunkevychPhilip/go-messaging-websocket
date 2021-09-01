package utils

import "net/http"

type ResponseWriter interface {
	GenerateError(w http.ResponseWriter, status int, msg string)
}

type Utils struct {
	ResponseWriter
}

func NewUtils(responseWriter ResponseWriter) *Utils {
	return &Utils{ResponseWriter: responseWriter}
}
