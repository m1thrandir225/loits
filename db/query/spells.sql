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
  THEN @book_id::VARCHAR(200) ELSE book_id END,
  
  element = CASE WHEN @element_do_update::boolean
  THEN @element::element ELSE element END,

  name = CASE WHEN @name_do_update::boolean
  THEN @name::VARCHAR(200) ELSE name END
WHERE
  id = @id 
RETURNING *;

-- name: DeleteSpell :exec
DELETE FROM spells
WHERE id = $1
RETURNING *;
