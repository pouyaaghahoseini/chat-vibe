package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"chatvibe/backend/internal/app"
	"chatvibe/backend/internal/config"
	"chatvibe/backend/internal/httpserver"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	application, err := app.New(ctx, cfg)
	if err != nil {
		log.Fatalf("initialize app: %v", err)
	}
	defer application.Close()

	handler := httpserver.NewHandler(application)
	router := httpserver.NewRouter(handler)

	server := &http.Server{
		Addr:              ":" + cfg.BackendPort,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("ChatVibe backend listening on :%s", cfg.BackendPort)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
