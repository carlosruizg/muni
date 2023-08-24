package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
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
	return []ent.Edge{
		edge.From("experts", Expert.Type).Ref("qualifications"),
	}
}

func (Qualification) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
