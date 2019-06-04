package redirector

import (
	"log"
	"net/http"

	"github.com/junnotantra/go-shortener/internal/config"

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

// HandleRoot will handle redirection for root URL
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	// check redirect config
	cfg := config.Get()
	if cfg.Redirect.BaseRedirect.Active {

		// redirect to configured url
		http.Redirect(w, r, cfg.Redirect.BaseRedirect.URL, http.StatusTemporaryRedirect)
	}

	// save statistic
	go func() {
		statReq := types.UpdateStatisticRequest{
			UniqueString: "/",
			IP:           r.Header.Get("X-Forwarded-For"),
			UserAgent:    r.Header.Get("User-Agent"),
		}
		err := statisticService.UpdateStatistic(statReq)
		if err != nil {
			log.Println(err)
		}
		return
	}()
	return
}

// HandleRedirect will get full URL from shortened URL data, update statistic and then redirect
func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}

	// get request data
	uniqueStr := chi.URLParam(r, "uniqueStr")

	// get full URL
	shortURL, err := shortenerService.GetShortURLInfo(uniqueStr)
	if err != nil {
		log.Println(err)
		if err == types.ErrNotFound {
			resp.SetError(err, http.StatusNotFound)
			resp.Render(w, r)
			return
		}
		resp.SetError(err, http.StatusInternalServerError)
		resp.Render(w, r)
		return
	}

	// save statistic
	go func() {
		statReq := types.UpdateStatisticRequest{
			UniqueString: uniqueStr,
			IP:           r.Header.Get("X-Forwarded-For"),
			UserAgent:    r.Header.Get("User-Agent"),
		}
		err = statisticService.UpdateStatistic(statReq)
		if err != nil {
			log.Println(err)
		}
		return
	}()

	// redirect to full url
	http.Redirect(w, r, shortURL.FullURL, http.StatusTemporaryRedirect)
	return
}
