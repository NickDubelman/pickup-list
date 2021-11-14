package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NBAPlayer holds the schema definition for the NBAPlayer entity.
type NBAPlayer struct {
	ent.Schema
}

// Fields of the NBAPlayer.
func (NBAPlayer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the NBAPlayer.
func (NBAPlayer) Edges() []ent.Edge {
	return []ent.Edge{
		// an NBA player can only be referenced by one user
		edge.From("user", User.Type). 
			Ref("nba_player").
			Unique().
			Annotations(entgql.Bind()),
	}
}
