package services

import (
	"errors"
	"quiz3/commons"
	"quiz3/middlewares"
	"quiz3/repositories"
	"quiz3/structs"
	"time"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Login(ctx *gin.Context) (result structs.LoginResponse, err error)
	SignUp(ctx *gin.Context) (err error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository,
	}
}

func (service *userService) Login(ctx *gin.Context) (result structs.LoginResponse, err error) {
	var userReq structs.LoginRequest

	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return
	}

	err = userReq.ValidateLogin()
	if err != nil {
		return
	}

	user, err := service.userRepository.Login(userReq)
	if err != nil {
		return
	}

	if commons.IsValueEmpty(user.ID) {
		err = errors.New("akun tidak valid")
		return
	}

	matches := commons.CheckPassword(user.Password, userReq.Password)
	if !matches {
		err = errors.New("username atau password yang dimasukkan salah")
		return
	}

	jwtToken, err := middlewares.GenerateJwtToken()
	if err != nil {
		return
	}

	structs.LoginRedis[jwtToken] = structs.UserLoginRedis{
		UserId:    0,
		Username:  user.Username,
		LoginAt:   time.Now(),
		ExpiredAt: time.Now().Add(time.Minute * 1),
	}

	result.Token = jwtToken

	return
}

func (service *userService) SignUp(ctx *gin.Context) (err error) {
	var userReq structs.SignUpRequest

	err = ctx.ShouldBind(&userReq)
	if err != nil {
		err = errors.New("parameter yang dimasukkan salah")
		return
	}

	err = userReq.ValidateSignUp()
	if err != nil {
		return
	}

	user, err := userReq.ConvertToModelForSignUp()
	if err != nil {
		return
	}

	err = service.userRepository.SignUp(user)
	if err != nil {
		err = errors.New("sign up gagal")
		return
	}

	return
}
