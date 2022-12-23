package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"playground/config"
	"playground/repository/postgres"
)

func main() {
	spec := config.InitSpecConfig()
	dbConfig := config.InitPostgresConfig(spec.Debug)
	_, err := postgres.NewPostgresDB(dbConfig)
	log.Info("After connecting to bd")
	if err != nil {
		panic("Error with creating postgres db")
	}

	http.HandleFunc("/", HelloHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", spec.Port), nil); err != nil {
		log.Fatal(err.Error())
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from go app")
}
