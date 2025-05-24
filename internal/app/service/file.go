package service

import (
	"context"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
	"io"
)

func (s *Service) UploadImage(
	ctx context.Context,
	reader io.Reader,
	size int64,
	contentType string,
) (string, error) {
	return s.storage.PutObject(ctx, reader, size, model.ImageBucketName, contentType)
}

func (s *Service) GetImage(ctx context.Context, objectId string) (io.Reader, string, error) {
	return s.storage.GetObject(ctx, objectId, model.ImageBucketName)
}
