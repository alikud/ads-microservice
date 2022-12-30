package postgres

import (
	"github.com/alikud/ads-microservice/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type OfferPostgres struct {
	db *pgxpool.Pool
}

func (o OfferPostgres) Create(offer domain.Offer) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (o OfferPostgres) GetAll() ([]domain.Offer, error) {
	//TODO implement me
	panic("implement me")
}

func (o OfferPostgres) GetById(id string) (domain.Offer, error) {
	//TODO implement me
	panic("implement me")
}

func (o OfferPostgres) Update(id string, offer domain.Offer) error {
	//TODO implement me
	panic("implement me")
}

func (o OfferPostgres) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewOfferPostgres(db *pgxpool.Pool) *OfferPostgres {
	return &OfferPostgres{db: db}
}
