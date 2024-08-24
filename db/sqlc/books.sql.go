// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: books.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSpellBook = `-- name: CreateSpellBook :one
INSERT INTO books (
  name,
  owner
) VALUES (
  $1,
  $2
) RETURNING id, name, owner, updated_at, created_at
`

type CreateSpellBookParams struct {
	Name  string      `json:"name"`
	Owner pgtype.UUID `json:"owner"`
}

func (q *Queries) CreateSpellBook(ctx context.Context, arg CreateSpellBookParams) (Book, error) {
	row := q.db.QueryRow(ctx, createSpellBook, arg.Name, arg.Owner)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Owner,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSpellBook = `-- name: DeleteSpellBook :exec
DELETE FROM books
WHERE id = $1 
RETURNING id, name, owner, updated_at, created_at
`

func (q *Queries) DeleteSpellBook(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteSpellBook, id)
	return err
}

const getSpellBookById = `-- name: GetSpellBookById :one
SELECT id, name, owner, updated_at, created_at FROM books
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSpellBookById(ctx context.Context, id pgtype.UUID) (Book, error) {
	row := q.db.QueryRow(ctx, getSpellBookById, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Owner,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserSpellBooks = `-- name: GetUserSpellBooks :many
SELECT id, name, owner, updated_at, created_at FROM books
WHERE owner = $1
`

func (q *Queries) GetUserSpellBooks(ctx context.Context, owner pgtype.UUID) ([]Book, error) {
	rows, err := q.db.Query(ctx, getUserSpellBooks, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Owner,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSpellBook = `-- name: UpdateSpellBook :one
UPDATE books
SET 
  name = CASE WHEN $1::boolean
  THEN $2::text ELSE name END,
  
  owner = CASE WHEN $3::boolean
  THEN $4::uuid ELSE owner END
WHERE  
  id = $5 
RETURNING id, name, owner, updated_at, created_at
`

type UpdateSpellBookParams struct {
	NameDoUpdate  bool        `json:"name_do_update"`
	Name          string      `json:"name"`
	OwnerDoUpdate bool        `json:"owner_do_update"`
	Owner         pgtype.UUID `json:"owner"`
	ID            pgtype.UUID `json:"id"`
}

func (q *Queries) UpdateSpellBook(ctx context.Context, arg UpdateSpellBookParams) (Book, error) {
	row := q.db.QueryRow(ctx, updateSpellBook,
		arg.NameDoUpdate,
		arg.Name,
		arg.OwnerDoUpdate,
		arg.Owner,
		arg.ID,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Owner,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
