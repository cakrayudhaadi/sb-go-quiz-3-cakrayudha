package routers

import (
	"net/http"
	"quiz3/commons"
	"quiz3/databases/connection"
	"quiz3/repositories"
	"quiz3/services"

	"github.com/gin-gonic/gin"
)

func userInitiator(router *gin.Engine) {
	api := router.Group("/api/users")
	{
		api.POST("/login", Login)
		api.POST("/signup", SignUp)
	}
}

func Login(ctx *gin.Context) {
	var (
		userRepo = repositories.NewUserRepository(connection.DBConnections)
		userSrv  = services.NewUserService(userRepo)
	)

	token, err := userSrv.Login(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "login berhasil dilakukan", token)
}

func SignUp(ctx *gin.Context) {
	var (
		userRepo = repositories.NewUserRepository(connection.DBConnections)
		userSrv  = services.NewUserService(userRepo)
	)

	err := userSrv.SignUp(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "akun berhasil dibuat")
}
