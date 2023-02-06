package storage

import (
	"my-vpn-shop/db"
	"time"
)

func (p *Postgres) GetSubscribers() ([]db.Subscriber, error) {
	rows, err := p.db.Query("select * from subscriber")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	subs := []db.Subscriber{}
	for rows.Next() {
		s := db.Subscriber{}
		if err := rows.Scan(&s.ID, &s.Name, &s.PayedAt); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}
	return subs, nil
}

func (p *Postgres) InsertSubscriber(id int64, name string) (db.Subscriber, error) {
	payedAt := time.Now()
	_, err := p.db.Exec(
		"insert into subscriber (id, name, payed_at) values ($1, $2, $3)",
		id, name, payedAt.Format("2006-01-02"),
	)
	if err != nil {
		return db.Subscriber{}, err
	}
	sub := db.Subscriber{ID: id, Name: name}
	return sub, nil
}

func (p *Postgres) DeleteSubscriber(id int64) error {
	_, err := p.db.Exec(
		"delete from subscriber where id=$1",
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
