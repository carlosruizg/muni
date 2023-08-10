// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/carlosruizg/muni/ent/expert"
)

// Expert is the model entity for the Expert schema.
type Expert struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExpertQuery when eager-loading is set.
	Edges        ExpertEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ExpertEdges holds the relations/edges for other nodes in the graph.
type ExpertEdges struct {
	// TaskResponses holds the value of the task_responses edge.
	TaskResponses []*LabellingTaskResponse `json:"task_responses,omitempty"`
	// Qualifications holds the value of the qualifications edge.
	Qualifications []*Qualification `json:"qualifications,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedTaskResponses  map[string][]*LabellingTaskResponse
	namedQualifications map[string][]*Qualification
}

// TaskResponsesOrErr returns the TaskResponses value or an error if the edge
// was not loaded in eager-loading.
func (e ExpertEdges) TaskResponsesOrErr() ([]*LabellingTaskResponse, error) {
	if e.loadedTypes[0] {
		return e.TaskResponses, nil
	}
	return nil, &NotLoadedError{edge: "task_responses"}
}

// QualificationsOrErr returns the Qualifications value or an error if the edge
// was not loaded in eager-loading.
func (e ExpertEdges) QualificationsOrErr() ([]*Qualification, error) {
	if e.loadedTypes[1] {
		return e.Qualifications, nil
	}
	return nil, &NotLoadedError{edge: "qualifications"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Expert) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case expert.FieldID:
			values[i] = new(sql.NullInt64)
		case expert.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Expert fields.
func (e *Expert) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case expert.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case expert.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Expert.
// This includes values selected through modifiers, order, etc.
func (e *Expert) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryTaskResponses queries the "task_responses" edge of the Expert entity.
func (e *Expert) QueryTaskResponses() *LabellingTaskResponseQuery {
	return NewExpertClient(e.config).QueryTaskResponses(e)
}

// QueryQualifications queries the "qualifications" edge of the Expert entity.
func (e *Expert) QueryQualifications() *QualificationQuery {
	return NewExpertClient(e.config).QueryQualifications(e)
}

// Update returns a builder for updating this Expert.
// Note that you need to call Expert.Unwrap() before calling this method if this Expert
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Expert) Update() *ExpertUpdateOne {
	return NewExpertClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Expert entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Expert) Unwrap() *Expert {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Expert is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Expert) String() string {
	var builder strings.Builder
	builder.WriteString("Expert(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedTaskResponses returns the TaskResponses named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Expert) NamedTaskResponses(name string) ([]*LabellingTaskResponse, error) {
	if e.Edges.namedTaskResponses == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedTaskResponses[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Expert) appendNamedTaskResponses(name string, edges ...*LabellingTaskResponse) {
	if e.Edges.namedTaskResponses == nil {
		e.Edges.namedTaskResponses = make(map[string][]*LabellingTaskResponse)
	}
	if len(edges) == 0 {
		e.Edges.namedTaskResponses[name] = []*LabellingTaskResponse{}
	} else {
		e.Edges.namedTaskResponses[name] = append(e.Edges.namedTaskResponses[name], edges...)
	}
}

// NamedQualifications returns the Qualifications named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Expert) NamedQualifications(name string) ([]*Qualification, error) {
	if e.Edges.namedQualifications == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedQualifications[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Expert) appendNamedQualifications(name string, edges ...*Qualification) {
	if e.Edges.namedQualifications == nil {
		e.Edges.namedQualifications = make(map[string][]*Qualification)
	}
	if len(edges) == 0 {
		e.Edges.namedQualifications[name] = []*Qualification{}
	} else {
		e.Edges.namedQualifications[name] = append(e.Edges.namedQualifications[name], edges...)
	}
}

// Experts is a parsable slice of Expert.
type Experts []*Expert
