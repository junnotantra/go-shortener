package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/junnotantra/go-shortener/internal/api/redirector"
	"github.com/junnotantra/go-shortener/internal/api/shortener"
	"github.com/junnotantra/go-shortener/internal/api/statistic"
)

func handler() http.Handler {
	r := chi.NewRouter()

	r.Get("/", redirector.HandleRoot)

	// shortener API
	r.Route("/shortener/v1/", func(r chi.Router) {
		r.Post("/new", shortener.HandleCreateShortURL)
		r.Get("/info/{uniqueStr}", shortener.GetShortURLInfo)
		r.Post("/update", shortener.HandleUpdateShortURL)
	})

	// statistic API
	r.Route("/statistic/v1", func(r chi.Router) {
		r.Get("/info/{uniqueStr}", statistic.HandleGetStatisticInfo)
	})

	// redirector
	r.Get("/{uniqueStr}", redirector.HandleRedirect)

	return r
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You know... for shortening"))
}
