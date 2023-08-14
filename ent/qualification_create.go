// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/carlosruizg/muni/ent/expert"
	"github.com/carlosruizg/muni/ent/labellingtask"
	"github.com/carlosruizg/muni/ent/qualification"
	"github.com/carlosruizg/muni/enums"
)

// QualificationCreate is the builder for creating a Qualification entity.
type QualificationCreate struct {
	config
	mutation *QualificationMutation
	hooks    []Hook
}

// SetValue sets the "value" field.
func (qc *QualificationCreate) SetValue(ev enums.QualificationValue) *QualificationCreate {
	qc.mutation.SetValue(ev)
	return qc
}

// AddTaskIDs adds the "tasks" edge to the LabellingTask entity by IDs.
func (qc *QualificationCreate) AddTaskIDs(ids ...int) *QualificationCreate {
	qc.mutation.AddTaskIDs(ids...)
	return qc
}

// AddTasks adds the "tasks" edges to the LabellingTask entity.
func (qc *QualificationCreate) AddTasks(l ...*LabellingTask) *QualificationCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return qc.AddTaskIDs(ids...)
}

// AddExpertIDs adds the "experts" edge to the Expert entity by IDs.
func (qc *QualificationCreate) AddExpertIDs(ids ...int) *QualificationCreate {
	qc.mutation.AddExpertIDs(ids...)
	return qc
}

// AddExperts adds the "experts" edges to the Expert entity.
func (qc *QualificationCreate) AddExperts(e ...*Expert) *QualificationCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return qc.AddExpertIDs(ids...)
}

// Mutation returns the QualificationMutation object of the builder.
func (qc *QualificationCreate) Mutation() *QualificationMutation {
	return qc.mutation
}

// Save creates the Qualification in the database.
func (qc *QualificationCreate) Save(ctx context.Context) (*Qualification, error) {
	return withHooks(ctx, qc.sqlSave, qc.mutation, qc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (qc *QualificationCreate) SaveX(ctx context.Context) *Qualification {
	v, err := qc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qc *QualificationCreate) Exec(ctx context.Context) error {
	_, err := qc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qc *QualificationCreate) ExecX(ctx context.Context) {
	if err := qc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qc *QualificationCreate) check() error {
	if _, ok := qc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Qualification.value"`)}
	}
	if v, ok := qc.mutation.Value(); ok {
		if err := qualification.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "Qualification.value": %w`, err)}
		}
	}
	return nil
}

func (qc *QualificationCreate) sqlSave(ctx context.Context) (*Qualification, error) {
	if err := qc.check(); err != nil {
		return nil, err
	}
	_node, _spec := qc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	qc.mutation.id = &_node.ID
	qc.mutation.done = true
	return _node, nil
}

func (qc *QualificationCreate) createSpec() (*Qualification, *sqlgraph.CreateSpec) {
	var (
		_node = &Qualification{config: qc.config}
		_spec = sqlgraph.NewCreateSpec(qualification.Table, sqlgraph.NewFieldSpec(qualification.FieldID, field.TypeInt))
	)
	if value, ok := qc.mutation.Value(); ok {
		_spec.SetField(qualification.FieldValue, field.TypeEnum, value)
		_node.Value = value
	}
	if nodes := qc.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   qualification.TasksTable,
			Columns: qualification.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(labellingtask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.ExpertsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   qualification.ExpertsTable,
			Columns: qualification.ExpertsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(expert.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// QualificationCreateBulk is the builder for creating many Qualification entities in bulk.
type QualificationCreateBulk struct {
	config
	builders []*QualificationCreate
}

// Save creates the Qualification entities in the database.
func (qcb *QualificationCreateBulk) Save(ctx context.Context) ([]*Qualification, error) {
	specs := make([]*sqlgraph.CreateSpec, len(qcb.builders))
	nodes := make([]*Qualification, len(qcb.builders))
	mutators := make([]Mutator, len(qcb.builders))
	for i := range qcb.builders {
		func(i int, root context.Context) {
			builder := qcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QualificationMutation)
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
					_, err = mutators[i+1].Mutate(root, qcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, qcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (qcb *QualificationCreateBulk) SaveX(ctx context.Context) []*Qualification {
	v, err := qcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qcb *QualificationCreateBulk) Exec(ctx context.Context) error {
	_, err := qcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qcb *QualificationCreateBulk) ExecX(ctx context.Context) {
	if err := qcb.Exec(ctx); err != nil {
		panic(err)
	}
}
