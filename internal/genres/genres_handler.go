package genres

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// CreateGenresHandler -
func CreateGenresHandler(srv *Service) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.StripSlashes)
	router.Post("/jointags", srv.joiningtags())
	router.Get("/list", srv.genresList(srv.stortagsPath))

	return router
}
