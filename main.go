package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"
)

var schema, err = graphql.NewSchema(graphql.SchemaConfig{
	Query:    QueryType,
	Mutation: MutationType,
})

func main() {
	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", "postgres://mike-work:''@localhost:5432/graphql?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// serve HTTP
	http.Handle("/graphql", h)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
