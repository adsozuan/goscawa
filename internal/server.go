package goscawa

import (
	"net/http"

	"adnotanumber.com/goscawa/internal/ademe"
	"adnotanumber.com/goscawa/internal/wastedisposal"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router               *gin.Engine
	logger               *AppLogger
	wasteDisposalService wastedisposal.Service
}

func NewServer(logger *AppLogger) *Server {
	s := &Server{}
	s.logger = logger
	s.router = gin.Default()

	s.wasteDisposalService = *wastedisposal.NewService(&ademe.FakeAdemeRetriever{})
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
