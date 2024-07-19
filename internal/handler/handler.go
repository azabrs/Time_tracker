package handler

import (
	"Time-tracker/internal/service/time_tracker"
	"github.com/go-chi/chi"
	"net/http"
)

type handler struct {
	Service time_tracker.Service
}

func NewHandler(service time_tracker.Service, mux *chi.Mux) {
	h := handler{
		service,
	}
	mux.MethodFunc(http.MethodGet, "/GetUser", h.GetUsers)
}
