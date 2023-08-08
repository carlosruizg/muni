// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/carlosruizg/muni/ent/labellingtask"
)

// LabellingTask is the model entity for the LabellingTask schema.
type LabellingTask struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LabellingTaskQuery when eager-loading is set.
	Edges        LabellingTaskEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LabellingTaskEdges holds the relations/edges for other nodes in the graph.
type LabellingTaskEdges struct {
	// Responses holds the value of the responses edge.
	Responses []*LabellingTaskResponse `json:"responses,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedResponses map[string][]*LabellingTaskResponse
}

// ResponsesOrErr returns the Responses value or an error if the edge
// was not loaded in eager-loading.
func (e LabellingTaskEdges) ResponsesOrErr() ([]*LabellingTaskResponse, error) {
	if e.loadedTypes[0] {
		return e.Responses, nil
	}
	return nil, &NotLoadedError{edge: "responses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LabellingTask) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case labellingtask.FieldID:
			values[i] = new(sql.NullInt64)
		case labellingtask.FieldTitle, labellingtask.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LabellingTask fields.
func (lt *LabellingTask) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case labellingtask.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lt.ID = int(value.Int64)
		case labellingtask.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				lt.Title = value.String
			}
		case labellingtask.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				lt.Description = value.String
			}
		default:
			lt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LabellingTask.
// This includes values selected through modifiers, order, etc.
func (lt *LabellingTask) Value(name string) (ent.Value, error) {
	return lt.selectValues.Get(name)
}

// QueryResponses queries the "responses" edge of the LabellingTask entity.
func (lt *LabellingTask) QueryResponses() *LabellingTaskResponseQuery {
	return NewLabellingTaskClient(lt.config).QueryResponses(lt)
}

// Update returns a builder for updating this LabellingTask.
// Note that you need to call LabellingTask.Unwrap() before calling this method if this LabellingTask
// was returned from a transaction, and the transaction was committed or rolled back.
func (lt *LabellingTask) Update() *LabellingTaskUpdateOne {
	return NewLabellingTaskClient(lt.config).UpdateOne(lt)
}

// Unwrap unwraps the LabellingTask entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lt *LabellingTask) Unwrap() *LabellingTask {
	_tx, ok := lt.config.driver.(*txDriver)
	if !ok {
		panic("ent: LabellingTask is not a transactional entity")
	}
	lt.config.driver = _tx.drv
	return lt
}

// String implements the fmt.Stringer.
func (lt *LabellingTask) String() string {
	var builder strings.Builder
	builder.WriteString("LabellingTask(")
	builder.WriteString(fmt.Sprintf("id=%v, ", lt.ID))
	builder.WriteString("title=")
	builder.WriteString(lt.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(lt.Description)
	builder.WriteByte(')')
	return builder.String()
}

// NamedResponses returns the Responses named value or an error if the edge was not
// loaded in eager-loading with this name.
func (lt *LabellingTask) NamedResponses(name string) ([]*LabellingTaskResponse, error) {
	if lt.Edges.namedResponses == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := lt.Edges.namedResponses[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (lt *LabellingTask) appendNamedResponses(name string, edges ...*LabellingTaskResponse) {
	if lt.Edges.namedResponses == nil {
		lt.Edges.namedResponses = make(map[string][]*LabellingTaskResponse)
	}
	if len(edges) == 0 {
		lt.Edges.namedResponses[name] = []*LabellingTaskResponse{}
	} else {
		lt.Edges.namedResponses[name] = append(lt.Edges.namedResponses[name], edges...)
	}
}

// LabellingTasks is a parsable slice of LabellingTask.
type LabellingTasks []*LabellingTask
