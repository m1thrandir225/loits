package api

import (
	"github.com/gin-gonic/gin"
	db "m1thrandir225/loits/db/sqlc"
	"m1thrandir225/loits/util"
)

type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		/*
		* Spells
		 */
		v1.GET("/spells/:id", server.getSpellById)
		v1.POST("/spells/", server.createSpell)
		//v1.PUT("/spells/:id", server.updateSpell)
		v1.DELETE("/spells/:id", server.deleteSpell)

		/*
		* Spell Books
		 */
		v1.POST("/books/", server.createSpellBook)
		v1.GET("/books/:id", server.getSpellBookById)
		v1.GET("/books/:id/spells", server.getSpellsByBook)
		v1.DELETE("/books/:id", server.deleteSpellBook)

		/*
		 * Magicians
		 */
		v1.POST("/register", server.register)
		v1.POST("/login", server.login)

		v1.GET("/magician/:id", server.getMagician)
		v1.PUT("/magician/:id", server.updateMagician)
		v1.DELETE("/magician/:id", server.deleteMagician)
	}
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
