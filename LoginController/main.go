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
	router.HandleFunc("/login",LoginEndpoint)
	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080",router))
}

