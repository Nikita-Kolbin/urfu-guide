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
	if section == nil {
		return fmt.Errorf("section is nil")
	}

	switch section.ContentType {
	case model.DefaultBlockContent:
		content, err := s.repo.GetDefaultBlock(ctx, section.ID)
		if err != nil {
			return err
		}
		section.Data = content
	case model.DropDownListBlockContent:
		block, err := s.repo.GetDropDownBlock(ctx, section.ID)
		if err != nil {
			return err
		}
		elements, err := s.repo.GetDropDownElements(ctx, block.ID)
		if err != nil {
			return err
		}
		block.Elements = elements
		section.Data = block
	case model.DefaultBlockListContent:
		blocks, err := s.repo.GetDefaultBlockList(ctx, section.ID)
		if err != nil {
			return err
		}
		section.Data = blocks
	case model.DropDownListBlockListContent:
		blocks, err := s.repo.GetDropDownBlockList(ctx, section.ID)
		if err != nil {
			return err
		}
		for _, block := range blocks {
			elements, err := s.repo.GetDropDownElements(ctx, block.ID)
			if err != nil {
				return err
			}
			block.Elements = elements
		}
		section.Data = blocks
	default:
		return fmt.Errorf("unknown content type: %s", section.ContentType)
	}

	return nil
}
