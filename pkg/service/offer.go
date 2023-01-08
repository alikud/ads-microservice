package service

import (
	"github.com/alikud/ads-microservice/domain"
	"github.com/alikud/ads-microservice/pkg/repository/postgres"
	"strings"
)

type OfferService struct {
	repo *postgres.Repository
}

func (o OfferService) GetAll(limit int, offset int, orderBy string) (*[]domain.Offer, error) {

	var orderData []string
	if orderBy == "" {
		orderData = []string{"created_at", "ASC"}
	} else {
		orderData = strings.Split(orderBy, ":")
	}

	offers, err := o.repo.GetAll(limit, offset, orderData[0], orderData[1])

	if err != nil {
		return nil, err
	}
	return offers, nil
}

func (o OfferService) GetById(offerId string) (*domain.Offer, error) {
	return o.repo.GetById(offerId)
}

func (o OfferService) Create(offer domain.Offer) (string, error) {
	id, err := o.repo.Create(offer)
	if err != nil {
		return "", err
	}
	return id, nil
}

func NewOfferService(repo *postgres.Repository) *OfferService {
	return &OfferService{repo: repo}
}
