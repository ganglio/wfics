package handlers

import (
	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ics", IcsHandler).Methods("GET")

	return router
}
