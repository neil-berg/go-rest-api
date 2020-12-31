package handlers

import "log"

// Handler is the basic shape of handlers that take a logger
type Handler struct {
	logger *log.Logger
}

// CreateHandler returns a new Handler object
func CreateHandler(logger *log.Logger) *Handler {
	return &Handler{logger}
}
