package shortener

import "github.com/junnotantra/go-shortener/internal/types"

type (
	// Resource interface define contract for resource level of shortener
	Resource interface {
		SaveShortURL(shortURL types.ShortURL, allowUpdate bool) error
		GetShortURLInfo(uniqueString string) (types.ShortURL, error)
	}

	// Service struct hold the required services and resource and function receiver for
	// service level of shortener
	Service struct {
		resource Resource
	}
)

// New function return reference of Service struct, never use Service Struct without new
func New(shortenerResource Resource) *Service {
	s := &Service{
		resource: shortenerResource,
	}

	return s
}
