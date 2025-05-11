package structs

import (
	"errors"
	"quiz3/commons"
	"time"
)

type Categories struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

type CategoriesRequest struct {
	Name string `json:"name"`
}

func (c *CategoriesRequest) ConvertCategoriesReqToCategories() Categories {
	return Categories{
		Name: c.Name,
	}
}

func (c *CategoriesRequest) ValidateCategory() (err error) {
	if commons.IsValueEmpty(c.Name) {
		return errors.New("name harus diisi")
	}

	return nil
}
