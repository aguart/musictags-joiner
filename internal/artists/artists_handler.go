package artists

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// CreateArtistsHandler -
func CreateArtistsHandler(srv *Service) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.StripSlashes)
	router.Get("/", srv.artistsList())
	return router
}
