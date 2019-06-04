package types

import "time"

type (
	// ShortURLStatus define shortURL status
	ShortURLStatus = int

	// ShortURL is data structure for shortened URL
	ShortURL struct {
		UniqueString string         `json:"unique_string"`
		FullURL      string         `json:"full_url"`
		Status       ShortURLStatus `json:"status"`
		CreatedTime  time.Time      `json:"created_time"`
		CreatedBy    int64          `json:"created_by"`
		UpdatedTime  time.Time      `json:"updated_time"`
		UpdatedBy    int64          `json:"updated_by"`
	}

	// CreateShortenURLRequest is request params to create new short URL
	CreateShortenURLRequest struct {
		FullURL      string
		UniqueString string
		UserID       int64
	}
)

const (
	// ShortURLStatusActive means short URL can be accessed
	ShortURLStatusActive ShortURLStatus = 1
	// ShortURLStatusModerated means short URL is temporarily moderated
	ShortURLStatusModerated ShortURLStatus = 2
	// ShortURLStatusInactive means short URL cannot be accessed anymore
	ShortURLStatusInactive ShortURLStatus = -1
)
