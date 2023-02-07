package storage

import "my-vpn-shop/db"

func (p *SQL) InsertAccessKey(id, name, accessUrl string, sub db.Subscriber) (db.AccessKey, error) {
	_, err := p.db.Exec(
		"insert into accesskey (id, name, access_url, subscriber_id) values ($1, $2, $3, $4)",
		id, name, accessUrl, sub.ID,
	)
	if err != nil {
		return db.AccessKey{}, err
	}
	key := db.AccessKey{ID: id, Name: name, AccessUrl: accessUrl, Subscriber: sub}
	return key, nil
}

func (p *SQL) GetKeyBySubId(id int64) (db.AccessKey, error) {
	rows, err := p.db.Query("select * from accesskey where subscriber_id=$1", id)
	if err != nil {
		return db.AccessKey{}, err
	}
	rows.Next()
	defer rows.Close()
	key := db.AccessKey{}
	if err := rows.Scan(&key.ID, &key.Name, &key.AccessUrl, &key.Subscriber.ID); err != nil {
		return db.AccessKey{}, err
	}
	return key, nil
}
