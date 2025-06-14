package repository

import (
	"context"
	"fmt"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/model"
)

func (r *Repository) GetVersion(ctx context.Context) (model.Version, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `SELECT version FROM version LIMIT 1`

	version := model.Version{}

	err := r.conn.GetContext(ctx, &version, query)
	if err != nil {
		return version, fmt.Errorf("[GetVersion] %w", err)
	}

	return version, nil
}
