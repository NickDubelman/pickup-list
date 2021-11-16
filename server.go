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
	"github.com/rs/cors"

	"github.com/NickDubelman/pickup-list/auth"
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

	graphQLServer := handler.NewDefaultServer(graph.NewSchema(client))

	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", auth.Middleware(cors.Default().Handler(graphQLServer)))

	auth.RouteHandlers(client, mux)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
