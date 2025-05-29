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
	SELECT id, section_id, address, timetable, phone, description 
	FROM default_blocks WHERE section_id = $1 LIMIT 1`

	res := &model.DefaultBlock{}

	err := r.conn.GetContext(ctx, res, query, sectionID)
	if err != nil {
		return nil, fmt.Errorf("[GetDefaultBlock] %w", err)
	}

	return res, nil
}
