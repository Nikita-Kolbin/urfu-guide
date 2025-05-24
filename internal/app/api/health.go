package api

import (
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
	"github.com/go-chi/render"
	"net/http"
)

// Health godoc
// @Summary  check health
// @Tags     admin
// @Accept   json
// @Produce  json
// @Success  200   {object}   model.HealthResponse
// @Failure  500   {object}   model.ErrorResponse
// @Router   /health [get]
func Health(_ service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, model.HealthResponse{Success: true})
	}
}
