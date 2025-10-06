package repository

import "context"

// FindByShortCode returns URL by short code or error if not found
func (r *Repository) FindByShortCode(ctx context.Context, shortCode string) (*URL, error) {
	const q = `SELECT id, original_url, short_code, created_at FROM short_links WHERE short_code = $1`
	var u URL
	row := r.store.DB.Master.QueryRowContext(ctx, q, shortCode)
	if err := row.Scan(&u.ID, &u.OriginalURL, &u.ShortCode, &u.CreatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

