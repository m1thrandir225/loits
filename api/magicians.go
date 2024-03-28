package api

import "github.com/gin-gonic/gin"

func (server *Server) register(ctx *gin.Context) {}

func (server *Server) login(ctx *gin.Context) {}

func (server *Server) changePassword(ctx *gin.Context) {}

func (server *Server) getMagician(ctx *gin.Context) {}

func (server *Server) deleteMagician(ctx *gin.Context) {}

func (server *Server) updateMagician(ctx *gin.Context) {}

func (server *Server) updateMagicianMagicRating(ctx *gin.Context) {}
