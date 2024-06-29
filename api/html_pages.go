package api

import (
	"m1thrandir225/loits/templates/layouts"
	"m1thrandir225/loits/templates/pages"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func renderTemplate(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func renderErrorPage(ctx *gin.Context, errorCode int) {
	err := renderTemplate(ctx, http.StatusNotFound, pages.ErrorPage(strconv.Itoa(errorCode)))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "There was an error rendering this page.")
	}
}


func (server *Server) renderHomePage(ctx *gin.Context) {
	authCookie, err  := ctx.Cookie("auth")

	pageData := layouts.PageData {
		Title: "Loits - Home",
		ActiveLink: "/",
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

func (server *Server) renderSpellsPage(ctx *gin.Context) {
	authCookie, err  := ctx.Cookie("auth")


	pageData := layouts.PageData {
		Title: "Loits - Your Spells",
		ActiveLink: "/spells",
		IsAuthenticated: true,
	}

	if err != nil {
		pageData.IsAuthenticated = false
		ctx.Redirect(http.StatusMovedPermanently, "/login")
	}
	println(authCookie)
	err = renderTemplate(ctx, http.StatusOK, pages.SpellsPage(pageData))

	if err != nil {
		renderErrorPage(ctx, http.StatusNotFound)
	}
}

func (server *Server) renderBooksPage(ctx *gin.Context) {
	authCookie, err  := ctx.Cookie("auth")

	pageData := layouts.PageData {
		Title: "Loits - Your Magic Books",
		ActiveLink: "/books",
		IsAuthenticated: true,
	}
	println(authCookie)
	if err != nil {
		pageData.IsAuthenticated = false
		ctx.Redirect(http.StatusMovedPermanently, "/login")
	}

	//TODO verify cookie


	err = renderTemplate(ctx, http.StatusOK, pages.BooksPage(pageData))

	if err != nil {
		renderErrorPage(ctx, http.StatusNotFound)
	}
}

func (server *Server) renderProfilePage(ctx *gin.Context) {
	authCookie, err  := ctx.Cookie("auth")

	pageData := layouts.PageData {
		Title: "Loits - My Profile",
		ActiveLink: "/profile",
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

func (server *Server) renderLoginPage(ctx *gin.Context) {
	pageData := layouts.PageData {
		Title: "Loits - Login",
		ActiveLink: "/login",
		IsAuthenticated: false,
	}
	err := renderTemplate(ctx, http.StatusOK, pages.LoginPage(pageData))

	if err != nil {
		renderErrorPage(ctx, http.StatusNotFound)
	}
}

func (server *Server) renderRegisterPage(ctx *gin.Context) {
	pageData := layouts.PageData {
		Title: "Loits - Register",
		ActiveLink: "/register",
		IsAuthenticated: false,
	}
	err := renderTemplate(ctx, http.StatusOK, pages.RegisterPage(pageData))

	if err != nil {
		renderErrorPage(ctx, http.StatusNotFound)
	}
}