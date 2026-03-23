package httpserver

import "net/http"

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/ready", handler.Ready)

	return mux
}
