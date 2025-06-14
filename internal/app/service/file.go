package service

import (
	"context"
	"fmt"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
	"io"
)

const defaultPath = "http://%s/api/file/%s" // Первый подстанвка - хостпорт, вторая id файла

func (s *Service) UploadFile(
	ctx context.Context,
	reader io.Reader,
	size int64,
	contentType string,
) (model.UploadFileResponse, error) {
	id, err := s.storage.PutObject(ctx, reader, size, model.DefaultBucketName, contentType)
	if err != nil {
		return model.UploadFileResponse{}, err
	}

	return model.UploadFileResponse{ID: id, URL: s.MakeURL(id)}, err
}

func (s *Service) GetFile(ctx context.Context, objectId string) (io.Reader, string, error) {
	return s.storage.GetObject(ctx, objectId, model.DefaultBucketName)
}

func (s *Service) MakeURL(id string) string {
	return fmt.Sprintf(defaultPath, s.serverHostPort, id)
}
