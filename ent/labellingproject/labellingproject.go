// Code generated by ent, DO NOT EDIT.

package labellingproject

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the labellingproject type in the database.
	Label = "labelling_project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIsPrivate holds the string denoting the is_private field in the database.
	FieldIsPrivate = "is_private"
	// FieldCallbackURL holds the string denoting the callback_url field in the database.
	FieldCallbackURL = "callback_url"
	// FieldWorkersPerTask holds the string denoting the workers_per_task field in the database.
	FieldWorkersPerTask = "workers_per_task"
	// Table holds the table name of the labellingproject in the database.
	Table = "labelling_projects"
)

// Columns holds all SQL columns for labellingproject fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldStatus,
	FieldDescription,
	FieldIsPrivate,
	FieldCallbackURL,
	FieldWorkersPerTask,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultIsPrivate holds the default value on creation for the "is_private" field.
	DefaultIsPrivate bool
	// DefaultWorkersPerTask holds the default value on creation for the "workers_per_task" field.
	DefaultWorkersPerTask int
)

// Status defines the type for the "status" enum field.
type Status string

// StatusDraft is the default value of the Status enum.
const DefaultStatus = StatusDraft

// Status values.
const (
	StatusDraft     Status = "draft"
	StatusRunning   Status = "running"
	StatusCancelled Status = "cancelled"
	StatusCompleted Status = "completed"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusDraft, StatusRunning, StatusCancelled, StatusCompleted:
		return nil
	default:
		return fmt.Errorf("labellingproject: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the LabellingProject queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByIsPrivate orders the results by the is_private field.
func ByIsPrivate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPrivate, opts...).ToFunc()
}

// ByCallbackURL orders the results by the callback_url field.
func ByCallbackURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCallbackURL, opts...).ToFunc()
}

// ByWorkersPerTask orders the results by the workers_per_task field.
func ByWorkersPerTask(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkersPerTask, opts...).ToFunc()
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}