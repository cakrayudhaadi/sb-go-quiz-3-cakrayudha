package structs

import (
	"errors"
	"quiz3/commons"
	"regexp"
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *LoginRequest) ValidateLogin() (err error) {
	if commons.IsValueEmpty(l.Username) {
		return errors.New("username harus diisi")
	}

	if commons.IsValueEmpty(l.Password) {
		return errors.New("password harus diisi")
	}

	return
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SignUpRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	ReTypePassword string `json:"re_type_password"`
}

func (s *SignUpRequest) ValidateSignUp() (err error) {
	if commons.IsValueEmpty(s.Username) {
		return errors.New("username harus diisi")
	}

	if commons.IsValueEmpty(s.Password) {

		return errors.New("password harus diisi")
	}

	if commons.IsValueEmpty(s.ReTypePassword) {
		return errors.New("retype password harus diisi")
	}

	if s.ReTypePassword != s.Password {
		return errors.New("password dan retype password tidak sama")
	}

	re := regexp.MustCompile(`^(.{8,})$`)
	if !re.MatchString(s.Password) {
		return errors.New("password harus mengandung minimal 8 karakter")
	}

	return nil
}

func (s *SignUpRequest) ConvertToModelForSignUp() (user User, err error) {
	hashedPassword, err := commons.HashPassword(s.Password)
	if err != nil {
		err = errors.New("hashing password gagal")
		return
	}

	return User{
		Username: s.Username,
		Password: hashedPassword,
	}, nil
}
