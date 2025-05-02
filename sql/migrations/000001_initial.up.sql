create table if not exists users (
	id           varchar(50) primary key,
	email        varchar(100) unique not null,
	name         varchar(100) not null,
	created_at   timestamp not null default (now() at time zone 'utc'),
	updated_at   timestamp not null default (now() at time zone 'utc')
)
