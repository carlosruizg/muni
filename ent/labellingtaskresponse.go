// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/carlosruizg/muni/ent/labellingtask"
	"github.com/carlosruizg/muni/ent/labellingtaskresponse"
)

// LabellingTaskResponse is the model entity for the LabellingTaskResponse schema.
type LabellingTaskResponse struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Response holds the value of the "response" field.
	Response string `json:"response,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LabellingTaskResponseQuery when eager-loading is set.
	Edges                    LabellingTaskResponseEdges `json:"edges"`
	labelling_task_responses *int
	selectValues             sql.SelectValues
}

// LabellingTaskResponseEdges holds the relations/edges for other nodes in the graph.
type LabellingTaskResponseEdges struct {
	// Task holds the value of the task edge.
	Task *LabellingTask `json:"task,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LabellingTaskResponseEdges) TaskOrErr() (*LabellingTask, error) {
	if e.loadedTypes[0] {
		if e.Task == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: labellingtask.Label}
		}
		return e.Task, nil
	}
	return nil, &NotLoadedError{edge: "task"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LabellingTaskResponse) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case labellingtaskresponse.FieldID:
			values[i] = new(sql.NullInt64)
		case labellingtaskresponse.FieldResponse:
			values[i] = new(sql.NullString)
		case labellingtaskresponse.ForeignKeys[0]: // labelling_task_responses
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LabellingTaskResponse fields.
func (ltr *LabellingTaskResponse) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case labellingtaskresponse.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ltr.ID = int(value.Int64)
		case labellingtaskresponse.FieldResponse:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field response", values[i])
			} else if value.Valid {
				ltr.Response = value.String
			}
		case labellingtaskresponse.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field labelling_task_responses", value)
			} else if value.Valid {
				ltr.labelling_task_responses = new(int)
				*ltr.labelling_task_responses = int(value.Int64)
			}
		default:
			ltr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LabellingTaskResponse.
// This includes values selected through modifiers, order, etc.
func (ltr *LabellingTaskResponse) Value(name string) (ent.Value, error) {
	return ltr.selectValues.Get(name)
}

// QueryTask queries the "task" edge of the LabellingTaskResponse entity.
func (ltr *LabellingTaskResponse) QueryTask() *LabellingTaskQuery {
	return NewLabellingTaskResponseClient(ltr.config).QueryTask(ltr)
}

// Update returns a builder for updating this LabellingTaskResponse.
// Note that you need to call LabellingTaskResponse.Unwrap() before calling this method if this LabellingTaskResponse
// was returned from a transaction, and the transaction was committed or rolled back.
func (ltr *LabellingTaskResponse) Update() *LabellingTaskResponseUpdateOne {
	return NewLabellingTaskResponseClient(ltr.config).UpdateOne(ltr)
}

// Unwrap unwraps the LabellingTaskResponse entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ltr *LabellingTaskResponse) Unwrap() *LabellingTaskResponse {
	_tx, ok := ltr.config.driver.(*txDriver)
	if !ok {
		panic("ent: LabellingTaskResponse is not a transactional entity")
	}
	ltr.config.driver = _tx.drv
	return ltr
}

// String implements the fmt.Stringer.
func (ltr *LabellingTaskResponse) String() string {
	var builder strings.Builder
	builder.WriteString("LabellingTaskResponse(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ltr.ID))
	builder.WriteString("response=")
	builder.WriteString(ltr.Response)
	builder.WriteByte(')')
	return builder.String()
}

// LabellingTaskResponses is a parsable slice of LabellingTaskResponse.
type LabellingTaskResponses []*LabellingTaskResponse
