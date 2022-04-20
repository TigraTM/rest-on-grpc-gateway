BEGIN;

CREATE SCHEMA "user";

CREATE TABLE IF NOT EXISTS "user".users
(
    id       bigserial primary key,
    name     text not null,
    email    text not null,
    password text not null
);

END;