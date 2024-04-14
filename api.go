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

	server := http.Server{
		Addr:    s.address,
		Handler: router,
	}

	log.Printf("Server has started at %s", s.address)

	return server.ListenAndServe()
}
