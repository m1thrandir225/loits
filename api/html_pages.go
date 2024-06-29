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


func (server *Server) renderHomepage(ctx *gin.Context) {
	pageData := layouts.PageData {
		Title: "Loits - Home",
		ActiveLink: "/",
	}
	err := renderTemplate(ctx, http.StatusOK, pages.HomePage(pageData))

	if err != nil {
		renderErrorPage(ctx, http.StatusNotFound)
	}
}

func (server *Server) renderSpellspage(ctx *gin.Context) {

	pageData := layouts.PageData {
		Title: "Your Spells",
		ActiveLink: "/spells",
	}

	err := renderTemplate(ctx, http.StatusOK, pages.SpellsPage(pageData))

	if err != nil {
		renderErrorPage(ctx, http.StatusNotFound)
	}
}

func (server *Server) renderBookspage(ctx *gin.Context) {
	pageData := layouts.PageData {
		Title: "Your Magic Books",
		ActiveLink: "/books",
	}
	err := renderTemplate(ctx, http.StatusOK, pages.BooksPage(pageData))

	if err != nil {
		renderErrorPage(ctx, http.StatusNotFound)
	}
}

func (server *Server) renderProfilepage(ctx *gin.Context) {

	pageData := layouts.PageData {
		Title: "My Profile",
		ActiveLink: "/profile",
	}
	err := renderTemplate(ctx, http.StatusOK, pages.ProfilePage(pageData))

	if err != nil {
		renderErrorPage(ctx, http.StatusNotFound)
	}
}

