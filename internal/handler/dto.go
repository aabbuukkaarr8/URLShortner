package handler

import "time"

type URL struct {
	ID          int64
	OriginalURL string
	ShortCode   string
	CreatedAt   time.Time
}

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	ShortURL string `json:"short_url"`
}

type RedirectParams struct {
	Code string `uri:"code" binding:"required"`
}
