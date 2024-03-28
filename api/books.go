package api

import (
	db "m1thrandir225/loits/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type createSpellBookRequest struct {
	Name  string      `json:"name" binding:"required,min=1"`
	Owner pgtype.UUID `json:"owner" binding:"required,min=1"`
}

type getSpellBookByIdRequest struct {
	ID pgtype.UUID `uri:"id" binding:"required,min=1"`
}

type getUserSpellBooksRequest struct {
	UserID pgtype.UUID `uri:"user_id" binding:"required,min=1"`
}

/**
* POST /books/
 */
func (server *Server) createSpellBook(ctx *gin.Context) {
	var req createSpellBookRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSpellBookParams{
		Name:  req.Name,
		Owner: req.Owner,
	}

	spellBook, err := server.store.CreateSpellBook(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, spellBook)
}

/**
* GET /books/{id}
 */

func (server *Server) getSpellBookById(ctx *gin.Context) {
	var req createSpellBookRequest

	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
}



/**
* GET /magician/{user_id}/books
 */
func (server *Server) getUserSpellBooks(ctx *gin.Context) {
	var req getUserSpellBooksRequest

	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	spellBooks, err := server.store.GetUserSpellBooks(ctx, req.UserID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, spellBooks)
}

/**
* DELETE /books/{id}
 */
func (server *Server) deleteSpellBook(ctx *gin.Context) {
	var req getSpellBookByIdRequest
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteSpellBook(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
