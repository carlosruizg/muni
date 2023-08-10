package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
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
	return []ent.Edge{
		edge.To("task_responses", LabellingTaskResponse.Type),
		edge.To("qualifications", Qualification.Type),
	}
}

func (Expert) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
