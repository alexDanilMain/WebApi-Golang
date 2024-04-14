package main

type APIServer struct {
	address string
}

func NewApiServer(address string) *APIServer {
	return &APIServer{
		address: address,
	}
}
