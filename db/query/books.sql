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

-- name: UpdateSpellBook :one 
UPDATE books
SET 
  name = CASE WHEN @name_do_update::boolean
  THEN @name::text ELSE name END,
  
  owner = CASE WHEN @owner_do_update::boolean
  THEN @owner::uuid ELSE owner END
WHERE  
  id = @id 
RETURNING *;

-- name: DeleteSpellBook :exec
DELETE FROM books
WHERE id = $1 
RETURNING *;
