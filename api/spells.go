package api

import db "m1thrandir225/loits/db/sqlc"

type createSpellRequest struct {
	name    string     `json:"name" binding:"required,min=6"`
	element db.Element `json:"element" binding:"required"`
}

type spellResponse struct {
	id      string     `json:"id"`
	name    string     `json:"name"`
	element db.Element `json:"element"`
}
