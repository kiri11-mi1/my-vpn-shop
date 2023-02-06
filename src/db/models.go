package db

import "time"

type Subscriber struct {
	ID      int64
	Name    string
	PayedAt time.Time
}

type AccessKey struct {
	ID         string
	Name       string
	AccessUrl  string
	Subscriber Subscriber
}
