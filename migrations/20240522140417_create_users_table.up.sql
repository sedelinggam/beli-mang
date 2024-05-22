CREATE TABLE users (
  id VARCHAR(26) PRIMARY KEY,
  username VARCHAR(30) UNIQUE NOT NULL,
  password VARCHAR(30) NOT NULL,
  email VARCHAR NOT NULL,
  created_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX users_id ON users (id);

CREATE INDEX users_username ON users (username);