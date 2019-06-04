package shortener

import (
	"encoding/json"
	"errors"

	"github.com/junnotantra/go-shortener/internal/types"
	bolt "go.etcd.io/bbolt"
)

const (
	// BucketShortURL is constant for short URL bucket name on DB
	BucketShortURL = "ShortURL"
)

// SaveShortURL will persist data to storage
func (r *Resource) SaveShortURL(shortURL types.ShortURL, allowUpdate bool) error {
	var (
		data []byte
		err  error
	)

	// marshall to json to be saved
	data, err = json.Marshal(shortURL)
	if err != nil {
		return err
	}

	// update transaction
	err = r.DB.Update(func(tx *bolt.Tx) error {
		// create bucket if not exist
		b, err := tx.CreateBucketIfNotExists([]byte(BucketShortURL))
		if err != nil {
			return err
		}

		// if not allowed to update, check if unique string already exist
		if !allowUpdate {
			exist := b.Get([]byte(shortURL.UniqueString))
			if exist != nil {
				return errors.New("key exist")
			}
		}

		// save to DB
		err = b.Put([]byte(shortURL.UniqueString), data)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// GetShortURLInfo will get short URL data from DB
func (r *Resource) GetShortURLInfo(uniqueString string) (types.ShortURL, error) {
	var (
		data         []byte
		shortURLData types.ShortURL
		err          error
	)

	// get data from DB
	err = r.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketShortURL))
		if b != nil {
			data = b.Get([]byte(uniqueString))
		}
		return nil
	})
	if err != nil {
		return shortURLData, err
	}

	// validate data
	if len(data) == 0 {
		return shortURLData, types.ErrNotFound
	}

	// unmarshall data
	err = json.Unmarshal(data, &shortURLData)
	if err != nil {
		return shortURLData, err
	}

	return shortURLData, nil
}
