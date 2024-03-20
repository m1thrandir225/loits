package api

import (
	"github.com/gin-gonic/gin"
	db "m1thrandir225/loits/db/sqlc"
	"net/http"
	"time"
)

type createMagicianRequest struct {
	Email         string         `json:"email" binding:"required,email"`
	OriginalName  string         `json:"original_name" binding:"required"`
	MagiciansName string         `json:"magicians_name" binding:"required,min=6"`
	MagicalRating db.MagicRating `json:"magical_rating" binding:"required"`
	Birthday      time.Time      `json:"birthday" binding:"required"`
}

type magicianResponse struct {
	Email         string             `json:"email"`
	OriginalName  string             `json:"original_name"`
	MagiciansName string             `json:"magicians_name"`
	Birthday      time.Time          `json:"birthday"`
	MagicalRating db.NullMagicRating `json:"magical_rating"`
	CreatedAt     time.Time          `json:"created_at"`
}

func newMagicianResponse(magician db.Magician) magicianResponse {
	return magicianResponse{
		Email:         "",
		OriginalName:  magician.OriginalName,
		MagiciansName: magician.MagicName,
		Birthday:      magician.Birthday,
		MagicalRating: magician.MagicalRating,
		CreatedAt:     magician.CreatedAt,
	}
}
func (server *Server) createMagician(ctx *gin.Context) {
	var req createMagicianRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
}
