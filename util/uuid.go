package util

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func StringToUUID(s string) (pgtype.UUID, error) {
	u, err := uuid.Parse(s)
	var pgUUID pgtype.UUID

	if err != nil {
		return pgUUID, err
	}

	pgUUID.Bytes = [16]byte(u)
	pgUUID.Valid = true

	return pgUUID, nil
}
