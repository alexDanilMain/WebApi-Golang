package main

import "net/http"

func setupRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /fizzbuzz/{number}", fizzBuzzHandler)
	router.HandleFunc("GET /users/{userID}", getUserHandler)
	router.HandleFunc("GET /users", getAllUsersHandler)
	router.HandleFunc("POST /users", postUserHandler)
	router.HandleFunc("POST /users/{userID}", noPostHandler)
}
