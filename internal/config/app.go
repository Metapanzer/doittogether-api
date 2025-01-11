package config

import (
	"DoItTogether/internal/delivery/http/handler"
	"DoItTogether/internal/delivery/http/router"
	"DoItTogether/internal/repository"
	"DoItTogether/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppConfig struct {
	DB     *gorm.DB
	App    *gin.Engine
	Config *viper.Viper
}

func (cfg *AppConfig) Init() {
	// Setup Repositories
	UserRepository := repository.NewUserRepository(cfg.DB)

	// Setup Use Cases
	UserUsecase := usecase.NewUserUsecase(UserRepository)

	// Setup Handlers
	UserHandler := handler.NewUserHandler(UserUsecase)

	// Setup Routes
	router := router.NewRouter(cfg.App, UserHandler)

	router.Setup()
}
