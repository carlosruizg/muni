package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/carlosruizg/muni/enums"
)

// Qualification holds the schema definition for the Qualification entity.
type Qualification struct {
	ent.Schema
}

// Fields of the Qualification.
func (Qualification) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("value").GoType(enums.QualificationValue(enums.Other)),
	}
}

// Edges of the Qualification.
func (Qualification) Edges() []ent.Edge {
	return nil
}
