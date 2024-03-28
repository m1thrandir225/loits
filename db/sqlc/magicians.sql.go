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

const updateMagician = `-- name: UpdateMagician :one
UPDATE magicians
SET 
  email = CASE WHEN $1::boolean
  THEN $2::text ELSE email END,

  original_name = CASE WHEN $3::boolean
  THEN $4::text ELSE original_name END,

  magic_name = CASE WHEN $5::boolean
  THEN $6::text ELSE magic_name END,

  birthday = CASE WHEN $7::boolean
  THEN  cast($8 as "timestamptz") ELSE birthday END
WHERE 
  id = $9 
RETURNING id, email, password, original_name, magic_name, birthday, magical_rating, updated_at, created_at
`

type UpdateMagicianParams struct {
	EmailDoUpdate        bool        `json:"email_do_update"`
	Email                string      `json:"email"`
	OriginalNameDoUpdate bool        `json:"original_name_do_update"`
	OriginalName         string      `json:"original_name"`
	MagicNameDoUpdate    bool        `json:"magic_name_do_update"`
	MagicName            string      `json:"magic_name"`
	BirthdayDoUpdate     bool        `json:"birthday_do_update"`
	Birthday             time.Time   `json:"birthday"`
	ID                   pgtype.UUID `json:"id"`
}

func (q *Queries) UpdateMagician(ctx context.Context, arg UpdateMagicianParams) (Magician, error) {
	row := q.db.QueryRow(ctx, updateMagician,
		arg.EmailDoUpdate,
		arg.Email,
		arg.OriginalNameDoUpdate,
		arg.OriginalName,
		arg.MagicNameDoUpdate,
		arg.MagicName,
		arg.BirthdayDoUpdate,
		arg.Birthday,
		arg.ID,
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

const updateMagicianRatin = `-- name: UpdateMagicianRatin :one
UPDATE magicians
SET magical_rating = $2
WHERE id = $1
RETURNING id, email, password, original_name, magic_name, birthday, magical_rating, updated_at, created_at
`

type UpdateMagicianRatinParams struct {
	ID            pgtype.UUID `json:"id"`
	MagicalRating MagicRating `json:"magical_rating"`
}

func (q *Queries) UpdateMagicianRatin(ctx context.Context, arg UpdateMagicianRatinParams) (Magician, error) {
	row := q.db.QueryRow(ctx, updateMagicianRatin, arg.ID, arg.MagicalRating)
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
