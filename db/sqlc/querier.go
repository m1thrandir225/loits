// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateMagician(ctx context.Context, arg CreateMagicianParams) (Magician, error)
	CreateSpell(ctx context.Context, arg CreateSpellParams) (Spell, error)
	CreateSpellBook(ctx context.Context, arg CreateSpellBookParams) (Book, error)
	DeleteMagician(ctx context.Context, id pgtype.UUID) error
	DeleteSpell(ctx context.Context, id pgtype.UUID) error
	DeleteSpellBook(ctx context.Context, id pgtype.UUID) error
	GetMagicianByEmail(ctx context.Context, email string) (Magician, error)
	GetMagicianById(ctx context.Context, id pgtype.UUID) (Magician, error)
	GetSpellBookById(ctx context.Context, id pgtype.UUID) (Book, error)
	GetSpellById(ctx context.Context, id pgtype.UUID) (Spell, error)
	GetSpellsByBook(ctx context.Context, bookID pgtype.UUID) ([]Spell, error)
	GetUserSpellBooks(ctx context.Context, owner pgtype.UUID) ([]Book, error)
	UpdateMagicalRating(ctx context.Context, arg UpdateMagicalRatingParams) (Magician, error)
	UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (Magician, error)
	UpdateSpell(ctx context.Context, arg UpdateSpellParams) (Spell, error)
}

var _ Querier = (*Queries)(nil)
