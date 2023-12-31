// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/carlosruizg/muni/ent/labellingtask"
	"github.com/carlosruizg/muni/ent/predicate"
)

// LabellingTaskUpdate is the builder for updating LabellingTask entities.
type LabellingTaskUpdate struct {
	config
	hooks    []Hook
	mutation *LabellingTaskMutation
}

// Where appends a list predicates to the LabellingTaskUpdate builder.
func (ltu *LabellingTaskUpdate) Where(ps ...predicate.LabellingTask) *LabellingTaskUpdate {
	ltu.mutation.Where(ps...)
	return ltu
}

// SetName sets the "name" field.
func (ltu *LabellingTaskUpdate) SetName(s string) *LabellingTaskUpdate {
	ltu.mutation.SetName(s)
	return ltu
}

// SetDescription sets the "description" field.
func (ltu *LabellingTaskUpdate) SetDescription(s string) *LabellingTaskUpdate {
	ltu.mutation.SetDescription(s)
	return ltu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ltu *LabellingTaskUpdate) SetNillableDescription(s *string) *LabellingTaskUpdate {
	if s != nil {
		ltu.SetDescription(*s)
	}
	return ltu
}

// ClearDescription clears the value of the "description" field.
func (ltu *LabellingTaskUpdate) ClearDescription() *LabellingTaskUpdate {
	ltu.mutation.ClearDescription()
	return ltu
}

// SetInstructions sets the "instructions" field.
func (ltu *LabellingTaskUpdate) SetInstructions(s string) *LabellingTaskUpdate {
	ltu.mutation.SetInstructions(s)
	return ltu
}

// Mutation returns the LabellingTaskMutation object of the builder.
func (ltu *LabellingTaskUpdate) Mutation() *LabellingTaskMutation {
	return ltu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ltu *LabellingTaskUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ltu.sqlSave, ltu.mutation, ltu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ltu *LabellingTaskUpdate) SaveX(ctx context.Context) int {
	affected, err := ltu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ltu *LabellingTaskUpdate) Exec(ctx context.Context) error {
	_, err := ltu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltu *LabellingTaskUpdate) ExecX(ctx context.Context) {
	if err := ltu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ltu *LabellingTaskUpdate) check() error {
	if v, ok := ltu.mutation.Name(); ok {
		if err := labellingtask.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "LabellingTask.name": %w`, err)}
		}
	}
	return nil
}

func (ltu *LabellingTaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ltu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(labellingtask.Table, labellingtask.Columns, sqlgraph.NewFieldSpec(labellingtask.FieldID, field.TypeInt))
	if ps := ltu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ltu.mutation.Name(); ok {
		_spec.SetField(labellingtask.FieldName, field.TypeString, value)
	}
	if value, ok := ltu.mutation.Description(); ok {
		_spec.SetField(labellingtask.FieldDescription, field.TypeString, value)
	}
	if ltu.mutation.DescriptionCleared() {
		_spec.ClearField(labellingtask.FieldDescription, field.TypeString)
	}
	if value, ok := ltu.mutation.Instructions(); ok {
		_spec.SetField(labellingtask.FieldInstructions, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ltu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{labellingtask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ltu.mutation.done = true
	return n, nil
}

// LabellingTaskUpdateOne is the builder for updating a single LabellingTask entity.
type LabellingTaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LabellingTaskMutation
}

// SetName sets the "name" field.
func (ltuo *LabellingTaskUpdateOne) SetName(s string) *LabellingTaskUpdateOne {
	ltuo.mutation.SetName(s)
	return ltuo
}

// SetDescription sets the "description" field.
func (ltuo *LabellingTaskUpdateOne) SetDescription(s string) *LabellingTaskUpdateOne {
	ltuo.mutation.SetDescription(s)
	return ltuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ltuo *LabellingTaskUpdateOne) SetNillableDescription(s *string) *LabellingTaskUpdateOne {
	if s != nil {
		ltuo.SetDescription(*s)
	}
	return ltuo
}

// ClearDescription clears the value of the "description" field.
func (ltuo *LabellingTaskUpdateOne) ClearDescription() *LabellingTaskUpdateOne {
	ltuo.mutation.ClearDescription()
	return ltuo
}

// SetInstructions sets the "instructions" field.
func (ltuo *LabellingTaskUpdateOne) SetInstructions(s string) *LabellingTaskUpdateOne {
	ltuo.mutation.SetInstructions(s)
	return ltuo
}

// Mutation returns the LabellingTaskMutation object of the builder.
func (ltuo *LabellingTaskUpdateOne) Mutation() *LabellingTaskMutation {
	return ltuo.mutation
}

// Where appends a list predicates to the LabellingTaskUpdate builder.
func (ltuo *LabellingTaskUpdateOne) Where(ps ...predicate.LabellingTask) *LabellingTaskUpdateOne {
	ltuo.mutation.Where(ps...)
	return ltuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ltuo *LabellingTaskUpdateOne) Select(field string, fields ...string) *LabellingTaskUpdateOne {
	ltuo.fields = append([]string{field}, fields...)
	return ltuo
}

// Save executes the query and returns the updated LabellingTask entity.
func (ltuo *LabellingTaskUpdateOne) Save(ctx context.Context) (*LabellingTask, error) {
	return withHooks(ctx, ltuo.sqlSave, ltuo.mutation, ltuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ltuo *LabellingTaskUpdateOne) SaveX(ctx context.Context) *LabellingTask {
	node, err := ltuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ltuo *LabellingTaskUpdateOne) Exec(ctx context.Context) error {
	_, err := ltuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltuo *LabellingTaskUpdateOne) ExecX(ctx context.Context) {
	if err := ltuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ltuo *LabellingTaskUpdateOne) check() error {
	if v, ok := ltuo.mutation.Name(); ok {
		if err := labellingtask.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "LabellingTask.name": %w`, err)}
		}
	}
	return nil
}

func (ltuo *LabellingTaskUpdateOne) sqlSave(ctx context.Context) (_node *LabellingTask, err error) {
	if err := ltuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(labellingtask.Table, labellingtask.Columns, sqlgraph.NewFieldSpec(labellingtask.FieldID, field.TypeInt))
	id, ok := ltuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LabellingTask.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ltuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, labellingtask.FieldID)
		for _, f := range fields {
			if !labellingtask.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != labellingtask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ltuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ltuo.mutation.Name(); ok {
		_spec.SetField(labellingtask.FieldName, field.TypeString, value)
	}
	if value, ok := ltuo.mutation.Description(); ok {
		_spec.SetField(labellingtask.FieldDescription, field.TypeString, value)
	}
	if ltuo.mutation.DescriptionCleared() {
		_spec.ClearField(labellingtask.FieldDescription, field.TypeString)
	}
	if value, ok := ltuo.mutation.Instructions(); ok {
		_spec.SetField(labellingtask.FieldInstructions, field.TypeString, value)
	}
	_node = &LabellingTask{config: ltuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ltuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{labellingtask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ltuo.mutation.done = true
	return _node, nil
}
