package schema

import (
	"demo/domain"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Note holds the schema definition for the Note entity.
type Note struct {
	ent.Schema
}

// Fields of the Note.
func (Note) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").NotEmpty().Immutable(),
		field.String("external_reference").GoType(&domain.ExternalReference{}).Nillable().Optional(),
	}
}

// Edges of the Note.
func (Note) Edges() []ent.Edge {
	return nil
}
