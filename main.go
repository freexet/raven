package main

import (
	"fmt"
	"log"

	"github.com/freexet/raven/repository"
)

func main() {
	db := repository.New()
	if db == nil {
		log.Fatal("Error connection to db")
	}

	fmt.Println("Hello World!")
}
