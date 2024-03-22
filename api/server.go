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

	router.GET("/spells/:name", server.getSpellByName)
	router.GET("/spells/:id", server.getSpellById)
	router.GET("/spells/:book_id", server.getSpellsByBook)
	router.POST("/spells", server.createSpell)
	router.PUT("/spells/:id/:element", server.updateSpellElement)
	router.PUT("/spells/:id/:name", server.updateSpellName)
	router.DELETE("/spells/:id", server.deleteSpell)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
