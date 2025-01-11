package mapper

import (
	"DoItTogether/internal/entity"
	"DoItTogether/internal/model"
)

func ToUserEntity(req *model.RegisterUserRequest) *entity.User {
	return &entity.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: req.Password,
		// AvatarFilename: req.AvatarFilename,
		// Role:           "user",
		// Token:          "",
	}
}
