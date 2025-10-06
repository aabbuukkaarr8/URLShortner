package service

import (
	"context"
	"errors"
	"strings"
	"time"
)

func (s *Service) ResolveAndTrack(ctx context.Context, shortCode string, userAgent string, ipAddress string) (*URL, error) {
	if strings.TrimSpace(shortCode) == "" {
		return nil, errors.New("short code must not be empty")
	}
	repoURL, err := s.repo.FindByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}
	_ = s.repo.RecordClick(ctx, shortCode, userAgent, ipAddress, time.Now().UTC())
	return &URL{
		ID:          repoURL.ID,
		OriginalURL: repoURL.OriginalURL,
		ShortCode:   repoURL.ShortCode,
		CreatedAt:   repoURL.CreatedAt,
	}, nil
}
