package shortener

import (
	"errors"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi"
	"github.com/junnotantra/go-shortener/internal/service/shortener"
	"github.com/junnotantra/go-shortener/internal/types"
	"github.com/junnotantra/go-shortener/internal/utils/response"
)

var shortenerService *shortener.Service

// Init will initialize shortener service
func Init(service *shortener.Service) {
	shortenerService = service
}

// HandleCreateShortURL will generate shortened URL and save it to DB
func HandleCreateShortURL(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.Render(w, r)

	// get request data
	fullURL := r.FormValue("full_url")
	if !govalidator.IsURL(fullURL) {
		log.Println("invalid fullURL", fullURL)
		resp.SetError(errors.New("invalid fullURL"), http.StatusBadRequest)
		return
	}
	customUniqueStr := r.FormValue("custom_unique_str")

	// prepare request to service
	serviceReq := types.CreateShortenURLRequest{
		FullURL:      fullURL,
		UniqueString: customUniqueStr,
	}

	// create short URL
	shortURL, err := shortenerService.CreateShortURL(serviceReq)
	if err != nil {
		log.Println(err)
		if err == types.ErrUnauthorized {
			resp.SetError(err, http.StatusUnauthorized)
			return
		}
		resp.SetError(err, http.StatusInternalServerError)
		return
	}

	resp.Data = shortURL
	return
}

// GetShortURLInfo will get short URL info from unique string
func GetShortURLInfo(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.Render(w, r)

	// get request data
	uniqueStr := chi.URLParam(r, "uniqueStr")

	// get full URL
	shortURL, err := shortenerService.GetShortURLInfo(uniqueStr)
	if err != nil {
		log.Println(err)
		resp.SetError(err, http.StatusInternalServerError)
		return
	}

	resp.Data = shortURL
	return
}

// HandleUpdateShortURL will generate shortened URL and save it to DB
func HandleUpdateShortURL(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.Render(w, r)

	// get request data
	fullURL := r.FormValue("full_url")
	if !govalidator.IsURL(fullURL) {
		log.Println("invalid fullURL", fullURL)
		resp.SetError(errors.New("invalid fullURL"), http.StatusBadRequest)
		return
	}
	customUniqueStr := r.FormValue("custom_unique_str")

	// prepare request to service
	serviceReq := types.CreateShortenURLRequest{
		FullURL:      fullURL,
		UniqueString: customUniqueStr,
	}

	// create short URL
	err := shortenerService.UpdateShortURL(serviceReq)
	if err != nil {
		log.Println(err)
		if err == types.ErrUnauthorized {
			resp.SetError(err, http.StatusUnauthorized)
			return
		}
		resp.SetError(err, http.StatusInternalServerError)
		return
	}

	resp.Data = true
	return
}
