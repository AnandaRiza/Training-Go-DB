package routes

import (
	"bcas/bookstore-go/internals/handlers"
	"bcas/bookstore-go/internals/middlewares"
	"bcas/bookstore-go/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitBookRouter(router *gin.Engine, db *sqlx.DB) {
	bookRouter := router.Group("/book")
	bookRepo := repositories.InitBookRepo(db)
	bookHandler := handlers.InitBookHandler(bookRepo)
	bookRouter.GET("", bookHandler.GetBooks)

	bookRouter.POST("/new", middlewares.CheckToken, bookHandler.CreateBooks)
	bookRouter.GET("/id", middlewares.CheckToken, bookHandler.GetBookById)
	bookRouter.DELETE("/id", middlewares.CheckToken, bookHandler.DeleteBookById)
	bookRouter.PATCH("/id", middlewares.CheckToken, bookHandler.UpdateBookById)

}
