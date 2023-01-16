package main

import (
	"fmt"
	"github.com/alikud/ads-microservice/config"
	"github.com/alikud/ads-microservice/pkg/handler"
	"github.com/alikud/ads-microservice/pkg/repository/postgres"
	"github.com/alikud/ads-microservice/pkg/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// @title Swagger  API
// @version 1.0
// @description This is a sample CRUD ads service
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
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
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", spec.Port)))

}
