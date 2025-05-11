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

type BookService interface {
	CreateBooks(ctx *gin.Context) (err error)
	GetAllBooks(ctx *gin.Context) (books []structs.Books, err error)
	GetBooks(ctx *gin.Context) (books structs.Books, err error)
	UpdateBooks(ctx *gin.Context) (err error)
	DeleteBooks(ctx *gin.Context) (err error)
}

type bookService struct {
	bookRepository repositories.BookRepository
}

func NewBookService(bookRepository repositories.BookRepository) BookService {
	return &bookService{
		bookRepository,
	}
}

func (service *bookService) CreateBooks(ctx *gin.Context) (err error) {
	var newBooks structs.Books

	newBooks, err = validateBooksReqAndConvertToBooks(ctx)
	if err != nil {
		return
	}

	loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	newBooks.CreatedBy = loginName
	newBooks.CreatedAt = time.Now()

	if newBooks.TotalPage > 100 {
		newBooks.Thickness = "Tebal"
	} else {
		newBooks.Thickness = "Tipis"
	}

	err = service.bookRepository.CreateBooks(newBooks)
	if err != nil {
		err = errors.New("data book gagal dibuat")
	}

	return
}

func (service *bookService) GetAllBooks(ctx *gin.Context) (books []structs.Books, err error) {
	books, err = service.bookRepository.GetAllBooks()
	if err != nil {
		err = errors.New("data book gagal diambil")
	} else if len(books) == 0 {
		err = errors.New("data book kosong")
	}

	return
}

func (service *bookService) GetBooks(ctx *gin.Context) (book structs.Books, err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	book, err = service.bookRepository.GetBooks(id)
	if book.ID == 0 {
		err = errors.New("data book tidak ada")
	} else if err != nil {
		err = errors.New("data book gagal diambil")
	}

	return
}

func (service *bookService) UpdateBooks(ctx *gin.Context) (err error) {
	var newBooks structs.Books
	id, _ := strconv.Atoi(ctx.Param("id"))

	newBooks, err = validateBooksReqAndConvertToBooks(ctx)
	if err != nil {
		return
	}

	oldBooks, err := service.GetBooks(ctx)
	if err != nil {
		err = errors.New("data book tidak ditemukan")
		return
	}
	newBooks.ID = id
	newBooks.CreatedBy = oldBooks.CreatedBy
	newBooks.CreatedAt = oldBooks.CreatedAt

	if newBooks.TotalPage > 100 {
		newBooks.Thickness = "Tebal"
	} else {
		newBooks.Thickness = "Tipis"
	}

	loginName, err := middlewares.GetUsernameFromToken(ctx)
	if err != nil {
		return
	}
	newBooks.ModifiedBy = loginName
	newBooks.ModifiedAt = time.Now()

	err = service.bookRepository.UpdateBooks(newBooks)
	if err != nil {
		err = errors.New("data book gagal diubah")
	}

	return
}

func (service *bookService) DeleteBooks(ctx *gin.Context) (err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err = service.GetBooks(ctx)
	if err != nil {
		err = errors.New("data book tidak ditemukan")
		return
	}

	err = service.bookRepository.DeleteBooks(id)
	if err != nil {
		err = errors.New("data book gagal dihapus")
	}

	return
}

func validateBooksReqAndConvertToBooks(ctx *gin.Context) (books structs.Books, err error) {
	var booksRequest structs.BooksRequest

	err = ctx.ShouldBindJSON(&booksRequest)
	if err != nil {
		err = errors.New("parameter yang dimasukkan salah")
		return
	}

	err = booksRequest.ValidateBooks()
	if err != nil {
		return
	}
	books = booksRequest.ConvertBooksReqToBooks()

	return
}
