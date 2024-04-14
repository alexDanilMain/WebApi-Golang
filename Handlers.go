package main

import (
	"net/http"
	"strconv"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	w.Write([]byte("User ID: " + userID))
}

func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gets all users"))
}

func postUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post a user"))
}

func noPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This route only has a GET request!"))
}

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	numberStr := r.PathValue("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}
	w.Write([]byte(FizzBuzz(number)))
}
