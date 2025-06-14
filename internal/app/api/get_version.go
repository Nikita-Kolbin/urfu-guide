package api

import (
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
	"github.com/Nikita-Kolbin/urfu-guide/internal/pkg/logger"
	"github.com/go-chi/render"
	"net/http"
)

// GetVersion godoc
// @Summary  get version
// @Tags     public
// @Accept   json
// @Produce  json
// @Success  200   {object}   model.Version
// @Failure  500   {object}   model.ErrorResponse
// @Router   /v1/version [get]
func (a *API) GetVersion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res, err := a.srv.GetVersion(ctx)
	if err != nil {
		logger.Error(ctx, "[GetVersion] failed to get version", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to get version"))
		return
	}

	logger.Info(ctx, "[GetVersion] version given success")

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}
