package api

import (
	db "m1thrandir225/loits/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type createSpellRequest struct {
	Name    string      `json:"name" binding:"required,min=6"`
	Element string      `json:"element" binding:"required"`
	BookID  pgtype.UUID `json:"book_id" binding:"required"`
}

type getSpellByNameRequest struct {
	Name string `uri:"name" binding:"required,min=1"`
}

type getSpellByIdRequest struct {
	ID pgtype.UUID `uri:"id" binding:"required,min=1"`
}

type getSpellsByBookRequest struct {
	BookID pgtype.UUID `uri:"book_id" binding:"required,min=1"`
}

type getSpellsByBookResponse struct {
	Spells []db.Spell
}

type moveSpellToNewBookRequest struct {
	ID     pgtype.UUID `uri:"id" binding:"required"`
	BookID pgtype.UUID `uri:"book_id" binding:"required"`
}

type updateSpellElementRequest struct {
	ID      pgtype.UUID `uri:"id" binding:"required"`
	Element db.Element  `uri:"element" binding:"required"`
}

type updateSpellNameRequest struct {
	ID   pgtype.UUID `uri:"id" binding:"required"`
	Name string      `uri:"name" binding:"required"`
}

/**
* POST /spells/
 */

func (server *Server) createSpell(ctx *gin.Context) {
	var req createSpellRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSpellParams{
		Name:    req.Name,
		Element: db.Element(req.Element),
		BookID:  req.BookID,
	}

	spell, err := server.store.CreateSpell(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, spell)
}

/**
* GET /spells/{id}/
 */

func (server *Server) getSpellById(ctx *gin.Context) {
	var req getSpellByIdRequest

	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	spell, err := server.store.GetSpellById(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, spell)

}

/**
* GET /spells/{name}/
 */
func (server *Server) getSpellByName(ctx *gin.Context) {
	var req getSpellByNameRequest

	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	spell, err := server.store.GetSpellByName(ctx, req.Name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, spell)
}

/**
* GET /spells/{book_id}
 */
func (server *Server) getSpellsByBook(ctx *gin.Context) {
	var req getSpellsByBookRequest

	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	spells, err := server.store.GetSpellsByBook(ctx, req.BookID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, spells)
}

/**
* PUT /spells/{id}/{book_id}
 */
func (server *Server) moveSpellToNewBook(ctx *gin.Context) {
	var req moveSpellToNewBookRequest

	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

		return
	}

	arg := db.MoveToNewBookParams{
		ID:     req.ID,
		BookID: req.BookID,
	}

	movedSpell, err := server.store.MoveToNewBook(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, movedSpell)
}

/**
* PUT /spells/{id}/{element}
 */
func (server *Server) updateSpellElement(ctx *gin.Context) {
	var req updateSpellElementRequest

	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

		return
	}

	arg := db.UpdateSpellElementParams{
		ID:      req.ID,
		Element: req.Element,
	}

	updatedSpell, err := server.store.UpdateSpellElement(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedSpell)
}

/**
* PUT /spells/{id}/{name}
 */

func (server *Server) updateSpellName(ctx *gin.Context) {
	var req updateSpellNameRequest

	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

		return
	}

	arg := db.UpdateSpellParams{
		ID:   req.ID,
		Name: req.Name,
	}

	updatedSpell, err := server.store.UpdateSpell(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedSpell)
}

/**
* DELETE /spells/{id}
 */
func (server *Server) deleteSpell(ctx *gin.Context) {
	var req getSpellByIdRequest
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteSpell(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
