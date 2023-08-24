package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// LabellingTask holds the schema definition for the LabellingTask entity.
type LabellingTask struct {
	ent.Schema
}

// Fields of the LabellingTask.
func (LabellingTask) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.String("instructions"),
	}
}

// Edges of the LabellingTask.
func (LabellingTask) Edges() []ent.Edge {
	return nil
}

func (LabellingTask) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
