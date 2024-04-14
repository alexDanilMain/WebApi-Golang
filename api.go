package main

import (
	"log"
	"net/http"
)

type APIServer struct {
	address string
}

func NewApiServer(address string) *APIServer {
	return &APIServer{
		address: address,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("GET /users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		w.Write([]byte("User ID: " + userID))
	})

	router.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Gets all users"))
	})

	router.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Post a user"))
	})

	router.HandleFunc("/users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This route only has a GET request!"))
	})

	server := http.Server{
		Addr:    s.address,
		Handler: RequestLoggerMiddleware(router),
	}

	log.Printf("Server has started at %s", s.address)

	return server.ListenAndServe()
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s, path %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	}
}
