package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"os"
)

// Specification Basic specification for start app with default values
type Specification struct {
	Debug bool   `default:"True"`
	Port  string `default:"8080"`
}

// InitSpecConfig Read specification from ENV
func InitSpecConfig() Specification {
	var s Specification
	err := envconfig.Process("spec", &s)
	if err != nil {
		log.Printf("Error with init specification config, start default: debug:True, port: 8080 %s", err.Error())
	}
	return s
}

// PostgresConfig for connection to PostgreSQL db
type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
	MinConns int32 `default:"10"`
	MaxConns int32 `default:"50"`
	TimeOut  int   `default:"15"`
}

// InitPostgresConfig read psql config from ENV , if err exit with code 1
func InitPostgresConfig(debug bool) PostgresConfig {
	if debug != true {
		var p PostgresConfig
		err := envconfig.Process("db", &p)
		if err != nil {
			log.Fatalf("Error with init psql config, %s", err.Error())
		}
		return p
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return PostgresConfig{User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
		MinConns: 10,
		MaxConns: 20,
		TimeOut:  5,
	}
}
