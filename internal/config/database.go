package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(viper *viper.Viper) (*gorm.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("User"),
		viper.GetString("Password"),
		viper.GetString("Host"),
		viper.GetString("Port"),
		viper.GetString("Name"),
	)
	dsn := connStr
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to database")
	return db, nil
}
