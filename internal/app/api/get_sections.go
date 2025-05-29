package api

import (
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
	"github.com/Nikita-Kolbin/urfu-guide/internal/pkg/logger"
	"github.com/go-chi/render"
	"net/http"
)

// GetSections godoc
// @Summary  check health
// @Tags     public
// @Accept   json
// @Produce  json
// @Success  200   {object}   model.LanguageList
// @Failure  500   {object}   model.ErrorResponse
// @Router   /v1/sections [get]
func (a *API) GetSections(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res, err := a.srv.GetSections(ctx)
	if err != nil {
		logger.Error(ctx, "[GetSections] failed to get sections", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to get sections"))
		return
	}

	logger.Info(ctx, "[GetSections] sections given success")

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}
