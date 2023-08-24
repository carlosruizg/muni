package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// LabellingProject holds the schema definition for the LabellingProject entity.
type LabellingProject struct {
	ent.Schema
}

// Fields of the LabellingProject.
func (LabellingProject) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Enum("status").Values("draft", "running", "cancelled", "completed").Default("draft"),
		field.String("description").Optional(),
		field.Bool("is_private").Default(false),
		field.String("callback_url").Optional(),
		field.Int("workers_per_task").Default(1),
	}
}

// Edges of the LabellingProject.
func (LabellingProject) Edges() []ent.Edge {
	return nil
}

func (LabellingProject) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
