package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Temperature holds the schema definition for the Temperature entity.
type Temperature struct {
	ent.Schema
}

// Fields of the Temperature.
func (Temperature) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id_oras"),
		field.Float("Valoare"),
		field.Time("Timestamp"),
	}
}
