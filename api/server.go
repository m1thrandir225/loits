package api

import (
	"fmt"
	db "m1thrandir225/loits/db/sqlc"
	"m1thrandir225/loits/token"
	"m1thrandir225/loits/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) verifyAuthCookie(cookie string) bool {

	payload, err := server.tokenMaker.VerifyToken(cookie)

	if err != nil {
		return false
	}

	err = payload.Valid()

	if err != nil {
		return false
	}

	return true
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.Static("/static", "./public")

	/**
	* HTML Pages
	 */
	router.GET("/", server.renderHomePage)
	router.GET("/profile", server.renderProfilePage)
	router.GET("/books", server.renderBooksPage)
	router.GET("/spells", server.renderSpellsPage)

	router.GET("/login", server.renderLoginPage)
	router.GET("/register", server.renderRegisterPage)

	//router.NoRoute(server.renderNotFoundPage)

	/**
	* Api Routes
	 */

	v1 := router.Group("/api/v1")

	authRoutes := v1.Group("/").Use(authMiddleware(server.tokenMaker))

	/**
	* Authentication
	 */
	v1.POST("/register", server.register)
	v1.POST("/login", server.login)

	authRoutes.POST("/change-password", server.changePassword)

	/**
	* Spells
	 */
	authRoutes.GET("/spells/:id", server.getSpellById)
	authRoutes.POST("/spells/", server.createSpell)
	authRoutes.PUT("/spells/:id", server.updateSpell)
	authRoutes.PUT("/spells/:id/move", server.updateSpellElement)
	authRoutes.DELETE("/spells/:id", server.deleteSpell)

	/**
	* Spell Books
	 */
	authRoutes.POST("/books/", server.createSpellBook)
	authRoutes.GET("/books/:id", server.getSpellBookById)
	authRoutes.GET("/books/:id/spells", server.getSpellsByBook)
	authRoutes.DELETE("/books/:id", server.deleteSpellBook)

	/**
	* Magicians
	 */
	authRoutes.GET("/magician/:id", server.getMagician)
	authRoutes.PUT("/magician/:id", server.updateMagician)
	authRoutes.PUT("/magician/:id/magic-rating", server.updateMagicianMagicRating)
	authRoutes.DELETE("/magician/:id", server.deleteMagician)
	authRoutes.GET("/magician/:id/books", server.getUserSpellBooks)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
