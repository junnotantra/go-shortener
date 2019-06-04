package statistic

import "github.com/junnotantra/go-shortener/internal/types"

type (
	// Resource interface define contract for resource level of statistic
	Resource interface {
		UpdateStatistic(statisticData types.StatisticData) error
		GetStatisticInfo(uniqueString string) (types.StatisticData, error)
	}

	// Service struct hold the required services and resource and function receiver for
	// service level of statistic
	Service struct {
		resource Resource
	}
)

// New function return reference of Service struct, never use Service Struct without new
func New(statisticResource Resource) *Service {
	s := &Service{
		resource: statisticResource,
	}

	return s
}
