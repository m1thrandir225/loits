-- name: CreateMagician :one
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

-- name: UpdateMagician :one
UPDATE magicians
SET 
  email = CASE WHEN @email_do_update::boolean
  THEN @email::text ELSE email END,

  original_name = CASE WHEN @original_name_do_update::boolean
  THEN @original_name::text ELSE original_name END,

  magic_name = CASE WHEN @magic_name_do_update::boolean
  THEN @magic_name::text ELSE magic_name END,

  birthday = CASE WHEN @birthday_do_update::boolean
  THEN  cast(@birthday as "timestamptz") ELSE birthday END
WHERE 
  id = @id 
RETURNING *;

-- name: UpdateMagicianRatin :one 
UPDATE magicians
SET magical_rating = $2
WHERE id = $1
RETURNING *;

-- name: UpdatePassword :one
UPDATE magicians
SET password = $2
WHERE id = $1
RETURNING *;

-- name: DeleteMagician :exec
DELETE FROM magicians
WHERE id = $1
RETURNING *;
