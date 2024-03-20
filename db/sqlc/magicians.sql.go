// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: magicians.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMagician = `-- name: CreateMagician :one
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
) RETURNING id, email, password, original_name, magic_name, birthday, magical_rating, updated_at, created_at
`

type CreateMagicianParams struct {
	Email         string      `json:"email"`
	Password      string      `json:"password"`
	OriginalName  string      `json:"original_name"`
	MagicName     string      `json:"magic_name"`
	Birthday      time.Time   `json:"birthday"`
	MagicalRating MagicRating `json:"magical_rating"`
}

func (q *Queries) CreateMagician(ctx context.Context, arg CreateMagicianParams) (Magician, error) {
	row := q.db.QueryRow(ctx, createMagician,
		arg.Email,
		arg.Password,
		arg.OriginalName,
		arg.MagicName,
		arg.Birthday,
		arg.MagicalRating,
	)
	var i Magician
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.OriginalName,
		&i.MagicName,
		&i.Birthday,
		&i.MagicalRating,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteMagician = `-- name: DeleteMagician :exec
DELETE FROM magicians
WHERE id = $1
RETURNING id, email, password, original_name, magic_name, birthday, magical_rating, updated_at, created_at
`

func (q *Queries) DeleteMagician(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteMagician, id)
	return err
}

const getMagicianByEmail = `-- name: GetMagicianByEmail :one
SELECT id, email, password, original_name, magic_name, birthday, magical_rating, updated_at, created_at
FROM magicians
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetMagicianByEmail(ctx context.Context, email string) (Magician, error) {
	row := q.db.QueryRow(ctx, getMagicianByEmail, email)
	var i Magician
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.OriginalName,
		&i.MagicName,
		&i.Birthday,
		&i.MagicalRating,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getMagicianById = `-- name: GetMagicianById :one
SELECT id, email, password, original_name, magic_name, birthday, magical_rating, updated_at, created_at
FROM magicians
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMagicianById(ctx context.Context, id pgtype.UUID) (Magician, error) {
	row := q.db.QueryRow(ctx, getMagicianById, id)
	var i Magician
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.OriginalName,
		&i.MagicName,
		&i.Birthday,
		&i.MagicalRating,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateMagicalRating = `-- name: UpdateMagicalRating :one
UPDATE magicians
SET magical_rating = $2
WHERE id = $1
RETURNING id, email, password, original_name, magic_name, birthday, magical_rating, updated_at, created_at
`

type UpdateMagicalRatingParams struct {
	ID            pgtype.UUID `json:"id"`
	MagicalRating MagicRating `json:"magical_rating"`
}

func (q *Queries) UpdateMagicalRating(ctx context.Context, arg UpdateMagicalRatingParams) (Magician, error) {
	row := q.db.QueryRow(ctx, updateMagicalRating, arg.ID, arg.MagicalRating)
	var i Magician
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.OriginalName,
		&i.MagicName,
		&i.Birthday,
		&i.MagicalRating,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updatePassword = `-- name: UpdatePassword :one
UPDATE magicians
SET password = $2
WHERE id = $1
RETURNING id, email, password, original_name, magic_name, birthday, magical_rating, updated_at, created_at
`

type UpdatePasswordParams struct {
	ID       pgtype.UUID `json:"id"`
	Password string      `json:"password"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (Magician, error) {
	row := q.db.QueryRow(ctx, updatePassword, arg.ID, arg.Password)
	var i Magician
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.OriginalName,
		&i.MagicName,
		&i.Birthday,
		&i.MagicalRating,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
