-- name: CreateMagician :exec
INSERT INTO magicians (
  email,
  password,
  original_name,
  magic_name,
  birthday,
  magical_rating
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
) RETURNING *;
-- name: GetMagicianByEmail :one
SELECT *
FROM magicians
WHERE email = $1 LIMIT 1;

-- name: GetMagicianById :one 
SELECT *
FROM magicians
WHERE id = $1 LIMIT 1;


-- name: UpdateMagicalRating :one
UPDATE magicians
SET magical_rating = $2
WHERE id = $1
RETURNING *;

-- name: UpdatePassword :one
UPDATE magicians
SET password = $2
WHERE id = $1
RETURNING *;

-- name: DeleteMagician :one
DELETE FROM magicians
WHERE id = $1
RETURNING *;
