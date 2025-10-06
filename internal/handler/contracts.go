package handler

import (
	"context"
	"github.com/aabbuukkaarr8/internal/repository"

	"github.com/aabbuukkaarr8/internal/service"
)

type Service interface {
	Create(ctx context.Context, originalURL string) (*service.URL, error)
	ResolveAndTrack(ctx context.Context, shortCode string, userAgent string, ipAddress string) (*service.URL, error)
	GetAnalytics(ctx context.Context, shortCode string, topN int) (*repository.AnalyticsResult, error)
}
