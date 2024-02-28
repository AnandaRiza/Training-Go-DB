package routes

import (
	"bcas/bookstore-go/internals/handlers"
	"bcas/bookstore-go/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitAuthRouter(router *gin.Engine, db *sqlx.DB) {

	authRouter := router.Group("/auth")
	authRepo := repositories.InitAuthRepo(db)
	authHandler := handlers.InitAuthHandler(authRepo)

	//bikin Rute
	// localhost:8000/auth/new
	authRouter.POST("/new", authHandler.Register)
	//localhost:8000/auth
	authRouter.POST("", authHandler.Login)
}
