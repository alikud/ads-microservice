package service

import (
	"fmt"
	"github.com/alikud/ads-microservice/domain"
	"github.com/alikud/ads-microservice/pkg/repository/postgres"
	"strings"
)

type OfferService struct {
	repo *postgres.Repository
}

func (o OfferService) GetAll(limit int, offset int, orderBy string) ([]domain.Offer, error) {

	var orderData []string
	if orderBy == "" {
		orderData = []string{"created_at", "ASC"}
	} else {
		orderData = strings.Split(orderBy, ":")
	}

	fmt.Println(orderData[0], orderData[1])

	offers, err := o.repo.GetAll(limit, offset+limit, orderData[0], orderData[1])

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
