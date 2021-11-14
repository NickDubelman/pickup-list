package schema

import (
	"entgo.io/ent"
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
	return nil
}
