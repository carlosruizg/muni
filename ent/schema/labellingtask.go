package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LabellingTask holds the schema definition for the LabellingTask entity.
type LabellingTask struct {
	ent.Schema
}

// Fields of the LabellingTask.
func (LabellingTask) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("description").Optional(),
	}
}

// Edges of the LabellingTask.
func (LabellingTask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("responses", LabellingTaskResponse.Type),
		edge.To("expert_requirements", Qualification.Type),
	}
}

func (LabellingTask) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
