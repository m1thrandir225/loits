package db_test

import (
	"context"
	db "m1thrandir225/loits/db/sqlc"
	"m1thrandir225/loits/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomMagician(t *testing.T) db.Magician {

	hashedPassword, err := util.HashPassword(util.RandomString(6))

	require.NoError(t, err)

	arg := db.CreateMagicianParams{
		Email:         util.RandomEmail(),
		Password:      hashedPassword,
		MagicName:     util.RandomOwner(),
		OriginalName:  util.RandomOwner(),
		MagicalRating: db.MagicRating(util.RandomMagicRating()),
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
