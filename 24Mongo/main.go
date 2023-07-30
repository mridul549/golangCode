package main

import (
	"fmt"
	"log"
	router "mongo/routes"
	"net/http"
)

func main() {
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Server Started")

}
