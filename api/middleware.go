package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func (s *Server) withErrorHandling(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				errorResponse := ErrorResponse{
					Error:   "Internal Server Error",
					Code:    http.StatusInternalServerError,
					Message: "An unexpected error occurred",
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errorResponse)
			}
		}()
		next(w, r)
	}
}

func (s *Server) withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next(w, r)
	}
}

func (s *Server) chainMiddleware(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
