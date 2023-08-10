package schema

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Qualification holds the schema definition for the Qualification entity.
type Qualification struct {
	ent.Schema
}

// Fields of the Qualification.
func (Qualification) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("value").GoType(QualificationValue(Other)),
	}
}

// Edges of the Qualification.
func (Qualification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tasks", LabellingTask.Type).Ref("expert_requirements"),
		edge.From("experts", Expert.Type).Ref("qualifications"),
	}
}

type QualificationValue string

const (
	Coder   QualificationValue = "CODER"
	Medical QualificationValue = "MEDICAL"
	Artist  QualificationValue = "ARTIST"
	STEM    QualificationValue = "STEM"
	Other   QualificationValue = "OTHER"
)

// Values provides list valid values for Enum.
func (QualificationValue) Values() (kinds []string) {
	for _, s := range []QualificationValue{Coder, Medical, Artist, STEM} {
		kinds = append(kinds, string(s))
	}
	return
}

// MarshalGQL implements graphql.Marshaler interface.
func (_g QualificationValue) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(string(_g)))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (_g *QualificationValue) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*_g = QualificationValue(str)

	switch *_g {
	case Coder, Medical, Artist, STEM, Other:
		return nil
	default:
		return fmt.Errorf("%s is not a valid Gender", str)
	}
}
