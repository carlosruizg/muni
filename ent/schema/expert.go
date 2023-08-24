package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Expert holds the schema definition for the Expert entity.
type Expert struct {
	ent.Schema
}

// Fields of the Expert.
func (Expert) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the Expert.
func (Expert) Edges() []ent.Edge {
	return nil
}
