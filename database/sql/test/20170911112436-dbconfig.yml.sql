-- +migrate Up
CREATE TABLE users (
	id integer PRIMARY KEY,
	user_name text UNIQUE,
	password text
);

CREATE TABLE tasks (
	id integer,
	content text
	place_id integer 
	important integer
);

CREATE TABLE places (
	id integer, 
	place_id integer,
	placename text,
	important integer
);

-- +migrate Down
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS task;
DROP TABLE IF EXISTS place;
