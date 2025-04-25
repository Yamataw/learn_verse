CREATE EXTENSION IF NOT EXISTS ulid;

CREATE TABLE  IF NOT EXISTS users (
                       id            ULID PRIMARY KEY DEFAULT gen_ulid(),
                       username TEXT    NOT NULL UNIQUE,
                       email text NOT NULL UNIQUE,
                       password_hash TEXT NOT NULL
);
