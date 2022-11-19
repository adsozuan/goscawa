package goscawa

import "net/http"

type Server struct {
	router *http.ServeMux
	logger *AppLogger
}

func NewServer(logger *AppLogger) *Server {
	s := &Server{}
	s.logger = logger
	s.router = http.NewServeMux()
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
