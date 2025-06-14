package service

import (
	"context"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
)

func (s *Service) GetVersion(ctx context.Context) (model.Version, error) {
	return s.repo.GetVersion(ctx)
}
