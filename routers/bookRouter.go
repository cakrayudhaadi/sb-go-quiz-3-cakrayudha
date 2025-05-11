package routers

import (
	"net/http"
	"quiz3/commons"
	"quiz3/databases/connection"
	"quiz3/middlewares"
	"quiz3/repositories"
	"quiz3/services"

	"github.com/gin-gonic/gin"
)

func bookInitiator(router *gin.Engine) {
	api := router.Group("/api/books")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateBooks)
		api.GET("", GetAllBooks)
		api.GET("/:id", GetBooks)
		api.PUT("/:id", UpdateBooks)
		api.DELETE("/:id", DeleteBooks)
	}
}

func CreateBooks(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBookRepository(connection.DBConnections)
		categorySrv  = services.NewBookService(categoryRepo)
	)

	err := categorySrv.CreateBooks(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil dibuat")
}

func GetAllBooks(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBookRepository(connection.DBConnections)
		categorySrv  = services.NewBookService(categoryRepo)
	)

	books, err := categorySrv.GetAllBooks(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data books berhasil diambil", books)
}

func GetBooks(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBookRepository(connection.DBConnections)
		categorySrv  = services.NewBookService(categoryRepo)
	)

	book, err := categorySrv.GetBooks(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data book berhasil diambil", book)
}

func UpdateBooks(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBookRepository(connection.DBConnections)
		categorySrv  = services.NewBookService(categoryRepo)
	)

	err := categorySrv.UpdateBooks(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil diubah")
}

func DeleteBooks(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewBookRepository(connection.DBConnections)
		categorySrv  = services.NewBookService(categoryRepo)
	)

	err := categorySrv.DeleteBooks(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data book berhasil dihapus")
}
