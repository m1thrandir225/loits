package db_test

import (
	"context"
	db "m1thrandir225/loits/db/sqlc"
	"m1thrandir225/loits/util"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func createRandomSpellBook(t *testing.T, owner db.Magician) db.Book {
	arg := db.CreateSpellBookParams{
		Name:  util.RandomString(6),
		Owner: owner.ID,
	}

	spellBook, err := testStore.CreateSpellBook(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, arg.Name, spellBook.Name)
	require.Equal(t, arg.Owner, spellBook.Owner)
	require.NotEmpty(t, spellBook.CreatedAt)
	require.NotEmpty(t, spellBook.ID)
	require.NotEmpty(t, spellBook.UpdatedAt)

	return spellBook
}

func TestCreateSpellBook(t *testing.T) {
	owner := createRandomMagician(t)
	createRandomSpellBook(t, owner)
}
func TestGetSpellBookById(t *testing.T) {
	owner := createRandomMagician(t)
	initialSpellBook := createRandomSpellBook(t, owner)

	spellBook, err := testStore.GetSpellBookById(context.Background(), initialSpellBook.ID)
	require.NoError(t, err)

	require.Equal(t, initialSpellBook.ID, spellBook.ID)
	require.Equal(t, initialSpellBook.Name, spellBook.Name)
	require.Equal(t, initialSpellBook.Owner, spellBook.Owner)
	require.WithinDuration(t, initialSpellBook.CreatedAt, spellBook.CreatedAt, time.Second)
	require.WithinDuration(t, initialSpellBook.UpdatedAt, spellBook.UpdatedAt, time.Second)
}

func TestGetUserSpellBooks(t *testing.T) {
	owner := createRandomMagician(t)

	for i := 0; i < 10; i++ {
		createRandomSpellBook(t, owner)
	}

	spellBoooks, err := testStore.GetUserSpellBooks(context.Background(), owner.ID)

	require.NoError(t, err)

	for i := 0; i < len(spellBoooks); i++ {
		spellBook := spellBoooks[i]

		require.Equal(t, spellBook.Owner, owner.ID)
		require.NotEmpty(t, spellBook.ID)
		require.NotEmpty(t, spellBook.Name)
		require.NotEmpty(t, spellBook.UpdatedAt)
		require.NotEmpty(t, spellBook.CreatedAt)
	}
}

func TestDeleteSpellBook(t *testing.T) {
	owner := createRandomMagician(t)
	initialSpellBook := createRandomSpellBook(t, owner)

	err := testStore.DeleteSpellBook(context.Background(), initialSpellBook.ID)

	require.NoError(t, err)

	spellBook, err := testStore.GetSpellBookById(context.Background(), initialSpellBook.ID)

	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, spellBook)
}
