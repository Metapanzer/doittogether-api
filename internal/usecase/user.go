package usecase

import (
	"DoItTogether/internal/entity"
	"DoItTogether/internal/model"
	"DoItTogether/internal/model/mapper"
	"DoItTogether/internal/repository"
)

type UserUsecase interface {
	Register(req *model.RegisterUserRequest) (*entity.User, error)
	GetUser(id int) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	UpdateUser(id int, user *entity.User) (*entity.User, error)
	DeleteUser(id int) error
}

type userUsecase struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		UserRepository: userRepository,
	}
}

func (us *userUsecase) Register(req *model.RegisterUserRequest) (*entity.User, error) {
	newUser := mapper.ToUserEntity(req)
	user, err := us.UserRepository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userUsecase) GetUser(id int) (*entity.User, error) {
	user, err := us.UserRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userUsecase) GetUsers() ([]entity.User, error) {
	users, err := us.UserRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *userUsecase) UpdateUser(id int, user *entity.User) (*entity.User, error) {
	user, err := us.UserRepository.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userUsecase) DeleteUser(id int) error {
	err := us.UserRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
