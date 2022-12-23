package main

import (
	"fmt"
	"github.com/alikud/ads-microservice/config"
	"github.com/alikud/ads-microservice/pkg/handler"
	"github.com/alikud/ads-microservice/pkg/repository/postgres"
	"github.com/alikud/ads-microservice/pkg/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	spec := config.InitSpecConfig()
	log.Infof("Init Specification config debug: %t port: %s", spec.Debug, spec.Port)

	dbConfig := config.InitPostgresConfig(spec.Debug)
	pool := postgres.NewPostgresDB(dbConfig)
	repository := postgres.NewRepository(pool)
	services := service.NewService(repository)

	e := echo.New()
	handlers := handler.NewHandler(services, e)
	handlers.InitRoutes()

	log.Infof("Init database connection in debug: %t mode", spec.Debug)

	http.HandleFunc("/", HelloHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", spec.Port)))

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from go app")
}
