CREATE TABLE admins (
  id VARCHAR(26) PRIMARY KEY,
  username VARCHAR(30) UNIQUE NOT NULL,
  password VARCHAR(30) NOT NULL,
  email VARCHAR NOT NULL,
  created_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX admins_id ON admins (id);

CREATE INDEX admins_username ON admins (username);