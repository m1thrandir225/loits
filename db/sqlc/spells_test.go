package db_test

import (
	"context"
	"log"
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

func TestUpdateBookId(t *testing.T) {
	owner := createRandomMagician(t)

	book := createRandomSpellBook(t, owner)

	initialSpell := createRandomSpell(t, book)

	newBook := createRandomSpellBook(t, owner)

	arg := db.UpdateSpellParams{
		BookIDDoUpdate:  true,
		BookID:          newBook.ID,
		ID:              initialSpell.ID,
		ElementDoUpdate: false,
		NameDoUpdate:    false,
	}

	updatedSpell, err := testStore.UpdateSpell(context.Background(), arg)

	log.Fatal(updatedSpell.Element)
	require.NoError(t, err)

	require.Equal(t, initialSpell.ID, updatedSpell.ID)
	require.Equal(t, initialSpell.Name, updatedSpell.Name)
	require.Equal(t, initialSpell.Element, updatedSpell.Element)
	require.WithinDuration(t, initialSpell.CreatedAt, updatedSpell.CreatedAt, time.Second)
	require.WithinDuration(t, initialSpell.UpdatedAt, updatedSpell.UpdatedAt, time.Second)

	require.NotEqual(t, book.ID, updatedSpell.BookID)
	require.Equal(t, newBook.ID, updatedSpell.BookID)
}

func TestUpdateSpellElement(t *testing.T) {
	owner := createRandomMagician(t)

	book := createRandomSpellBook(t, owner)

	initialSpell := createRandomSpell(t, book)

	newElement := util.RandomElement()

	arg := db.UpdateSpellParams{
		ElementDoUpdate: true,
		Element:         db.Element(newElement),
		ID:              initialSpell.ID,
	}

	updatedSpell, err := testStore.UpdateSpell(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, initialSpell.ID, updatedSpell.ID)
	require.Equal(t, initialSpell.Name, updatedSpell.Name)
	require.Equal(t, initialSpell.BookID, updatedSpell.BookID)
	require.WithinDuration(t, initialSpell.CreatedAt, updatedSpell.CreatedAt, time.Second)
	require.WithinDuration(t, initialSpell.UpdatedAt, updatedSpell.UpdatedAt, time.Second)

	require.NotEqual(t, initialSpell.Element, updatedSpell.Element)
	require.Equal(t, newElement, updatedSpell.Element)
}

func TestUpdateSpellName(t *testing.T) {
	owner := createRandomMagician(t)

	book := createRandomSpellBook(t, owner)

	initialSpell := createRandomSpell(t, book)

	newName := util.RandomString(6)

	arg := db.UpdateSpellParams{
		NameDoUpdate:    true,
		Name:            newName,
		ElementDoUpdate: false,
		ID:              initialSpell.ID,
	}

	updatedSpell, err := testStore.UpdateSpell(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, initialSpell.ID, updatedSpell.ID)
	require.Equal(t, initialSpell.Element, updatedSpell.Element)
	require.Equal(t, initialSpell.BookID, updatedSpell.BookID)
	require.WithinDuration(t, initialSpell.CreatedAt, updatedSpell.CreatedAt, time.Second)
	require.WithinDuration(t, initialSpell.UpdatedAt, updatedSpell.UpdatedAt, time.Second)

	require.NotEqual(t, initialSpell.Name, updatedSpell.Name)
	require.Equal(t, newName, updatedSpell.Name)
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
