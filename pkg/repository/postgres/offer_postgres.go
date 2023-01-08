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
	//TODO implement me
	panic("implement me")
}

func (o OfferPostgres) GetAll(limit int, offset int, orderBy string, orderedType string) ([]domain.Offer, error) {
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
			fmt.Println(err)
		}
		offers = append(offers, offer)
	}

	if err := rows.Err(); err != nil {
		log.Error(err.Error())
	}
	return offers, nil

}

func (o OfferPostgres) GetById(id string, fields ...string) (domain.Offer, error) {
	var offer domain.Offer
	q := "SELECT title, description, photo_url, price "

	for _, j := range fields {
		q += fmt.Sprintf(", %s", j)
	}
	q += "FROM Offers WHERE id=$1;"
	row := o.Db.QueryRow(context.Background(), q,
		id)
	err := row.Scan(&offer.Title, &offer.Description, pq.Array(&offer.PhotoUrl), &offer.Price)

	if err != nil {
		log.Error(err.Error())
		return offer, err
	}

	return offer, nil
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
