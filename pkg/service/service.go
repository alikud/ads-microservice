package service

import (
	"github.com/alikud/ads-microservice/pkg/repository/postgres"
)

type Service struct {
}

func NewService(repo *postgres.Repository) *Service {
	return &Service{}
}
