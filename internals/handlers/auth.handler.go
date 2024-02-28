package handlers

import (
	"bcas/bookstore-go/internals/models"
	"bcas/bookstore-go/internals/repositories"
	"log"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	*repositories.AuthRepo
}

func InitAuthHandler(a *repositories.AuthRepo) *AuthHandler {
	return &AuthHandler{a}
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	// ambil body
	body := models.AuthModel{}
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, err := a.FindByEmail(body)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//pengecekan duplicate email
	if len(result) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email is already registered!",
		})
		return
	}

	hash, err := argon2id.CreateHash(body.Password, argon2id.DefaultParams)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := a.SaveUser(models.AuthModel{
		Email:    body.Email,
		Password: hash,
	}); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success Registered!",
	})
}

func (a *AuthHandler) Login(ctx *gin.Context) {}
