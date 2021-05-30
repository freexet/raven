package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/freexet/raven/auth"
	"github.com/freexet/raven/graph"
	"github.com/freexet/raven/graph/generated"
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

	port := os.Getenv("PORT")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
