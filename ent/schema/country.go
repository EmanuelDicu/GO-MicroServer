package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.String("Nume_tara").Unique(),
		field.Float("Latitudine"),
		field.Float("Longitudine"),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Cities", City.Type).
			StorageKey(edge.Column("id_tara")),
	}
}
