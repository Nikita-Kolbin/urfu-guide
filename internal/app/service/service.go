package service

import (
	"context"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
	"io"
	"time"
)

type repository interface {
	GetLanguages(ctx context.Context) (model.LanguageList, error)
	GetSections(ctx context.Context, languageCode string) (model.SectionList, error)

	GetDefaultBlock(ctx context.Context, sectionID int) (*model.DefaultBlock, error)
	GetDefaultBlockList(ctx context.Context, sectionID int) (model.DefaultBlockList, error)

	GetDropDownBlock(ctx context.Context, sectionID int) (*model.DropDownBlock, error)
	GetDropDownBlockList(ctx context.Context, sectionID int) (model.DropDownBlockList, error)
	GetDropDownElements(ctx context.Context, blockID int) (model.DropDownElementList, error)
}

type objectStorage interface {
	PutObject(ctx context.Context, reader io.Reader, size int64, bucketName, contentType string) (_ string, err error)
	GetObject(ctx context.Context, objectId, bucketName string) (io.Reader, string, error)
}

type cache interface {
	Set(ctx context.Context, prefix, key, val string, exp time.Duration) error
	Get(ctx context.Context, prefix, key string) (string, error)
}

type Service struct {
	jwtSecret string
	repo      repository
	storage   objectStorage
	cache     cache
}

func New(repo repository, storage objectStorage, cache cache, jwtSecret string) *Service {
	return &Service{
		jwtSecret: jwtSecret,
		repo:      repo,
		storage:   storage,
		cache:     cache,
	}
}
