package repository

import (
	"context"
	"fmt"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
)

func (r *Repository) GetLanguages(ctx context.Context) (model.LanguageList, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `SELECT code FROM languages`

	res := make(model.LanguageList, 0)

	err := r.conn.SelectContext(ctx, &res, query)
	if err != nil {
		return nil, fmt.Errorf("[GetLanguages] %w", err)
	}

	return res, nil
}
