-- +migrate Up
CREATE TABLE user (
	id integer PRIMARY KEY,
	name text UNIQUE,
	password text
);

CREATE TABLE task (
	id integer,
	content text
);

CREATE TABLE place (
	id integer, 
	placename text,
	important integer
);

-- +migrate Down
