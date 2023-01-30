CREATE TABLE users (
	id bigserial not null primary key,
	email varchar not null unique,
	encrypted_pass varchar not null
);
