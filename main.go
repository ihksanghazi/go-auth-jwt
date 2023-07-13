package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ihksanghazi/go-auth-jwt/controllers/authcontroller"
	"github.com/ihksanghazi/go-auth-jwt/models"
)

func main() {
	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/login",authcontroller.Login).Methods("POST")
	r.HandleFunc("/register",authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout",authcontroller.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", r))
}