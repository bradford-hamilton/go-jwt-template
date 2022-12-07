package server

import (
	"os"
	"time"

	"github.com/bradford-hamilton/go-jwt-template/internal/storage"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type API struct {
	baseURL string
	db      storage.GoJWTDb
	Mux     *chi.Mux
}

func New(db storage.GoJWTDb) *API {
	r := chi.NewRouter()
	r.Use(
		corsMiddleware().Handler,
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.StripSlashes,
		middleware.Recoverer,
		middleware.Timeout(30*time.Second),
	)

	baseURL := "http://localhost:4000"
	if os.Getenv("GO_JWT_TEMPLATE_ENVIRONMENT") == "production" {
		baseURL = "TODO"
	}

	api := &API{baseURL: baseURL, db: db, Mux: r}
	api.initializeRoutes()

	return api
}

func (a *API) initializeRoutes() {
	a.Mux.Get("/ping", a.ping)
}
