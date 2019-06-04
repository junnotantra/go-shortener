package statistic

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/junnotantra/go-shortener/internal/service/shortener"
	"github.com/junnotantra/go-shortener/internal/service/statistic"
	"github.com/junnotantra/go-shortener/internal/types"
	"github.com/junnotantra/go-shortener/internal/utils/response"
)

var (
	shortenerService *shortener.Service
	statisticService *statistic.Service
)

// Init will initialize shortener service
func Init(shortener *shortener.Service, statistic *statistic.Service) {
	shortenerService = shortener
	statisticService = statistic
}

// HandleGetStatisticInfo will get statistic data
func HandleGetStatisticInfo(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.Render(w, r)

	// get request data
	uniqueStr := chi.URLParam(r, "uniqueStr")

	// get full URL
	shortURL, err := shortenerService.GetShortURLInfo(uniqueStr)
	if err != nil {
		log.Println(err)
		if err == types.ErrNotFound {
			resp.SetError(err, http.StatusNotFound)
			return
		}
		resp.SetError(err, http.StatusInternalServerError)
		return
	}

	// get statistic data
	statistic, err := statisticService.GetStatisticInfo(uniqueStr)
	if err != nil && err != types.ErrNotFound {
		log.Println(err)
		resp.SetError(err, http.StatusInternalServerError)
		return
	}

	// prepare response
	data := struct {
		Item      types.ShortURL      `json:"item"`
		Statistic types.StatisticData `json:"statistic"`
	}{
		Item:      shortURL,
		Statistic: statistic,
	}

	resp.Data = data
	return
}
