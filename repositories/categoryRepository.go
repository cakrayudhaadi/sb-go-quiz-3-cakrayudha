package repositories

import (
	"quiz3/structs"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategories() (categories []structs.Categories, err error)
	GetCategories(id int) (category structs.Categories, err error)
	CreateCategories(category structs.Categories) (err error)
	UpdateCategories(category structs.Categories) (err error)
	DeleteCategories(id int) (err error)
	GetBooksByCategories(id int) (books []structs.Books, err error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(database *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: database,
	}
}

func (repo *categoryRepository) GetAllCategories() (categories []structs.Categories, err error) {
	err = repo.db.Find(&categories).Error
	return
}

func (repo *categoryRepository) GetCategories(id int) (category structs.Categories, err error) {
	err = repo.db.Where("id = ?", id).Find(&category).Error
	return
}

func (repo *categoryRepository) CreateCategories(category structs.Categories) (err error) {
	err = repo.db.Create(&category).Error
	return
}

func (repo *categoryRepository) UpdateCategories(category structs.Categories) (err error) {
	err = repo.db.Model(&structs.Categories{}).Where("id = ?", category.ID).Updates(category).Error
	return
}

func (repo *categoryRepository) DeleteCategories(id int) (err error) {
	err = repo.db.Where("id = ?", id).Delete(&structs.Categories{}).Error
	return
}

func (repo *categoryRepository) GetBooksByCategories(id int) (books []structs.Books, err error) {
	err = repo.db.Table("books").Where("categories.id = ?", id).
		Joins("Join categories on books.category_id = categories.id").Find(&books).Error
	return
}
