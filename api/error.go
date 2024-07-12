package api

import (
	"m1thrandir225/loits/templates/pages"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func renderErrorPage(ctx *gin.Context, errorCode int) {
	err := renderTemplate(ctx, http.StatusNotFound, pages.ErrorPage(strconv.Itoa(errorCode)))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "There was an error rendering this page.")
	}
}


func (server *Server) renderNotFoundPage(ctx *gin.Context) {
	renderErrorPage(ctx, http.StatusNotFound)
}