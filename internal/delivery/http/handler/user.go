package handler

import (
	"DoItTogether/internal/model"
	"DoItTogether/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (uh *UserHandler) Register(c *gin.Context) {
	var req model.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
	}

	user, err := uh.UserUsecase.Register(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
		})
	}
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    user,
	})

}
