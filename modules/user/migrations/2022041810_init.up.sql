BEGIN;

CREATE SCHEMA "user";

CREATE TABLE IF NOT EXISTS "user".users
(
    id              bigserial primary key,
    created_at      timestamp not null default now(),
    updated_at      timestamp not null default now(),
    name            text  not null,
    email           text  not null,
    password_hash   bytea not null,

    unique(email)
);

END;