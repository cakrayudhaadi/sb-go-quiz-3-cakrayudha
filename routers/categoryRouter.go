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

func categoryInitiator(router *gin.Engine) {
	api := router.Group("/api/categories")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateCategories)
		api.GET("", GetAllCategories)
		api.GET("/:id", GetCategories)
		api.PUT("/:id", UpdateCategories)
		api.DELETE("/:id", DeleteCategories)
		api.GET("/:id/books", GetBooksByCategories)
	}
}

func CreateCategories(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewCategoryRepository(connection.DBConnections)
		categorySrv  = services.NewCategoryService(categoryRepo)
	)

	err := categorySrv.CreateCategories(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data categories berhasil dibuat")
}

func GetAllCategories(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewCategoryRepository(connection.DBConnections)
		categorySrv  = services.NewCategoryService(categoryRepo)
	)

	categories, err := categorySrv.GetAllCategories(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data category berhasil diambil", categories)
}

func GetCategories(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewCategoryRepository(connection.DBConnections)
		categorySrv  = services.NewCategoryService(categoryRepo)
	)

	category, err := categorySrv.GetCategories(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data category berhasil diambil", category)
}

func UpdateCategories(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewCategoryRepository(connection.DBConnections)
		categorySrv  = services.NewCategoryService(categoryRepo)
	)

	err := categorySrv.UpdateCategories(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data category berhasil diubah")
}

func DeleteCategories(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewCategoryRepository(connection.DBConnections)
		categorySrv  = services.NewCategoryService(categoryRepo)
	)

	err := categorySrv.DeleteCategories(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "data category berhasil dihapus")
}

func GetBooksByCategories(ctx *gin.Context) {
	var (
		categoryRepo = repositories.NewCategoryRepository(connection.DBConnections)
		categorySrv  = services.NewCategoryService(categoryRepo)
	)

	books, err := categorySrv.GetBooksByCategories(ctx)
	if err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	commons.ResponseSuccessWithData(ctx, http.StatusOK, "data books berhasil diambil", books)
}
