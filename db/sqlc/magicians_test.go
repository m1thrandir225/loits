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

func createRandomMagician(t *testing.T) db.Magician {

	hashedPassword, err := util.HashPassword(util.RandomString(6))

	require.NoError(t, err)

	arg := db.CreateMagicianParams{
		Email:         util.RandomEmail(),
		Password:      hashedPassword,
		MagicName:     util.RandomMagicianName(),
		OriginalName:  util.RandomMagicianName(),
		MagicalRating: util.RandomMagicRating(),
		Birthday:      util.RandomDate(),
	}

	magician, err := testStore.CreateMagician(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, magician.ID)
	require.Equal(t, magician.MagicalRating, arg.MagicalRating)
	require.Equal(t, magician.MagicName, arg.MagicName)
	require.Equal(t, magician.OriginalName, arg.OriginalName)
	require.Equal(t, magician.Email, arg.Email)
	require.WithinDuration(t, magician.Birthday, arg.Birthday, time.Second)

	require.NotEmpty(t, magician.CreatedAt)
	require.NotEmpty(t, magician.UpdatedAt)

	return magician
}

func TestCreateMagician(t *testing.T) {
	createRandomMagician(t)
}

func TestGetMagicianByEmail(t *testing.T) {
	initialMagician := createRandomMagician(t)

	magician, err := testStore.GetMagicianByEmail(context.Background(), initialMagician.Email)

	require.NoError(t, err)

	require.Equal(t, initialMagician.ID, magician.ID)
	require.Equal(t, initialMagician.Email, magician.Email)
	require.Equal(t, initialMagician.Password, magician.Password)
	require.Equal(t, initialMagician.MagicName, magician.MagicName)
	require.Equal(t, initialMagician.OriginalName, magician.OriginalName)
	require.Equal(t, initialMagician.MagicalRating, magician.MagicalRating)
	require.WithinDuration(t, initialMagician.Birthday, magician.Birthday, time.Second)
	require.WithinDuration(t, initialMagician.CreatedAt, magician.CreatedAt, time.Second)
	require.WithinDuration(t, initialMagician.UpdatedAt, magician.UpdatedAt, time.Second)
}

func TestUpdateMagicianRating(t *testing.T) {
	initialMagician := createRandomMagician(t)

	newMagicalRating := util.RandomMagicalRatingButNotGiven(initialMagician.MagicalRating)

	arg := db.UpdateMagicianRatinParams{
		ID:                    initialMagician.ID,
		MagicalRating:         db.MagicRating(newMagicalRating),
	}

	magician, err := testStore.UpdateMagicianRatin(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, initialMagician.ID, magician.ID)
	require.Equal(t, initialMagician.Email, magician.Email)
	require.Equal(t, initialMagician.Password, magician.Password)
	require.Equal(t, initialMagician.MagicName, magician.MagicName)
	require.Equal(t, initialMagician.OriginalName, magician.OriginalName)
	require.Equal(t, db.MagicRating(newMagicalRating), magician.MagicalRating)
	require.WithinDuration(t, initialMagician.Birthday, magician.Birthday, time.Second)
	require.WithinDuration(t, initialMagician.CreatedAt, magician.CreatedAt, time.Second)
	require.WithinDuration(t, initialMagician.UpdatedAt, magician.UpdatedAt, time.Second)
}

func TestUpdateEmail(t *testing.T) {
	initialMagician := createRandomMagician(t)

	newEmail := util.RandomEmail()

	arg := db.UpdateMagicianParams{
		BirthdayDoUpdate:      false,
		OriginalNameDoUpdate:  false,
		MagicNameDoUpdate:     false,
		ID:                    initialMagician.ID,
		EmailDoUpdate:         true,
		Email:                 newEmail,
	}

	magician, err := testStore.UpdateMagician(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, initialMagician.ID, magician.ID)
	require.Equal(t, initialMagician.Password, magician.Password)
	require.Equal(t, initialMagician.MagicName, magician.MagicName)
	require.Equal(t, initialMagician.OriginalName, magician.OriginalName)
	require.Equal(t, initialMagician.MagicalRating, magician.MagicalRating)
	require.WithinDuration(t, initialMagician.Birthday, magician.Birthday, time.Second)
	require.WithinDuration(t, initialMagician.CreatedAt, magician.CreatedAt, time.Second)
	require.WithinDuration(t, initialMagician.UpdatedAt, magician.UpdatedAt, time.Second)

	require.Equal(t, newEmail, magician.Email)
	require.NotEqual(t, initialMagician.Email, magician.Email)
}

