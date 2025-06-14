package repository

import (
	"context"
	"fmt"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
)

func (r *Repository) GetSections(ctx context.Context, languageCode string) (model.SectionList, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT id, language_code, content_type, title, icon, position 
	FROM sections WHERE language_code = $1 ORDER BY position`

	res := make(model.SectionList, 0)

	err := r.conn.SelectContext(ctx, &res, query, languageCode)
	if err != nil {
		return nil, fmt.Errorf("[GetSections] %w", err)
	}

	return res, nil
}
