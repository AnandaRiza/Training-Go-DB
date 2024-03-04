package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	//buat routenya
	router.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Help Wot;fd")
	})

	InitAuthRouter(router, db)
	InitBookRouter(router, db)

	return router
}
