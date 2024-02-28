package pkg

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitServer(router *gin.Engine) *http.Server {
	address := "localhost:4000"
	server := &http.Server{
		Addr:         address,
		Handler:      router,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 10,
		ReadTimeout:  time.Second * 10,
	}
	return server
}
