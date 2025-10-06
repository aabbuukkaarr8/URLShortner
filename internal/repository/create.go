package repository

import (
	"context"
	"time"
)

func (r *Repository) Create(ctx context.Context, shortCode string, originalURL string) (*URL, error) {
	now := time.Now().UTC()

	const query = `
		INSERT INTO short_links (original_url, short_code, created_at)
		VALUES ($1, $2, $3)
		RETURNING id, original_url, short_code, created_at
	`

	var u URL
	row := r.store.DB.Master.QueryRowContext(ctx, query, originalURL, shortCode, now)
	if err := row.Scan(&u.ID, &u.OriginalURL, &u.ShortCode, &u.CreatedAt); err != nil {
		return nil, err
	}

	return &u, nil
}
