// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/carlosruizg/muni/ent/labellingtask"
	"github.com/carlosruizg/muni/ent/labellingtaskresponse"
)

// LabellingTaskResponseCreate is the builder for creating a LabellingTaskResponse entity.
type LabellingTaskResponseCreate struct {
	config
	mutation *LabellingTaskResponseMutation
	hooks    []Hook
}

// SetResponse sets the "response" field.
func (ltrc *LabellingTaskResponseCreate) SetResponse(s string) *LabellingTaskResponseCreate {
	ltrc.mutation.SetResponse(s)
	return ltrc
}

// SetTaskID sets the "task" edge to the LabellingTask entity by ID.
func (ltrc *LabellingTaskResponseCreate) SetTaskID(id int) *LabellingTaskResponseCreate {
	ltrc.mutation.SetTaskID(id)
	return ltrc
}

// SetNillableTaskID sets the "task" edge to the LabellingTask entity by ID if the given value is not nil.
func (ltrc *LabellingTaskResponseCreate) SetNillableTaskID(id *int) *LabellingTaskResponseCreate {
	if id != nil {
		ltrc = ltrc.SetTaskID(*id)
	}
	return ltrc
}

// SetTask sets the "task" edge to the LabellingTask entity.
func (ltrc *LabellingTaskResponseCreate) SetTask(l *LabellingTask) *LabellingTaskResponseCreate {
	return ltrc.SetTaskID(l.ID)
}

// Mutation returns the LabellingTaskResponseMutation object of the builder.
func (ltrc *LabellingTaskResponseCreate) Mutation() *LabellingTaskResponseMutation {
	return ltrc.mutation
}

// Save creates the LabellingTaskResponse in the database.
func (ltrc *LabellingTaskResponseCreate) Save(ctx context.Context) (*LabellingTaskResponse, error) {
	return withHooks(ctx, ltrc.sqlSave, ltrc.mutation, ltrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ltrc *LabellingTaskResponseCreate) SaveX(ctx context.Context) *LabellingTaskResponse {
	v, err := ltrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ltrc *LabellingTaskResponseCreate) Exec(ctx context.Context) error {
	_, err := ltrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltrc *LabellingTaskResponseCreate) ExecX(ctx context.Context) {
	if err := ltrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ltrc *LabellingTaskResponseCreate) check() error {
	if _, ok := ltrc.mutation.Response(); !ok {
		return &ValidationError{Name: "response", err: errors.New(`ent: missing required field "LabellingTaskResponse.response"`)}
	}
	return nil
}

func (ltrc *LabellingTaskResponseCreate) sqlSave(ctx context.Context) (*LabellingTaskResponse, error) {
	if err := ltrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ltrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ltrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ltrc.mutation.id = &_node.ID
	ltrc.mutation.done = true
	return _node, nil
}

func (ltrc *LabellingTaskResponseCreate) createSpec() (*LabellingTaskResponse, *sqlgraph.CreateSpec) {
	var (
		_node = &LabellingTaskResponse{config: ltrc.config}
		_spec = sqlgraph.NewCreateSpec(labellingtaskresponse.Table, sqlgraph.NewFieldSpec(labellingtaskresponse.FieldID, field.TypeInt))
	)
	if value, ok := ltrc.mutation.Response(); ok {
		_spec.SetField(labellingtaskresponse.FieldResponse, field.TypeString, value)
		_node.Response = value
	}
	if nodes := ltrc.mutation.TaskIDs(); len(nodes) > 0 {
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
		_node.labelling_task_responses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LabellingTaskResponseCreateBulk is the builder for creating many LabellingTaskResponse entities in bulk.
type LabellingTaskResponseCreateBulk struct {
	config
	builders []*LabellingTaskResponseCreate
}

// Save creates the LabellingTaskResponse entities in the database.
func (ltrcb *LabellingTaskResponseCreateBulk) Save(ctx context.Context) ([]*LabellingTaskResponse, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ltrcb.builders))
	nodes := make([]*LabellingTaskResponse, len(ltrcb.builders))
	mutators := make([]Mutator, len(ltrcb.builders))
	for i := range ltrcb.builders {
		func(i int, root context.Context) {
			builder := ltrcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LabellingTaskResponseMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ltrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ltrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ltrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ltrcb *LabellingTaskResponseCreateBulk) SaveX(ctx context.Context) []*LabellingTaskResponse {
	v, err := ltrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ltrcb *LabellingTaskResponseCreateBulk) Exec(ctx context.Context) error {
	_, err := ltrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltrcb *LabellingTaskResponseCreateBulk) ExecX(ctx context.Context) {
	if err := ltrcb.Exec(ctx); err != nil {
		panic(err)
	}
}