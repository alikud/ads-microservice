package service

import (
	"github.com/alikud/ads-microservice/domain"
	"github.com/alikud/ads-microservice/pkg/repository/postgres"
)

type Offer interface {
	GetAll(limit int, offset int, orderBy string) ([]domain.Offer, error)
	GetById(offerId string) (domain.Offer, error)
	Create(offer domain.Offer) (string, error)
}

type Service struct {
	Offer
}

func NewService(repo *postgres.Repository) *Service {
	return &Service{Offer: NewOfferService(repo)}
}
