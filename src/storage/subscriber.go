package storage

import (
	"my-vpn-shop/db"
)

func (p *Postgres) GetSubscribers() ([]db.Subscriber, error) {
	rows, err := p.db.Query("select * from Subscriber")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	subs := []db.Subscriber{}
	for rows.Next() {
		s := db.Subscriber{}
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}
	return subs, nil
}

func (p *Postgres) InsertSubscriber(id int64, name string) (db.Subscriber, error) {
	_, err := p.db.Exec(
		"insert into subscriber (id, name) values ($1, $2)",
		id, name,
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
