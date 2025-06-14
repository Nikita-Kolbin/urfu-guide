package repository

import (
	"context"
	"fmt"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
)

func (r *Repository) GetDefaultBlock(ctx context.Context, sectionID int) (*model.DefaultBlock, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT id, section_id, address, timetable, phone, description, link, image, position
	FROM default_blocks WHERE section_id = $1 LIMIT 1`

	res := &model.DefaultBlock{}

	err := r.conn.GetContext(ctx, res, query, sectionID)
	if err != nil {
		return nil, fmt.Errorf("[GetDefaultBlock] %w", err)
	}

	return res, nil
}

func (r *Repository) GetDefaultBlockList(ctx context.Context, sectionID int) (model.DefaultBlockList, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT id, section_id, address, timetable, phone, description, link, image, position
	FROM default_blocks WHERE section_id = $1 ORDER BY position`

	res := make(model.DefaultBlockList, 0)

	err := r.conn.SelectContext(ctx, &res, query, sectionID)
	if err != nil {
		return nil, fmt.Errorf("[GetDefaultBlockList] %w", err)
	}

	return res, nil
}
