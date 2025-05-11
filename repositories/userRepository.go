package repositories

import (
	"errors"
	"quiz3/structs"

	"gorm.io/gorm"
)

type UserRepository interface {
	Login(user structs.LoginRequest) (result structs.User, err error)
	SignUp(user structs.User) (err error)
	DeleteUser(user structs.User) (err error)
	GetListUser() (users []structs.User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepository{
		db: database,
	}
}

func (repo *userRepository) Login(user structs.LoginRequest) (result structs.User, err error) {
	err = repo.db.Where("username = ?", user.Username).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return result, err
	}

	return result, nil
}

func (repo *userRepository) SignUp(user structs.User) (err error) {
	err = repo.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) DeleteUser(user structs.User) (err error) {
	err = repo.db.Where("username = ?", user.Username).Delete(&structs.User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) GetListUser() (users []structs.User, err error) {
	err = repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
