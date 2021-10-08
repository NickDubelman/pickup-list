package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// List holds the schema definition for the List entity.
type List struct {
	ent.Schema
}

// Fields of the List.
func (List) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the List.
func (List) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", User.Type).
			Unique().
			Required().
			Annotations(entgql.Bind()),

		edge.To("users", User.Type).
			Annotations(entgql.Bind()),
	}
}
