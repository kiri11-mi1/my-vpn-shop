package storage

import (
	"my-vpn-shop/db"
	"time"
)

func (s *SQL) GetSubscribers() ([]db.Subscriber, error) {
	rows, err := s.db.Query("select * from subscriber")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	subs := []db.Subscriber{}
	for rows.Next() {
		sub := db.Subscriber{}
		if err := rows.Scan(&sub.ID, &sub.Name, &sub.PayedAt); err != nil {
			return nil, err
		}
		subs = append(subs, sub)
	}
	return subs, nil
}

func (s *SQL) InsertSubscriber(id int64, name string) (db.Subscriber, error) {
	payedAt := time.Now()
	_, err := s.db.Exec(
		"insert into subscriber (id, name, payed_at) values ($1, $2, $3)",
		id, name, payedAt.Format("2006-01-02"),
	)
	if err != nil {
		return db.Subscriber{}, err
	}
	sub := db.Subscriber{ID: id, Name: name, PayedAt: payedAt}
	return sub, nil
}

func (s *SQL) DeleteSubscriber(id int64) error {
	_, err := s.db.Exec(
		"delete from subscriber where id=$1",
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQL) GetSubByID(id int64) (db.Subscriber, error) {
	rows, err := s.db.Query("select * from subscriber where id=$1", id)
	if err != nil {
		return db.Subscriber{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return db.Subscriber{}, ErrSubNotFound
	}
	sub := db.Subscriber{}
	if err := rows.Scan(&sub.ID, &sub.Name, &sub.PayedAt); err != nil {
		return db.Subscriber{}, err
	}
	return sub, nil
}

func (s *SQL) UpdateSubscriberPayedAt(id int64) error {
	payedAt := time.Now()
	_, err := s.db.Exec(
		"update subscriber set payed_at = $1 where id=$2",
		payedAt.Format("2006-01-02"), id,
	)
	if err != nil {
		return err
	}
	return nil
}
