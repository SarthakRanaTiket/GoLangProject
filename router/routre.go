package router

import (
	"github.com/SarthakRanaTiket/projectname/middleware"
	"github.com/gorilla/mux"
)

//routes

func Router() *mux.Router{

	router := mux.NewRouter()

	router.HandleFunc("/api/user", middleware.GetAllUsers).Methods("GET", "OPTIONS")

	return router
}