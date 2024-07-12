package api

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)


func renderTemplate(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}