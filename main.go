package main

import (
	"fmt"
	"log"

	"github.com/freexet/raven/auth"
	"github.com/freexet/raven/repository"
	env "github.com/joho/godotenv"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	auth.AutoMigrate(db)
}

func main() {
	env.Load()

	db := repository.New()
	if db == nil {
		log.Fatal("Error connection to db")
	}
	migrate(db)

	fmt.Println("Hello World!")
}
