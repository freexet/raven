package main

import (
	"fmt"
	"log"

	"github.com/freexet/raven/repository"
	env "github.com/joho/godotenv"
)

func main() {
	env.Load()

	db := repository.New()
	if db == nil {
		log.Fatal("Error connection to db")
	}

	fmt.Println("Hello World!")
}
