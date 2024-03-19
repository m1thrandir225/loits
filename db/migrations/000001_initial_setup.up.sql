CREATE TYPE element as ENUM ('water', 'fire', 'earth', 'wind', 'electricity', 'metal');
CREATE TYPE magic_rating as ENUM('S', 'A', 'B', 'C', 'D', 'F');

CREATE TABLE magicians (
  id UUID PRIMARY KEY,
  original_name TEXT NOT NULL,
  magic_name TEXT NOT NULL UNIQUE,
  birthday TIMESTAMPTZ NOT NULL,
  magical_rating magic_rating DEFAULT 'F',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE spells (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  element element NOT NULL,
  owner UUID references magicians(id)
);
