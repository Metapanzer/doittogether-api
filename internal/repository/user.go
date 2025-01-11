package repository

import (
	"DoItTogether/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByID(id int) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	UpdateUser(id int, user *entity.User) (*entity.User, error)
	DeleteUser(id int) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) CreateUser(user *entity.User) (*entity.User, error) {
	if err := ur.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) GetUserByID(id int) (*entity.User, error) {
	var user entity.User
	if err := ur.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) GetUsers() ([]entity.User, error) {
	var users []entity.User
	if err := ur.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := ur.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) UpdateUser(id int, user *entity.User) (*entity.User, error) {
	if err := ur.DB.Model(&entity.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) DeleteUser(id int) error {
	if err := ur.DB.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
