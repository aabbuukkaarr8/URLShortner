package service

import "github.com/aabbuukkaarr8/internal/repository"
import "context"

func (s *Service) GetAnalytics(ctx context.Context, shortCode string, topN int) (*repository.AnalyticsResult, error) {
	return s.repo.GetAnalytics(ctx, shortCode, topN)
}
