package ports

import (
	"DoItTogether/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	CreateUser(c *gin.Context, user *domain.User) (*domain.User, error)
	GetUserByID(c *gin.Context, id int) (*domain.User, error)
	GetUsers(c *gin.Context) ([]domain.User, error)
	GetUserByEmail(c *gin.Context, email string) (*domain.User, error)
	UpdateUser(c *gin.Context, id int, user *domain.User) (*domain.User, error)
	DeleteUser(c *gin.Context, id int) error
}

type UserService interface {
	Register(c *gin.Context, user *domain.User) (*domain.User, error)
	GetUser(c *gin.Context, id int) (*domain.User, error)
	GetUsers(c *gin.Context) ([]domain.User, error)
	UpdateUser(c *gin.Context, id int, user *domain.User) (*domain.User, error)
	DeleteUser(c *gin.Context, id int) error
}
