package main

import (
	"DoItTogether/internal/adapter/config"
	"DoItTogether/internal/adapter/storage/mysql"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	mysql.New(cfg.DB)

}
