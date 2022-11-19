package goscawa

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) routes() {
	api := s.router.Group("/api")
	v1 := api.Group("v1")
	v1.Handle("POST", "/nearest", s.handleNearest)
}

func (s *Server) handleNearest(c *gin.Context) {
	var req FindNearestRequest
	if c.ShouldBind(&req) == nil {
		log.Println(req)
	}
	rsp := FindNearestResponse{
		Id:   10,
		Name: "Chez lui",
	}
	c.JSON(http.StatusOK, rsp)

}
