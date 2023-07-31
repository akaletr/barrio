package server

import (
	"net/http"

	"cmd/main/main.go/internal/entity/user"
	"cmd/main/main.go/internal/handlers"
	"cmd/main/main.go/internal/mw"

	"github.com/go-chi/chi/v5"
)

type server struct {
	srv *http.Server
}

func (s *server) Start() error {
	return s.srv.ListenAndServe()
}

func New() Server {
	router := chi.NewRouter()

	u := user.New()

	router.Use(mw.Hello)
	router.Get("/", handlers.Test)

	router.Route("/user", func(router chi.Router) {
		router.Post("/create", u.Create)
		router.Get("/{id}", user.Get)
	})

	return &server{
		srv: &http.Server{
			Addr:    ":8080",
			Handler: router,
		}}
}
