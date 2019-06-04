package types

import "time"

type (
	// StatisticData is data structure for statistic data
	StatisticData struct {
		UniqueString string               `json:"unique_string"`
		Clicks       int64                `json:"clicks"`
		DailyData    []StatisticDailyData `json:"daily_data"`
	}

	// StatisticDailyData is statistic data per day
	StatisticDailyData struct {
		Date   time.Time `json:"date"`
		Clicks int64     `json:"clicks"`
	}

	// UpdateStatisticRequest is request param to update statistic of shortURL from a single request
	UpdateStatisticRequest struct {
		UniqueString string
		IP           string
		UserAgent    string
	}
)
