package api

import (
	db "m1thrandir225/loits/db/sqlc"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type getMagicianByIdRequest struct {
	ID pgtype.UUID `uri:"id" binding:"required,min=1"`
}

type createMagicianRequest struct {
	OriginalName string `json:"original_name" binding:"required"`
	MagicName   string `json:"magic_name" binding:"required"`
	Email       string `json:"email" binding:"required, email"`
	Password   string `json:"password" binding:"required"`
	Birthday time.Time `json:"birthday" binding:"required"`
	MagicRating db.MagicRating `json:"magic_rating", binding:"required"`
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

type loginRequest struct {
	Email string `json:"magic_name", binding:"required"`
	Password string `json:"password", binding:"required"`
}

//TODO: implement authentication middleware
func (server *Server) register(ctx *gin.Context) {
	var req createMagicianRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMagicianParams {
		Email: req.Email,
		Birthday: req.Birthday,
		OriginalName: req.OriginalName,
		Password: req.Password,
		MagicName: req.MagicName,
		MagicalRating: req.MagicRating,
	}

	new, err := server.store.CreateMagician(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, new)
}

//TODO: implement authentication middleware
func (server *Server) login(ctx *gin.Context) {
	var req loginRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	magician, err := server.store.GetMagicianByEmail(ctx, req.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if magician.Password == req.Password {
		ctx.JSON(http.StatusOK, "You are logged in!")
	}

	ctx.JSON(http.StatusUnauthorized, "Invalid password")
}


//TODO: implement authentication middleware
func (server *Server) changePassword(ctx *gin.Context) {
	var uriBind getMagicianByIdRequest
	var req changePasswordRequest

	if err := ctx.ShouldBindUri(&uriBind); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	magician, err := server.store.GetMagicianById(ctx, uriBind.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if magician.Password != req.OldPassword {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.UpdatePasswordParams{
		ID: uriBind.ID,
		Password: req.NewPassword,
	}

	updated, err := server.store.UpdatePassword(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

func (server *Server) getMagician(ctx *gin.Context) {
	var uriBind getMagicianByIdRequest

	if err := ctx.ShouldBindUri(&uriBind); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	magician, err := server.store.GetMagicianById(ctx, uriBind.ID);

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return;
	}
	ctx.JSON(http.StatusOK, magician)
}



func (server *Server) updateMagician(ctx *gin.Context) {
	var uriBind getMagicianByIdRequest
	var req updateMagicianRequest

	if err := ctx.BindUri(&uriBind); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateMagicianParams {
		ID: uriBind.ID,
		EmailDoUpdate: req.EmailDoUpdate,
		Email: req.Email,
		BirthdayDoUpdate: req.BirthdayDoUpdate,
		Birthday: req.Birthday,
		OriginalNameDoUpdate: req.OriginalNameDoUpdate,
		OriginalName: req.OriginalName,
		MagicNameDoUpdate: req.MagicNameDoUpdate,
		MagicName: req.MagicName,
	}

	updated, err := server.store.UpdateMagician(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

/** 
* PUT /magicians/{id}/rating
*/
func (server *Server) updateMagicianMagicRating(ctx *gin.Context) {
	var uriBind getMagicianByIdRequest
	var req updateMagicianMagicRatingRequest

	if err := ctx.BindUri(&uriBind); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateMagicianRatinParams {
		ID: uriBind.ID,
		MagicalRating: req.MagicRating,
	}

	updated, err := server.store.UpdateMagicianRatin(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

/** 
* DELETE /magicians/{id}
*/
func (server *Server) deleteMagician(ctx *gin.Context) {
	var uriBind getMagicianByIdRequest

	if err := ctx.BindUri(&uriBind); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteMagician(ctx, uriBind.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.Status(http.StatusOK)
}