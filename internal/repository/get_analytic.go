package repository

import (
	"context"
	"database/sql"
)

var ErrShortCodeNotFound = sql.ErrNoRows

func (r *Repository) GetAnalytics(ctx context.Context, shortCode string, topN int) (*AnalyticsResult, error) {
	result := &AnalyticsResult{
		ShortCode: shortCode,
	}

	// Проверяем, что shortCode существует в short_links
	var exists bool
	err := r.store.DB.QueryRowContext(ctx,
		"SELECT EXISTS(SELECT 1 FROM short_links WHERE LOWER(short_code)=LOWER($1))",
		shortCode,
	).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrShortCodeNotFound
	}

	// 1. Общее количество кликов
	if err := r.store.DB.QueryRowContext(ctx,
		"SELECT COUNT(*) FROM clicks WHERE LOWER(short_code)=LOWER($1)",
		shortCode,
	).Scan(&result.TotalClicks); err != nil {
		return nil, err
	}

	// 2. Клики по дням
	rows, err := r.store.DB.QueryContext(ctx,
		`SELECT DATE(timestamp), COUNT(*) 
		 FROM clicks 
		 WHERE LOWER(short_code)=LOWER($1)
		 GROUP BY DATE(timestamp)
		 ORDER BY DATE(timestamp)`,
		shortCode,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var d DailyClicks
		if err := rows.Scan(&d.Date, &d.ClickCount); err != nil {
			return nil, err
		}
		result.ClicksByDay = append(result.ClicksByDay, d)
	}

	// 3. Топ User-Agent
	rowsUA, err := r.store.DB.QueryContext(ctx,
		`SELECT user_agent, COUNT(*) 
		 FROM clicks 
		 WHERE LOWER(short_code)=LOWER($1)
		 GROUP BY user_agent
		 ORDER BY COUNT(*) DESC
		 LIMIT $2`,
		shortCode, topN,
	)
	if err != nil {
		return nil, err
	}
	defer rowsUA.Close()

	result.TopUserAgents = make(map[string]int)
	for rowsUA.Next() {
		var ua string
		var cnt int
		if err := rowsUA.Scan(&ua, &cnt); err != nil {
			return nil, err
		}
		result.TopUserAgents[ua] = cnt
	}

	return result, nil
}
