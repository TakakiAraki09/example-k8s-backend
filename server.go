package main

import (
	// "context"

	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/TakakiAraki09/example-k8s-backend/constants"
	"github.com/TakakiAraki09/example-k8s-backend/internal/database"
	"github.com/TakakiAraki09/example-k8s-backend/internal/generated"
	"github.com/TakakiAraki09/example-k8s-backend/resolver"
	"github.com/TakakiAraki09/example-k8s-backend/service"
	"github.com/TakakiAraki09/example-k8s-backend/utils"

	// "github.com/TakakiAraki09/example-k8s-backend/internal/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv(constants.EnvPort)
	if port == "" {
		port = constants.DefaultPort
	}

	log.Printf("db access")
	db, err := sql.Open("mysql", utils.CreateDBUrlString(utils.DatabaseMetadata{
		User:     os.Getenv(constants.EnvDbUser),
		Password: os.Getenv(constants.EnvDbPassword),
		Host:     os.Getenv(constants.EnvDbHost),
		Port:     os.Getenv(constants.EnvDbPort),
		Table:    "example_database",
	}))
	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Println("Waiting for DB to be ready... retrying")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("DB not ready after retries:", err)
	}
	defer db.Close()

	queries := database.New(db)
	server := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		Service: service.Service{
			Queries: queries,
		},
	}}))
	server.AddTransport(transport.Options{})
	server.AddTransport(transport.GET{})
	server.AddTransport(transport.POST{})
	server.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	server.Use(extension.Introspection{})
	server.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