func TestUpdateOriginalName(t *testing.T) {
	initialMagician := createRandomMagician(t)

	newName := util.RandomString(8)
	arg := db.UpdateMagicianParams{
		ID:                   initialMagician.ID,
		OriginalNameDoUpdate: true,
		OriginalName:         newName,
	}

	magician, err := testStore.UpdateMagician(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, initialMagician.ID, magician.ID)
	require.Equal(t, initialMagician.Email, magician.Email)
	require.Equal(t, initialMagician.Password, magician.Password)
	require.Equal(t, initialMagician.MagicName, magician.MagicName)
	require.Equal(t, initialMagician.MagicalRating, magician.MagicalRating)
	require.WithinDuration(t, initialMagician.Birthday, magician.Birthday, time.Second)
	require.WithinDuration(t, initialMagician.CreatedAt, magician.CreatedAt, time.Second)
	require.WithinDuration(t, initialMagician.UpdatedAt, magician.UpdatedAt, time.Second)

	require.NotEqual(t, initialMagician.OriginalName, magician.OriginalName)
	require.Equal(t, newName, magician.OriginalName)
}

func TestUpdateBirthday(t *testing.T) {
	initialMagician := createRandomMagician(t)

	newBirthday := util.RandomDate()

	arg := db.UpdateMagicianParams{
		ID:               initialMagician.ID,
		BirthdayDoUpdate: true,
		Birthday:         newBirthday,
	}

	magician, err := testStore.UpdateMagician(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, initialMagician.ID, magician.ID)
	require.Equal(t, initialMagician.Email, magician.Email)
	require.Equal(t, initialMagician.Password, magician.Password)
	require.Equal(t, initialMagician.MagicName, magician.MagicName)
	require.Equal(t, initialMagician.OriginalName, magician.OriginalName)
	require.Equal(t, initialMagician.MagicalRating, magician.MagicalRating)
	require.WithinDuration(t, initialMagician.CreatedAt, magician.CreatedAt, time.Second)
	require.WithinDuration(t, initialMagician.UpdatedAt, magician.UpdatedAt, time.Second)

	require.NotEqual(t, initialMagician.Birthday, magician.Birthday)
	require.Equal(t, newBirthday, magician.Birthday)
}

func TestUpdateMagicianPassword(t *testing.T) {
	initialMagician := createRandomMagician(t)

	newPassword, err := util.HashPassword(util.RandomString(12))

	require.NoError(t, err)

	arg := db.UpdatePasswordParams{
		ID:       initialMagician.ID,
		Password: newPassword,
	}

	magician, err := testStore.UpdatePassword(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, initialMagician.ID, magician.ID)
	require.Equal(t, initialMagician.Email, magician.Email)
	require.Equal(t, newPassword, magician.Password)
	require.Equal(t, initialMagician.MagicName, magician.MagicName)
	require.Equal(t, initialMagician.OriginalName, magician.OriginalName)
	require.Equal(t, initialMagician.MagicalRating, magician.MagicalRating)
	require.WithinDuration(t, initialMagician.Birthday, magician.Birthday, time.Second)
	require.WithinDuration(t, initialMagician.CreatedAt, magician.CreatedAt, time.Second)
	require.WithinDuration(t, initialMagician.UpdatedAt, magician.UpdatedAt, time.Second)
}

func TestDeleteMagician(t *testing.T) {
	initialMagician := createRandomMagician(t)

	err := testStore.DeleteMagician(context.Background(), initialMagician.ID)

	require.NoError(t, err)

	deletedMagician, err := testStore.GetMagicianById(context.Background(), initialMagician.ID)

	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, deletedMagician)
}
