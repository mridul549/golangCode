package router

import (
	controller "mongo/controllers"

	"github.com/gorilla/mux"
)

func Router () *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/createMovie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/udpateMovies", controller.MarkAsWatched).Methods("PATCH")
	router.HandleFunc("/api/deleteMovie", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteAllMovies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}