// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/carlosruizg/muni/ent/expert"
	"github.com/carlosruizg/muni/ent/predicate"
)

// ExpertUpdate is the builder for updating Expert entities.
type ExpertUpdate struct {
	config
	hooks    []Hook
	mutation *ExpertMutation
}

// Where appends a list predicates to the ExpertUpdate builder.
func (eu *ExpertUpdate) Where(ps ...predicate.Expert) *ExpertUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *ExpertUpdate) SetName(s string) *ExpertUpdate {
	eu.mutation.SetName(s)
	return eu
}

// Mutation returns the ExpertMutation object of the builder.
func (eu *ExpertUpdate) Mutation() *ExpertMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *ExpertUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *ExpertUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *ExpertUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *ExpertUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *ExpertUpdate) check() error {
	if v, ok := eu.mutation.Name(); ok {
		if err := expert.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Expert.name": %w`, err)}
		}
	}
	return nil
}

func (eu *ExpertUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(expert.Table, expert.Columns, sqlgraph.NewFieldSpec(expert.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(expert.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{expert.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// ExpertUpdateOne is the builder for updating a single Expert entity.
type ExpertUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExpertMutation
}

// SetName sets the "name" field.
func (euo *ExpertUpdateOne) SetName(s string) *ExpertUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// Mutation returns the ExpertMutation object of the builder.
func (euo *ExpertUpdateOne) Mutation() *ExpertMutation {
	return euo.mutation
}

// Where appends a list predicates to the ExpertUpdate builder.
func (euo *ExpertUpdateOne) Where(ps ...predicate.Expert) *ExpertUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *ExpertUpdateOne) Select(field string, fields ...string) *ExpertUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Expert entity.
func (euo *ExpertUpdateOne) Save(ctx context.Context) (*Expert, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *ExpertUpdateOne) SaveX(ctx context.Context) *Expert {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *ExpertUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *ExpertUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *ExpertUpdateOne) check() error {
	if v, ok := euo.mutation.Name(); ok {
		if err := expert.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Expert.name": %w`, err)}
		}
	}
	return nil
}

func (euo *ExpertUpdateOne) sqlSave(ctx context.Context) (_node *Expert, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(expert.Table, expert.Columns, sqlgraph.NewFieldSpec(expert.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Expert.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, expert.FieldID)
		for _, f := range fields {
			if !expert.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != expert.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(expert.FieldName, field.TypeString, value)
	}
	_node = &Expert{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{expert.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
