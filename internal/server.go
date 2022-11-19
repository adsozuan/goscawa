package goscawa

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	logger *AppLogger
}

func NewServer(logger *AppLogger) *Server {
	s := &Server{}
	s.logger = logger
	s.router = gin.Default()
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
