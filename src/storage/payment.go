package storage

import (
	"my-vpn-shop/db"
	"time"
)

func (p *Postgres) InsertPayment(amount float64, sub db.Subscriber) (db.Payment, error) {
	created_at := time.Now()
	_, err := p.db.Exec(
		"insert into Payment (amount, created_at, subscriber_id) values ($1, $2, $3)",
		amount, created_at.Format("2006-01-02"), sub.ID,
	)
	if err != nil {
		return db.Payment{}, err
	}
	payment := db.Payment{
		Amount:     amount,
		Subscriber: sub,
		CreatedAt:  created_at,
	}
	return payment, nil
}
