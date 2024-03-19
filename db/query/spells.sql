-- name: CreateSpell :exec
INSERT INTO spells (
  name,
  element,
  owner
) VALUES (
  $1,
  $2,
  $3
) RETURNING *;

-- name: GetSpellByName :one
SELECT *
FROM spells
WHERE name = $1 LIMIT 1;

-- name: GetSpellById :one
SELECT *
FROM spells
WHERE id = $1 LIMIT 1;
