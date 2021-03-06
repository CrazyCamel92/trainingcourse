package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	)

//testing using shell curl command curl -X POST -H "Content-Type: application/json" -d '{"username":"erez","password":"mypassword"}' http://localhost:8080/login


func main() {
	router:= mux.NewRouter()
	endpointsManager := NewEndpointsManager();
	router.PathPrefix("/").Handler(http.StripPrefix("/",http.FileServer(http.Dir("assets/")))).Methods("GET")
	router.HandleFunc("/login",endpointsManager.LoginEndpoint).Methods("POST")
	router.HandleFunc("/register",endpointsManager.RegisterEndpoint).Methods("POST")
	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080",router))
}

