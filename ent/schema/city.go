package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// City holds the schema definition for the City entity.
type City struct {
	ent.Schema
}

// Fields of the City.
func (City) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id_tara"),
		field.String("Nume_oras").Unique(),
		field.Float("Latitudine"),
		field.Float("Longitudine"),
	}
}

// Edges of the City.
func (City) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Temperatures", Temperature.Type).
			StorageKey(edge.Column("id_oras")),
	}
}
