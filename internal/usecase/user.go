package usecase

import (
	"DoItTogether/internal/entity"
	"DoItTogether/internal/repository"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	Register(c *gin.Context, user *entity.User) (*entity.User, error)
	GetUser(c *gin.Context, id int) (*entity.User, error)
	GetUsers(c *gin.Context) ([]entity.User, error)
	UpdateUser(c *gin.Context, id int, user *entity.User) (*entity.User, error)
	DeleteUser(c *gin.Context, id int) error
}

type userUsecase struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		UserRepository: userRepository,
	}
}

func (us *userUsecase) Register(c *gin.Context, user *entity.User) (*entity.User, error) {
	user, err := us.UserRepository.CreateUser(c, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userUsecase) GetUser(c *gin.Context, id int) (*entity.User, error) {
	user, err := us.UserRepository.GetUserByID(c, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userUsecase) GetUsers(c *gin.Context) ([]entity.User, error) {
	users, err := us.UserRepository.GetUsers(c)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *userUsecase) UpdateUser(c *gin.Context, id int, user *entity.User) (*entity.User, error) {
	user, err := us.UserRepository.UpdateUser(c, id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userUsecase) DeleteUser(c *gin.Context, id int) error {
	err := us.UserRepository.DeleteUser(c, id)
	if err != nil {
		return err
	}
	return nil
}
