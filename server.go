package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/TakakiAraki09/k8s-lesson/constants"
	"github.com/TakakiAraki09/k8s-lesson/database"
	"github.com/TakakiAraki09/k8s-lesson/graph"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	err := godotenv.Load(".env.local")
	ctx := context.Background()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv(constants.EnvPort)
	if port == "" {
		port = constants.DefaultPort
	}
	db, err := sql.Open("mysql", "root:1f2d1e2e67df@/example_database?parseTime=true")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	queries := database.New(db)
	// CreateExampleContentParams

	// queries.CreateExampleContent(ctx,
	// 	database.CreateExampleContentParams{
	// 		Title: sql.NullString{String: "Hello World", Valid: true},
	// 		Hoge:  sql.NullString{String: "Hoge", Valid: true},
	// 	})
	result, err := queries.ExampleOne(ctx)
	if err != nil {
		log.Fatal("hogehoge")
		log.Fatal(err)
	}

	if result.Title.Valid {
		log.Printf(result.Title.String)
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
