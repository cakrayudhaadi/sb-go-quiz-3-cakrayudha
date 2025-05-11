package services

import (
	"errors"
	"quiz3/middlewares"
	"quiz3/repositories"
	"quiz3/structs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	CreateCategories(ctx *gin.Context) (err error)
	GetAllCategories(ctx *gin.Context) (categories []structs.Categories, err error)
	GetCategories(ctx *gin.Context) (category structs.Categories, err error)
	UpdateCategories(ctx *gin.Context) (err error)
	DeleteCategories(ctx *gin.Context) (err error)
	GetBooksByCategories(ctx *gin.Context) (books []structs.Books, err error)
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository,
	}
}

func (service *categoryService) CreateCategories(ctx *gin.Context) (err error) {
	var newCategories structs.Categories

	newCategories, err = validateCategoriesReqAndConvertToCategories(ctx)
	if err != nil {
		return
	}

	loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	newCategories.CreatedBy = loginName
	newCategories.CreatedAt = time.Now()

	err = service.categoryRepository.CreateCategories(newCategories)
	if err != nil {
		err = errors.New("data category gagal dibuat")
	}

	return
}

func (service *categoryService) GetAllCategories(ctx *gin.Context) (categories []structs.Categories, err error) {
	categories, err = service.categoryRepository.GetAllCategories()
	if err != nil {
		err = errors.New("data category gagal diambil")
	} else if len(categories) == 0 {
		err = errors.New("data category kosong")
	}

	return
}

func (service *categoryService) GetCategories(ctx *gin.Context) (category structs.Categories, err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	category, err = service.categoryRepository.GetCategories(id)
	if category.ID == 0 {
		err = errors.New("data category tidak ada")
	} else if err != nil {
		err = errors.New("data category gagal diambil")
	}

	return
}

func (service *categoryService) UpdateCategories(ctx *gin.Context) (err error) {
	var newCategories structs.Categories
	id, _ := strconv.Atoi(ctx.Param("id"))

	newCategories, err = validateCategoriesReqAndConvertToCategories(ctx)
	if err != nil {
		return
	}

	oldCategories, err := service.GetCategories(ctx)
	if err != nil {
		err = errors.New("data category tidak ditemukan")
		return
	}
	newCategories.ID = id
	newCategories.CreatedBy = oldCategories.CreatedBy
	newCategories.CreatedAt = oldCategories.CreatedAt

	loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	newCategories.ModifiedBy = loginName
	newCategories.ModifiedAt = time.Now()

	err = service.categoryRepository.UpdateCategories(newCategories)
	if err != nil {
		err = errors.New("data category gagal diubah")
	}

	return
}

func (service *categoryService) DeleteCategories(ctx *gin.Context) (err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err = service.GetCategories(ctx)
	if err != nil {
		err = errors.New("data category tidak ditemukan")
		return
	}

	err = service.categoryRepository.DeleteCategories(id)
	if err != nil {
		err = errors.New("data category gagal dihapus")
	}

	return
}

func (service *categoryService) GetBooksByCategories(ctx *gin.Context) (books []structs.Books, err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	books, err = service.categoryRepository.GetBooksByCategories(id)
	if err != nil {
		err = errors.New("data book gagal diambil")
	} else if len(books) == 0 {
		err = errors.New("data book kosong")
	}

	return
}

func validateCategoriesReqAndConvertToCategories(ctx *gin.Context) (categories structs.Categories, err error) {
	var categoriesRequest structs.CategoriesRequest
	err = ctx.ShouldBindJSON(&categoriesRequest)
	if err != nil {
		err = errors.New("parameter yang dimasukkan salah")
		return
	}

	err = categoriesRequest.ValidateCategory()
	if err != nil {
		return
	}
	categories = categoriesRequest.ConvertCategoriesReqToCategories()

	return
}
