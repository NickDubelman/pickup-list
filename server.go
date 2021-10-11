package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"

	"github.com/NickDubelman/pickup-list/db"
	"github.com/NickDubelman/pickup-list/db/migrate"
	"github.com/NickDubelman/pickup-list/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Establish connection to our db
	dsn := "root@tcp(localhost:3381)/pickup_list?parseTime=true"
	client, err := db.Open(dialect.MySQL, dsn)
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	ctx := context.Background()
	err = client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true))
	if err != nil {
		log.Fatal("running schema migration", err)
	}

	// Create a user
	user, err := client.User.Create().
		SetRealName("Nick Dubelman").
		SetNbaName("Alex Caruso").
		SetEmail("ndubelman@gmail.com").
		Save(ctx)
	if err != nil {
		log.Println("could not create user:", err)
	} else {
		log.Println("created user:", user)
	}

	graphQLServer := handler.NewDefaultServer(graph.NewSchema(client))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graphQLServer)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
