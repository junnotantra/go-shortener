package shortener

import (
	"math/rand"
	"time"

	"github.com/junnotantra/go-shortener/internal/config"
	"github.com/junnotantra/go-shortener/internal/types"
)

// GenerateUniqueString will generate shortened URL from a full URL
func (s *Service) GenerateUniqueString(fullURL string) (string, error) {
	// get config
	cfg := config.Get()
	charset := cfg.Shortener.Charset
	length := cfg.Shortener.UniqueStringLength

	// seed random
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// get chars
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b), nil
}

// CreateShortURL will generate shortened URL from a full URL
// it will persist data to storage and return the generated shortened URL
func (s *Service) CreateShortURL(req types.CreateShortenURLRequest) (string, error) {
	var (
		uniqueString string
		now          time.Time
		shortURL     types.ShortURL
		err          error
	)

	// if custom unique string not supplied, generate new unique string
	// else validate supplied unique string and use it
	if req.UniqueString == "" {
		uniqueString, err = s.GenerateUniqueString(req.FullURL)
		if err != nil {
			return "", err
		}
	} else {
		uniqueString = req.UniqueString
	}

	// construct shortURL data
	now = time.Now()
	shortURL = types.ShortURL{
		FullURL:      req.FullURL,
		UniqueString: uniqueString,
		Status:       types.ShortURLStatusActive,
		CreatedTime:  now,
		CreatedBy:    req.UserID,
		UpdatedTime:  now,
		UpdatedBy:    req.UserID,
	}

	// save to DB
	err = s.resource.SaveShortURL(shortURL, false)
	if err != nil {
		return "", err
	}
	return shortURL.UniqueString, err
}

// UpdateShortURL will update shor URL data on storage
func (s *Service) UpdateShortURL(req types.CreateShortenURLRequest) error {
	var (
		shortURL types.ShortURL
		now      time.Time
		err      error
	)

	// get old data
	shortURL, err = s.resource.GetShortURLInfo(req.UniqueString)
	if err != nil {
		return err
	}

	// validate user
	if req.UserID == 0 || req.UserID != shortURL.CreatedBy {
		return types.ErrUnauthorized
	}

	// update shortURL data
	now = time.Now()
	shortURL.FullURL = req.FullURL
	shortURL.UpdatedTime = now
	shortURL.UpdatedBy = req.UserID

	// save to DB
	err = s.resource.SaveShortURL(shortURL, true)
	if err != nil {
		return err
	}
	return err
}

// GetShortURLInfo will get the short URL data from DB
func (s *Service) GetShortURLInfo(uniqueString string) (types.ShortURL, error) {
	var (
		shortURL types.ShortURL
		err      error
	)

	// get data
	shortURL, err = s.resource.GetShortURLInfo(uniqueString)
	if err != nil {
		return shortURL, err
	}

	return shortURL, nil
}
