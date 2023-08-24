// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/carlosruizg/muni/ent/expert"
	"github.com/carlosruizg/muni/ent/predicate"
)

// ExpertDelete is the builder for deleting a Expert entity.
type ExpertDelete struct {
	config
	hooks    []Hook
	mutation *ExpertMutation
}

// Where appends a list predicates to the ExpertDelete builder.
func (ed *ExpertDelete) Where(ps ...predicate.Expert) *ExpertDelete {
	ed.mutation.Where(ps...)
	return ed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ed *ExpertDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ed.sqlExec, ed.mutation, ed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ed *ExpertDelete) ExecX(ctx context.Context) int {
	n, err := ed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ed *ExpertDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(expert.Table, sqlgraph.NewFieldSpec(expert.FieldID, field.TypeInt))
	if ps := ed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ed.mutation.done = true
	return affected, err
}

// ExpertDeleteOne is the builder for deleting a single Expert entity.
type ExpertDeleteOne struct {
	ed *ExpertDelete
}

// Where appends a list predicates to the ExpertDelete builder.
func (edo *ExpertDeleteOne) Where(ps ...predicate.Expert) *ExpertDeleteOne {
	edo.ed.mutation.Where(ps...)
	return edo
}

// Exec executes the deletion query.
func (edo *ExpertDeleteOne) Exec(ctx context.Context) error {
	n, err := edo.ed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{expert.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (edo *ExpertDeleteOne) ExecX(ctx context.Context) {
	if err := edo.Exec(ctx); err != nil {
		panic(err)
	}
}