package graph

import (
	"github.com/99designs/gqlgen/graphql"

	"github.com/NickDubelman/pickup-list/db"
	"github.com/NickDubelman/pickup-list/graph/generated"
)

type Resolver struct{ client *db.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *db.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client},
	})
}

// TODO: this exists temporarily, until I've implemented auth
const testUser = "ndubelman@gmail.com"
