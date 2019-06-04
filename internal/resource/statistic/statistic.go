package statistic

import (
	"encoding/json"

	"github.com/junnotantra/go-shortener/internal/types"
	bolt "go.etcd.io/bbolt"
)

const (
	// BucketStatistic is constant for statistic bucket name on DB
	BucketStatistic = "Statistic"
)

// UpdateStatistic will update statistic data for a unique string
func (r *Resource) UpdateStatistic(statisticData types.StatisticData) error {
	var (
		data []byte
		err  error
	)

	// marshall to json to be saved
	data, err = json.Marshal(statisticData)
	if err != nil {
		return err
	}

	// update transaction
	err = r.DB.Update(func(tx *bolt.Tx) error {
		// create bucket if not exist
		b, err := tx.CreateBucketIfNotExists([]byte(BucketStatistic))
		if err != nil {
			return err
		}

		// save to DB
		err = b.Put([]byte(statisticData.UniqueString), data)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// GetStatisticInfo will get statistic data from DB
func (r *Resource) GetStatisticInfo(uniqueString string) (types.StatisticData, error) {
	var (
		data          []byte
		statisticData types.StatisticData
		err           error
	)

	// get data from DB
	err = r.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketStatistic))
		if b != nil {
			data = b.Get([]byte(uniqueString))
		}
		return nil
	})
	if err != nil {
		return statisticData, err
	}

	// validate data
	if len(data) == 0 {
		return statisticData, types.ErrNotFound
	}

	// unmarshall data
	err = json.Unmarshal(data, &statisticData)
	if err != nil {
		return statisticData, err
	}

	return statisticData, nil
}
