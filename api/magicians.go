package api

import (
	db "m1thrandir225/loits/db/sqlc"
	"m1thrandir225/loits/templates/layouts"
	"m1thrandir225/loits/templates/pages"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type getMagicianByIdRequest struct {
	ID pgtype.UUID `uri:"id" binding:"required,min=1"`
}

type createMagicianRequest struct {
	OriginalName string         `json:"original_name" binding:"required"`
	MagicName    string         `json:"magic_name" binding:"required"`
	Email        string         `json:"email" binding:"required"`
	Password     string         `json:"password" binding:"required"`
	Birthday     string         `json:"birthday" binding:"required"`
	MagicRating  db.MagicRating `json:"magic_rating", binding:"required"`
}

type updateMagicianRequest struct {
	Email                string    `json:"email" binding:"required"`
	EmailDoUpdate        bool      `json:"email_do_update"`
	OriginalName         string    `json:"original_name"`
	OriginalNameDoUpdate bool      `json:"original_name_do_update"`
	MagicName            string    `json:"magic_name"`
	MagicNameDoUpdate    bool      `json:"magic_name_do_update"`
	Birthday             time.Time `json:"birthday"`
	BirthdayDoUpdate     bool      `json:"birthday_do_update"`
}

type updateMagicianMagicRatingRequest struct {
	MagicRating db.MagicRating `json:"magic_rating" binding:"required"`
}

type changePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type loginRequest struct {
	Email    string `json:"email", binding:"required"`
	Password string `json:"password", binding:"required"`
}

type createMagicianResponse struct {
	Magician    db.Magician `json:"user"`
	AccessToken string      `json:"access_token"`
}

type loginResponse struct {
	AccessToken  string         `json:"access_token"`
	OriginalName string         `json:"original_name"`
	MagicName    string         `json:"magic_name" binding:"required"`
	Email        string         `json:"email" binding:"required"`
	Birthday     string         `json:"birthday" binding:"required"`
	MagicRating  db.MagicRating `json:"magic_rating", binding:"required"`
}

const layout = "Jan 2, 2006 at 3:04pm (MST)"

// TODO: implement authentication middleware
func (server *Server) register(ctx *gin.Context) {
	var req createMagicianRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	birthdayDate, err := time.Parse(layout, req.Birthday)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateMagicianParams{
		Email:         req.Email,
		Birthday:      birthdayDate,
		OriginalName:  req.OriginalName,
		Password:      req.Password,
		MagicName:     req.MagicName,
		MagicalRating: req.MagicRating,
	}

	new, err := server.store.CreateMagician(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	token, err := server.tokenMaker.CreateToken(new.Email, server.config.AccessTokenDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := createMagicianResponse{
		Magician:    new,
		AccessToken: token,
	}

	ctx.JSON(http.StatusOK, response)
}

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

	if magician.Password != req.Password {
		ctx.JSON(http.StatusUnauthorized, "Invalid password")
		return
	}

	token, err := server.tokenMaker.CreateToken(magician.Email, server.config.AccessTokenDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := loginResponse{
		OriginalName: magician.OriginalName,
		Email:        magician.Email,
		MagicRating:  magician.MagicalRating,
		Birthday:     magician.Birthday.Format(layout),
		MagicName:    magician.MagicName,
		AccessToken:  token,
	}

	ctx.JSON(http.StatusOK, response)

}

// TODO: implement authentication middleware
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
		ID:       uriBind.ID,
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

	magician, err := server.store.GetMagicianById(ctx, uriBind.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
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

	arg := db.UpdateMagicianParams{
		ID:                   uriBind.ID,
		EmailDoUpdate:        req.EmailDoUpdate,
		Email:                req.Email,
		BirthdayDoUpdate:     req.BirthdayDoUpdate,
		Birthday:             req.Birthday,
		OriginalNameDoUpdate: req.OriginalNameDoUpdate,
		OriginalName:         req.OriginalName,
		MagicNameDoUpdate:    req.MagicNameDoUpdate,
		MagicName:            req.MagicName,
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

	arg := db.UpdateMagicianRatinParams{
		ID:            uriBind.ID,
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

func (server *Server) renderLoginPage(ctx *gin.Context) {
	pageData := layouts.PageData{
		Title:           "Loits - Login",
		ActiveLink:      "/login",
		IsAuthenticated: false,
	}
	err := renderTemplate(ctx, http.StatusOK, pages.LoginPage(pageData))

	if err != nil {
		renderErrorPage(ctx, http.StatusNotFound)
	}
}

func (server *Server) renderRegisterPage(ctx *gin.Context) {
	pageData := layouts.PageData{
		Title:           "Loits - Register",
		ActiveLink:      "/register",
		IsAuthenticated: false,
	}
	err := renderTemplate(ctx, http.StatusOK, pages.RegisterPage(pageData))

	if err != nil {
		renderErrorPage(ctx, http.StatusNotFound)
	}
}

func (server *Server) renderProfilePage(ctx *gin.Context) {
	authCookie, err := ctx.Cookie("auth")

	pageData := layouts.PageData{
		Title:           "Loits - My Profile",
		ActiveLink:      "/profile",
		IsAuthenticated: true,
	}

	if err != nil {
		pageData.IsAuthenticated = false
		ctx.Redirect(http.StatusMovedPermanently, "/login")
	}
	println(authCookie)
	err = renderTemplate(ctx, http.StatusOK, pages.ProfilePage(pageData))

	if err != nil {
		renderErrorPage(ctx, http.StatusNotFound)
	}
}

func (server *Server) renderHomePage(ctx *gin.Context) {
	authCookie, err := ctx.Cookie("auth")

	pageData := layouts.PageData{
		Title:           "Loits - Home",
		ActiveLink:      "/",
		IsAuthenticated: true,
	}

	println(authCookie)

	if err != nil {
		pageData.IsAuthenticated = false
		ctx.Redirect(http.StatusMovedPermanently, "/login")
	}

	err = renderTemplate(ctx, http.StatusOK, pages.HomePage(pageData))

	if err != nil {
		renderErrorPage(ctx, http.StatusNotFound)
	}
}
