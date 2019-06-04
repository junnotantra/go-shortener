package statistic

import (
	"time"

	"github.com/junnotantra/go-shortener/internal/types"
	"github.com/junnotantra/go-shortener/internal/utils/datetime"
)

// UpdateStatistic will update statistic data for a unique string
func (s *Service) UpdateStatistic(req types.UpdateStatisticRequest) error {
	var (
		statisticData types.StatisticData
		newDailyData  []types.StatisticDailyData
		now           time.Time
		err           error
	)

	y, m, d := time.Now().Date()
	now = time.Date(y, m, d, 0, 0, 0, 0, time.Local)

	// get old data
	statisticData, err = s.resource.GetStatisticInfo(req.UniqueString)
	if err != nil && err != types.ErrNotFound {
		return err
	}

	// update statistic data
	statisticData.UniqueString = req.UniqueString
	statisticData.Clicks++

	// update daily data
	dateUpdated := false
	for _, v := range statisticData.DailyData {
		// update current date data
		if datetime.DateEqual(v.Date, now) {
			dateUpdated = true
			v.Clicks++
		}

		// only save last 30 days
		if now.Sub(v.Date) < (time.Hour * 24 * 30) {
			newDailyData = append(newDailyData, v)
		}
	}

	// if no today data, create new
	if !dateUpdated {
		newDailyData = append(newDailyData, types.StatisticDailyData{
			Clicks: 1,
			Date:   now,
		})
	}

	// replace with new daily data
	statisticData.DailyData = newDailyData

	// save to DB
	err = s.resource.UpdateStatistic(statisticData)
	if err != nil {
		return err
	}
	return err
}

// GetStatisticInfo will get the short URL data from DB
func (s *Service) GetStatisticInfo(uniqueString string) (types.StatisticData, error) {
	var (
		statistic types.StatisticData
		err       error
	)

	// get data
	statistic, err = s.resource.GetStatisticInfo(uniqueString)
	if err != nil {
		return statistic, err
	}

	return statistic, nil
}
