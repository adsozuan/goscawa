package goscawa

import (
	"log"
	"net/http"

	"adnotanumber.com/goscawa/internal/wastedisposal"
	"github.com/gin-gonic/gin"
)

func (s *Server) routes() {
	api := s.router.Group("/api")
	v1 := api.Group("v1")
	v1.Handle("POST", "/nearest", s.handleNearest)
}

func (s *Server) handleNearest(c *gin.Context) {
	var req wastedisposal.FindNearestRequest
	if c.ShouldBind(&req) == nil {
		log.Println(req)
	}
	rsp := s.wasteDisposalService.FindNearest(req)
	c.JSON(http.StatusOK, rsp)
}
