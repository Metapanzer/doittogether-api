package repository

import (
	"DoItTogether/internal/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(c *gin.Context, user *entity.User) (*entity.User, error)
	GetUserByID(c *gin.Context, id int) (*entity.User, error)
	GetUsers(c *gin.Context) ([]entity.User, error)
	GetUserByEmail(c *gin.Context, email string) (*entity.User, error)
	UpdateUser(c *gin.Context, id int, user *entity.User) (*entity.User, error)
	DeleteUser(c *gin.Context, id int) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) CreateUser(c *gin.Context, user *entity.User) (*entity.User, error) {
	if err := ur.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) GetUserByID(c *gin.Context, id int) (*entity.User, error) {
	var user entity.User
	if err := ur.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) GetUsers(c *gin.Context) ([]entity.User, error) {
	var users []entity.User
	if err := ur.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetUserByEmail(c *gin.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := ur.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) UpdateUser(c *gin.Context, id int, user *entity.User) (*entity.User, error) {
	if err := ur.DB.Model(&entity.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) DeleteUser(c *gin.Context, id int) error {
	if err := ur.DB.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
