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

type getSpellByIdRequest struct {
	ID pgtype.UUID `uri:"id" binding:"required,min=1"`
}

type getSpellsByBookRequest struct {
	BookID pgtype.UUID `uri:"id" binding:"required,min=1"`
}


type updateSpellElementRequest struct {
	ID      pgtype.UUID `uri:"id" binding:"required"`
	Element db.Element  `uri:"element" binding:"required"`
}

type updateSpellRequest struct {
	ID      pgtype.UUID `uri:"id" binding:"required"`
	Name    string      `json:"name"`
	NameDoUpdate bool    `json:"name_do_update"`
	BookID pgtype.UUID `json:"book_id"`
	BookIDDoUpdate bool `json:"book_id_do_update"`
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
* POST /spells/{id}
 */

 func (server *Server) updateSpell(ctx *gin.Context) {
	var uriBind getSpellByIdRequest
	var req updateSpellRequest

	if err := ctx.ShouldBindUri(&uriBind); err != nil { 
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateSpellParams{
		ID: uriBind.ID,
		Name:    req.Name,
		BookIDDoUpdate: req.BookIDDoUpdate,
		NameDoUpdate: req.NameDoUpdate,
		BookID:  req.BookID,
	}

	spell, err := server.store.UpdateSpell(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, spell)
}

/**
* POST /spells/{id}/element
*/
func (server *Server) updateSpellElement(ctx *gin.Context) {
	var uriBind getSpellByIdRequest
	var req updateSpellElementRequest

	if err := ctx.ShouldBindUri(&uriBind); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateSpellElementParams {
		Element: db.Element(req.Element),
		ID: uriBind.ID,
	}

	spell, err := server.store.UpdateSpellElement(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, spell)
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
