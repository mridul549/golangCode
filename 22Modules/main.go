package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func greeter() {
	fmt.Println("Hello there mod users")
}

func serverHome (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series</h1>"))
}

func main() {
	fmt.Println("Welcome to mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}


