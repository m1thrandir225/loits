-- name: CreateSpell :one
INSERT INTO spells (
  name,
  element,
  book_id
) VALUES (
  $1,
  $2,
  $3
) RETURNING *;

-- name: GetSpellById :one
SELECT *
FROM spells
WHERE id = $1 LIMIT 1;

-- name: GetSpellsByBook :many
SELECT *
FROM spells
WHERE book_id = $1;

-- name: UpdateSpell :one 
UPDATE spells 
SET 
  book_id = CASE WHEN @book_id_do_update::boolean
  THEN @book_id::uuid ELSE book_id END,
  
  name = CASE WHEN @name_do_update::boolean
  THEN @name::text ELSE name END
WHERE
  id = @id 
RETURNING *;

-- name: UpdateSpellElement :one

UPDATE spells
SET element = $2
WHERE id = $1
RETURNING *;

-- name: DeleteSpell :exec
DELETE FROM spells
WHERE id = $1
RETURNING *;
