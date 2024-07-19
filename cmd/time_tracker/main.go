package main

import (
	"Time-tracker/internal/config"
	"Time-tracker/internal/handler"
	"Time-tracker/internal/repository/time_tracker_db"
	"Time-tracker/internal/service/time_tracker"
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	cfg := config.LoadConfig(ctx)
	db, err := time_tracker_db.NewDB(*cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	mux := chi.NewRouter()
	service := time_tracker.New(&db)
	handler.NewHandler(service, mux)
	httpServer := http.Server{
		Addr:    fmt.Sprintf(cfg.Server.Host + ":" + cfg.Server.Port),
		Handler: mux,
	}
	log.Println("Server start")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
