package main

import (
	"EmailN/internal/domain/campaing"
	"EmailN/internal/endpoints"
	"EmailN/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaingService := campaing.ServiceImp{
		Repository: &database.CampaingRepository{},
	}
	handler := endpoints.Handler{
		CampaingService: &campaingService,
	}
	r.Post("/campaings", endpoints.HandlerError(handler.CampaingPost))
	r.Get("/campaings/{id}", endpoints.HandlerError(handler.CampaingGetById))

	http.ListenAndServe(":3000", r)
}
