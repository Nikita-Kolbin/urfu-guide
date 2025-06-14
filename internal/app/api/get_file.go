package api

import (
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
	"github.com/Nikita-Kolbin/urfu-guide/internal/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

// GetFile   godoc
// @Summary  Get file by id
// @Tags     file
// @Accept   json
// @Produce  json
// @Param    file-id  path  string  true  "file id"
// @Success  200   {object}   []byte
// @Failure  404   {object}   model.ErrorResponse
// @Failure  500   {object}   model.ErrorResponse
// @Router /file/{file-id} [get]
func (a *API) GetFile(w http.ResponseWriter, r *http.Request) {
	const (
		contentTypeHeader = "content-type"
		fileIdURLParam    = "file-id"
	)

	ctx := r.Context()

	fileID := chi.URLParam(r, fileIdURLParam)
	reader, contentType, err := a.srv.GetFile(ctx, fileID)
	if err != nil {
		logger.Error(ctx, "[GetFile] failed to get file", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to get file"))
		return
	}

	buffer, err := io.ReadAll(reader)
	if err != nil {
		logger.Error(ctx, "[GetFile] failed create buffer", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to get file"))
		return
	}

	_, err = w.Write(buffer)
	if err != nil {
		logger.Error(ctx, "[GetFile] failed to write response", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to get file"))
		return
	}

	w.Header().Set(contentTypeHeader, contentType)

	logger.Info(ctx, "[GetFile] file given", "id", fileID)

	render.Status(r, http.StatusOK)
}
