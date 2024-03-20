-- name: CreateSpellBook :one
INSERT INTO books (
  name,
  owner
) VALUES (
  $1,
  $2
) RETURNING *;

-- name: GetSpellBookById :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;


-- name: GetUserSpellBooks :many
SELECT * FROM books
WHERE owner = $1;


-- name: DeleteSpellBook :exec
DELETE FROM books
WHERE id = $1 
RETURNING *;
