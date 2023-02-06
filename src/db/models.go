package db

import "time"

type Subscriber struct {
	ID   int64
	Name string
}

type Payment struct {
	CreatedAt  time.Time
	Amount     float64
	Subscriber Subscriber
}

type AccessKey struct {
	ID         string
	Name       string
	AccessUrl  string
	Subscriber Subscriber
}
