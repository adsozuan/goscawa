package goscawa

import (
	"fmt"
	"net/http"
	"os"
)

func (s *Server) routes() {
	s.router.Handle("/", s.handleOla())
}

func (s *Server) handleOla() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		name := r.FormValue("name")
		s.logger.Info(fmt.Sprintf("GET / name=%s", name))
		fmt.Fprintf(w, "Hello %s from %s", name, hostname)
	}

}
