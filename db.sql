create table AccessKey(
    id varchar(128) primary key,
    name varchar(128),
    access_url varchar(256),
    subscriber_id integer references Subscriber(id) on delete cascade
);

create table Subscriber(
    id integer primary key,
    name varchar(128),
    payed_at date
);
