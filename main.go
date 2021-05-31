package main

import (
	"context"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/freexet/raven/auth"
	"github.com/freexet/raven/graph"
	"github.com/freexet/raven/graph/generated"
	"github.com/freexet/raven/repository"
	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	auth.AutoMigrate(db)
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GinContextToContextMiddleware(a auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("auth", a)
		ctx := context.WithValue(c.Request.Context(), "ginCtx", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func main() {
	env.Load()

	port := os.Getenv("PORT")

	r, err := repository.New()
	if err != nil {
		log.Fatal("Error connection to db")
	}
	migrate(r.GetDB())

	a := auth.NewService(r)

	e := gin.Default()
	e.Use(GinContextToContextMiddleware(a))
	e.POST("/graphql", graphqlHandler())
	e.GET("/", playgroundHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(e.Run(":" + port))
}
