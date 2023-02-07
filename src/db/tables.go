package db

var QUERY = `
	create table if not exists Subscriber(
		id integer primary key,
		name varchar(128),
		payed_at date
	);
	create table if not exists AccessKey(
		id varchar(128) primary key,
		name varchar(128),
		access_url varchar(256),
		subscriber_id integer references Subscriber(id) on delete cascade
	);
`
