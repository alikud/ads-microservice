package postgres

import (
	"context"
	"fmt"
	"github.com/alikud/ads-microservice/domain"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type OfferPostgres struct {
	Db *pgxpool.Pool
}

func (o OfferPostgres) Create(offer domain.Offer) (string, error) {
	var id string
	query := "INSERT INTO Offers (title, description, photo_url, price) VALUES ($1, $2, $3, $4) RETURNING ID"
	err := o.Db.QueryRow(context.Background(), query,
		offer.Title, offer.Description, pq.Array(offer.PhotoUrl), offer.Price).Scan(&id)

	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	return id, nil
}

func (o OfferPostgres) GetAll(limit int, offset int, orderBy string, orderedType string) (*[]domain.Offer, error) {
	rows, err := o.Db.Query(context.Background(),
		fmt.Sprintf("SELECT title, description, photo_url, price FROM Offers ORDER BY %s %s offset %d limit %d",
			orderBy, orderedType, offset, limit))
	if err != nil {
		log.Error(err.Error())
	}
	defer rows.Close()
	var offers []domain.Offer
	for rows.Next() {
		var offer domain.Offer
		err := rows.Scan(&offer.Title, &offer.Description, pq.Array(&offer.PhotoUrl), &offer.Price)
		if err != nil {
			log.Error(err.Error())
		}
		offers = append(offers, offer)
	}

	if err := rows.Err(); err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return &offers, nil

}

func (o OfferPostgres) GetById(id string) (*domain.Offer, error) {
	var offer domain.Offer
	query := "SELECT title, description, photo_url, price from Offers where ID=$1"
	row := o.Db.QueryRow(context.Background(), query,
		id)
	err := row.Scan(&offer.Title, &offer.Description, pq.Array(&offer.PhotoUrl), &offer.Price)

	if err != nil {
		log.Error(err.Error())
		return &offer, err
	}

	return &offer, nil
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
	return &OfferPostgres{Db: db}
}
