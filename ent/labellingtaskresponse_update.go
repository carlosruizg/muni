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
	"github.com/carlosruizg/muni/ent/labellingtaskresponse"
	"github.com/carlosruizg/muni/ent/predicate"
)

// LabellingTaskResponseUpdate is the builder for updating LabellingTaskResponse entities.
type LabellingTaskResponseUpdate struct {
	config
	hooks    []Hook
	mutation *LabellingTaskResponseMutation
}

// Where appends a list predicates to the LabellingTaskResponseUpdate builder.
func (ltru *LabellingTaskResponseUpdate) Where(ps ...predicate.LabellingTaskResponse) *LabellingTaskResponseUpdate {
	ltru.mutation.Where(ps...)
	return ltru
}

// SetResponse sets the "response" field.
func (ltru *LabellingTaskResponseUpdate) SetResponse(s string) *LabellingTaskResponseUpdate {
	ltru.mutation.SetResponse(s)
	return ltru
}

// SetTaskID sets the "task" edge to the LabellingTask entity by ID.
func (ltru *LabellingTaskResponseUpdate) SetTaskID(id int) *LabellingTaskResponseUpdate {
	ltru.mutation.SetTaskID(id)
	return ltru
}

// SetNillableTaskID sets the "task" edge to the LabellingTask entity by ID if the given value is not nil.
func (ltru *LabellingTaskResponseUpdate) SetNillableTaskID(id *int) *LabellingTaskResponseUpdate {
	if id != nil {
		ltru = ltru.SetTaskID(*id)
	}
	return ltru
}

// SetTask sets the "task" edge to the LabellingTask entity.
func (ltru *LabellingTaskResponseUpdate) SetTask(l *LabellingTask) *LabellingTaskResponseUpdate {
	return ltru.SetTaskID(l.ID)
}

// Mutation returns the LabellingTaskResponseMutation object of the builder.
func (ltru *LabellingTaskResponseUpdate) Mutation() *LabellingTaskResponseMutation {
	return ltru.mutation
}

// ClearTask clears the "task" edge to the LabellingTask entity.
func (ltru *LabellingTaskResponseUpdate) ClearTask() *LabellingTaskResponseUpdate {
	ltru.mutation.ClearTask()
	return ltru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ltru *LabellingTaskResponseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ltru.sqlSave, ltru.mutation, ltru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ltru *LabellingTaskResponseUpdate) SaveX(ctx context.Context) int {
	affected, err := ltru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ltru *LabellingTaskResponseUpdate) Exec(ctx context.Context) error {
	_, err := ltru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltru *LabellingTaskResponseUpdate) ExecX(ctx context.Context) {
	if err := ltru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ltru *LabellingTaskResponseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(labellingtaskresponse.Table, labellingtaskresponse.Columns, sqlgraph.NewFieldSpec(labellingtaskresponse.FieldID, field.TypeInt))
	if ps := ltru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ltru.mutation.Response(); ok {
		_spec.SetField(labellingtaskresponse.FieldResponse, field.TypeString, value)
	}
	if ltru.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   labellingtaskresponse.TaskTable,
			Columns: []string{labellingtaskresponse.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(labellingtask.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ltru.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   labellingtaskresponse.TaskTable,
			Columns: []string{labellingtaskresponse.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(labellingtask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ltru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{labellingtaskresponse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ltru.mutation.done = true
	return n, nil
}

// LabellingTaskResponseUpdateOne is the builder for updating a single LabellingTaskResponse entity.
type LabellingTaskResponseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LabellingTaskResponseMutation
}

// SetResponse sets the "response" field.
func (ltruo *LabellingTaskResponseUpdateOne) SetResponse(s string) *LabellingTaskResponseUpdateOne {
	ltruo.mutation.SetResponse(s)
	return ltruo
}

// SetTaskID sets the "task" edge to the LabellingTask entity by ID.
func (ltruo *LabellingTaskResponseUpdateOne) SetTaskID(id int) *LabellingTaskResponseUpdateOne {
	ltruo.mutation.SetTaskID(id)
	return ltruo
}

// SetNillableTaskID sets the "task" edge to the LabellingTask entity by ID if the given value is not nil.
func (ltruo *LabellingTaskResponseUpdateOne) SetNillableTaskID(id *int) *LabellingTaskResponseUpdateOne {
	if id != nil {
		ltruo = ltruo.SetTaskID(*id)
	}
	return ltruo
}

// SetTask sets the "task" edge to the LabellingTask entity.
func (ltruo *LabellingTaskResponseUpdateOne) SetTask(l *LabellingTask) *LabellingTaskResponseUpdateOne {
	return ltruo.SetTaskID(l.ID)
}

// Mutation returns the LabellingTaskResponseMutation object of the builder.
func (ltruo *LabellingTaskResponseUpdateOne) Mutation() *LabellingTaskResponseMutation {
	return ltruo.mutation
}

// ClearTask clears the "task" edge to the LabellingTask entity.
func (ltruo *LabellingTaskResponseUpdateOne) ClearTask() *LabellingTaskResponseUpdateOne {
	ltruo.mutation.ClearTask()
	return ltruo
}

// Where appends a list predicates to the LabellingTaskResponseUpdate builder.
func (ltruo *LabellingTaskResponseUpdateOne) Where(ps ...predicate.LabellingTaskResponse) *LabellingTaskResponseUpdateOne {
	ltruo.mutation.Where(ps...)
	return ltruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ltruo *LabellingTaskResponseUpdateOne) Select(field string, fields ...string) *LabellingTaskResponseUpdateOne {
	ltruo.fields = append([]string{field}, fields...)
	return ltruo
}

// Save executes the query and returns the updated LabellingTaskResponse entity.
func (ltruo *LabellingTaskResponseUpdateOne) Save(ctx context.Context) (*LabellingTaskResponse, error) {
	return withHooks(ctx, ltruo.sqlSave, ltruo.mutation, ltruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ltruo *LabellingTaskResponseUpdateOne) SaveX(ctx context.Context) *LabellingTaskResponse {
	node, err := ltruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ltruo *LabellingTaskResponseUpdateOne) Exec(ctx context.Context) error {
	_, err := ltruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltruo *LabellingTaskResponseUpdateOne) ExecX(ctx context.Context) {
	if err := ltruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ltruo *LabellingTaskResponseUpdateOne) sqlSave(ctx context.Context) (_node *LabellingTaskResponse, err error) {
	_spec := sqlgraph.NewUpdateSpec(labellingtaskresponse.Table, labellingtaskresponse.Columns, sqlgraph.NewFieldSpec(labellingtaskresponse.FieldID, field.TypeInt))
	id, ok := ltruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LabellingTaskResponse.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ltruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, labellingtaskresponse.FieldID)
		for _, f := range fields {
			if !labellingtaskresponse.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != labellingtaskresponse.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ltruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ltruo.mutation.Response(); ok {
		_spec.SetField(labellingtaskresponse.FieldResponse, field.TypeString, value)
	}
	if ltruo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   labellingtaskresponse.TaskTable,
			Columns: []string{labellingtaskresponse.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(labellingtask.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ltruo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   labellingtaskresponse.TaskTable,
			Columns: []string{labellingtaskresponse.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(labellingtask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LabellingTaskResponse{config: ltruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ltruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{labellingtaskresponse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ltruo.mutation.done = true
	return _node, nil
}