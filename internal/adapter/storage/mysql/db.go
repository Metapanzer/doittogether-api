package mysql

import (
	"DoItTogether/internal/adapter/config"
	"DoItTogether/internal/core/domain"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Conn *gorm.DB
}

func New(config *config.DB) (*DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
	dsn := connStr
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&domain.User{}, &domain.Campaign{}, &domain.CampaignImage{}, &domain.Transaction{})

	fmt.Println("Connected to database")
	return &DB{Conn: db}, nil
}
