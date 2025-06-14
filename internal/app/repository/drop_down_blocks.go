package repository

import (
	"context"
	"fmt"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
)

func (r *Repository) GetDropDownBlock(ctx context.Context, sectionID int) (*model.DropDownBlock, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT id, section_id, title, sub_title, icon, position
	FROM drop_down_blocks WHERE section_id = $1 LIMIT 1`

	res := &model.DropDownBlock{}

	err := r.conn.GetContext(ctx, res, query, sectionID)
	if err != nil {
		return nil, fmt.Errorf("[GetDropDownBlock] %w", err)
	}

	return res, nil
}

func (r *Repository) GetDropDownBlockList(ctx context.Context, sectionID int) (model.DropDownBlockList, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT id, section_id, title, sub_title, icon, position
	FROM drop_down_blocks WHERE section_id = $1 ORDER BY position`

	res := make(model.DropDownBlockList, 0)

	err := r.conn.SelectContext(ctx, &res, query, sectionID)
	if err != nil {
		return nil, fmt.Errorf("[GetDropDownBlockList] %w", err)
	}

	return res, nil
}

func (r *Repository) GetDropDownElements(ctx context.Context, blockID int) (model.DropDownElementList, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT id, block_id, title, description, link, position
	FROM drop_down_elements WHERE block_id = $1 ORDER BY position`

	res := make(model.DropDownElementList, 0)

	err := r.conn.SelectContext(ctx, &res, query, blockID)
	if err != nil {
		return nil, fmt.Errorf("[GetDropDownElements] %w", err)
	}

	return res, nil
}
