CREATE TYPE element as ENUM ('water', 'fire', 'earth', 'wind', 'electricity', 'metal');
CREATE TYPE magic_rating as ENUM('S', 'A', 'B', 'C', 'D', 'F');

CREATE TABLE magicians (
  id UUID PRIMARY KEY,
  email VARCHAR(64) NOT NULL UNIQUE,
  password VARCHAR(32) NOT NULL,
  original_name VARCHAR(32) NOT NULL,
  magic_name VARCHAR(32) NOT NULL,
  birthday TIMESTAMPTZ NOT NULL,
  magical_rating magic_rating DEFAULT 'F',
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE books (
  id UUID PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  owner UUID REFERENCES magicians(id) ON DELETE CASCADE, 
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE spells (
 id UUID PRIMARY KEY,
  name VARCHAR(64) NOT NULL UNIQUE,
  element element NOT NULL,
  book_id UUID REFERENCES books(id) ON DELETE CASCADE,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

