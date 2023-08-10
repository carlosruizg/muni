// Code generated by ent, DO NOT EDIT.

package qualification

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/99designs/gqlgen/graphql"
	"github.com/carlosruizg/muni/ent/schema"
)

const (
	// Label holds the string label denoting the qualification type in the database.
	Label = "qualification"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// EdgeTasks holds the string denoting the tasks edge name in mutations.
	EdgeTasks = "tasks"
	// EdgeExperts holds the string denoting the experts edge name in mutations.
	EdgeExperts = "experts"
	// Table holds the table name of the qualification in the database.
	Table = "qualifications"
	// TasksTable is the table that holds the tasks relation/edge. The primary key declared below.
	TasksTable = "labelling_task_expert_requirements"
	// TasksInverseTable is the table name for the LabellingTask entity.
	// It exists in this package in order to avoid circular dependency with the "labellingtask" package.
	TasksInverseTable = "labelling_tasks"
	// ExpertsTable is the table that holds the experts relation/edge. The primary key declared below.
	ExpertsTable = "expert_qualifications"
	// ExpertsInverseTable is the table name for the Expert entity.
	// It exists in this package in order to avoid circular dependency with the "expert" package.
	ExpertsInverseTable = "experts"
)

// Columns holds all SQL columns for qualification fields.
var Columns = []string{
	FieldID,
	FieldValue,
}

var (
	// TasksPrimaryKey and TasksColumn2 are the table columns denoting the
	// primary key for the tasks relation (M2M).
	TasksPrimaryKey = []string{"labelling_task_id", "qualification_id"}
	// ExpertsPrimaryKey and ExpertsColumn2 are the table columns denoting the
	// primary key for the experts relation (M2M).
	ExpertsPrimaryKey = []string{"expert_id", "qualification_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// ValueValidator is a validator for the "value" field enum values. It is called by the builders before save.
func ValueValidator(v schema.QualificationValue) error {
	switch v {
	case "CODER", "MEDICAL", "ARTIST", "STEM":
		return nil
	default:
		return fmt.Errorf("qualification: invalid enum value for value field: %q", v)
	}
}

// OrderOption defines the ordering options for the Qualification queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByTasksCount orders the results by tasks count.
func ByTasksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTasksStep(), opts...)
	}
}

// ByTasks orders the results by tasks terms.
func ByTasks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTasksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByExpertsCount orders the results by experts count.
func ByExpertsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExpertsStep(), opts...)
	}
}

// ByExperts orders the results by experts terms.
func ByExperts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExpertsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTasksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TasksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TasksTable, TasksPrimaryKey...),
	)
}
func newExpertsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExpertsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ExpertsTable, ExpertsPrimaryKey...),
	)
}

var (
	// schema.QualificationValue must implement graphql.Marshaler.
	_ graphql.Marshaler = (*schema.QualificationValue)(nil)
	// schema.QualificationValue must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*schema.QualificationValue)(nil)
)