package structs

import (
	"errors"
	"quiz3/commons"
	"time"
)

type Books struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	Price       int       `json:"price"`
	TotalPage   int       `json:"total_page"`
	Thickness   string    `json:"thickness"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	ModifiedAt  time.Time `json:"modified_at"`
	ModifiedBy  string    `json:"modified_by"`
}

type BooksRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	ReleaseYear int    `json:"release_year"`
	Price       int    `json:"price"`
	TotalPage   int    `json:"total_page"`
	CategoryID  int    `json:"category_id"`
}

func (b *BooksRequest) ConvertBooksReqToBooks() Books {
	return Books{
		Title:       b.Title,
		Description: b.Description,
		ImageURL:    b.ImageURL,
		ReleaseYear: b.ReleaseYear,
		Price:       b.Price,
		TotalPage:   b.TotalPage,
		CategoryID:  b.CategoryID,
	}
}

func (b *BooksRequest) ValidateBooks() (err error) {
	if commons.IsValueEmpty(b.Title) {
		return errors.New("title harus diisi")
	}
	if commons.IsValueEmpty(b.Description) {
		return errors.New("description harus diisi")
	}
	if commons.IsValueEmpty(b.ImageURL) {
		return errors.New("image_url harus diisi")
	}
	if commons.IsValueEmpty(b.ReleaseYear) {
		return errors.New("release_year harus diisi")
	}
	if commons.IsValueEmpty(b.Price) {
		return errors.New("price harus diisi")
	}
	if commons.IsValueEmpty(b.TotalPage) {
		return errors.New("v harus diisi")
	}
	if commons.IsValueEmpty(b.CategoryID) {
		return errors.New("category_id harus diisi")
	}

	if b.ReleaseYear < 1980 || b.ReleaseYear > 2024 {
		return errors.New("release_year harus antara 1980 dan 2024")
	}

	return nil
}
