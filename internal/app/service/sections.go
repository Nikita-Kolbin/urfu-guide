package service

import (
	"context"
	"fmt"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
)

func (s *Service) GetSections(ctx context.Context) ([]*model.Language, error) {
	languages, err := s.repo.GetLanguages(ctx)
	if err != nil {
		return nil, err
	}

	for _, lang := range languages {
		sections, err := s.repo.GetSections(ctx, lang.Code)
		if err != nil {
			return nil, err
		}
		lang.Sections = sections

		for _, section := range lang.Sections {
			err := s.addContent(ctx, section)
			if err != nil {
				return nil, err
			}
		}
	}

	return languages, nil
}

func (s *Service) addContent(ctx context.Context, section *model.Section) error {
	switch section.ContentType {
	case model.DefaultBlockContent:
		content, err := s.repo.GetDefaultBlock(ctx, section.ID)
		if err != nil {
			return err
		}
		section.DefaultBlock = content
	default:
		return fmt.Errorf("unknown content type: %s", section.ContentType)
	}
	return nil
}
