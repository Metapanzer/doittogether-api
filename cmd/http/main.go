package main

import (
	"DoItTogether/internal/adapter/config"
	"DoItTogether/internal/adapter/storage/mysql"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// Connect to MySQL
	mysql.New(cfg.DB)

}
