package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/saad-karim/graphql-learning/graph"
	"github.com/saad-karim/graphql-learning/graph/generated"
	"github.com/saad-karim/graphql-learning/internal/database"
)

const defaultPort = "7070"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	localDB := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, dbport, user, pass, name)
	db, err := sqlx.Connect("postgres", localDB)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Event: &database.EventDatabase{DB: db},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
