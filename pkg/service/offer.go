package service

import (
	"fmt"
	"github.com/alikud/ads-microservice/domain"
	"github.com/alikud/ads-microservice/pkg/repository/postgres"
)

type OfferService struct {
	repo *postgres.Repository
}

func (o OfferService) GetAll(limit int) ([]domain.Offer, error) {
	offers, err := o.repo.GetAll(limit)
	if err != nil {
		fmt.Println(err)
	}
	return offers, nil
}

func (o OfferService) GetById(offerId string) (domain.Offer, error) {
	//TODO implement me
	panic("implement me")
}

func (o OfferService) Create(offer domain.Offer) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewOfferService(repo *postgres.Repository) *OfferService {
	return &OfferService{repo: repo}
}
