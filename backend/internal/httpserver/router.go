package httpserver

import (
	"net/http"

	"chatvibe/backend/internal/middleware"
)

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/ready", handler.Ready)

	return middleware.Logging(mux)
}
