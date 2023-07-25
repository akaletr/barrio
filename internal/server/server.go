package server

import "net/http"

type server struct {
	srv *http.Server
}

func (s *server) Start() error {
	return s.srv.ListenAndServe()
}

func New() Server {
	return &server{
		srv: &http.Server{
			Addr: "",
		}}
}
