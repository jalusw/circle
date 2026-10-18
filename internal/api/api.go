package api

import (
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func New(mux *http.ServeMux) *Server {
	return &Server{mux: mux}
}

func (s *Server) Routes() {
	s.mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
}
