package main

import (
	"fmt"
	"net/http"
	"wfical/config"
	"wfical/handlers"

	log "github.com/sirupsen/logrus"
)

func main() {
	router := handlers.Setup()

	log.WithFields(log.Fields{
		"port": config.GetEnv().Port,
	}).Info("Starting server")

	err := http.ListenAndServe(fmt.Sprintf(":%d", config.GetEnv().Port), router)
	if err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}
