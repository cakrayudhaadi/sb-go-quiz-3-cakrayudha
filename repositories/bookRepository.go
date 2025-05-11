package repositories

import (
	"quiz3/structs"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBooks() (books []structs.Books, err error)
	GetBooks(id int) (book structs.Books, err error)
	CreateBooks(book structs.Books) error
	UpdateBooks(book structs.Books) error
	DeleteBooks(id int) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(database *gorm.DB) BookRepository {
	return &bookRepository{
		db: database,
	}
}

func (repo *bookRepository) GetAllBooks() (books []structs.Books, err error) {
	err = repo.db.Find(&books).Error
	return
}

func (repo *bookRepository) GetBooks(id int) (book structs.Books, err error) {
	err = repo.db.Where("id = ?", id).Find(&book).Error
	return
}

func (repo *bookRepository) CreateBooks(book structs.Books) error {
	err := repo.db.Create(&book).Error
	return err
}

func (repo *bookRepository) UpdateBooks(book structs.Books) error {
	err := repo.db.Model(&structs.Books{}).Where("id = ?", book.ID).Updates(book).Error
	return err
}

func (repo *bookRepository) DeleteBooks(id int) error {
	err := repo.db.Where("id = ?", id).Delete(&structs.Books{}).Error
	return err
}
