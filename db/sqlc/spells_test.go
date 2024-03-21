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

func createRandomSpell(t *testing.T, spellBook db.Book) db.Spell {
	arg := db.CreateSpellParams{
		Name:    util.RandomString(6),
		Element: util.RandomElement(),
		BookID:  spellBook.ID,
	}

	spell, err := testStore.CreateSpell(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, spell.BookID, arg.BookID)
	require.Equal(t, spell.Name, arg.Name)
	require.Equal(t, spell.Element, arg.Element)
	require.NotEmpty(t, spell.ID)
	require.NotEmpty(t, spell.CreatedAt)
	require.NotEmpty(t, spell.UpdatedAt)

	return spell
}

func TestCreateSpell(t *testing.T) {
	owner := createRandomMagician(t)
	book := createRandomSpellBook(t, owner)

	createRandomSpell(t, book)
}

func TestGetSpellByName(t *testing.T) {
	owner := createRandomMagician(t)
	book := createRandomSpellBook(t, owner)

	initialSpell := createRandomSpell(t, book)

	spell, err := testStore.GetSpellByName(context.Background(), initialSpell.Name)

	require.NoError(t, err)

	require.Equal(t, spell.ID, initialSpell.ID)
	require.Equal(t, spell.Name, initialSpell.Name)
	require.Equal(t, spell.BookID, initialSpell.BookID)
	require.Equal(t, spell.Element, initialSpell.Element)

	require.WithinDuration(t, spell.CreatedAt, initialSpell.CreatedAt, time.Second)
	require.WithinDuration(t, spell.UpdatedAt, initialSpell.UpdatedAt, time.Second)
}

func TestGetSpellById(t *testing.T) {
	owner := createRandomMagician(t)
	book := createRandomSpellBook(t, owner)

	initialSpell := createRandomSpell(t, book)

	spell, err := testStore.GetSpellById(context.Background(), initialSpell.ID)

	require.NoError(t, err)

	require.Equal(t, spell.ID, initialSpell.ID)
	require.Equal(t, spell.Name, initialSpell.Name)
	require.Equal(t, spell.BookID, initialSpell.BookID)
	require.Equal(t, spell.Element, initialSpell.Element)

	require.WithinDuration(t, spell.CreatedAt, initialSpell.CreatedAt, time.Second)
	require.WithinDuration(t, spell.UpdatedAt, initialSpell.UpdatedAt, time.Second)
}
func TestGetSpellsByBook(t *testing.T) {
	owner := createRandomMagician(t)
	book := createRandomSpellBook(t, owner)

	for i := 0; i < 10; i++ {
		createRandomSpell(t, book)
	}

	spells, err := testStore.GetSpellsByBook(context.Background(), book.ID)
	require.NoError(t, err)

	for i := 0; i < len(spells); i++ {
		spell := spells[i]

		require.NotEmpty(t, spell.ID)
		require.NotEmpty(t, spell.Name)
		require.Equal(t, spell.BookID, book.ID)
		require.NotEmpty(t, spell.Element)

		require.NotEmpty(t, spell.CreatedAt)
		require.NotEmpty(t, spell.UpdatedAt)
	}

}

func TestMoveToNewBook(t *testing.T) {
	owner := createRandomMagician(t)

	book := createRandomSpellBook(t, owner)

	spell := createRandomSpell(t, book)

	newBook := createRandomSpellBook(t, owner)

	arg := db.MoveToNewBookParams{
		ID:     spell.ID,
		BookID: newBook.ID,
	}

	movedSpell, err := testStore.MoveToNewBook(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, spell.ID, movedSpell.ID)
	require.Equal(t, spell.Name, movedSpell.Name)
	require.Equal(t, spell.Element, movedSpell.Element)
	require.Equal(t, movedSpell.BookID, newBook.ID)

	require.WithinDuration(t, spell.CreatedAt, movedSpell.CreatedAt, time.Second)
	require.WithinDuration(t, spell.UpdatedAt, movedSpell.UpdatedAt, time.Second)
}

func TestUpdateSpellElement(t *testing.T) {
	owner := createRandomMagician(t)

	book := createRandomSpellBook(t, owner)

	initialSpell := createRandomSpell(t, book)

	arg := db.UpdateSpellElementParams{
		ID:      initialSpell.ID,
		Element: util.RandomElement(),
	}

	updatedSpell, err := testStore.UpdateSpellElement(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, updatedSpell.ID, initialSpell.ID)
	require.NotEqual(t, updatedSpell.Element, initialSpell.Element)
	require.Equal(t, updatedSpell.Name, initialSpell.Name)
	require.Equal(t, updatedSpell.BookID, initialSpell.BookID)
	require.WithinDuration(t, updatedSpell.CreatedAt, initialSpell.CreatedAt, time.Second)
	require.WithinDuration(t, updatedSpell.UpdatedAt, initialSpell.UpdatedAt, time.Second)
}

func TestUpdateSpellName(t *testing.T) {
	owner := createRandomMagician(t)

	book := createRandomSpellBook(t, owner)

	initialSpell := createRandomSpell(t, book)

	arg := db.UpdateSpellParams{
		ID:   initialSpell.ID,
		Name: util.RandomString(6),
	}

	updatedSpell, err := testStore.UpdateSpell(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, updatedSpell.ID, initialSpell.ID)
	require.Equal(t, updatedSpell.Element, initialSpell.Element)
	require.NotEqual(t, updatedSpell.Name, initialSpell.Name)
	require.Equal(t, updatedSpell.BookID, initialSpell.BookID)
	require.WithinDuration(t, updatedSpell.CreatedAt, initialSpell.CreatedAt, time.Second)
	require.WithinDuration(t, updatedSpell.UpdatedAt, initialSpell.UpdatedAt, time.Second)
}

func TestDeleteSpell(t *testing.T) {
	owner := createRandomMagician(t)

	book := createRandomSpellBook(t, owner)

	spell := createRandomSpell(t, book)

	err := testStore.DeleteSpell(context.Background(), spell.ID)

	require.NoError(t, err)

	deletedSpell, err := testStore.GetSpellById(context.Background(), spell.ID)

	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, deletedSpell)
}
