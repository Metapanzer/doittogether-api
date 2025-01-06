package main

import (
	"DoItTogether/internal/config"
	"DoItTogether/internal/entity"
	"DoItTogether/internal/repository"
	"DoItTogether/internal/usecase"
	"log"
)

func main() {
	// Load configuration
	viperConfig := config.LoadConfig()

	// Connect to MySQL
	DBConn, err := config.NewDatabase(viperConfig)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	DBConn.AutoMigrate(&entity.User{}, &entity.Campaign{}, &entity.CampaignImage{}, &entity.Transaction{})

	// Initialize repository
	UserRepository := repository.NewUserRepository(DBConn)

	// Initialize service
	UserUsecase := usecase.NewUserUsecase(UserRepository)

}
