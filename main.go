package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

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

	db, err := gorm.Open("postgres", "host=127.0.0.1 user=mike-work dbname=graphql sslmode=disable password=''")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// serve HTTP
	http.Handle("/graphql", h)
	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
