package service

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

const shortCodeAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (s *Service) Create(ctx context.Context, originalURL string) (*URL, error) {
	trimmed := strings.TrimSpace(originalURL)
	if trimmed == "" {
		return nil, errors.New("original URL must not be empty")
	}

	shortCode, err := generateShortCode(8)
	if err != nil {
		return nil, err
	}

	repoURL, err := s.repo.Create(ctx, shortCode, trimmed)
	if err != nil {
		return nil, err
	}

	return &URL{
		ID:          repoURL.ID,
		OriginalURL: repoURL.OriginalURL,
		ShortCode:   repoURL.ShortCode,
		CreatedAt:   repoURL.CreatedAt,
	}, nil
}

func generateShortCode(length int) (string, error) {
	if length <= 0 {
		length = 8
	}
	var b strings.Builder
	b.Grow(length)
	alphabetLen := big.NewInt(int64(len(shortCodeAlphabet)))
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, alphabetLen)
		if err != nil {
			return "", err
		}
		b.WriteByte(shortCodeAlphabet[n.Int64()])
	}
	return b.String(), nil
}
