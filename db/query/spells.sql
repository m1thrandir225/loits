-- name: CreateSpell :exec
INSERT INTO spells (
  name,
  element,
  book_id
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

-- name: MoveToNewBook :exec
UPDATE spells
SET book_id = $2
WHERE id = $1
RETURNING *;

-- name: UpdateSpellElement :exec
UPDATE spells
SET element = $2
WHERE id = $1
RETURNING *;

-- name: UpdateSpell :exec
UPDATE spells
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteSpell :exec
DELETE FROM spells
WHERE id = $1
RETURNING *;
