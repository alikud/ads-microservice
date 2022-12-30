package postgres

import (
	"context"
	"fmt"
	"github.com/alikud/ads-microservice/config"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"net/url"
)

// NewPostgresDB create *pgxpool instance and return it, or error and exit with status code 1
func NewPostgresDB(config config.PostgresConfig) *pgxpool.Pool {
	ctx := context.Background()
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s&connect_timeout=%d",
		"postgres",
		url.QueryEscape(config.User),
		url.QueryEscape(config.Password),
		config.Host,
		config.Port,
		config.DBName,
		config.SSLMode,
		config.TimeOut)

	poolConfig, _ := pgxpool.ParseConfig(connStr)
	poolConfig.MinConns = config.MinConns
	poolConfig.MaxConns = config.MaxConns
	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := healthCheck(pool); err != nil {
		log.Fatal(err.Error())
	}
	return pool
}

func healthCheck(conn *pgxpool.Pool) error {
	_, err := conn.Exec(context.Background(), ";")
	if err != nil {
		return err
	}

	return nil
}
