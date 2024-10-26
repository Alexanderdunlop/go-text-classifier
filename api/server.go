package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}
	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.router.HandleFunc("/api/classify", s.handleClassify).Methods("POST")
	s.router.HandleFunc("/api/health", s.handleHealth).Methods("GET")
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
