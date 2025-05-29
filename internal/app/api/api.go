package api

import (
	"context"
	"fmt"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"

	"net/http"

	_ "github.com/Nikita-Kolbin/urfu-guide/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type service interface {
	GetSections(ctx context.Context) ([]*model.Language, error)
}

type API struct {
	srv service
}

func New(_ context.Context, srv service, address string) http.Handler {
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*", "https://*", "http://*", "http://127.0.0.1:3000"}, // TODO: Настроить корсы
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders: []string{"*", "Accept", "Authorization", "Content-Type", "X-Token"},
	}))

	// swagger
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", address)),
	))

	api := API{srv: srv}

	// handlers
	router.Get("/api/health", api.Health)
	router.Get("/api/v1/sections", api.GetSections)

	return router
}
