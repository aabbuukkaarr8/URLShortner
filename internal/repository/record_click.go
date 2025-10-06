package repository

import (
	"context"
	"time"
)

func (r *Repository) RecordClick(ctx context.Context, shortCode, userAgent, ip string, clickedAt time.Time) error {
	if shortCode == "" {
		return nil
	}

	_, err := r.store.DB.ExecContext(ctx,
		`INSERT INTO clicks (short_code, user_agent, ip_address, timestamp)
		 VALUES ($1, $2, $3, $4)`,
		shortCode, userAgent, ip, clickedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
