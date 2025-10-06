package service

import (
	"context"
	"github.com/aabbuukkaarr8/internal/repository"
	"time"
)

type Repository interface {
	Create(ctx context.Context, shortCode string, originalURL string) (*repository.URL, error)
	FindByShortCode(ctx context.Context, shortCode string) (*repository.URL, error)
	RecordClick(ctx context.Context, shortCode string, userAgent, ip string, clickedAt time.Time) error
	GetAnalytics(ctx context.Context, shortCode string, topN int) (*repository.AnalyticsResult, error)
}
