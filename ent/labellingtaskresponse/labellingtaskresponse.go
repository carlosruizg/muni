// Code generated by ent, DO NOT EDIT.

package labellingtaskresponse

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the labellingtaskresponse type in the database.
	Label = "labelling_task_response"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldResponse holds the string denoting the response field in the database.
	FieldResponse = "response"
	// Table holds the table name of the labellingtaskresponse in the database.
	Table = "labelling_task_responses"
)

// Columns holds all SQL columns for labellingtaskresponse fields.
var Columns = []string{
	FieldID,
	FieldResponse,
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

// OrderOption defines the ordering options for the LabellingTaskResponse queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByResponse orders the results by the response field.
func ByResponse(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResponse, opts...).ToFunc()
}
