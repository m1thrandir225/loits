// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: spells.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createSpell = `-- name: CreateSpell :exec
INSERT INTO spells (
  name,
  element,
  book_id
) VALUES (
  $1,
  $2,
  $3
) RETURNING id, name, element, book_id, updated_at, created_at
`

type CreateSpellParams struct {
	Name    string      `json:"name"`
	Element Element     `json:"element"`
	BookID  pgtype.UUID `json:"book_id"`
}

func (q *Queries) CreateSpell(ctx context.Context, arg CreateSpellParams) error {
	_, err := q.db.Exec(ctx, createSpell, arg.Name, arg.Element, arg.BookID)
	return err
}

const deleteSpell = `-- name: DeleteSpell :exec
DELETE FROM spells
WHERE id = $1
RETURNING id, name, element, book_id, updated_at, created_at
`

func (q *Queries) DeleteSpell(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteSpell, id)
	return err
}

const getSpellById = `-- name: GetSpellById :one
SELECT id, name, element, book_id, updated_at, created_at
FROM spells
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSpellById(ctx context.Context, id uuid.UUID) (Spell, error) {
	row := q.db.QueryRow(ctx, getSpellById, id)
	var i Spell
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Element,
		&i.BookID,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getSpellByName = `-- name: GetSpellByName :one
SELECT id, name, element, book_id, updated_at, created_at
FROM spells
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetSpellByName(ctx context.Context, name string) (Spell, error) {
	row := q.db.QueryRow(ctx, getSpellByName, name)
	var i Spell
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Element,
		&i.BookID,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const moveToNewBook = `-- name: MoveToNewBook :exec
UPDATE spells
SET book_id = $2
WHERE id = $1
RETURNING id, name, element, book_id, updated_at, created_at
`

type MoveToNewBookParams struct {
	ID     uuid.UUID   `json:"id"`
	BookID pgtype.UUID `json:"book_id"`
}

func (q *Queries) MoveToNewBook(ctx context.Context, arg MoveToNewBookParams) error {
	_, err := q.db.Exec(ctx, moveToNewBook, arg.ID, arg.BookID)
	return err
}

const updateSpell = `-- name: UpdateSpell :exec
UPDATE spells
SET name = $2
WHERE id = $1
RETURNING id, name, element, book_id, updated_at, created_at
`

type UpdateSpellParams struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (q *Queries) UpdateSpell(ctx context.Context, arg UpdateSpellParams) error {
	_, err := q.db.Exec(ctx, updateSpell, arg.ID, arg.Name)
	return err
}

const updateSpellElement = `-- name: UpdateSpellElement :exec
UPDATE spells
SET element = $2
WHERE id = $1
RETURNING id, name, element, book_id, updated_at, created_at
`

type UpdateSpellElementParams struct {
	ID      uuid.UUID `json:"id"`
	Element Element   `json:"element"`
}

func (q *Queries) UpdateSpellElement(ctx context.Context, arg UpdateSpellElementParams) error {
	_, err := q.db.Exec(ctx, updateSpellElement, arg.ID, arg.Element)
	return err
}
