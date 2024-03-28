package api

import (
	db "m1thrandir225/loits/db/sqlc"
	"time"

	"github.com/gin-gonic/gin"
)

type getMagicianByIdRequest struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type createMagicianRequest struct {
	OriginalName string `json:"original_name" binding:"required"`
	MagicName   string `json:"magic_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Birthday time.Time `json:"birthday" binding:"required"`
}

type updateMagicianRequest struct {
	Email string `json:"email" binding:"required"`
	EmailDoUpdate bool `json:"email_do_update"`
	OriginalName string `json:"original_name"`
	OriginalNameDoUpdate bool `json:"original_name_do_update"`
	MagicName string `json:"magic_name"`
	MagicNameDoUpdate bool `json:"magic_name_do_update"`
	Birthday time.Time `json:"birthday"`
	BirthdayDoUpdate bool `json:"birthday_do_update"`
}

type updateMagicianMagicRatingRequest struct {
	MagicRating db.MagicRating `json:"magic_rating" binding:"required"`
}


type changePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

//TODO: implement auth
func (server *Server) register(ctx *gin.Context) {}

//TODO: implement auth
func (server *Server) login(ctx *gin.Context) {}


//TODO: implement auth
func (server *Server) changePassword(ctx *gin.Context) {
	var uriBind getMagicianByIdRequest
	var req changePasswordRequest
}

func (server *Server) getMagician(ctx *gin.Context) {
	var uriBind getMagicianByIdRequest

}

func (server *Server) deleteMagician(ctx *gin.Context) {
	var uriBind getMagicianByIdRequest
}

func (server *Server) updateMagician(ctx *gin.Context) {
	var uriBind getMagicianByIdRequest
	var req updateMagicianRequest
}

func (server *Server) updateMagicianMagicRating(ctx *gin.Context) {
	var uriBind getMagicianByIdRequest
	var req updateMagicianMagicRatingRequest
}
