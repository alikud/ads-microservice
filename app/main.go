package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

type Options struct {
	Debug    bool   `envconfig:"DEBUG" default:"true"`
	HTTPPort string `envconfig:"HTTP_PORT" default:"8080"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
	Timeout  int64  `envconfig:"TIME_OUT" default:"15"`

	PostgresqlUser     string `envconfig:"POSTGRESQL_USER" required:"true" default:"saas"`
	PostgresqlPass     string `envconfig:"POSTGRESQL_PASSWORD" required:"true" default:"saas123"`
	PostgresqlHost     string `envconfig:"POSTGRESQL_HOST" default:"0.0.0.0"`
	PostgresqlPort     string `envconfig:"POSTGRESQL_PORT" default:"5432"`
	PostgresqlDatabase string `envconfig:"POSTGRESQL_DATABASE" default:"saas"`
}

func InitLogger(logLevel string) error {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return fmt.Errorf("failed to parse log level. %v", err)
	}
	logrus.SetLevel(lvl)
	return nil
}

func main() {
	// Config
	var cfg Options
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Logger
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	lvl, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	logrus.SetLevel(lvl)

	// Syscall
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		call := <-c
		log.Printf("system call:%+v", call)
		cancel()
	}()

	// Http
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(time.Duration(cfg.Timeout) * time.Second))
	mux.Use(middleware.Heartbeat("/health"))
	mux.Handle("/metrics", promhttp.Handler())

	// sub := subscription.New(db)
	// mux.Get("/v1/saas/subscription", sub.GetByUUID)

	httpServer := &http.Server{Addr: fmt.Sprintf(":%s", cfg.HTTPPort), Handler: mux}

	go func() {
		if cErr := httpServer.ListenAndServe(); cErr != nil {
			logrus.Fatalf("failed to serve: %v", err)
		}
	}()

	logrus.Infof("Started on port %s", cfg.HTTPPort)
	<-ctx.Done()
	logrus.Info("Stopped")

	err = httpServer.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
	}

}
