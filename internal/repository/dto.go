package repository

import "time"

type URL struct {
	ID          int64
	OriginalURL string
	ShortCode   string
	CreatedAt   time.Time
}

type DailyClicks struct {
	Date       time.Time `json:"date"`
	ClickCount int       `json:"click_count"`
}

type AnalyticsResult struct {
	ShortCode     string         `json:"short_code"`
	TotalClicks   int            `json:"total_clicks"`
	ClicksByDay   []DailyClicks  `json:"clicks_by_day"`
	TopUserAgents map[string]int `json:"top_user_agents"`
}
