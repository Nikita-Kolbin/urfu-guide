package api

import (
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
	"github.com/Nikita-Kolbin/urfu-guide/internal/pkg/logger"
	"github.com/go-chi/render"
	"net/http"
)

// UploadFile godoc
// @Summary Upload file
// @Tags     file
// @Accept   jpeg
// @Produce  json
// @Param file formData file true "file"
// @Success      200   {object}   model.UploadFileResponse
// @Failure      400   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router /file/upload [post]
func (a *API) UploadFile(w http.ResponseWriter, r *http.Request) {
	const contentTypeHeader = "content-type"

	ctx := r.Context()

	file, info, err := r.FormFile("file")
	if err != nil {
		logger.Error(ctx, "[UploadFile] can't get file from form", "err", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("haven't file in form"))
		return
	}

	contentType := info.Header.Get(contentTypeHeader)
	resp, err := a.srv.UploadFile(ctx, file, info.Size, contentType)
	if err != nil {
		logger.Error(ctx, "[UploadFile] can't upload file", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("can't upload file"))
		return
	}

	logger.Info(ctx, "[UploadFile] file uploaded", "id", resp.ID)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}
