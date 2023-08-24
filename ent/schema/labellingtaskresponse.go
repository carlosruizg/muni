package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// LabellingTaskResponse holds the schema definition for the LabellingTaskResponse entity.
type LabellingTaskResponse struct {
	ent.Schema
}

// Fields of the LabellingTaskResponse.
func (LabellingTaskResponse) Fields() []ent.Field {
	return []ent.Field{
		field.String("response"),
	}
}

// Edges of the LabellingTaskResponse.
func (LabellingTaskResponse) Edges() []ent.Edge {
	return nil
}
