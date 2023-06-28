package postgres

import (
	"github.com/alikud/ads-microservice/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

//Base crud interface
type Offer interface {
	Create(offer domain.Offer) (string, error)
	GetAll(limit int, offset int, orderBy string, orderedType string) ([]domain.Offer, error)
	GetById(id string) (domain.Offer, error)
	Update(id string, offer domain.Offer) error
	Delete(id string) error
}

type Repository struct {
	Offer
}

func NewRepository(Db *pgxpool.Pool) *Repository {
	return &Repository{Offer: NewOfferPostgres(Db)}
}
