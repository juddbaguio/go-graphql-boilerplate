package main

import (
	"context"
	"go-graphql-boilerplate/config"
	"go-graphql-boilerplate/domain/data"
	"go-graphql-boilerplate/domain/usecase"
	"go-graphql-boilerplate/graph"
	"go-graphql-boilerplate/infra"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	container, err := config.SetupConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	mongoClient, err := infra.ConnectMongo(ctx, &container.Mongo)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	userRepo := data.NewUserRepo(mongoClient)
	todoRepo := data.NewTodoRepo(mongoClient)

	resolver := usecase.New(userRepo, todoRepo)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
